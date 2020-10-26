// VCS relates to anything involving version control, such as persisting
// the state of the journal on a VCS service.

package vcs

// Versioner captures behavior that relates to version control.
type Versioner interface {
	// Create an entry from version control.
	Commit(path interface{}) error
}
