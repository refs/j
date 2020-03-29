package journal

// Options are configurable parameters for the Journal.
type Options struct {
	Home string
}

// Option just implements the functional options pattern.
type Option func(*Options)

// Home set J's home directory.
func Home(h string) Option {
	return func(o *Options) {
		o.Home = h
	}
}
