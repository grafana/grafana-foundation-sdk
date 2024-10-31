// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryPanel] = (*LibraryPanelBuilder)(nil)

type LibraryPanelBuilder struct {
	internal *LibraryPanel
	errors   map[string]cog.BuildErrors
}

func NewLibraryPanelBuilder() *LibraryPanelBuilder {
	resource := &LibraryPanel{}
	builder := &LibraryPanelBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LibraryPanelBuilder) Build() (LibraryPanel, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryPanel{}, err
	}

	return *builder.internal, nil
}

// Folder UID
func (builder *LibraryPanelBuilder) FolderUid(folderUid string) *LibraryPanelBuilder {
	builder.internal.FolderUid = &folderUid

	return builder
}

// Library element UID
func (builder *LibraryPanelBuilder) Uid(uid string) *LibraryPanelBuilder {
	builder.internal.Uid = uid

	return builder
}

// Panel name (also saved in the model)
func (builder *LibraryPanelBuilder) Name(name string) *LibraryPanelBuilder {
	builder.internal.Name = name

	return builder
}

// Panel description
func (builder *LibraryPanelBuilder) Description(description string) *LibraryPanelBuilder {
	builder.internal.Description = &description

	return builder
}

// The panel type (from inside the model)
func (builder *LibraryPanelBuilder) Type(typeArg string) *LibraryPanelBuilder {
	builder.internal.Type = typeArg

	return builder
}

// Dashboard version when this was saved (zero if unknown)
func (builder *LibraryPanelBuilder) SchemaVersion(schemaVersion uint16) *LibraryPanelBuilder {
	builder.internal.SchemaVersion = &schemaVersion

	return builder
}

// panel version, incremented each time the dashboard is updated.
func (builder *LibraryPanelBuilder) Version(version int64) *LibraryPanelBuilder {
	builder.internal.Version = version

	return builder
}

// TODO: should be the same panel schema defined in dashboard
// Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
func (builder *LibraryPanelBuilder) Model(model cog.Builder[LibrarypanelLibraryPanelModel]) *LibraryPanelBuilder {
	modelResource, err := model.Build()
	if err != nil {
		builder.errors["model"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Model = modelResource

	return builder
}

// Object storage metadata
func (builder *LibraryPanelBuilder) Meta(meta cog.Builder[LibraryElementDTOMeta]) *LibraryPanelBuilder {
	metaResource, err := meta.Build()
	if err != nil {
		builder.errors["meta"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Meta = &metaResource

	return builder
}

func (builder *LibraryPanelBuilder) applyDefaults() {
}
