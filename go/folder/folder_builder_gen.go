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

func NewFolderBuilder(title string) *FolderBuilder {
	resource := &Folder{}
	builder := &FolderBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Title = title

	return builder
}

func (builder *FolderBuilder) Build() (Folder, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Folder", err)...)
	}

	if len(errs) != 0 {
		return Folder{}, errs
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
