package journal

// Options are configurable parameters for the Journal.
type Options struct {
	// Core
	Home string

	// Logs
	Color bool
	Level string
}

// Option just implements the functional options pattern.
type Option func(*Options)

// Home set J's home directory.
func Home(h string) Option {
	return func(o *Options) {
		o.Home = h
	}
}

// Color toggles J's colored logs.
func Color(c bool) Option {
	return func(o *Options) {
		o.Color = c
	}
}

// Level set J's log level.
func Level(l string) Option {
	return func(o *Options) {
		o.Level = l
	}
}
