// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	"time"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryElementDTOMeta] = (*LibraryElementDTOMetaBuilder)(nil)

type LibraryElementDTOMetaBuilder struct {
	internal *LibraryElementDTOMeta
	errors   cog.BuildErrors
}

func NewLibraryElementDTOMetaBuilder() *LibraryElementDTOMetaBuilder {
	resource := NewLibraryElementDTOMeta()
	builder := &LibraryElementDTOMetaBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) Build() (LibraryElementDTOMeta, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryElementDTOMeta{}, err
	}

	if len(builder.errors) > 0 {
		return LibraryElementDTOMeta{}, cog.MakeBuildErrors("librarypanel.libraryElementDTOMeta", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LibraryElementDTOMetaBuilder) RecordError(path string, err error) *LibraryElementDTOMetaBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *LibraryElementDTOMetaBuilder) FolderName(folderName string) *LibraryElementDTOMetaBuilder {
	builder.internal.FolderName = folderName

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) FolderUid(folderUid string) *LibraryElementDTOMetaBuilder {
	builder.internal.FolderUid = folderUid

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) ConnectedDashboards(connectedDashboards int64) *LibraryElementDTOMetaBuilder {
	builder.internal.ConnectedDashboards = connectedDashboards

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) Created(created time.Time) *LibraryElementDTOMetaBuilder {
	builder.internal.Created = created

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) Updated(updated time.Time) *LibraryElementDTOMetaBuilder {
	builder.internal.Updated = updated

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) CreatedBy(createdBy cog.Builder[LibraryElementDTOMetaUser]) *LibraryElementDTOMetaBuilder {
	createdByResource, err := createdBy.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.CreatedBy = createdByResource

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) UpdatedBy(updatedBy cog.Builder[LibraryElementDTOMetaUser]) *LibraryElementDTOMetaBuilder {
	updatedByResource, err := updatedBy.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.UpdatedBy = updatedByResource

	return builder
}
