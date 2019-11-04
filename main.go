// Binary occt_bazel generates a Bazel BUILD file for compiling
// OpenCascade/OCCT. The OCCT project contains a bunch of directories under the
// 'src' toplevel directory such as 'Draw', 'Geom2d', and 'BRep' that contain
// source files. These directories (referred to as packages/pkgs from here
// onward) reference each other via #includes. This binary scans all source
// files to determine which packages #include files from each other, forming a
// dependency graph. Any packages that form a cycle in the dependency graph must
// be placed in the same library (otherwise there would be cycles in the bazel
// build graph). Tarjan's algorithm is used to find these cycles or
// "strongly-connected components" and group them for placement into the same
// cc_library.
//
// This binary is split into two subcommands: "depgraph" and "genbazel".
//
// * "depgraph" scans the OCCT sources and builds a dependency graph, which is
//   emitted in json form. This may take several minutes, but should only need
//   to be done once.
// * "genbazel" reads the json file from "depgraph" and uses it to generate a
//   Bazel BUILD file. You may need to adjust bazel.go and run it again.
//
// Note: the generated BUILD file has external dependencies as specified by the
// `extraDeps` variable in bazel.go. You'll need to add these to your WORKSPACE
// file - see the one in this repo as an example.
//
// Example usage:
//   $ occt_bazel depgraph --srcdir /path/to/occt/src > depgraph.json
//   $ occt_bazel genbazel --depgraph depgraph.json > /path/to/occt/src/BUILD
package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&genDepgraphSubcommand{cache: make(map[string]string)}, "")
	subcommands.Register(&genBazelSubcommand{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
