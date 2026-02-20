// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package librarypanel

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryElementDTOMetaUser] = (*LibraryElementDTOMetaUserBuilder)(nil)

type LibraryElementDTOMetaUserBuilder struct {
	internal *LibraryElementDTOMetaUser
	errors   cog.BuildErrors
}

func NewLibraryElementDTOMetaUserBuilder() *LibraryElementDTOMetaUserBuilder {
	resource := NewLibraryElementDTOMetaUser()
	builder := &LibraryElementDTOMetaUserBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LibraryElementDTOMetaUserBuilder) Build() (LibraryElementDTOMetaUser, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryElementDTOMetaUser{}, err
	}

	if len(builder.errors) > 0 {
		return LibraryElementDTOMetaUser{}, cog.MakeBuildErrors("librarypanel.libraryElementDTOMetaUser", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LibraryElementDTOMetaUserBuilder) RecordError(path string, err error) *LibraryElementDTOMetaUserBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
