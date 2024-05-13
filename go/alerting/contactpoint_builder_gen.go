// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ContactPoint] = (*ContactPointBuilder)(nil)

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
type ContactPointBuilder struct {
	internal *ContactPoint
	errors   map[string]cog.BuildErrors
}

func NewContactPointBuilder() *ContactPointBuilder {
	resource := &ContactPoint{}
	builder := &ContactPointBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ContactPointBuilder) Build() (ContactPoint, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ContactPoint", err)...)
	}

	if len(errs) != 0 {
		return ContactPoint{}, errs
	}

	return *builder.internal, nil
}

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
func (builder *ContactPointBuilder) DisableResolveMessage(disableResolveMessage bool) *ContactPointBuilder {
	builder.internal.DisableResolveMessage = &disableResolveMessage

	return builder
}

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
func (builder *ContactPointBuilder) Name(name string) *ContactPointBuilder {
	builder.internal.Name = &name

	return builder
}

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
func (builder *ContactPointBuilder) Provenance(provenance string) *ContactPointBuilder {
	builder.internal.Provenance = &provenance

	return builder
}

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
func (builder *ContactPointBuilder) Settings(settings Json) *ContactPointBuilder {
	builder.internal.Settings = settings

	return builder
}

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
func (builder *ContactPointBuilder) Type(typeArg ContactPointType) *ContactPointBuilder {
	builder.internal.Type = typeArg

	return builder
}

// EmbeddedContactPoint is the contact point type that is used
// by grafanas embedded alertmanager implementation.
func (builder *ContactPointBuilder) Uid(uid string) *ContactPointBuilder {
	if !(len([]rune(uid)) >= 1) {
		builder.errors["uid"] = cog.MakeBuildErrors("uid", errors.New("len([]rune(uid)) must be >= 1"))
		return builder
	}
	if !(len([]rune(uid)) <= 40) {
		builder.errors["uid"] = cog.MakeBuildErrors("uid", errors.New("len([]rune(uid)) must be <= 40"))
		return builder
	}
	builder.internal.Uid = &uid

	return builder
}

func (builder *ContactPointBuilder) applyDefaults() {
}
