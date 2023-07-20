// Package testutil provides convenience utilities for tests within this repository
package testutil

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/common-repo/common-repo-v1/pkg/gitutil"
	"github.com/common-repo/common-repo-v1/pkg/repos"
	. "github.com/onsi/gomega"
)

// InlineYaml returns a bytestring from the given heredoc string
func InlineYaml(doc string) []byte {
	return []byte(heredoc.Doc(doc))
}

// LocalRepo returns the local repo or raises an error
func LocalRepo() (repo *repos.Repo) {
	// This is a reimplementation of GetLocalRepo for testing purposes
	path, err := gitutil.FindLocalRepoPath()
	Expect(err).ShouldNot(HaveOccurred())

	repo, err = repos.New(path)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(repo).ShouldNot(BeNil())
	return
}
