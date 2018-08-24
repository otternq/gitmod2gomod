package main

import (
	git "gopkg.in/src-d/go-git.v4"
	gitplumming "gopkg.in/src-d/go-git.v4/plumbing"
)

func getSubmodulePathAndHash(gitModule *git.Submodule) (string, gitplumming.Hash, error) {
	var (
		err error

		gitPath            string
		gitHash            gitplumming.Hash
		gitSubmoduleStatus *git.SubmoduleStatus
	)

	if gitSubmoduleStatus, err = gitModule.Status(); err != nil {
		return "", gitplumming.Hash{}, err
	}

	gitPath = gitSubmoduleStatus.Path
	gitHash = gitSubmoduleStatus.Current

	return gitPath, gitHash, nil
}

func getSubmodulesFromRepositoryPath(gitRepositoryPath string) (git.Submodules, error) {
	var (
		err error

		gitRepository *git.Repository
		gitWorktree   *git.Worktree
	)

	if gitRepository, err = git.PlainOpen(gitRepositoryPath); err != nil {
		return nil, err
	}

	if gitWorktree, err = gitRepository.Worktree(); err != nil {
		return nil, err
	}

	return gitWorktree.Submodules()
}
