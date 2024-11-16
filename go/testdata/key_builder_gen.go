// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Key] = (*KeyBuilder)(nil)

type KeyBuilder struct {
	internal *Key
	errors   map[string]cog.BuildErrors
}

func NewKeyBuilder() *KeyBuilder {
	resource := NewKey()
	builder := &KeyBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *KeyBuilder) Build() (Key, error) {
	if err := builder.internal.Validate(); err != nil {
		return Key{}, err
	}

	return *builder.internal, nil
}

func (builder *KeyBuilder) Tick(tick float64) *KeyBuilder {
	builder.internal.Tick = tick

	return builder
}

func (builder *KeyBuilder) Type(typeArg string) *KeyBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *KeyBuilder) Uid(uid string) *KeyBuilder {
	builder.internal.Uid = &uid

	return builder
}
