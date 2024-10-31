// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SLOQuery] = (*SLOQueryBuilder)(nil)

// SLO sub-query properties.
type SLOQueryBuilder struct {
	internal *SLOQuery
	errors   map[string]cog.BuildErrors
}

func NewSLOQueryBuilder() *SLOQueryBuilder {
	resource := &SLOQuery{}
	builder := &SLOQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *SLOQueryBuilder) Build() (SLOQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return SLOQuery{}, err
	}

	return *builder.internal, nil
}

// GCP project to execute the query against.
func (builder *SLOQueryBuilder) ProjectName(projectName string) *SLOQueryBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

// Alignment function to be used. Defaults to ALIGN_MEAN.
func (builder *SLOQueryBuilder) PerSeriesAligner(perSeriesAligner string) *SLOQueryBuilder {
	builder.internal.PerSeriesAligner = &perSeriesAligner

	return builder
}

// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
func (builder *SLOQueryBuilder) AlignmentPeriod(alignmentPeriod string) *SLOQueryBuilder {
	builder.internal.AlignmentPeriod = &alignmentPeriod

	return builder
}

// SLO selector.
func (builder *SLOQueryBuilder) SelectorName(selectorName string) *SLOQueryBuilder {
	builder.internal.SelectorName = selectorName

	return builder
}

// ID for the service the SLO is in.
func (builder *SLOQueryBuilder) ServiceId(serviceId string) *SLOQueryBuilder {
	builder.internal.ServiceId = serviceId

	return builder
}

// Name for the service the SLO is in.
func (builder *SLOQueryBuilder) ServiceName(serviceName string) *SLOQueryBuilder {
	builder.internal.ServiceName = serviceName

	return builder
}

// ID for the SLO.
func (builder *SLOQueryBuilder) SloId(sloId string) *SLOQueryBuilder {
	builder.internal.SloId = sloId

	return builder
}

// Name of the SLO.
func (builder *SLOQueryBuilder) SloName(sloName string) *SLOQueryBuilder {
	builder.internal.SloName = sloName

	return builder
}

// SLO goal value.
func (builder *SLOQueryBuilder) Goal(goal float64) *SLOQueryBuilder {
	builder.internal.Goal = &goal

	return builder
}

// Specific lookback period for the SLO.
func (builder *SLOQueryBuilder) LookbackPeriod(lookbackPeriod string) *SLOQueryBuilder {
	builder.internal.LookbackPeriod = &lookbackPeriod

	return builder
}

func (builder *SLOQueryBuilder) applyDefaults() {
}
