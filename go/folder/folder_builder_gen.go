// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package folder

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Folder] = (*FolderBuilder)(nil)

// TODO:
// common metadata will soon support setting the parent folder in the metadata
type FolderBuilder struct {
	internal *Folder
	errors   map[string]cog.BuildErrors
}

func NewFolderBuilder() *FolderBuilder {
	resource := &Folder{}
	builder := &FolderBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *FolderBuilder) Build() (Folder, error) {
	if err := builder.internal.Validate(); err != nil {
		return Folder{}, err
	}

	return *builder.internal, nil
}

// Unique folder id. (will be k8s name)
func (builder *FolderBuilder) Uid(uid string) *FolderBuilder {
	builder.internal.Uid = uid

	return builder
}

// Folder title
func (builder *FolderBuilder) Title(title string) *FolderBuilder {
	builder.internal.Title = title

	return builder
}

// Description of the folder.
func (builder *FolderBuilder) Description(description string) *FolderBuilder {
	builder.internal.Description = &description

	return builder
}

func (builder *FolderBuilder) applyDefaults() {
}
