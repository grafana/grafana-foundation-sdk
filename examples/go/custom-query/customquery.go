package main

import (
	"encoding/json"

	"github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type CustomQuery struct {
	// RefId and Hide are expected on all queries
	RefId *string `json:"refId,omitempty"`
	Hide  *bool   `json:"hide,omitempty"`

	// Query is specific to the CustomQuery type
	Query string `json:"query,omitempty"`
}

// Let cog know that CustomQuery is a Dataquery variant
func (resource CustomQuery) ImplementsDataqueryVariant() {}

func CustomQueryVariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "custom", // datasource plugin ID
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := &CustomQuery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

// Compile-time check to ensure that CustomQueryBuilder indeed is
// a builder for cogvariants.Dataquery
var _ cog.Builder[cogvariants.Dataquery] = (*CustomQueryBuilder)(nil)

type CustomQueryBuilder struct {
	internal *CustomQuery
}

func NewCustomQueryBuilder(query string) *CustomQueryBuilder {
	return &CustomQueryBuilder{
		internal: &CustomQuery{Query: query},
	}
}

func (builder *CustomQueryBuilder) Build() (cogvariants.Dataquery, error) {
	return *builder.internal, nil
}

func (builder *CustomQueryBuilder) RefId(refId string) *CustomQueryBuilder {
	builder.internal.RefId = &refId
	return builder
}

func (builder *CustomQueryBuilder) Hide(hide bool) *CustomQueryBuilder {
	builder.internal.Hide = &hide
	return builder
}
