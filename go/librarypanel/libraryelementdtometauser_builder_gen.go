// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryElementDTOMetaUser] = (*LibraryElementDTOMetaUserBuilder)(nil)

type LibraryElementDTOMetaUserBuilder struct {
	internal *LibraryElementDTOMetaUser
	errors   map[string]cog.BuildErrors
}

func NewLibraryElementDTOMetaUserBuilder() *LibraryElementDTOMetaUserBuilder {
	resource := &LibraryElementDTOMetaUser{}
	builder := &LibraryElementDTOMetaUserBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LibraryElementDTOMetaUserBuilder) Build() (LibraryElementDTOMetaUser, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryElementDTOMetaUser{}, err
	}

	return *builder.internal, nil
}

func (builder *LibraryElementDTOMetaUserBuilder) Id(id int64) *LibraryElementDTOMetaUserBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *LibraryElementDTOMetaUserBuilder) Name(name string) *LibraryElementDTOMetaUserBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *LibraryElementDTOMetaUserBuilder) AvatarUrl(avatarUrl string) *LibraryElementDTOMetaUserBuilder {
	builder.internal.AvatarUrl = avatarUrl

	return builder
}

func (builder *LibraryElementDTOMetaUserBuilder) applyDefaults() {
}
