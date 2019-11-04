package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"bitbucket.org/creachadair/stringset"
	"github.com/google/subcommands"
)

type genDepgraphSubcommand struct {
	srcdir string

	cacheMu sync.RWMutex
	cache   map[string]string
}

func (*genDepgraphSubcommand) Name() string { return "depgraph" }
func (*genDepgraphSubcommand) Synopsis() string {
	return "Scan the OCCT --srcdir and determine dependencies based on #includes. Write the resulting dependency graph to stdout in json format."
}
func (*genDepgraphSubcommand) Usage() string {
	return "depgraph --srcdir /path/to/occt/src"
}
func (s *genDepgraphSubcommand) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.srcdir, "srcdir", "", "Path to 'src' subdirectory of OCCT project")
}
func (s *genDepgraphSubcommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if s.srcdir == "" {
		log.Printf("--srcdir is required")
		s.Usage()
		return subcommands.ExitFailure
	}
	log.Printf("Using srcdir: %q\n", s.srcdir)

	graph, err := s.buildDepGraph()
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(graph, "", "  ")
	if err != nil {
		log.Fatalf("marshaling graph: %v", err)
	}
	if _, err := os.Stdout.Write(b); err != nil {
		log.Fatalf("Error writing depgraph to stdout")
	}

	return subcommands.ExitSuccess
}

// buildDepGraph scans the 'src' dir of occt and builds a dependency graph of the form:
// {
//   "dirA": ["dirB", "dirC"],
//   "dirD": ["dirC"]
// }
func (cmd *genDepgraphSubcommand) buildDepGraph() (map[string][]string, error) {
	// get all dirs under 'src'
	var pkgDirs []string
	err := filepath.Walk(cmd.srcdir, func(path string, info os.FileInfo, err error) error {
		if path == cmd.srcdir || path == "" {
			return nil
		}
		if info.IsDir() {
			pkgDirs = append(pkgDirs, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// for each pkg, record a set of which pkgs it #includes things from
	graph := make(map[string][]string)
	var graphMu sync.Mutex
	var wg sync.WaitGroup
	for _, d := range pkgDirs {
		if d == "" {
			continue
		}
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			deps, err := cmd.findPackageDeps(d)
			if err != nil {
				log.Printf("ignoring dependency scan error: %v", err)
			} else {
				rel, err := filepath.Rel(cmd.srcdir, d)
				if err != nil {
					log.Printf("can't relativize: %v\n", err)
					return
				}
				// if rel == "" {
				// 	log.Printf("ignoring empty reldir")
				// }
				depsSlice := deps.Elements()
				graphMu.Lock()
				graph[rel] = depsSlice
				graphMu.Unlock()
			}
		}(d)
	}
	wg.Wait()

	return graph, nil
}

func (cmd *genDepgraphSubcommand) cachedPackageFromInclude(filename string) (string, error) {
	// try cache first
	cmd.cacheMu.RLock()
	pkg := cmd.cache[filename]
	cmd.cacheMu.RUnlock()
	if pkg != "" {
		return pkg, nil
	}

	// call actual implementation
	var err error
	pkg, err = cmd.packageFromInclude(filename)
	if err != nil {
		return "", err
	}

	// save to cache for future use
	cmd.cacheMu.Lock()
	cmd.cache[filename] = pkg
	cmd.cacheMu.Unlock()

	return pkg, nil
}

// packageFromInclude scans the files under srcdir to determine which package
// the given file comes from.
func (cmd *genDepgraphSubcommand) packageFromInclude(filename string) (string, error) {
	glob := filepath.Join(cmd.srcdir, "*", filename)
	m, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}

	if len(m) != 1 {
		if len(m) == 0 {
			// If there are no matches, this file probably doesn't come from
			// occt. Could be a stdlib file or something else.
			return "", nil
		} else {
			return "", fmt.Errorf("Ambiguous match: glob=%v, m=%v", glob, m)
		}
	}

	rel, err := filepath.Rel(cmd.srcdir, m[0])
	if err != nil {
		return "", err
	}
	d, _ := filepath.Split(rel)

	return strings.TrimSuffix(d, "/"), nil
}

var includePattern = regexp.MustCompile("#include (\"|<)(?P<path>.+)(\"|>)")

// findPackageDeps determines the diretories this dir depends on by scanning for
// #includes in the contained files. Example findPackageDeps("src/dirA") ->
// ["src/dirB", "src/dirC"]
func (cmd *genDepgraphSubcommand) findPackageDeps(dir string) (stringset.Set, error) {
	var deps stringset.Set
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fi := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, fi.Name()))
		if err != nil {
			return nil, err
		}
		for _, match := range includePattern.FindAllStringSubmatch(string(b), -1) {
			include := match[2]
			dep, err := cmd.cachedPackageFromInclude(include)
			if err != nil {
				return nil, err
			}
			if dep == "" {
				continue
			}
			deps.Add(dep)
		}
	}

	return deps, nil
}
