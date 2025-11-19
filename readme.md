# Go Git Cmd Wrapper


It's a simple wrapper around `git` command.

Import `github.com/kumose-go/xgit/git`.

```go
// clone
output, err := git.Clone(clone.Repository("https://github.com/ldez/prm"))
// with debug option
output, err := git.Clone(clone.Repository("https://github.com/ldez/prm"), git.Debug)
output, err := git.Clone(clone.Repository("https://github.com/ldez/prm"), git.Debugger(true))

// fetch
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"))
output, err = git.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

// add a remote
output, err = git.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/ldez/prm"))

// --- global options ---
output, err := git.Clone(global.UpperC("/tmp"), clone.Repository("https://github.com/ldez/prm"))
```

More examples: [Documentation](https://pkg.go.dev/github.com/kumose-go/xgit/git)
