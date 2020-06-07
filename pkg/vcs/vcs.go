// VCS relates to anything involving version control, such as persisting
// the state of the journal on a VCS service.

package vcs

// VCS captures behavior that relates to version control
type VCS interface {
	// Create an entry from version control.
	Create(entry interface{}) error

	// Update an entry from version control.
	Update(entry interface{}) error

	// Delete an entry from version control.
	Delete(entry interface{}) error
}
