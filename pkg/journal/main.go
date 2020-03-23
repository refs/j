package journal

// Journaling defines journaling behavior
type Journaling interface {
	// Open creates a new journal entry or opens the current one
	Open()
}
