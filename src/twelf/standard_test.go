package twelf_test

import (
	"bytes"
	"log"

	. "github.com/jmalloc/twelf/src/twelf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StandardLogger", func() {
	var (
		buffer bytes.Buffer
		target *log.Logger = log.New(&buffer, "", 0)
		logger *StandardLogger
	)

	BeforeEach(func() {
		buffer.Reset()

		logger = &StandardLogger{
			Target: target,
		}
	})

	Context("when debug logging is enabled", func() {
		BeforeEach(func() {
			logger.CaptureDebug = true
		})

		Describe("Debug", func() {
			It("writes a formatted message to the target", func() {
				logger.Debug("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("DebugString", func() {
			It("writes a message to the target", func() {
				logger.Debug("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})
		})
	})
})
