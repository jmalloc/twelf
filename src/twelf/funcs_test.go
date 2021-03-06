package twelf_test

import (
	. "github.com/jmalloc/twelf/src/twelf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logging Functions", func() {
	var (
		logger          BufferedLogger
		originalDefault Logger
	)

	BeforeEach(func() {
		logger.CaptureDebug = true
		logger.Reset()
	})

	Context("when passed a nil logger", func() {
		BeforeSuite(func() {
			originalDefault = DefaultLogger
			DefaultLogger = &logger
		})

		AfterSuite(func() {
			DefaultLogger = originalDefault
		})

		Describe("Log", func() {
			It("forwards to the default logger", func() {
				Log(nil, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("LogString", func() {
			It("forwards to the default logger", func() {
				LogString(nil, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})
		})

		Describe("Debug", func() {
			It("forwards to the default logger", func() {
				Debug(nil, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: true,
				}))
			})
		})

		Describe("DebugString", func() {
			It("forwards to the default logger", func() {
				DebugString(nil, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: true,
				}))
			})
		})

		Describe("IsDebug", func() {
			It("forwards to the default logger", func() {
				Expect(IsDebug(nil)).To(BeTrue())
				logger.CaptureDebug = false
				Expect(IsDebug(nil)).To(BeFalse())
			})
		})
	})

	Context("when passed a non-nil logger", func() {
		Describe("Log", func() {
			It("forwards to the given logger", func() {
				Log(&logger, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: false,
				}))
			})
		})

		Describe("LogString", func() {
			It("forwards to the given logger", func() {
				LogString(&logger, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: false,
				}))
			})
		})

		Describe("Debug", func() {
			It("forwards to the given logger", func() {
				Debug(&logger, "message <%s>", "arg")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "message <arg>",
					IsDebug: true,
				}))
			})
		})

		Describe("DebugString", func() {
			It("forwards to the given logger", func() {
				DebugString(&logger, "<message>")

				Expect(logger.Messages()).To(ConsistOf(BufferedLogMessage{
					Message: "<message>",
					IsDebug: true,
				}))
			})
		})

		Describe("IsDebug", func() {
			It("forwards to the given logger", func() {
				Expect(IsDebug(&logger)).To(BeTrue())
				logger.CaptureDebug = false
				Expect(IsDebug(&logger)).To(BeFalse())
			})
		})
	})
})
