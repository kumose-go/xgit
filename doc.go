/*
Package xgit A simple wrapper around `git` command.

	import (
		// ...
		"github.com/kumose-go/xgit"
		// ...
		"github.com/kumose-go/xgit/clone"
		"github.com/kumose-go/xgit/config"
		"github.com/kumose-go/xgit/fetch"
		"github.com/kumose-go/xgit/remote"
	)

	// clone
	output, err := xgit.Clone(clone.Repository("https://github.com/ldez/gcg"))
	// with debug option
	output, err := xgit.Clone(clone.Repository("https://github.com/ldez/gcg"), xgit.Debug)
	output, err := xgit.Clone(clone.Repository("https://github.com/ldez/gcg"), xgit.Debugger(true))

	// fetch
	output, err = xgit.Fetch(fetch.NoTags, fetch.Remote("upstream"))
	output, err = xgit.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

	// add a remote
	output, err = xgit.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/ldez/gcg"))
*/
package xgit
