package journal

// Journaling define actions on a journal.
type Journaling interface {
	// Open opens the journal's current entry.
	Open() error

	// Commit latest changes to an entry to version control.
	Close() error
}
