package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

	"bitbucket.org/creachadair/stringset"
	"github.com/google/subcommands"
	"github.com/looplab/tarjan"
)

// extraDeps provides a way to manually add dependencies to certain occt
// "packages". These should correspond to external deps defined in your
// WORKSPACE file or elsewhere in your bazel tree.
//
// WARNING: this list is definitely not complete and you may need to modify it,
// depending on which parts of occt you are using.
var extraDeps = map[string][]string{
	"Font":      []string{"@org_freetype_freetype2//:freetype2"},
	"Draw":      []string{"@tcl"},
	"Adaptor2d": []string{"@org_freedesktop_fontconfig//:fontconfig"},
}

type genBazelSubcommand struct {
	depgraphFile string
}

func (*genBazelSubcommand) Name() string { return "genbazel" }
func (*genBazelSubcommand) Synopsis() string {
	return "Reads the dependency graph from a json file and uses it to generate a Bazel BUILD file"
}
func (*genBazelSubcommand) Usage() string {
	return "genbazel --depgraph graph.json"
}
func (s *genBazelSubcommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.depgraphFile, "depgraph", "", "Path to json-formatted depgraph, which is generated by the depgraph subcommand")
}
func (s *genBazelSubcommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if s.depgraphFile == "" {
		log.Fatal("Missing --depgraph")
	}

	b, err := ioutil.ReadFile(s.depgraphFile)
	if err != nil {
		log.Fatal(err)
	}
	g := make(map[string][]string)
	if err := json.Unmarshal(b, &g); err != nil {
		log.Fatal(err)
	}

	// convert graph to a type tarjan.Connections accepts
	graph := make(map[interface{}][]interface{})
	for pkg, deps := range g {
		for _, d := range deps {
			graph[pkg] = append(graph[pkg], d)
		}
	}

	// get a list of strongly connected components using tarjan's algorithm. The
	// result looks like [["dir1", "dir2"], ["dir3"]]. The idea is that we can't
	// have circular dependencies in the bazel libraries we create, so any
	// cycles in our dep graph must be placed in the same cc_library.
	var components [][]interface{}
	components = tarjan.Connections(graph)

	// create an occtLib for each strongly-connected component
	var libs []*occtLib
	for _, comp := range components {
		sort.Slice(comp, func(i, j int) bool {
			return comp[i].(string) < comp[j].(string)
		})
		// Name the library based on the first package/subdir included. Note
		// that this gives terrible names for the larger ones because the first
		// package isn't representative of the library as a whole. TODO: do
		// better
		lib := &occtLib{
			name: comp[0].(string),
			deps: stringset.New(),
		}
		// convert from []interface{} to []string
		for _, pkg := range comp {
			lib.pkgs = append(lib.pkgs, pkg.(string))
		}
		libs = append(libs, lib)
	}

	pkgToLib := make(map[string]*occtLib)
	for _, lib := range libs {
		for _, pkg := range lib.pkgs {
			pkgToLib[pkg] = lib
		}
	}

	// record dependencies to each lib's deps
	for _, lib := range libs {
		for _, pkg := range lib.pkgs {
			// inter-package dependencies
			for _, dep := range graph[pkg] {
				depLib := pkgToLib[dep.(string)]
				if depLib.name == lib.name {
					// filter out self-edges
					continue
				}
				lib.deps.Add(":" + depLib.name)
			}

			// manually-specified extra deps
			for _, e := range extraDeps[pkg] {
				lib.deps.Add(e)
			}
		}
	}

	// Write BUILD file to stdout
	fmt.Println(genAllBazel(libs))
	return subcommands.ExitSuccess
}

type occtLib struct {
	name string
	pkgs []string
	deps stringset.Set
}

func genAllBazel(libs []*occtLib) string {
	// sort alphabetically for consistent output
	sort.Slice(libs, func(i, j int) bool {
		return libs[i].name < libs[j].name
	})

	allBzl := "# AUTOGENERATED: this file was generated by github.com/justbuchanan/occt_bazel\n"
	allBzl += "package(default_visibility = [\"//visibility:public\"])\n\n"
	for _, l := range libs {
		allBzl += genBazelLib(l)
	}
	return allBzl
}

func genBazelLib(l *occtLib) string {
	bzl := "cc_library(\n"
	bzl += fmt.Sprintf("    name = \"%s\",\n", l.name)

	// srcs
	bzl += fmt.Sprintf("    srcs = glob([\n")
	for _, pkg := range l.pkgs {
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.c"))
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.cpp"))
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.cxx"))
	}
	bzl += "    ]),\n"

	// hdrs
	bzl += fmt.Sprintf("    hdrs = glob([\n")
	for _, pkg := range l.pkgs {
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.gxx"))
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.pxx"))
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.hxx"))
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.lxx"))
		bzl += fmt.Sprintf("        \"%s\",\n", filepath.Join(pkg, "*.h"))
	}
	bzl += "    ]),\n"

	// includes
	bzl += fmt.Sprintf("    includes = [\n")
	for _, pkg := range l.pkgs {
		bzl += fmt.Sprintf("        \"%s/\",\n", filepath.Join(pkg))
	}
	bzl += "    ],\n"

	// deps
	bzl += fmt.Sprintf("    deps = [\n")
	for _, dep := range l.deps.Elements() {
		bzl += fmt.Sprintf("        \"%s\",\n", dep)
	}
	bzl += "    ],\n)\n\n"

	return bzl
}
