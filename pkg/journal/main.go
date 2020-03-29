package journal

// Journaling captures journaling behavior.
type Journaling interface {
	// Open opens the journal.
	Open() error
}
