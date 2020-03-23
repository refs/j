package journal

// Journaling captures journaling behavior.
type Journaling interface {
	// Open opens a journal entry.
	Open()
}
