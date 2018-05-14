package twelf_test

import (
	. "github.com/jmalloc/twelf/src/twelf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BufferedLogger", func() {
	var (
		logger *BufferedLogger
	)

	BeforeEach(func() {
		logger = &BufferedLogger{}
	})

	Context("when debug logging is enabled", func() {
		BeforeEach(func() {
			logger.CaptureDebug = true
		})

		Describe("Log", func() {
			It("stores a formatted message", func() {
				logger.Log("message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("LogString", func() {
			It("stores a message", func() {
				logger.LogString("<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})

			It("does not substitute placeholders", func() {
				logger.LogString("<%s>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<%s>",
					IsDebug: false,
				}))
			})
		})

		Describe("Debug", func() {
			It("stores a formatted message", func() {
				logger.Debug("message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: true,
				}))
			})
		})

		Describe("DebugString", func() {
			It("stores a message", func() {
				logger.DebugString("<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: true,
				}))
			})

			It("does not substitute placeholders", func() {
				logger.DebugString("<%s>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<%s>",
					IsDebug: true,
				}))
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
			It("stores a formatted message", func() {
				logger.Log("message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("LogString", func() {
			It("stores a message", func() {
				logger.LogString("<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})

			It("does not substitute placeholders", func() {
				logger.LogString("<%s>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<%s>",
					IsDebug: false,
				}))
			})
		})

		Describe("Debug", func() {
			It("does not produce any output", func() {
				logger.Debug("message <%s>", "arg")

				Expect(logger.Messages()).To(BeEmpty())
			})
		})

		Describe("DebugString", func() {
			It("does not produce any output", func() {
				logger.DebugString("<message>")

				Expect(logger.Messages()).To(BeEmpty())
			})
		})

		Describe("IsDebug", func() {
			It("returns false", func() {
				Expect(logger.IsDebug()).To(BeFalse())
			})
		})
	})

	Describe("Reset", func() {
		It("clears messages from the logger", func() {
			logger.LogString("<message>")
			logger.Reset()

			Expect(logger.Messages()).To(BeEmpty())
		})
	})

	Describe("TakeMessages", func() {
		It("returns and clears messages from the logger", func() {
			logger.LogString("<message>")

			m := logger.TakeMessages()
			Expect(m).To(HaveLen(1))
			Expect(logger.Messages()).To(BeEmpty())
		})
	})

	Describe("Flush", func() {
		dest := &BufferedLogger{
			CaptureDebug: true,
		}

		BeforeEach(func() {
			logger.CaptureDebug = true
			dest.Reset()
		})

		It("logs captures messages to the destination logger", func() {
			logger.LogString("<message>")
			logger.FlushTo(dest)

			Expect(logger.Messages()).To(BeEmpty())
			Expect(dest.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: false,
			}))
		})

		It("logs debug messages as debug messages", func() {
			logger.DebugString("<message>")
			logger.FlushTo(dest)

			Expect(logger.Messages()).To(BeEmpty())
			Expect(dest.Messages()).To(ConsistOf(BufferedLogMessage{
				Message: "<message>",
				IsDebug: true,
			}))
		})
	})
})
