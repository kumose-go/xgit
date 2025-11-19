package xgit_test

import (
	"context"
	"fmt"
	"strings"

	"github.com/kumose-go/xgit/add"
	"github.com/kumose-go/xgit/branch"
	"github.com/kumose-go/xgit/checkout"
	"github.com/kumose-go/xgit/clone"
	"github.com/kumose-go/xgit/commit"
	"github.com/kumose-go/xgit/config"
	"github.com/kumose-go/xgit/fetch"
	"github.com/kumose-go/xgit"
	ginit "github.com/kumose-go/xgit/init"
	"github.com/kumose-go/xgit/lsfiles"
	"github.com/kumose-go/xgit/merge"
	"github.com/kumose-go/xgit/notes"
	"github.com/kumose-go/xgit/pull"
	"github.com/kumose-go/xgit/push"
	"github.com/kumose-go/xgit/rebase"
	"github.com/kumose-go/xgit/remote"
	"github.com/kumose-go/xgit/reset"
	"github.com/kumose-go/xgit/revparse"
	"github.com/kumose-go/xgit/stash"
	"github.com/kumose-go/xgit/status"
	"github.com/kumose-go/xgit/tag"
	"github.com/kumose-go/xgit/types"
	"github.com/kumose-go/xgit/worktree"
)

