package reset

import "github.com/kumose-go/xgit/types"

// HyphenHyphen add `--`
func HyphenHyphen(g *types.Cmd) {
	g.AddOptions("--")
}

// Path <paths>...
func Path(values ...string) types.Option {
	return func(g *types.Cmd) {
		for _, value := range values {
			g.AddOptions(value)
		}
	}
}

// Commit [<commit>]
func Commit(hash string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(hash)
	}
}

// TreeIsh [<tree-ish>]
func TreeIsh(hash string) types.Option {
	return func(g *types.Cmd) {
		g.AddOptions(hash)
	}
}
