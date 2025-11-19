package status

import "github.com/kumose-go/xgit/types"

// PathSpec <pathspec>...
// See the pathspec entry in gitglossary(7).
func PathSpec(paths ...string) types.Option {
	return func(g *types.Cmd) {
		for _, path := range paths {
			g.AddOptions(path)
		}
	}
}
