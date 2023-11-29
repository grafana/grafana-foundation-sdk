// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryElementDTOMeta] = (*LibraryElementDTOMetaBuilder)(nil)

type LibraryElementDTOMetaBuilder struct {
	internal *LibraryElementDTOMeta
	errors   map[string]cog.BuildErrors
}

func NewLibraryElementDTOMetaBuilder() *LibraryElementDTOMetaBuilder {
	resource := &LibraryElementDTOMeta{}
	builder := &LibraryElementDTOMetaBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) Build() (LibraryElementDTOMeta, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("LibraryElementDTOMeta", err)...)
	}

	if len(errs) != 0 {
		return LibraryElementDTOMeta{}, errs
	}

	return *builder.internal, nil
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

func (builder *LibraryElementDTOMetaBuilder) Created(created string) *LibraryElementDTOMetaBuilder {
	builder.internal.Created = created

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) Updated(updated string) *LibraryElementDTOMetaBuilder {
	builder.internal.Updated = updated

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) CreatedBy(createdBy cog.Builder[LibraryElementDTOMetaUser]) *LibraryElementDTOMetaBuilder {
	createdByResource, err := createdBy.Build()
	if err != nil {
		builder.errors["createdBy"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.CreatedBy = createdByResource

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) UpdatedBy(updatedBy cog.Builder[LibraryElementDTOMetaUser]) *LibraryElementDTOMetaBuilder {
	updatedByResource, err := updatedBy.Build()
	if err != nil {
		builder.errors["updatedBy"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.UpdatedBy = updatedByResource

	return builder
}

func (builder *LibraryElementDTOMetaBuilder) applyDefaults() {
}
