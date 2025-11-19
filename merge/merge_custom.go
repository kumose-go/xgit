package merge

import "github.com/kumose-go/xgit/types"

// Commits <commit>...
// Commits, usually other branch heads, to merge into our branch.
// Specifying more than one commit will create a merge with more than two parents (affectionately called an Octopus merge).
func Commits(values ...string) types.Option {
	return func(g *types.Cmd) {
		for _, value := range values {
			g.AddOptions(value)
		}
	}
}