func ExampleInit() {
	out, _ := xgit.Init(ginit.Bare, ginit.Quiet, ginit.Directory("foobar"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExampleInitWithContext() {
	out, _ := xgit.InitWithContext(context.Background(), ginit.Bare, ginit.Quiet, ginit.Directory("foobar"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git init --bare --quiet foobar
}

func ExamplePush() {
	out, _ := xgit.Push(push.All, push.FollowTags, push.ReceivePack("aaa"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePushWithContext() {
	out, _ := xgit.PushWithContext(context.Background(), push.All, push.FollowTags, push.ReceivePack("aaa"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git push --all --follow-tags --receive-pack=aaa
}

func ExamplePull() {
	out, _ := xgit.Pull(pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExamplePullWithContext() {
	out, _ := xgit.PullWithContext(context.Background(), pull.All, pull.Force, pull.Repository("upstream"), pull.Refspec("master"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git pull --all --force upstream master
}

func ExampleClone() {
	out, _ := xgit.Clone(clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleCloneWithContext() {
	out, _ := xgit.CloneWithContext(context.Background(), clone.Repository("git@github.com:ldez/go-git-cmd-wrapper.git"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git clone git@github.com:ldez/go-git-cmd-wrapper.git
}

func ExampleRemote() {
	out, _ := xgit.Remote(remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleRemoteWithContext() {
	out, _ := xgit.RemoteWithContext(context.Background(), remote.Add("upstream", "git@github.com:johndoe/go-git-cmd-wrapper.git"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git remote add upstream git@github.com:johndoe/go-git-cmd-wrapper.git
}

func ExampleFetch() {
	out, _ := xgit.Fetch(fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleFetchWithContext() {
	out, _ := xgit.FetchWithContext(context.Background(), fetch.NoTags, fetch.Remote("upstream"), fetch.RefSpec("myBranchName"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git fetch --no-tags upstream myBranchName
}

func ExampleRebase() {
	out, _ := xgit.Rebase(rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleRebaseWithContext() {
	out, _ := xgit.RebaseWithContext(context.Background(), rebase.PreserveMerges, rebase.Branch(fmt.Sprintf("%s/%s", "upstream", "master")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rebase --preserve-merges upstream/master
}

func ExampleCheckout() {
	out, _ := xgit.Checkout(checkout.NewBranch("myBranchName"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleCheckoutWithContext() {
	out, _ := xgit.CheckoutWithContext(context.Background(), checkout.NewBranch("myBranchName"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git checkout -b myBranchName
}

func ExampleConfig() {
	out, _ := xgit.Config(config.Entry("rebase.autoSquash", "true"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleConfigWithContext() {
	out, _ := xgit.ConfigWithContext(context.Background(), config.Entry("rebase.autoSquash", "true"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git config rebase.autoSquash true
}

func ExampleBranch() {
	out, _ := xgit.Branch(branch.DeleteForce, branch.BranchName("myBranch"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleBranchWithContext() {
	out, _ := xgit.BranchWithContext(context.Background(), branch.DeleteForce, branch.BranchName("myBranch"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git branch -D myBranch
}

func ExampleRevParse() {
	out, _ := xgit.RevParse(revparse.AbbrevRef(""), revparse.Args("HEAD"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleRevParseWithContext() {
	out, _ := xgit.RevParseWithContext(context.Background(), revparse.AbbrevRef(""), revparse.Args("HEAD"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git rev-parse --abbrev-ref HEAD
}

func ExampleReset() {
	out, _ := xgit.Reset(reset.Soft, reset.Commit("e41f083"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleResetWithContext() {
	out, _ := xgit.ResetWithContext(context.Background(), reset.Soft, reset.Commit("e41f083"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git reset --soft e41f083
}

func ExampleCommit() {
	out, _ := xgit.Commit(commit.Amend, commit.Message("foo"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git commit --amend --message=foo
}

func ExampleCommitWithContext() {
	out, _ := xgit.CommitWithContext(context.Background(), commit.Amend, commit.Message("foo"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git commit --amend --message=foo
}

func ExampleAdd() {
	out, _ := xgit.Add(add.All, xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git add --all
}

func ExampleAddWithContext() {
	out, _ := xgit.AddWithContext(context.Background(), add.All, xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git add --all
}

func ExampleMerge() {
	out, _ := xgit.Merge(merge.Squash, merge.Commits("myBranch"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleMergeWithContext() {
	out, _ := xgit.MergeWithContext(context.Background(), merge.Squash, merge.Commits("myBranch"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git merge --squash myBranch
}

func ExampleWorktree() {
	out, _ := xgit.Worktree(worktree.Add("v1.0", "origin/v1.0"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func ExampleWorktreeWithContext() {
	out, _ := xgit.WorktreeWithContext(context.Background(), worktree.Add("v1.0", "origin/v1.0"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git worktree add v1.0 origin/v1.0
}

func ExampleTag() {
	out, _ := xgit.Tag(tag.List, xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git tag --list
}

func ExampleTagWithContext() {
	out, _ := xgit.TagWithContext(context.Background(), tag.List, xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git tag --list
}

func ExampleStatus() {
	out, _ := xgit.Status(status.Short, status.Branch, xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git status --short --branch
}

func ExampleStatusWithContext() {
	out, _ := xgit.StatusWithContext(context.Background(), status.Short, status.Branch, xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git status --short --branch
}

func ExampleLsFiles() {
	out, _ := xgit.LsFiles(lsfiles.Z, lsfiles.ExcludeStandard, lsfiles.Others, lsfiles.Cached, xgit.CmdExecutor(cmdExecutorMock))

	// Notes: to parse the output you can use `fmt.Println(strings.Split(out, "\x00"))`

	fmt.Println(out)
	// Output: git ls-files -z --exclude-standard --others --cached
}

func ExampleLsFilesWithContext() {
	out, _ := xgit.LsFilesWithContext(context.Background(), lsfiles.Z, lsfiles.ExcludeStandard, lsfiles.Others, lsfiles.Cached, xgit.CmdExecutor(cmdExecutorMock))

	// Notes: to parse the output you can use `fmt.Println(strings.Split(out, "\x00"))`

	fmt.Println(out)
	// Output: git ls-files -z --exclude-standard --others --cached
}

func ExampleNotes_list() {
	out, _ := xgit.Notes(notes.List(""), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes list
}

func ExampleNotes_list_ref() {
	out, _ := xgit.Notes(notes.Ref("c9718bfd46a7261d1120ac2e50ef6b298bb2394a"), notes.List(""), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes --ref c9718bfd46a7261d1120ac2e50ef6b298bb2394a list
}

func ExampleNotes_add() {
	out, _ := xgit.Notes(notes.Add("", notes.Message("foo")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes add --message=foo
}

func ExampleNotes_copy() {
	out, _ := xgit.Notes(notes.Copy(notes.Object("cb17b52c17fb36a807f135245725dee88603cc08", "c9718bfd46a7261d1120ac2e50ef6b298bb2394a")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes copy cb17b52c17fb36a807f135245725dee88603cc08 c9718bfd46a7261d1120ac2e50ef6b298bb2394a
}

func ExampleNotes_append() {
	out, _ := xgit.Notes(notes.Append("", notes.Message("foo")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes append --message=foo
}

func ExampleNotes_edit() {
	out, _ := xgit.Notes(notes.Edit("", notes.Message("foo")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes edit --message=foo
}

func ExampleNotes_show() {
	out, _ := xgit.Notes(notes.Show(""), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes show
}

func ExampleNotes_merge() {
	out, _ := xgit.Notes(notes.Merge(notes.Commit, notes.NotesRef("cb17b52c17fb36a807f135245725dee88603cc08")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes merge --commit cb17b52c17fb36a807f135245725dee88603cc08
}

func ExampleNotes_remove() {
	out, _ := xgit.Notes(notes.Remove("", notes.Stdin), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes remove --stdin
}

func ExampleNotes_prune() {
	out, _ := xgit.Notes(notes.Prune(notes.Verbose), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes prune --verbose
}

func ExampleNotes_getRef() {
	out, _ := xgit.Notes(notes.GetRef(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes get-ref
}

func ExampleNotesWithContext_list() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.List(""), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes list
}

func ExampleNotesWithContext_list_ref() {
	out, _ := xgit.Notes(notes.Ref("c9718bfd46a7261d1120ac2e50ef6b298bb2394a"), notes.List(""), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes --ref c9718bfd46a7261d1120ac2e50ef6b298bb2394a list
}

func ExampleNotesWithContext_add() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Add("", notes.Message("foo")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes add --message=foo
}

func ExampleNotesWithContext_copy() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Copy(notes.Stdin), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes copy --stdin
}

func ExampleNotesWithContext_append() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Append("", notes.Message("foo")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes append --message=foo
}

func ExampleNotesWithContext_edit() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Edit("", notes.Message("foo")), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes edit --message=foo
}

func ExampleNotesWithContext_show() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Show(""), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes show
}

func ExampleNotesWithContext_merge() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Merge(notes.Commit), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes merge --commit
}

func ExampleNotesWithContext_remove() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Remove("", notes.Stdin), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes remove --stdin
}

func ExampleNotesWithContext_prune() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.Prune(notes.Verbose), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes prune --verbose
}

func ExampleNotesWithContext_getRef() {
	out, _ := xgit.NotesWithContext(context.Background(), notes.GetRef(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git notes get-ref
}

func ExampleStash_push() {
	out, _ := xgit.Stash(stash.Push("foo", stash.All), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash push --all foo
}

func ExampleStash_save() {
	out, _ := xgit.Stash(stash.Save("foo", stash.Patch), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash save --patch foo
}

func ExampleStash_list() {
	out, _ := xgit.Stash(stash.List(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash list
}

func ExampleStash_show() {
	out, _ := xgit.Stash(stash.Show("stash@{1}", stash.IncludeUntracked), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash show --include-untracked stash@{1}
}

func ExampleStash_pop() {
	out, _ := xgit.Stash(stash.Pop("stash@{1}", stash.Quiet), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash pop --quiet stash@{1}
}

func ExampleStash_apply() {
	out, _ := xgit.Stash(stash.Apply("stash@{1}", stash.Quiet), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash apply --quiet stash@{1}
}

func ExampleStash_branch() {
	out, _ := xgit.Stash(stash.Branch("foo", "stash@{1}"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash branch foo stash@{1}
}

func ExampleStash_clear() {
	out, _ := xgit.Stash(stash.Clear(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash clear
}

func ExampleStash_drop() {
	out, _ := xgit.Stash(stash.Drop("stash@{1}", stash.Quiet), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash drop --quiet stash@{1}
}

func ExampleStash_create() {
	out, _ := xgit.Stash(stash.Create(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash create
}

func ExampleStash_store() {
	out, _ := xgit.Stash(stash.Store(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash store
}

func ExampleStashWithContext_push() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Push("foo", stash.All), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash push --all foo
}

func ExampleStashWithContext_save() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Save("foo", stash.Patch), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash save --patch foo
}

func ExampleStashWithContext_list() {
	out, _ := xgit.StashWithContext(context.Background(), stash.List(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash list
}

func ExampleStashWithContext_show() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Show("stash@{1}", stash.IncludeUntracked), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash show --include-untracked stash@{1}
}

func ExampleStashWithContext_pop() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Pop("stash@{1}", stash.Quiet), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash pop --quiet stash@{1}
}

func ExampleStashWithContext_apply() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Apply("stash@{1}", stash.Quiet), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash apply --quiet stash@{1}
}

func ExampleStashWithContext_branch() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Branch("foo", "stash@{1}"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash branch foo stash@{1}
}

func ExampleStashWithContext_clear() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Clear(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash clear
}

func ExampleStashWithContext_drop() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Drop("stash@{1}", stash.Quiet), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash drop --quiet stash@{1}
}

func ExampleStashWithContext_create() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Create(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash create
}

func ExampleStashWithContext_store() {
	out, _ := xgit.StashWithContext(context.Background(), stash.Store(), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Println(out)
	// Output: git stash store
}

func ExampleRaw() {
	out, _ := xgit.Raw("stash", xgit.CmdExecutor(cmdExecutorMock), func(g *types.Cmd) {
		g.AddOptions("list")
		g.AddOptions("--pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'")
	})

	fmt.Println(out)
	// Output: git stash list --pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'
}

func ExampleRawWithContext() {
	out, _ := xgit.RawWithContext(context.Background(), "stash", xgit.CmdExecutor(cmdExecutorMock), func(g *types.Cmd) {
		g.AddOptions("list")
		g.AddOptions("--pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'")
	})

	fmt.Println(out)
	// Output: git stash list --pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'
}

func ExampleRawWithContext_baseOptions() {
	out, _ := xgit.RawWithContext(
		context.Background(),
		"stash",
		func(g *types.Cmd) {
			g.AddOptions("list")
			g.AddOptions("--pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'")
		},
		func(g *types.Cmd) {
			g.AddBaseOptions("-C")
			g.AddBaseOptions("<your_path>")
		},
		xgit.CmdExecutor(cmdExecutorMock),
	)

	fmt.Println(out)
	// Output: git -C <your_path> stash list --pretty=format:'%Cblue%gd%Creset%Cred:%Creset %C(yellow)%s%Creset'
}

func ExampleCond() {
	param := false
	out, _ := xgit.Push(push.All, xgit.Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Print(out)

	param = true
	out, _ = xgit.Push(push.All, xgit.Cond(param, push.DryRun), push.FollowTags, push.ReceivePack("aaa"), xgit.CmdExecutor(cmdExecutorMock))

	fmt.Print(out)

	// Output:
	// git push --all --follow-tags --receive-pack=aaa
	// git push --all --dry-run --follow-tags --receive-pack=aaa
}

func cmdExecutorMock(_ context.Context, name string, _ bool, args ...string) (string, error) {
	return fmt.Sprintln(name, strings.Join(args, " ")), nil
}
