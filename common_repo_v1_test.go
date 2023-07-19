package common_repo_v1_test

import (
	"testing"

	. "github.com/common-repo/common-repo-v1"

	. "github.com/onsi/gomega"
	"github.com/shakefu/goblin"
	// . "github.com/common-repo/common-repo-v1/internal/testutil"
)

func TestCommonRepo(t *testing.T) {
	// Initialize the Goblin test suite
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) }) // Gomega hook

	g.Describe("commonrepo", func() {
		g.Describe("Upstreams", func() {
			g.It("returns the upstreams", func() {
				cr, _ := NewFrom("testdata/fixtures/local/single.yml", ".")
				ups, err := cr.Upstreams()
				Expect(err).ToNot(HaveOccurred())
				Expect(len(ups)).To(Equal(2))
				Expect(ups[1].String()).To(HavePrefix("./testdata/fixtures/local/single.yml@"))
				Expect(ups[0].String()).To(HavePrefix("./@"))
			})
		})
	})
}
