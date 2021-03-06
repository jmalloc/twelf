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

		Describe("Log", func() {
			It("writes a formatted message to the target", func() {
				logger.Log("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("LogString", func() {
			It("writes a message to the target", func() {
				logger.LogString("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})

			It("does not substitute placeholders", func() {
				logger.LogString("<%s>")

				Expect(buffer.String()).To(Equal("<%s>\n"))
			})
		})

		Describe("Debug", func() {
			It("writes a formatted message to the target", func() {
				logger.Debug("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("DebugString", func() {
			It("writes a message to the target", func() {
				logger.DebugString("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})

			It("does not substitute placeholders", func() {
				logger.DebugString("<%s>")

				Expect(buffer.String()).To(Equal("<%s>\n"))
			})
		})

		Describe("IsDebug", func() {
			It("returns true", func() {
				Expect(logger.IsDebug()).To(BeTrue())
			})
		})
	})

	Context("when debug logging is disabled", func() {
		Describe("Log", func() {
			It("writes a formatted message to the target", func() {
				logger.Log("message <%s>", "arg")

				Expect(buffer.String()).To(Equal("message <arg>\n"))
			})
		})

		Describe("LogString", func() {
			It("writes a message to the target", func() {
				logger.LogString("<message>")

				Expect(buffer.String()).To(Equal("<message>\n"))
			})
		})

		Describe("Debug", func() {
			It("does not produce any output", func() {
				logger.Debug("message <%s>", "arg")

				Expect(buffer.String()).To(Equal(""))
			})
		})

		Describe("DebugString", func() {
			It("does not produce any output", func() {
				logger.DebugString("<message>")

				Expect(buffer.String()).To(Equal(""))
			})
		})

		Describe("IsDebug", func() {
			It("returns false", func() {
				Expect(logger.IsDebug()).To(BeFalse())
			})
		})
	})
})
