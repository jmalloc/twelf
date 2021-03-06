package twelf

var (
	// DefaultLogger is a logger that only logs non-debug messages.
	DefaultLogger Logger = &StandardLogger{CaptureDebug: false}

	// DebugLogger is a logger that logs both debug and non-debug messages.
	DebugLogger Logger = &StandardLogger{CaptureDebug: true}

	// SilentLogger is a logger that does not log any messages.
	SilentLogger Logger = DiscardLogger{}
)

// Logger is an interface for writing log messages.
type Logger interface {
	// Log writes an application log message formatted according to a format
	// specifier.
	//
	// It should be ussed for messages that are intended for people responsible for
	// operating the application, such as the end-user or operations staff.
	//
	// fmt is the format specifier, as per fmt.Printf(), etc.
	Log(fmt string, v ...interface{})

	// LogString writes a pre-formatted application log message.
	//
	// It should be ussed for messages that are intended for people responsible for
	// operating the application, such as the end-user or operations staff.
	LogString(s string)

	// Debug writes a debug log message formatted according to a format
	// specifier.
	//
	// If IsDebug() returns false, no logging is performed.
	//
	// It should be used for messages that are intended for the software developers
	// that maintain the application.
	//
	// fmt is the format specifier, as per fmt.Printf(), etc.
	Debug(fmt string, v ...interface{})

	// DebugString writes a pre-formatted debug log message.
	//
	// If IsDebug() returns false, no logging is performed.
	//
	// It should be used for messages that are intended for the software developers
	// that maintain the application.
	DebugString(s string)

	// IsDebug returns true if this logger will perform debug logging.
	//
	// Generally the application should just call Debug() or DebugString() without
	// calling IsDebug(), however it can be used to check if debug logging is
	// necessary before executing expensive code that is only used to obtain debug
	// information.
	IsDebug() bool
}
