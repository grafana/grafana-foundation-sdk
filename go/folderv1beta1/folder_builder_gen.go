// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package folderv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[Folder] = (*FolderBuilder)(nil)

type FolderBuilder struct {
	internal *Folder
	errors   cog.BuildErrors
}

func NewFolderBuilder(title string) *FolderBuilder {
	resource := NewFolder()
	builder := &FolderBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Title = title

	return builder
}

// Creates a resource manifest from a Folder.
func Manifest(name string, folder cog.Builder[Folder]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("folder.grafana.app/v1beta1")

	builder.Kind("Folder")

	builder.Metadata(resource.Named(name))
	folderResource, err := folder.Build()
	if err != nil {
		builder.RecordError("Manifest(folder)", err)
		return builder
	}
	builder.Spec(folderResource)

	return builder
}

func (builder *FolderBuilder) Build() (Folder, error) {
	if err := builder.internal.Validate(); err != nil {
		return Folder{}, err
	}

	if len(builder.errors) > 0 {
		return Folder{}, cog.MakeBuildErrors("folderv1beta1.folder", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FolderBuilder) RecordError(path string, err error) *FolderBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *FolderBuilder) Title(title string) *FolderBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *FolderBuilder) Description(description string) *FolderBuilder {
	builder.internal.Description = &description

	return builder
}
