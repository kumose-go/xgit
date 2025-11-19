# Go Git Cmd Wrapper


It's a simple wrapper around `git` command.

Import `github.com/kumose-go/xgit`.

```go
// clone
output, err := xgit.Clone(clone.Repository("https://github.com/ldez/prm"))
// with debug option
output, err := xgit.Clone(clone.Repository("https://github.com/ldez/prm"), xgit.Debug)
output, err := xgit.Clone(clone.Repository("https://github.com/ldez/prm"), xgit.Debugger(true))

// fetch
output, err = xgit.Fetch(fetch.NoTags, fetch.Remote("upstream"))
output, err = xgit.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("master"))

// add a remote
output, err = xgit.Remote(remote.Add, remote.Name("upstream"), remote.URL("https://github.com/ldez/prm"))

// --- global options ---
output, err := xgit.Clone(global.UpperC("/tmp"), clone.Repository("https://github.com/ldez/prm"))
```

More examples: [Documentation](https://pkg.go.dev/github.com/kumose-go/xgit)
