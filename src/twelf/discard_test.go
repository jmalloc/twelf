package twelf_test

import (
	. "github.com/jmalloc/twelf/src/twelf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiscardLogger", func() {
	Describe("IsDebug", func() {
		It("returns false", func() {
			Expect(DiscardLogger{}.IsDebug()).To(BeFalse())
		})
	})
})
