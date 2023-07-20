package files_test

import (
	"testing"

	. "github.com/common-repo/common-repo-v1/pkg/files"

	. "github.com/common-repo/common-repo-v1/internal/testutil"
	. "github.com/onsi/gomega"
	"github.com/shakefu/goblin"
)

func TestFiles(t *testing.T) {
	// Initialize the Goblin test suite
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) }) // Gomega hook

	g.Describe("List", func() {
		g.It("works", func() {
			repo := LocalRepo()
			files, err := List(repo.FS())
			Expect(err).ToNot(HaveOccurred())
			Expect(files).To(ContainElement(".commonrepo.yaml"))
		})
	})
}
