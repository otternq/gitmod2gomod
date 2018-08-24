package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	gitplumming "gopkg.in/src-d/go-git.v4/plumbing"
)

var goModTemplate = template.Must(template.New("package").Parse(`module {{ .PackageName }}

require ({{ range $i, $SharedModule := .SharedModules }}
  {{ $SharedModule.URL }}  {{ $SharedModule.Hash }}{{ end }}
)
`))

type SharedModule struct {
	URL  string
	Hash string
}

func convertGitModulesToSharedModule(gitSubmodules git.Submodules) ([]SharedModule, error) {
	var (
		sharedModules = []SharedModule{}
	)

	for _, gitModule := range gitSubmodules {
		var (
			err     error
			gitPath string
			gitHash gitplumming.Hash
			goURL   string
		)

		if gitPath, gitHash, err = getSubmodulePathAndHash(gitModule); err != nil {
			log.Println("unable to get path and status:", err.Error())
			continue
		}

		if goURL, err = gitPathToGoURL(gitPath); err != nil {
			log.Println("unable to convert git path to go url", err.Error())
			continue
		}

		sharedModules = append(sharedModules, SharedModule{
			URL:  goURL,
			Hash: fmt.Sprintf("%s", gitHash),
		})
	}

	return sharedModules, nil
}

func gitPathToGoURL(gitPath string) (string, error) {
	var (
		directoryComponents []string
	)

	directoryComponents = strings.Split(gitPath, "/")

	if len(directoryComponents) < 1 {
		return "", fmt.Errorf("path can't be in vendor directory")
	}

	if directoryComponents[0] != "vendor" {
		return "", fmt.Errorf("path isn't in the vendor directory")
	}

	return strings.Join(directoryComponents[1:], "/"), nil
}

func writeGoModFile(writer io.Writer, packageName string, sharedModules []SharedModule) error {
	var goModInfo = struct {
		PackageName   string
		SharedModules []SharedModule
	}{
		PackageName:   packageName,
		SharedModules: sharedModules,
	}

	return goModTemplate.Execute(writer, goModInfo)
}
