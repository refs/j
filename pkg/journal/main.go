package journal

// Journaling define actions on a journal.
type Journaling interface {
	// Open opens the journal's current entry.
	Open() error
}
