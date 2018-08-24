package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	git "gopkg.in/src-d/go-git.v4"
)

func run(c *cli.Context) error {
	var (
		err               error
		goPath            string
		goSourcePath      string
		gitRepositoryPath = c.String("repo-path")
		gitSubmodules     git.Submodules
		sharedModules     []SharedModule
	)

	if len(gitRepositoryPath) == 0 {
		return fmt.Errorf("repo-path must be supplied")
	}

	goPath = os.Getenv("GOPATH")

	if len(goPath) == 0 {
		return fmt.Errorf("unable to find GOPATH")
	}

	goSourcePath = fmt.Sprintf("%s/", filepath.Join(goPath, "src"))

	if gitSubmodules, err = getSubmodulesFromRepositoryPath(gitRepositoryPath); err != nil {
		return err
	}

	if sharedModules, err = convertGitModulesToSharedModule(gitSubmodules); err != nil {
		return err
	}

	projectPath := strings.Replace(gitRepositoryPath, goSourcePath, "", 1)

	return writeGoModFile(os.Stdout, projectPath, sharedModules)
}
