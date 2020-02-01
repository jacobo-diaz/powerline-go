package main

import (
	"fmt"
	pwl "github.com/justjanne/powerline-go/powerline"
	"strconv"
	"strings"
)

func addRepoStats(nChanges int, symbol string) string{
	if nChanges > 0 {
		return fmt.Sprintf(" %d%s", nChanges, symbol )
	} else {
		return ""
	}
}

func segmentGitCompact(p *powerline) {
	if len(p.ignoreRepos) > 0 {
		out, err := runGitCommand("git", "rev-parse", "--show-toplevel")
		if err != nil {
			return
		}
		out = strings.TrimSpace(out)
		if p.ignoreRepos[out] {
			return
		}
	}

	out, err := runGitCommand("git", "status", "--porcelain", "-b", "--ignore-submodules")
	if err != nil {
		return
	}

	status := strings.Split(out, "\n")
	stats := parseGitStats(status)
	branchInfo := parseGitBranchInfo(status)
	var branch string

	if branchInfo["local"] != "" {
		ahead, _ := strconv.ParseInt(branchInfo["ahead"], 10, 32)
		stats.ahead = int(ahead)

		behind, _ := strconv.ParseInt(branchInfo["behind"], 10, 32)
		stats.behind = int(behind)

		branch = branchInfo["local"]
	} else {
		branch = getGitDetachedBranch(p)
	}

	var foreground, background uint8
	if stats.dirty() {
		foreground = p.theme.RepoDirtyFg
		background = p.theme.RepoDirtyBg
	} else {
		foreground = p.theme.RepoCleanFg
		background = p.theme.RepoCleanBg
	}

	out, err = runGitCommand("git", "rev-list", "-g", "refs/stash")
	if err == nil && len(out) > 0 {
		stats.stashed = len(strings.Split(out, "\n")) - 1
	}

	p.appendSegment("git-branch", pwl.Segment{
		Content: fmt.Sprintf("%s %s", p.symbolTemplates.RepoBranch, branch ) +
		addRepoStats(stats.ahead, p.symbolTemplates.RepoAhead ) +
		addRepoStats(stats.behind, p.symbolTemplates.RepoBehind ) +
		addRepoStats(stats.staged, p.symbolTemplates.RepoStaged ) +
		addRepoStats(stats.notStaged, p.symbolTemplates.RepoNotStaged ) +
		addRepoStats(stats.untracked, p.symbolTemplates.RepoUntracked ) +
		addRepoStats(stats.conflicted, p.symbolTemplates.RepoConflicted ) +
		addRepoStats(stats.stashed, p.symbolTemplates.RepoStashed )	,
		Foreground: foreground,
		Background: background,
	})
}