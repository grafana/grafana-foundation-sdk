// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package folder

// TODO:
// common metadata will soon support setting the parent folder in the metadata
type Folder struct {
	// Unique folder id. (will be k8s name)
	Uid string `json:"uid"`
	// Folder title
	Title string `json:"title"`
	// Description of the folder.
	Description *string `json:"description,omitempty"`
}
