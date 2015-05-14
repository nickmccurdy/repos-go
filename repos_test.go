package repos_test

import (
	// . "github.com/nicolasmccurdy/repos-go"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Repos", func() {
	Describe(".IsClean", func() {
		Context("with an untracked file", func() {
			PIt("returns false")
		})

		Context("with a modified file in the working tree", func() {
			PIt("returns false")
		})

		Context("with a modified file in the stage", func() {
			PIt("returns false")
		})

		Context("with neither uncommitted nor untracked changes", func() {
			PIt("returns true")
		})

		Context("with no commits", func() {
			Context("and no untracked files", func() {
				PIt("returns true")
			})

			Context("and an untracked file", func() {
				PIt("returns false")
			})

			Context("and an empty directory", func() {
				PIt("returns true")
			})
		})
	})

	Describe(".List", func() {
		Context("by default", func() {
			PIt("lists git repositories in the given directory")
		})

		Context("with recursion", func() {
			PIt("lists git repositories in the given directory and subdirectories")
		})

		Context("with the \"clean\" filter", func() {
			PIt("lists clean repositories")
		})

		Context("with the \"dirty\" filter", func() {
			PIt("lists dirty repositories")
		})
	})
})
