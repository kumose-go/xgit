package fetch

import (
	"github.com/kumose-go/xgit/types"
)

// Remote name.
func Remote(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// Group name.
func Group(name string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(name)
	}
}

// RefSpec name.
func RefSpec(ref string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(ref)
	}
}
