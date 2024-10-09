// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type BucketAggregation = DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested

type MetricAggregation = CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics

type BucketAggregationType string

const (
	BucketAggregationTypeTerms         BucketAggregationType = "terms"
	BucketAggregationTypeFilters       BucketAggregationType = "filters"
	BucketAggregationTypeGeohashGrid   BucketAggregationType = "geohash_grid"
	BucketAggregationTypeDateHistogram BucketAggregationType = "date_histogram"
	BucketAggregationTypeHistogram     BucketAggregationType = "histogram"
	BucketAggregationTypeNested        BucketAggregationType = "nested"
)

type BaseBucketAggregation struct {
	Id       string                `json:"id"`
	Type     BucketAggregationType `json:"type"`
	Settings any                   `json:"settings,omitempty"`
}

func (resource BaseBucketAggregation) Equals(other BaseBucketAggregation) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Settings, other.Settings) {
		return false
	}

	return true
}

type BucketAggregationWithField struct {
	Field    *string               `json:"field,omitempty"`
	Id       string                `json:"id"`
	Type     BucketAggregationType `json:"type"`
	Settings any                   `json:"settings,omitempty"`
}

func (resource BucketAggregationWithField) Equals(other BucketAggregationWithField) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Settings, other.Settings) {
		return false
	}

	return true
}

type DateHistogram struct {
	Field    *string                             `json:"field,omitempty"`
	Id       string                              `json:"id"`
	Type     string                              `json:"type"`
	Settings *ElasticsearchDateHistogramSettings `json:"settings,omitempty"`
}

func (resource DateHistogram) Equals(other DateHistogram) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

type DateHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
	TrimEdges   *string `json:"trimEdges,omitempty"`
	Offset      *string `json:"offset,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}

func (resource DateHistogramSettings) Equals(other DateHistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.TrimEdges == nil && other.TrimEdges != nil || resource.TrimEdges != nil && other.TrimEdges == nil {
		return false
	}

	if resource.TrimEdges != nil {
		if *resource.TrimEdges != *other.TrimEdges {
			return false
		}
	}
	if resource.Offset == nil && other.Offset != nil || resource.Offset != nil && other.Offset == nil {
		return false
	}

	if resource.Offset != nil {
		if *resource.Offset != *other.Offset {
			return false
		}
	}
	if resource.TimeZone == nil && other.TimeZone != nil || resource.TimeZone != nil && other.TimeZone == nil {
		return false
	}

	if resource.TimeZone != nil {
		if *resource.TimeZone != *other.TimeZone {
			return false
		}
	}

	return true
}

type Histogram struct {
	Field    *string                         `json:"field,omitempty"`
	Id       string                          `json:"id"`
	Type     string                          `json:"type"`
	Settings *ElasticsearchHistogramSettings `json:"settings,omitempty"`
}

func (resource Histogram) Equals(other Histogram) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

type HistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
}

func (resource HistogramSettings) Equals(other HistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}

	return true
}

type TermsOrder string

const (
	TermsOrderDesc TermsOrder = "desc"
	TermsOrderAsc  TermsOrder = "asc"
)

type Nested struct {
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Settings any     `json:"settings,omitempty"`
}

func (resource Nested) Equals(other Nested) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Settings, other.Settings) {
		return false
	}

	return true
}

type Terms struct {
	Field    *string                     `json:"field,omitempty"`
	Id       string                      `json:"id"`
	Type     string                      `json:"type"`
	Settings *ElasticsearchTermsSettings `json:"settings,omitempty"`
}

func (resource Terms) Equals(other Terms) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

type TermsSettings struct {
	Order       *TermsOrder `json:"order,omitempty"`
	Size        *string     `json:"size,omitempty"`
	MinDocCount *string     `json:"min_doc_count,omitempty"`
	OrderBy     *string     `json:"orderBy,omitempty"`
	Missing     *string     `json:"missing,omitempty"`
}

func (resource TermsSettings) Equals(other TermsSettings) bool {
	if resource.Order == nil && other.Order != nil || resource.Order != nil && other.Order == nil {
		return false
	}

	if resource.Order != nil {
		if *resource.Order != *other.Order {
			return false
		}
	}
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if *resource.OrderBy != *other.OrderBy {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type Filters struct {
	Id       string                        `json:"id"`
	Type     string                        `json:"type"`
	Settings *ElasticsearchFiltersSettings `json:"settings,omitempty"`
}

func (resource Filters) Equals(other Filters) bool {
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

type Filter struct {
	Query string `json:"query"`
	Label string `json:"label"`
}

func (resource Filter) Equals(other Filter) bool {
	if resource.Query != other.Query {
		return false
	}
	if resource.Label != other.Label {
		return false
	}

	return true
}

type FiltersSettings struct {
	Filters []Filter `json:"filters,omitempty"`
}

func (resource FiltersSettings) Equals(other FiltersSettings) bool {

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}

	return true
}

type GeoHashGrid struct {
	Field    *string                           `json:"field,omitempty"`
	Id       string                            `json:"id"`
	Type     string                            `json:"type"`
	Settings *ElasticsearchGeoHashGridSettings `json:"settings,omitempty"`
}

func (resource GeoHashGrid) Equals(other GeoHashGrid) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}

	return true
}

type GeoHashGridSettings struct {
	Precision *string `json:"precision,omitempty"`
}

func (resource GeoHashGridSettings) Equals(other GeoHashGridSettings) bool {
	if resource.Precision == nil && other.Precision != nil || resource.Precision != nil && other.Precision == nil {
		return false
	}

	if resource.Precision != nil {
		if *resource.Precision != *other.Precision {
			return false
		}
	}

	return true
}

type PipelineMetricAggregationType string

const (
	PipelineMetricAggregationTypeMovingAvg     PipelineMetricAggregationType = "moving_avg"
	PipelineMetricAggregationTypeMovingFn      PipelineMetricAggregationType = "moving_fn"
	PipelineMetricAggregationTypeDerivative    PipelineMetricAggregationType = "derivative"
	PipelineMetricAggregationTypeSerialDiff    PipelineMetricAggregationType = "serial_diff"
	PipelineMetricAggregationTypeCumulativeSum PipelineMetricAggregationType = "cumulative_sum"
	PipelineMetricAggregationTypeBucketScript  PipelineMetricAggregationType = "bucket_script"
)

type MetricAggregationType = StringOrPipelineMetricAggregationType

type BaseMetricAggregation struct {
	Type MetricAggregationType `json:"type"`
	Id   string                `json:"id"`
	Hide *bool                 `json:"hide,omitempty"`
}

func (resource BaseMetricAggregation) Equals(other BaseMetricAggregation) bool {
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type PipelineVariable struct {
	Name        string `json:"name"`
	PipelineAgg string `json:"pipelineAgg"`
}

func (resource PipelineVariable) Equals(other PipelineVariable) bool {
	if resource.Name != other.Name {
		return false
	}
	if resource.PipelineAgg != other.PipelineAgg {
		return false
	}

	return true
}

type MetricAggregationWithField struct {
	Field *string               `json:"field,omitempty"`
	Type  MetricAggregationType `json:"type"`
	Id    string                `json:"id"`
	Hide  *bool                 `json:"hide,omitempty"`
}

func (resource MetricAggregationWithField) Equals(other MetricAggregationWithField) bool {
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type MetricAggregationWithMissingSupport struct {
	Settings *ElasticsearchMetricAggregationWithMissingSupportSettings `json:"settings,omitempty"`
	Type     MetricAggregationType                                     `json:"type"`
	Id       string                                                    `json:"id"`
	Hide     *bool                                                     `json:"hide,omitempty"`
}

func (resource MetricAggregationWithMissingSupport) Equals(other MetricAggregationWithMissingSupport) bool {
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type InlineScript = StringOrElasticsearchInlineScript

type MetricAggregationWithInlineScript struct {
	Settings *ElasticsearchMetricAggregationWithInlineScriptSettings `json:"settings,omitempty"`
	Type     MetricAggregationType                                   `json:"type"`
	Id       string                                                  `json:"id"`
	Hide     *bool                                                   `json:"hide,omitempty"`
}

func (resource MetricAggregationWithInlineScript) Equals(other MetricAggregationWithInlineScript) bool {
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Count struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Hide *bool  `json:"hide,omitempty"`
}

func (resource Count) Equals(other Count) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Average struct {
	Type     string                        `json:"type"`
	Field    *string                       `json:"field,omitempty"`
	Id       string                        `json:"id"`
	Settings *ElasticsearchAverageSettings `json:"settings,omitempty"`
	Hide     *bool                         `json:"hide,omitempty"`
}

func (resource Average) Equals(other Average) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Sum struct {
	Type     string                    `json:"type"`
	Field    *string                   `json:"field,omitempty"`
	Id       string                    `json:"id"`
	Settings *ElasticsearchSumSettings `json:"settings,omitempty"`
	Hide     *bool                     `json:"hide,omitempty"`
}

func (resource Sum) Equals(other Sum) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Max struct {
	Type     string                    `json:"type"`
	Field    *string                   `json:"field,omitempty"`
	Id       string                    `json:"id"`
	Settings *ElasticsearchMaxSettings `json:"settings,omitempty"`
	Hide     *bool                     `json:"hide,omitempty"`
}

func (resource Max) Equals(other Max) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Min struct {
	Type     string                    `json:"type"`
	Field    *string                   `json:"field,omitempty"`
	Id       string                    `json:"id"`
	Settings *ElasticsearchMinSettings `json:"settings,omitempty"`
	Hide     *bool                     `json:"hide,omitempty"`
}

func (resource Min) Equals(other Min) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type ExtendedStatMetaType string

const (
	ExtendedStatMetaTypeAvg                     ExtendedStatMetaType = "avg"
	ExtendedStatMetaTypeMin                     ExtendedStatMetaType = "min"
	ExtendedStatMetaTypeMax                     ExtendedStatMetaType = "max"
	ExtendedStatMetaTypeSum                     ExtendedStatMetaType = "sum"
	ExtendedStatMetaTypeCount                   ExtendedStatMetaType = "count"
	ExtendedStatMetaTypeStdDeviation            ExtendedStatMetaType = "std_deviation"
	ExtendedStatMetaTypeStdDeviationBoundsUpper ExtendedStatMetaType = "std_deviation_bounds_upper"
	ExtendedStatMetaTypeStdDeviationBoundsLower ExtendedStatMetaType = "std_deviation_bounds_lower"
)

type ExtendedStat struct {
	Label string               `json:"label"`
	Value ExtendedStatMetaType `json:"value"`
}

func (resource ExtendedStat) Equals(other ExtendedStat) bool {
	if resource.Label != other.Label {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

type ExtendedStats struct {
	Type     string                              `json:"type"`
	Settings *ElasticsearchExtendedStatsSettings `json:"settings,omitempty"`
	Field    *string                             `json:"field,omitempty"`
	Id       string                              `json:"id"`
	Meta     any                                 `json:"meta,omitempty"`
	Hide     *bool                               `json:"hide,omitempty"`
}

func (resource ExtendedStats) Equals(other ExtendedStats) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.Meta, other.Meta) {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Percentiles struct {
	Type     string                            `json:"type"`
	Field    *string                           `json:"field,omitempty"`
	Id       string                            `json:"id"`
	Settings *ElasticsearchPercentilesSettings `json:"settings,omitempty"`
	Hide     *bool                             `json:"hide,omitempty"`
}

func (resource Percentiles) Equals(other Percentiles) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type UniqueCount struct {
	Type     string                            `json:"type"`
	Field    *string                           `json:"field,omitempty"`
	Id       string                            `json:"id"`
	Settings *ElasticsearchUniqueCountSettings `json:"settings,omitempty"`
	Hide     *bool                             `json:"hide,omitempty"`
}

func (resource UniqueCount) Equals(other UniqueCount) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type RawDocument struct {
	Type     string                            `json:"type"`
	Id       string                            `json:"id"`
	Settings *ElasticsearchRawDocumentSettings `json:"settings,omitempty"`
	Hide     *bool                             `json:"hide,omitempty"`
}

func (resource RawDocument) Equals(other RawDocument) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type RawData struct {
	Type     string                        `json:"type"`
	Id       string                        `json:"id"`
	Settings *ElasticsearchRawDataSettings `json:"settings,omitempty"`
	Hide     *bool                         `json:"hide,omitempty"`
}

func (resource RawData) Equals(other RawData) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Logs struct {
	Type     string                     `json:"type"`
	Id       string                     `json:"id"`
	Settings *ElasticsearchLogsSettings `json:"settings,omitempty"`
	Hide     *bool                      `json:"hide,omitempty"`
}

func (resource Logs) Equals(other Logs) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Rate struct {
	Type     string                     `json:"type"`
	Field    *string                    `json:"field,omitempty"`
	Id       string                     `json:"id"`
	Settings *ElasticsearchRateSettings `json:"settings,omitempty"`
	Hide     *bool                      `json:"hide,omitempty"`
}

func (resource Rate) Equals(other Rate) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type BasePipelineMetricAggregation struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Hide        *bool   `json:"hide,omitempty"`
}

func (resource BasePipelineMetricAggregation) Equals(other BasePipelineMetricAggregation) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type PipelineMetricAggregationWithMultipleBucketPaths struct {
	PipelineVariables []PipelineVariable    `json:"pipelineVariables,omitempty"`
	Type              MetricAggregationType `json:"type"`
	Id                string                `json:"id"`
	Hide              *bool                 `json:"hide,omitempty"`
}

func (resource PipelineMetricAggregationWithMultipleBucketPaths) Equals(other PipelineMetricAggregationWithMultipleBucketPaths) bool {

	if len(resource.PipelineVariables) != len(other.PipelineVariables) {
		return false
	}

	for i1 := range resource.PipelineVariables {
		if !resource.PipelineVariables[i1].Equals(other.PipelineVariables[i1]) {
			return false
		}
	}
	if !resource.Type.Equals(other.Type) {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type MovingAverageModel string

const (
	MovingAverageModelSimple      MovingAverageModel = "simple"
	MovingAverageModelLinear      MovingAverageModel = "linear"
	MovingAverageModelEwma        MovingAverageModel = "ewma"
	MovingAverageModelHolt        MovingAverageModel = "holt"
	MovingAverageModelHoltWinters MovingAverageModel = "holt_winters"
)

type MovingAverageModelOption struct {
	Label string             `json:"label"`
	Value MovingAverageModel `json:"value"`
}

func (resource MovingAverageModelOption) Equals(other MovingAverageModelOption) bool {
	if resource.Label != other.Label {
		return false
	}
	if resource.Value != other.Value {
		return false
	}

	return true
}

type BaseMovingAverageModelSettings struct {
	Model   MovingAverageModel `json:"model"`
	Window  string             `json:"window"`
	Predict string             `json:"predict"`
}

func (resource BaseMovingAverageModelSettings) Equals(other BaseMovingAverageModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

type MovingAverageSimpleModelSettings struct {
	Model   string `json:"model"`
	Window  string `json:"window"`
	Predict string `json:"predict"`
}

func (resource MovingAverageSimpleModelSettings) Equals(other MovingAverageSimpleModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

type MovingAverageLinearModelSettings struct {
	Model   string `json:"model"`
	Window  string `json:"window"`
	Predict string `json:"predict"`
}

func (resource MovingAverageLinearModelSettings) Equals(other MovingAverageLinearModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

type MovingAverageEWMAModelSettings struct {
	Model    string                                               `json:"model"`
	Settings *ElasticsearchMovingAverageEWMAModelSettingsSettings `json:"settings,omitempty"`
	Window   string                                               `json:"window"`
	Minimize bool                                                 `json:"minimize"`
	Predict  string                                               `json:"predict"`
}

func (resource MovingAverageEWMAModelSettings) Equals(other MovingAverageEWMAModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Minimize != other.Minimize {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

type MovingAverageHoltModelSettings struct {
	Model    string                                              `json:"model"`
	Settings ElasticsearchMovingAverageHoltModelSettingsSettings `json:"settings"`
	Window   string                                              `json:"window"`
	Minimize bool                                                `json:"minimize"`
	Predict  string                                              `json:"predict"`
}

func (resource MovingAverageHoltModelSettings) Equals(other MovingAverageHoltModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if !resource.Settings.Equals(other.Settings) {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Minimize != other.Minimize {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

type MovingAverageHoltWintersModelSettings struct {
	Model    string                                                     `json:"model"`
	Settings ElasticsearchMovingAverageHoltWintersModelSettingsSettings `json:"settings"`
	Window   string                                                     `json:"window"`
	Minimize bool                                                       `json:"minimize"`
	Predict  string                                                     `json:"predict"`
}

func (resource MovingAverageHoltWintersModelSettings) Equals(other MovingAverageHoltWintersModelSettings) bool {
	if resource.Model != other.Model {
		return false
	}
	if !resource.Settings.Equals(other.Settings) {
		return false
	}
	if resource.Window != other.Window {
		return false
	}
	if resource.Minimize != other.Minimize {
		return false
	}
	if resource.Predict != other.Predict {
		return false
	}

	return true
}

// #MovingAverage's settings are overridden in types.ts
type MovingAverage struct {
	PipelineAgg *string        `json:"pipelineAgg,omitempty"`
	Field       *string        `json:"field,omitempty"`
	Type        string         `json:"type"`
	Id          string         `json:"id"`
	Settings    map[string]any `json:"settings,omitempty"`
	Hide        *bool          `json:"hide,omitempty"`
}

func (resource MovingAverage) Equals(other MovingAverage) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}

	if len(resource.Settings) != len(other.Settings) {
		return false
	}

	for key1 := range resource.Settings {
		// is DeepEqual good enough here?
		if !reflect.DeepEqual(resource.Settings[key1], other.Settings[key1]) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type MovingFunction struct {
	PipelineAgg *string                              `json:"pipelineAgg,omitempty"`
	Field       *string                              `json:"field,omitempty"`
	Type        string                               `json:"type"`
	Id          string                               `json:"id"`
	Settings    *ElasticsearchMovingFunctionSettings `json:"settings,omitempty"`
	Hide        *bool                                `json:"hide,omitempty"`
}

func (resource MovingFunction) Equals(other MovingFunction) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type Derivative struct {
	PipelineAgg *string                          `json:"pipelineAgg,omitempty"`
	Field       *string                          `json:"field,omitempty"`
	Type        string                           `json:"type"`
	Id          string                           `json:"id"`
	Settings    *ElasticsearchDerivativeSettings `json:"settings,omitempty"`
	Hide        *bool                            `json:"hide,omitempty"`
}

func (resource Derivative) Equals(other Derivative) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type SerialDiff struct {
	PipelineAgg *string                          `json:"pipelineAgg,omitempty"`
	Field       *string                          `json:"field,omitempty"`
	Type        string                           `json:"type"`
	Id          string                           `json:"id"`
	Settings    *ElasticsearchSerialDiffSettings `json:"settings,omitempty"`
	Hide        *bool                            `json:"hide,omitempty"`
}

func (resource SerialDiff) Equals(other SerialDiff) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type CumulativeSum struct {
	PipelineAgg *string                             `json:"pipelineAgg,omitempty"`
	Field       *string                             `json:"field,omitempty"`
	Type        string                              `json:"type"`
	Id          string                              `json:"id"`
	Settings    *ElasticsearchCumulativeSumSettings `json:"settings,omitempty"`
	Hide        *bool                               `json:"hide,omitempty"`
}

func (resource CumulativeSum) Equals(other CumulativeSum) bool {
	if resource.PipelineAgg == nil && other.PipelineAgg != nil || resource.PipelineAgg != nil && other.PipelineAgg == nil {
		return false
	}

	if resource.PipelineAgg != nil {
		if *resource.PipelineAgg != *other.PipelineAgg {
			return false
		}
	}
	if resource.Field == nil && other.Field != nil || resource.Field != nil && other.Field == nil {
		return false
	}

	if resource.Field != nil {
		if *resource.Field != *other.Field {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type BucketScript struct {
	Type              string                             `json:"type"`
	PipelineVariables []PipelineVariable                 `json:"pipelineVariables,omitempty"`
	Id                string                             `json:"id"`
	Settings          *ElasticsearchBucketScriptSettings `json:"settings,omitempty"`
	Hide              *bool                              `json:"hide,omitempty"`
}

func (resource BucketScript) Equals(other BucketScript) bool {
	if resource.Type != other.Type {
		return false
	}

	if len(resource.PipelineVariables) != len(other.PipelineVariables) {
		return false
	}

	for i1 := range resource.PipelineVariables {
		if !resource.PipelineVariables[i1].Equals(other.PipelineVariables[i1]) {
			return false
		}
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type TopMetrics struct {
	Type     string                           `json:"type"`
	Id       string                           `json:"id"`
	Settings *ElasticsearchTopMetricsSettings `json:"settings,omitempty"`
	Hide     *bool                            `json:"hide,omitempty"`
}

func (resource TopMetrics) Equals(other TopMetrics) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}

	return true
}

type PipelineMetricAggregation = MovingAverageOrDerivativeOrCumulativeSumOrBucketScript

type MetricAggregationWithSettings = BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics

type Dataquery struct {
	// Alias pattern
	Alias *string `json:"alias,omitempty"`
	// Lucene query
	Query *string `json:"query,omitempty"`
	// Name of time field
	TimeField *string `json:"timeField,omitempty"`
	// List of bucket aggregations
	BucketAggs []BucketAggregation `json:"bucketAggs,omitempty"`
	// List of metric aggregations
	Metrics []MetricAggregation `json:"metrics,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "elasticsearch"
}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "elasticsearch",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery Dataquery
			if cast, ok := input.(*Dataquery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(Dataquery)
			}
			return DataqueryConverter(dataquery)
		},
	}
}

func (resource Dataquery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(Dataquery)
	if !ok {
		return false
	}
	if resource.Alias == nil && other.Alias != nil || resource.Alias != nil && other.Alias == nil {
		return false
	}

	if resource.Alias != nil {
		if *resource.Alias != *other.Alias {
			return false
		}
	}
	if resource.Query == nil && other.Query != nil || resource.Query != nil && other.Query == nil {
		return false
	}

	if resource.Query != nil {
		if *resource.Query != *other.Query {
			return false
		}
	}
	if resource.TimeField == nil && other.TimeField != nil || resource.TimeField != nil && other.TimeField == nil {
		return false
	}

	if resource.TimeField != nil {
		if *resource.TimeField != *other.TimeField {
			return false
		}
	}

	if len(resource.BucketAggs) != len(other.BucketAggs) {
		return false
	}

	for i1 := range resource.BucketAggs {
		if !resource.BucketAggs[i1].Equals(other.BucketAggs[i1]) {
			return false
		}
	}

	if len(resource.Metrics) != len(other.Metrics) {
		return false
	}

	for i1 := range resource.Metrics {
		if !resource.Metrics[i1].Equals(other.Metrics[i1]) {
			return false
		}
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}

	return true
}

type ElasticsearchDateHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
	TrimEdges   *string `json:"trimEdges,omitempty"`
	Offset      *string `json:"offset,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}

func (resource ElasticsearchDateHistogramSettings) Equals(other ElasticsearchDateHistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.TrimEdges == nil && other.TrimEdges != nil || resource.TrimEdges != nil && other.TrimEdges == nil {
		return false
	}

	if resource.TrimEdges != nil {
		if *resource.TrimEdges != *other.TrimEdges {
			return false
		}
	}
	if resource.Offset == nil && other.Offset != nil || resource.Offset != nil && other.Offset == nil {
		return false
	}

	if resource.Offset != nil {
		if *resource.Offset != *other.Offset {
			return false
		}
	}
	if resource.TimeZone == nil && other.TimeZone != nil || resource.TimeZone != nil && other.TimeZone == nil {
		return false
	}

	if resource.TimeZone != nil {
		if *resource.TimeZone != *other.TimeZone {
			return false
		}
	}

	return true
}

type ElasticsearchHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
}

func (resource ElasticsearchHistogramSettings) Equals(other ElasticsearchHistogramSettings) bool {
	if resource.Interval == nil && other.Interval != nil || resource.Interval != nil && other.Interval == nil {
		return false
	}

	if resource.Interval != nil {
		if *resource.Interval != *other.Interval {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}

	return true
}

type ElasticsearchTermsSettings struct {
	Order       *TermsOrder `json:"order,omitempty"`
	Size        *string     `json:"size,omitempty"`
	MinDocCount *string     `json:"min_doc_count,omitempty"`
	OrderBy     *string     `json:"orderBy,omitempty"`
	Missing     *string     `json:"missing,omitempty"`
}

func (resource ElasticsearchTermsSettings) Equals(other ElasticsearchTermsSettings) bool {
	if resource.Order == nil && other.Order != nil || resource.Order != nil && other.Order == nil {
		return false
	}

	if resource.Order != nil {
		if *resource.Order != *other.Order {
			return false
		}
	}
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}
	if resource.MinDocCount == nil && other.MinDocCount != nil || resource.MinDocCount != nil && other.MinDocCount == nil {
		return false
	}

	if resource.MinDocCount != nil {
		if *resource.MinDocCount != *other.MinDocCount {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if *resource.OrderBy != *other.OrderBy {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchFiltersSettings struct {
	Filters []Filter `json:"filters,omitempty"`
}

func (resource ElasticsearchFiltersSettings) Equals(other ElasticsearchFiltersSettings) bool {

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if !resource.Filters[i1].Equals(other.Filters[i1]) {
			return false
		}
	}

	return true
}

type ElasticsearchGeoHashGridSettings struct {
	Precision *string `json:"precision,omitempty"`
}

func (resource ElasticsearchGeoHashGridSettings) Equals(other ElasticsearchGeoHashGridSettings) bool {
	if resource.Precision == nil && other.Precision != nil || resource.Precision != nil && other.Precision == nil {
		return false
	}

	if resource.Precision != nil {
		if *resource.Precision != *other.Precision {
			return false
		}
	}

	return true
}

type ElasticsearchMetricAggregationWithMissingSupportSettings struct {
	Missing *string `json:"missing,omitempty"`
}

func (resource ElasticsearchMetricAggregationWithMissingSupportSettings) Equals(other ElasticsearchMetricAggregationWithMissingSupportSettings) bool {
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchInlineScript struct {
	Inline *string `json:"inline,omitempty"`
}

func (resource ElasticsearchInlineScript) Equals(other ElasticsearchInlineScript) bool {
	if resource.Inline == nil && other.Inline != nil || resource.Inline != nil && other.Inline == nil {
		return false
	}

	if resource.Inline != nil {
		if *resource.Inline != *other.Inline {
			return false
		}
	}

	return true
}

type ElasticsearchMetricAggregationWithInlineScriptSettings struct {
	Script *InlineScript `json:"script,omitempty"`
}

func (resource ElasticsearchMetricAggregationWithInlineScriptSettings) Equals(other ElasticsearchMetricAggregationWithInlineScriptSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}

	return true
}

type ElasticsearchAverageSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

func (resource ElasticsearchAverageSettings) Equals(other ElasticsearchAverageSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchSumSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

func (resource ElasticsearchSumSettings) Equals(other ElasticsearchSumSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchMaxSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

func (resource ElasticsearchMaxSettings) Equals(other ElasticsearchMaxSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchMinSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
}

func (resource ElasticsearchMinSettings) Equals(other ElasticsearchMinSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchExtendedStatsSettings struct {
	Script  *InlineScript `json:"script,omitempty"`
	Missing *string       `json:"missing,omitempty"`
	Sigma   *string       `json:"sigma,omitempty"`
}

func (resource ElasticsearchExtendedStatsSettings) Equals(other ElasticsearchExtendedStatsSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}
	if resource.Sigma == nil && other.Sigma != nil || resource.Sigma != nil && other.Sigma == nil {
		return false
	}

	if resource.Sigma != nil {
		if *resource.Sigma != *other.Sigma {
			return false
		}
	}

	return true
}

type ElasticsearchPercentilesSettings struct {
	Script   *InlineScript `json:"script,omitempty"`
	Missing  *string       `json:"missing,omitempty"`
	Percents []string      `json:"percents,omitempty"`
}

func (resource ElasticsearchPercentilesSettings) Equals(other ElasticsearchPercentilesSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	if len(resource.Percents) != len(other.Percents) {
		return false
	}

	for i1 := range resource.Percents {
		if resource.Percents[i1] != other.Percents[i1] {
			return false
		}
	}

	return true
}

type ElasticsearchUniqueCountSettings struct {
	PrecisionThreshold *string `json:"precision_threshold,omitempty"`
	Missing            *string `json:"missing,omitempty"`
}

func (resource ElasticsearchUniqueCountSettings) Equals(other ElasticsearchUniqueCountSettings) bool {
	if resource.PrecisionThreshold == nil && other.PrecisionThreshold != nil || resource.PrecisionThreshold != nil && other.PrecisionThreshold == nil {
		return false
	}

	if resource.PrecisionThreshold != nil {
		if *resource.PrecisionThreshold != *other.PrecisionThreshold {
			return false
		}
	}
	if resource.Missing == nil && other.Missing != nil || resource.Missing != nil && other.Missing == nil {
		return false
	}

	if resource.Missing != nil {
		if *resource.Missing != *other.Missing {
			return false
		}
	}

	return true
}

type ElasticsearchRawDocumentSettings struct {
	Size *string `json:"size,omitempty"`
}

func (resource ElasticsearchRawDocumentSettings) Equals(other ElasticsearchRawDocumentSettings) bool {
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}

	return true
}

type ElasticsearchRawDataSettings struct {
	Size *string `json:"size,omitempty"`
}

func (resource ElasticsearchRawDataSettings) Equals(other ElasticsearchRawDataSettings) bool {
	if resource.Size == nil && other.Size != nil || resource.Size != nil && other.Size == nil {
		return false
	}

	if resource.Size != nil {
		if *resource.Size != *other.Size {
			return false
		}
	}

	return true
}

type ElasticsearchLogsSettings struct {
	Limit *string `json:"limit,omitempty"`
}

func (resource ElasticsearchLogsSettings) Equals(other ElasticsearchLogsSettings) bool {
	if resource.Limit == nil && other.Limit != nil || resource.Limit != nil && other.Limit == nil {
		return false
	}

	if resource.Limit != nil {
		if *resource.Limit != *other.Limit {
			return false
		}
	}

	return true
}

type ElasticsearchRateSettings struct {
	Unit *string `json:"unit,omitempty"`
	Mode *string `json:"mode,omitempty"`
}

func (resource ElasticsearchRateSettings) Equals(other ElasticsearchRateSettings) bool {
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
			return false
		}
	}
	if resource.Mode == nil && other.Mode != nil || resource.Mode != nil && other.Mode == nil {
		return false
	}

	if resource.Mode != nil {
		if *resource.Mode != *other.Mode {
			return false
		}
	}

	return true
}

type ElasticsearchMovingAverageEWMAModelSettingsSettings struct {
	Alpha *string `json:"alpha,omitempty"`
}

func (resource ElasticsearchMovingAverageEWMAModelSettingsSettings) Equals(other ElasticsearchMovingAverageEWMAModelSettingsSettings) bool {
	if resource.Alpha == nil && other.Alpha != nil || resource.Alpha != nil && other.Alpha == nil {
		return false
	}

	if resource.Alpha != nil {
		if *resource.Alpha != *other.Alpha {
			return false
		}
	}

	return true
}

type ElasticsearchMovingAverageHoltModelSettingsSettings struct {
	Alpha *string `json:"alpha,omitempty"`
	Beta  *string `json:"beta,omitempty"`
}

func (resource ElasticsearchMovingAverageHoltModelSettingsSettings) Equals(other ElasticsearchMovingAverageHoltModelSettingsSettings) bool {
	if resource.Alpha == nil && other.Alpha != nil || resource.Alpha != nil && other.Alpha == nil {
		return false
	}

	if resource.Alpha != nil {
		if *resource.Alpha != *other.Alpha {
			return false
		}
	}
	if resource.Beta == nil && other.Beta != nil || resource.Beta != nil && other.Beta == nil {
		return false
	}

	if resource.Beta != nil {
		if *resource.Beta != *other.Beta {
			return false
		}
	}

	return true
}

type ElasticsearchMovingAverageHoltWintersModelSettingsSettings struct {
	Alpha  *string `json:"alpha,omitempty"`
	Beta   *string `json:"beta,omitempty"`
	Gamma  *string `json:"gamma,omitempty"`
	Period *string `json:"period,omitempty"`
	Pad    *bool   `json:"pad,omitempty"`
}

func (resource ElasticsearchMovingAverageHoltWintersModelSettingsSettings) Equals(other ElasticsearchMovingAverageHoltWintersModelSettingsSettings) bool {
	if resource.Alpha == nil && other.Alpha != nil || resource.Alpha != nil && other.Alpha == nil {
		return false
	}

	if resource.Alpha != nil {
		if *resource.Alpha != *other.Alpha {
			return false
		}
	}
	if resource.Beta == nil && other.Beta != nil || resource.Beta != nil && other.Beta == nil {
		return false
	}

	if resource.Beta != nil {
		if *resource.Beta != *other.Beta {
			return false
		}
	}
	if resource.Gamma == nil && other.Gamma != nil || resource.Gamma != nil && other.Gamma == nil {
		return false
	}

	if resource.Gamma != nil {
		if *resource.Gamma != *other.Gamma {
			return false
		}
	}
	if resource.Period == nil && other.Period != nil || resource.Period != nil && other.Period == nil {
		return false
	}

	if resource.Period != nil {
		if *resource.Period != *other.Period {
			return false
		}
	}
	if resource.Pad == nil && other.Pad != nil || resource.Pad != nil && other.Pad == nil {
		return false
	}

	if resource.Pad != nil {
		if *resource.Pad != *other.Pad {
			return false
		}
	}

	return true
}

type ElasticsearchMovingFunctionSettings struct {
	Window *string       `json:"window,omitempty"`
	Script *InlineScript `json:"script,omitempty"`
	Shift  *string       `json:"shift,omitempty"`
}

func (resource ElasticsearchMovingFunctionSettings) Equals(other ElasticsearchMovingFunctionSettings) bool {
	if resource.Window == nil && other.Window != nil || resource.Window != nil && other.Window == nil {
		return false
	}

	if resource.Window != nil {
		if *resource.Window != *other.Window {
			return false
		}
	}
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}
	if resource.Shift == nil && other.Shift != nil || resource.Shift != nil && other.Shift == nil {
		return false
	}

	if resource.Shift != nil {
		if *resource.Shift != *other.Shift {
			return false
		}
	}

	return true
}

type ElasticsearchDerivativeSettings struct {
	Unit *string `json:"unit,omitempty"`
}

func (resource ElasticsearchDerivativeSettings) Equals(other ElasticsearchDerivativeSettings) bool {
	if resource.Unit == nil && other.Unit != nil || resource.Unit != nil && other.Unit == nil {
		return false
	}

	if resource.Unit != nil {
		if *resource.Unit != *other.Unit {
			return false
		}
	}

	return true
}

type ElasticsearchSerialDiffSettings struct {
	Lag *string `json:"lag,omitempty"`
}

func (resource ElasticsearchSerialDiffSettings) Equals(other ElasticsearchSerialDiffSettings) bool {
	if resource.Lag == nil && other.Lag != nil || resource.Lag != nil && other.Lag == nil {
		return false
	}

	if resource.Lag != nil {
		if *resource.Lag != *other.Lag {
			return false
		}
	}

	return true
}

type ElasticsearchCumulativeSumSettings struct {
	Format *string `json:"format,omitempty"`
}

func (resource ElasticsearchCumulativeSumSettings) Equals(other ElasticsearchCumulativeSumSettings) bool {
	if resource.Format == nil && other.Format != nil || resource.Format != nil && other.Format == nil {
		return false
	}

	if resource.Format != nil {
		if *resource.Format != *other.Format {
			return false
		}
	}

	return true
}

type ElasticsearchBucketScriptSettings struct {
	Script *InlineScript `json:"script,omitempty"`
}

func (resource ElasticsearchBucketScriptSettings) Equals(other ElasticsearchBucketScriptSettings) bool {
	if resource.Script == nil && other.Script != nil || resource.Script != nil && other.Script == nil {
		return false
	}

	if resource.Script != nil {
		if !resource.Script.Equals(*other.Script) {
			return false
		}
	}

	return true
}

type ElasticsearchTopMetricsSettings struct {
	Order   *string  `json:"order,omitempty"`
	OrderBy *string  `json:"orderBy,omitempty"`
	Metrics []string `json:"metrics,omitempty"`
}

func (resource ElasticsearchTopMetricsSettings) Equals(other ElasticsearchTopMetricsSettings) bool {
	if resource.Order == nil && other.Order != nil || resource.Order != nil && other.Order == nil {
		return false
	}

	if resource.Order != nil {
		if *resource.Order != *other.Order {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if *resource.OrderBy != *other.OrderBy {
			return false
		}
	}

	if len(resource.Metrics) != len(other.Metrics) {
		return false
	}

	for i1 := range resource.Metrics {
		if resource.Metrics[i1] != other.Metrics[i1] {
			return false
		}
	}

	return true
}

type DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested struct {
	DateHistogram *DateHistogram `json:"DateHistogram,omitempty"`
	Histogram     *Histogram     `json:"Histogram,omitempty"`
	Terms         *Terms         `json:"Terms,omitempty"`
	Filters       *Filters       `json:"Filters,omitempty"`
	GeoHashGrid   *GeoHashGrid   `json:"GeoHashGrid,omitempty"`
	Nested        *Nested        `json:"Nested,omitempty"`
}

func (resource DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) MarshalJSON() ([]byte, error) {
	if resource.DateHistogram != nil {
		return json.Marshal(resource.DateHistogram)
	}
	if resource.Histogram != nil {
		return json.Marshal(resource.Histogram)
	}
	if resource.Terms != nil {
		return json.Marshal(resource.Terms)
	}
	if resource.Filters != nil {
		return json.Marshal(resource.Filters)
	}
	if resource.GeoHashGrid != nil {
		return json.Marshal(resource.GeoHashGrid)
	}
	if resource.Nested != nil {
		return json.Marshal(resource.Nested)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "date_histogram":
		var dateHistogram DateHistogram
		if err := json.Unmarshal(raw, &dateHistogram); err != nil {
			return err
		}

		resource.DateHistogram = &dateHistogram
		return nil
	case "filters":
		var filters Filters
		if err := json.Unmarshal(raw, &filters); err != nil {
			return err
		}

		resource.Filters = &filters
		return nil
	case "geohash_grid":
		var geoHashGrid GeoHashGrid
		if err := json.Unmarshal(raw, &geoHashGrid); err != nil {
			return err
		}

		resource.GeoHashGrid = &geoHashGrid
		return nil
	case "histogram":
		var histogram Histogram
		if err := json.Unmarshal(raw, &histogram); err != nil {
			return err
		}

		resource.Histogram = &histogram
		return nil
	case "nested":
		var nested Nested
		if err := json.Unmarshal(raw, &nested); err != nil {
			return err
		}

		resource.Nested = &nested
		return nil
	case "terms":
		var terms Terms
		if err := json.Unmarshal(raw, &terms); err != nil {
			return err
		}

		resource.Terms = &terms
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) Equals(other DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) bool {
	if resource.DateHistogram == nil && other.DateHistogram != nil || resource.DateHistogram != nil && other.DateHistogram == nil {
		return false
	}

	if resource.DateHistogram != nil {
		if !resource.DateHistogram.Equals(*other.DateHistogram) {
			return false
		}
	}
	if resource.Histogram == nil && other.Histogram != nil || resource.Histogram != nil && other.Histogram == nil {
		return false
	}

	if resource.Histogram != nil {
		if !resource.Histogram.Equals(*other.Histogram) {
			return false
		}
	}
	if resource.Terms == nil && other.Terms != nil || resource.Terms != nil && other.Terms == nil {
		return false
	}

	if resource.Terms != nil {
		if !resource.Terms.Equals(*other.Terms) {
			return false
		}
	}
	if resource.Filters == nil && other.Filters != nil || resource.Filters != nil && other.Filters == nil {
		return false
	}

	if resource.Filters != nil {
		if !resource.Filters.Equals(*other.Filters) {
			return false
		}
	}
	if resource.GeoHashGrid == nil && other.GeoHashGrid != nil || resource.GeoHashGrid != nil && other.GeoHashGrid == nil {
		return false
	}

	if resource.GeoHashGrid != nil {
		if !resource.GeoHashGrid.Equals(*other.GeoHashGrid) {
			return false
		}
	}
	if resource.Nested == nil && other.Nested != nil || resource.Nested != nil && other.Nested == nil {
		return false
	}

	if resource.Nested != nil {
		if !resource.Nested.Equals(*other.Nested) {
			return false
		}
	}

	return true
}

type CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics struct {
	Count          *Count          `json:"Count,omitempty"`
	MovingAverage  *MovingAverage  `json:"MovingAverage,omitempty"`
	Derivative     *Derivative     `json:"Derivative,omitempty"`
	CumulativeSum  *CumulativeSum  `json:"CumulativeSum,omitempty"`
	BucketScript   *BucketScript   `json:"BucketScript,omitempty"`
	SerialDiff     *SerialDiff     `json:"SerialDiff,omitempty"`
	RawData        *RawData        `json:"RawData,omitempty"`
	RawDocument    *RawDocument    `json:"RawDocument,omitempty"`
	UniqueCount    *UniqueCount    `json:"UniqueCount,omitempty"`
	Percentiles    *Percentiles    `json:"Percentiles,omitempty"`
	ExtendedStats  *ExtendedStats  `json:"ExtendedStats,omitempty"`
	Min            *Min            `json:"Min,omitempty"`
	Max            *Max            `json:"Max,omitempty"`
	Sum            *Sum            `json:"Sum,omitempty"`
	Average        *Average        `json:"Average,omitempty"`
	MovingFunction *MovingFunction `json:"MovingFunction,omitempty"`
	Logs           *Logs           `json:"Logs,omitempty"`
	Rate           *Rate           `json:"Rate,omitempty"`
	TopMetrics     *TopMetrics     `json:"TopMetrics,omitempty"`
}

func (resource CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error) {
	if resource.Count != nil {
		return json.Marshal(resource.Count)
	}
	if resource.MovingAverage != nil {
		return json.Marshal(resource.MovingAverage)
	}
	if resource.Derivative != nil {
		return json.Marshal(resource.Derivative)
	}
	if resource.CumulativeSum != nil {
		return json.Marshal(resource.CumulativeSum)
	}
	if resource.BucketScript != nil {
		return json.Marshal(resource.BucketScript)
	}
	if resource.SerialDiff != nil {
		return json.Marshal(resource.SerialDiff)
	}
	if resource.RawData != nil {
		return json.Marshal(resource.RawData)
	}
	if resource.RawDocument != nil {
		return json.Marshal(resource.RawDocument)
	}
	if resource.UniqueCount != nil {
		return json.Marshal(resource.UniqueCount)
	}
	if resource.Percentiles != nil {
		return json.Marshal(resource.Percentiles)
	}
	if resource.ExtendedStats != nil {
		return json.Marshal(resource.ExtendedStats)
	}
	if resource.Min != nil {
		return json.Marshal(resource.Min)
	}
	if resource.Max != nil {
		return json.Marshal(resource.Max)
	}
	if resource.Sum != nil {
		return json.Marshal(resource.Sum)
	}
	if resource.Average != nil {
		return json.Marshal(resource.Average)
	}
	if resource.MovingFunction != nil {
		return json.Marshal(resource.MovingFunction)
	}
	if resource.Logs != nil {
		return json.Marshal(resource.Logs)
	}
	if resource.Rate != nil {
		return json.Marshal(resource.Rate)
	}
	if resource.TopMetrics != nil {
		return json.Marshal(resource.TopMetrics)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "avg":
		var average Average
		if err := json.Unmarshal(raw, &average); err != nil {
			return err
		}

		resource.Average = &average
		return nil
	case "bucket_script":
		var bucketScript BucketScript
		if err := json.Unmarshal(raw, &bucketScript); err != nil {
			return err
		}

		resource.BucketScript = &bucketScript
		return nil
	case "cardinality":
		var uniqueCount UniqueCount
		if err := json.Unmarshal(raw, &uniqueCount); err != nil {
			return err
		}

		resource.UniqueCount = &uniqueCount
		return nil
	case "count":
		var count Count
		if err := json.Unmarshal(raw, &count); err != nil {
			return err
		}

		resource.Count = &count
		return nil
	case "cumulative_sum":
		var cumulativeSum CumulativeSum
		if err := json.Unmarshal(raw, &cumulativeSum); err != nil {
			return err
		}

		resource.CumulativeSum = &cumulativeSum
		return nil
	case "derivative":
		var derivative Derivative
		if err := json.Unmarshal(raw, &derivative); err != nil {
			return err
		}

		resource.Derivative = &derivative
		return nil
	case "extended_stats":
		var extendedStats ExtendedStats
		if err := json.Unmarshal(raw, &extendedStats); err != nil {
			return err
		}

		resource.ExtendedStats = &extendedStats
		return nil
	case "logs":
		var logs Logs
		if err := json.Unmarshal(raw, &logs); err != nil {
			return err
		}

		resource.Logs = &logs
		return nil
	case "max":
		var max Max
		if err := json.Unmarshal(raw, &max); err != nil {
			return err
		}

		resource.Max = &max
		return nil
	case "min":
		var min Min
		if err := json.Unmarshal(raw, &min); err != nil {
			return err
		}

		resource.Min = &min
		return nil
	case "moving_avg":
		var movingAverage MovingAverage
		if err := json.Unmarshal(raw, &movingAverage); err != nil {
			return err
		}

		resource.MovingAverage = &movingAverage
		return nil
	case "moving_fn":
		var movingFunction MovingFunction
		if err := json.Unmarshal(raw, &movingFunction); err != nil {
			return err
		}

		resource.MovingFunction = &movingFunction
		return nil
	case "percentiles":
		var percentiles Percentiles
		if err := json.Unmarshal(raw, &percentiles); err != nil {
			return err
		}

		resource.Percentiles = &percentiles
		return nil
	case "rate":
		var rate Rate
		if err := json.Unmarshal(raw, &rate); err != nil {
			return err
		}

		resource.Rate = &rate
		return nil
	case "raw_data":
		var rawData RawData
		if err := json.Unmarshal(raw, &rawData); err != nil {
			return err
		}

		resource.RawData = &rawData
		return nil
	case "raw_document":
		var rawDocument RawDocument
		if err := json.Unmarshal(raw, &rawDocument); err != nil {
			return err
		}

		resource.RawDocument = &rawDocument
		return nil
	case "serial_diff":
		var serialDiff SerialDiff
		if err := json.Unmarshal(raw, &serialDiff); err != nil {
			return err
		}

		resource.SerialDiff = &serialDiff
		return nil
	case "sum":
		var sum Sum
		if err := json.Unmarshal(raw, &sum); err != nil {
			return err
		}

		resource.Sum = &sum
		return nil
	case "top_metrics":
		var topMetrics TopMetrics
		if err := json.Unmarshal(raw, &topMetrics); err != nil {
			return err
		}

		resource.TopMetrics = &topMetrics
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Equals(other CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) bool {
	if resource.Count == nil && other.Count != nil || resource.Count != nil && other.Count == nil {
		return false
	}

	if resource.Count != nil {
		if !resource.Count.Equals(*other.Count) {
			return false
		}
	}
	if resource.MovingAverage == nil && other.MovingAverage != nil || resource.MovingAverage != nil && other.MovingAverage == nil {
		return false
	}

	if resource.MovingAverage != nil {
		if !resource.MovingAverage.Equals(*other.MovingAverage) {
			return false
		}
	}
	if resource.Derivative == nil && other.Derivative != nil || resource.Derivative != nil && other.Derivative == nil {
		return false
	}

	if resource.Derivative != nil {
		if !resource.Derivative.Equals(*other.Derivative) {
			return false
		}
	}
	if resource.CumulativeSum == nil && other.CumulativeSum != nil || resource.CumulativeSum != nil && other.CumulativeSum == nil {
		return false
	}

	if resource.CumulativeSum != nil {
		if !resource.CumulativeSum.Equals(*other.CumulativeSum) {
			return false
		}
	}
	if resource.BucketScript == nil && other.BucketScript != nil || resource.BucketScript != nil && other.BucketScript == nil {
		return false
	}

	if resource.BucketScript != nil {
		if !resource.BucketScript.Equals(*other.BucketScript) {
			return false
		}
	}
	if resource.SerialDiff == nil && other.SerialDiff != nil || resource.SerialDiff != nil && other.SerialDiff == nil {
		return false
	}

	if resource.SerialDiff != nil {
		if !resource.SerialDiff.Equals(*other.SerialDiff) {
			return false
		}
	}
	if resource.RawData == nil && other.RawData != nil || resource.RawData != nil && other.RawData == nil {
		return false
	}

	if resource.RawData != nil {
		if !resource.RawData.Equals(*other.RawData) {
			return false
		}
	}
	if resource.RawDocument == nil && other.RawDocument != nil || resource.RawDocument != nil && other.RawDocument == nil {
		return false
	}

	if resource.RawDocument != nil {
		if !resource.RawDocument.Equals(*other.RawDocument) {
			return false
		}
	}
	if resource.UniqueCount == nil && other.UniqueCount != nil || resource.UniqueCount != nil && other.UniqueCount == nil {
		return false
	}

	if resource.UniqueCount != nil {
		if !resource.UniqueCount.Equals(*other.UniqueCount) {
			return false
		}
	}
	if resource.Percentiles == nil && other.Percentiles != nil || resource.Percentiles != nil && other.Percentiles == nil {
		return false
	}

	if resource.Percentiles != nil {
		if !resource.Percentiles.Equals(*other.Percentiles) {
			return false
		}
	}
	if resource.ExtendedStats == nil && other.ExtendedStats != nil || resource.ExtendedStats != nil && other.ExtendedStats == nil {
		return false
	}

	if resource.ExtendedStats != nil {
		if !resource.ExtendedStats.Equals(*other.ExtendedStats) {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if !resource.Min.Equals(*other.Min) {
			return false
		}
	}
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if !resource.Max.Equals(*other.Max) {
			return false
		}
	}
	if resource.Sum == nil && other.Sum != nil || resource.Sum != nil && other.Sum == nil {
		return false
	}

	if resource.Sum != nil {
		if !resource.Sum.Equals(*other.Sum) {
			return false
		}
	}
	if resource.Average == nil && other.Average != nil || resource.Average != nil && other.Average == nil {
		return false
	}

	if resource.Average != nil {
		if !resource.Average.Equals(*other.Average) {
			return false
		}
	}
	if resource.MovingFunction == nil && other.MovingFunction != nil || resource.MovingFunction != nil && other.MovingFunction == nil {
		return false
	}

	if resource.MovingFunction != nil {
		if !resource.MovingFunction.Equals(*other.MovingFunction) {
			return false
		}
	}
	if resource.Logs == nil && other.Logs != nil || resource.Logs != nil && other.Logs == nil {
		return false
	}

	if resource.Logs != nil {
		if !resource.Logs.Equals(*other.Logs) {
			return false
		}
	}
	if resource.Rate == nil && other.Rate != nil || resource.Rate != nil && other.Rate == nil {
		return false
	}

	if resource.Rate != nil {
		if !resource.Rate.Equals(*other.Rate) {
			return false
		}
	}
	if resource.TopMetrics == nil && other.TopMetrics != nil || resource.TopMetrics != nil && other.TopMetrics == nil {
		return false
	}

	if resource.TopMetrics != nil {
		if !resource.TopMetrics.Equals(*other.TopMetrics) {
			return false
		}
	}

	return true
}

type StringOrPipelineMetricAggregationType struct {
	String                        *string                        `json:"String,omitempty"`
	PipelineMetricAggregationType *PipelineMetricAggregationType `json:"PipelineMetricAggregationType,omitempty"`
}

func (resource StringOrPipelineMetricAggregationType) Equals(other StringOrPipelineMetricAggregationType) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.PipelineMetricAggregationType == nil && other.PipelineMetricAggregationType != nil || resource.PipelineMetricAggregationType != nil && other.PipelineMetricAggregationType == nil {
		return false
	}

	if resource.PipelineMetricAggregationType != nil {
		if *resource.PipelineMetricAggregationType != *other.PipelineMetricAggregationType {
			return false
		}
	}

	return true
}

type StringOrElasticsearchInlineScript struct {
	String                    *string                    `json:"String,omitempty"`
	ElasticsearchInlineScript *ElasticsearchInlineScript `json:"ElasticsearchInlineScript,omitempty"`
}

func (resource StringOrElasticsearchInlineScript) Equals(other StringOrElasticsearchInlineScript) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.ElasticsearchInlineScript == nil && other.ElasticsearchInlineScript != nil || resource.ElasticsearchInlineScript != nil && other.ElasticsearchInlineScript == nil {
		return false
	}

	if resource.ElasticsearchInlineScript != nil {
		if !resource.ElasticsearchInlineScript.Equals(*other.ElasticsearchInlineScript) {
			return false
		}
	}

	return true
}

type MovingAverageOrDerivativeOrCumulativeSumOrBucketScript struct {
	MovingAverage *MovingAverage `json:"MovingAverage,omitempty"`
	Derivative    *Derivative    `json:"Derivative,omitempty"`
	CumulativeSum *CumulativeSum `json:"CumulativeSum,omitempty"`
	BucketScript  *BucketScript  `json:"BucketScript,omitempty"`
}

func (resource MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) MarshalJSON() ([]byte, error) {
	if resource.MovingAverage != nil {
		return json.Marshal(resource.MovingAverage)
	}
	if resource.Derivative != nil {
		return json.Marshal(resource.Derivative)
	}
	if resource.CumulativeSum != nil {
		return json.Marshal(resource.CumulativeSum)
	}
	if resource.BucketScript != nil {
		return json.Marshal(resource.BucketScript)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "bucket_script":
		var bucketScript BucketScript
		if err := json.Unmarshal(raw, &bucketScript); err != nil {
			return err
		}

		resource.BucketScript = &bucketScript
		return nil
	case "cumulative_sum":
		var cumulativeSum CumulativeSum
		if err := json.Unmarshal(raw, &cumulativeSum); err != nil {
			return err
		}

		resource.CumulativeSum = &cumulativeSum
		return nil
	case "derivative":
		var derivative Derivative
		if err := json.Unmarshal(raw, &derivative); err != nil {
			return err
		}

		resource.Derivative = &derivative
		return nil
	case "moving_avg":
		var movingAverage MovingAverage
		if err := json.Unmarshal(raw, &movingAverage); err != nil {
			return err
		}

		resource.MovingAverage = &movingAverage
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) Equals(other MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) bool {
	if resource.MovingAverage == nil && other.MovingAverage != nil || resource.MovingAverage != nil && other.MovingAverage == nil {
		return false
	}

	if resource.MovingAverage != nil {
		if !resource.MovingAverage.Equals(*other.MovingAverage) {
			return false
		}
	}
	if resource.Derivative == nil && other.Derivative != nil || resource.Derivative != nil && other.Derivative == nil {
		return false
	}

	if resource.Derivative != nil {
		if !resource.Derivative.Equals(*other.Derivative) {
			return false
		}
	}
	if resource.CumulativeSum == nil && other.CumulativeSum != nil || resource.CumulativeSum != nil && other.CumulativeSum == nil {
		return false
	}

	if resource.CumulativeSum != nil {
		if !resource.CumulativeSum.Equals(*other.CumulativeSum) {
			return false
		}
	}
	if resource.BucketScript == nil && other.BucketScript != nil || resource.BucketScript != nil && other.BucketScript == nil {
		return false
	}

	if resource.BucketScript != nil {
		if !resource.BucketScript.Equals(*other.BucketScript) {
			return false
		}
	}

	return true
}

type BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics struct {
	BucketScript   *BucketScript   `json:"BucketScript,omitempty"`
	CumulativeSum  *CumulativeSum  `json:"CumulativeSum,omitempty"`
	Derivative     *Derivative     `json:"Derivative,omitempty"`
	SerialDiff     *SerialDiff     `json:"SerialDiff,omitempty"`
	RawData        *RawData        `json:"RawData,omitempty"`
	RawDocument    *RawDocument    `json:"RawDocument,omitempty"`
	UniqueCount    *UniqueCount    `json:"UniqueCount,omitempty"`
	Percentiles    *Percentiles    `json:"Percentiles,omitempty"`
	ExtendedStats  *ExtendedStats  `json:"ExtendedStats,omitempty"`
	Min            *Min            `json:"Min,omitempty"`
	Max            *Max            `json:"Max,omitempty"`
	Sum            *Sum            `json:"Sum,omitempty"`
	Average        *Average        `json:"Average,omitempty"`
	MovingAverage  *MovingAverage  `json:"MovingAverage,omitempty"`
	MovingFunction *MovingFunction `json:"MovingFunction,omitempty"`
	Logs           *Logs           `json:"Logs,omitempty"`
	Rate           *Rate           `json:"Rate,omitempty"`
	TopMetrics     *TopMetrics     `json:"TopMetrics,omitempty"`
}

func (resource BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error) {
	if resource.BucketScript != nil {
		return json.Marshal(resource.BucketScript)
	}
	if resource.CumulativeSum != nil {
		return json.Marshal(resource.CumulativeSum)
	}
	if resource.Derivative != nil {
		return json.Marshal(resource.Derivative)
	}
	if resource.SerialDiff != nil {
		return json.Marshal(resource.SerialDiff)
	}
	if resource.RawData != nil {
		return json.Marshal(resource.RawData)
	}
	if resource.RawDocument != nil {
		return json.Marshal(resource.RawDocument)
	}
	if resource.UniqueCount != nil {
		return json.Marshal(resource.UniqueCount)
	}
	if resource.Percentiles != nil {
		return json.Marshal(resource.Percentiles)
	}
	if resource.ExtendedStats != nil {
		return json.Marshal(resource.ExtendedStats)
	}
	if resource.Min != nil {
		return json.Marshal(resource.Min)
	}
	if resource.Max != nil {
		return json.Marshal(resource.Max)
	}
	if resource.Sum != nil {
		return json.Marshal(resource.Sum)
	}
	if resource.Average != nil {
		return json.Marshal(resource.Average)
	}
	if resource.MovingAverage != nil {
		return json.Marshal(resource.MovingAverage)
	}
	if resource.MovingFunction != nil {
		return json.Marshal(resource.MovingFunction)
	}
	if resource.Logs != nil {
		return json.Marshal(resource.Logs)
	}
	if resource.Rate != nil {
		return json.Marshal(resource.Rate)
	}
	if resource.TopMetrics != nil {
		return json.Marshal(resource.TopMetrics)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "avg":
		var average Average
		if err := json.Unmarshal(raw, &average); err != nil {
			return err
		}

		resource.Average = &average
		return nil
	case "bucket_script":
		var bucketScript BucketScript
		if err := json.Unmarshal(raw, &bucketScript); err != nil {
			return err
		}

		resource.BucketScript = &bucketScript
		return nil
	case "cardinality":
		var uniqueCount UniqueCount
		if err := json.Unmarshal(raw, &uniqueCount); err != nil {
			return err
		}

		resource.UniqueCount = &uniqueCount
		return nil
	case "cumulative_sum":
		var cumulativeSum CumulativeSum
		if err := json.Unmarshal(raw, &cumulativeSum); err != nil {
			return err
		}

		resource.CumulativeSum = &cumulativeSum
		return nil
	case "derivative":
		var derivative Derivative
		if err := json.Unmarshal(raw, &derivative); err != nil {
			return err
		}

		resource.Derivative = &derivative
		return nil
	case "extended_stats":
		var extendedStats ExtendedStats
		if err := json.Unmarshal(raw, &extendedStats); err != nil {
			return err
		}

		resource.ExtendedStats = &extendedStats
		return nil
	case "logs":
		var logs Logs
		if err := json.Unmarshal(raw, &logs); err != nil {
			return err
		}

		resource.Logs = &logs
		return nil
	case "max":
		var max Max
		if err := json.Unmarshal(raw, &max); err != nil {
			return err
		}

		resource.Max = &max
		return nil
	case "min":
		var min Min
		if err := json.Unmarshal(raw, &min); err != nil {
			return err
		}

		resource.Min = &min
		return nil
	case "moving_avg":
		var movingAverage MovingAverage
		if err := json.Unmarshal(raw, &movingAverage); err != nil {
			return err
		}

		resource.MovingAverage = &movingAverage
		return nil
	case "moving_fn":
		var movingFunction MovingFunction
		if err := json.Unmarshal(raw, &movingFunction); err != nil {
			return err
		}

		resource.MovingFunction = &movingFunction
		return nil
	case "percentiles":
		var percentiles Percentiles
		if err := json.Unmarshal(raw, &percentiles); err != nil {
			return err
		}

		resource.Percentiles = &percentiles
		return nil
	case "rate":
		var rate Rate
		if err := json.Unmarshal(raw, &rate); err != nil {
			return err
		}

		resource.Rate = &rate
		return nil
	case "raw_data":
		var rawData RawData
		if err := json.Unmarshal(raw, &rawData); err != nil {
			return err
		}

		resource.RawData = &rawData
		return nil
	case "raw_document":
		var rawDocument RawDocument
		if err := json.Unmarshal(raw, &rawDocument); err != nil {
			return err
		}

		resource.RawDocument = &rawDocument
		return nil
	case "serial_diff":
		var serialDiff SerialDiff
		if err := json.Unmarshal(raw, &serialDiff); err != nil {
			return err
		}

		resource.SerialDiff = &serialDiff
		return nil
	case "sum":
		var sum Sum
		if err := json.Unmarshal(raw, &sum); err != nil {
			return err
		}

		resource.Sum = &sum
		return nil
	case "top_metrics":
		var topMetrics TopMetrics
		if err := json.Unmarshal(raw, &topMetrics); err != nil {
			return err
		}

		resource.TopMetrics = &topMetrics
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Equals(other BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) bool {
	if resource.BucketScript == nil && other.BucketScript != nil || resource.BucketScript != nil && other.BucketScript == nil {
		return false
	}

	if resource.BucketScript != nil {
		if !resource.BucketScript.Equals(*other.BucketScript) {
			return false
		}
	}
	if resource.CumulativeSum == nil && other.CumulativeSum != nil || resource.CumulativeSum != nil && other.CumulativeSum == nil {
		return false
	}

	if resource.CumulativeSum != nil {
		if !resource.CumulativeSum.Equals(*other.CumulativeSum) {
			return false
		}
	}
	if resource.Derivative == nil && other.Derivative != nil || resource.Derivative != nil && other.Derivative == nil {
		return false
	}

	if resource.Derivative != nil {
		if !resource.Derivative.Equals(*other.Derivative) {
			return false
		}
	}
	if resource.SerialDiff == nil && other.SerialDiff != nil || resource.SerialDiff != nil && other.SerialDiff == nil {
		return false
	}

	if resource.SerialDiff != nil {
		if !resource.SerialDiff.Equals(*other.SerialDiff) {
			return false
		}
	}
	if resource.RawData == nil && other.RawData != nil || resource.RawData != nil && other.RawData == nil {
		return false
	}

	if resource.RawData != nil {
		if !resource.RawData.Equals(*other.RawData) {
			return false
		}
	}
	if resource.RawDocument == nil && other.RawDocument != nil || resource.RawDocument != nil && other.RawDocument == nil {
		return false
	}

	if resource.RawDocument != nil {
		if !resource.RawDocument.Equals(*other.RawDocument) {
			return false
		}
	}
	if resource.UniqueCount == nil && other.UniqueCount != nil || resource.UniqueCount != nil && other.UniqueCount == nil {
		return false
	}

	if resource.UniqueCount != nil {
		if !resource.UniqueCount.Equals(*other.UniqueCount) {
			return false
		}
	}
	if resource.Percentiles == nil && other.Percentiles != nil || resource.Percentiles != nil && other.Percentiles == nil {
		return false
	}

	if resource.Percentiles != nil {
		if !resource.Percentiles.Equals(*other.Percentiles) {
			return false
		}
	}
	if resource.ExtendedStats == nil && other.ExtendedStats != nil || resource.ExtendedStats != nil && other.ExtendedStats == nil {
		return false
	}

	if resource.ExtendedStats != nil {
		if !resource.ExtendedStats.Equals(*other.ExtendedStats) {
			return false
		}
	}
	if resource.Min == nil && other.Min != nil || resource.Min != nil && other.Min == nil {
		return false
	}

	if resource.Min != nil {
		if !resource.Min.Equals(*other.Min) {
			return false
		}
	}
	if resource.Max == nil && other.Max != nil || resource.Max != nil && other.Max == nil {
		return false
	}

	if resource.Max != nil {
		if !resource.Max.Equals(*other.Max) {
			return false
		}
	}
	if resource.Sum == nil && other.Sum != nil || resource.Sum != nil && other.Sum == nil {
		return false
	}

	if resource.Sum != nil {
		if !resource.Sum.Equals(*other.Sum) {
			return false
		}
	}
	if resource.Average == nil && other.Average != nil || resource.Average != nil && other.Average == nil {
		return false
	}

	if resource.Average != nil {
		if !resource.Average.Equals(*other.Average) {
			return false
		}
	}
	if resource.MovingAverage == nil && other.MovingAverage != nil || resource.MovingAverage != nil && other.MovingAverage == nil {
		return false
	}

	if resource.MovingAverage != nil {
		if !resource.MovingAverage.Equals(*other.MovingAverage) {
			return false
		}
	}
	if resource.MovingFunction == nil && other.MovingFunction != nil || resource.MovingFunction != nil && other.MovingFunction == nil {
		return false
	}

	if resource.MovingFunction != nil {
		if !resource.MovingFunction.Equals(*other.MovingFunction) {
			return false
		}
	}
	if resource.Logs == nil && other.Logs != nil || resource.Logs != nil && other.Logs == nil {
		return false
	}

	if resource.Logs != nil {
		if !resource.Logs.Equals(*other.Logs) {
			return false
		}
	}
	if resource.Rate == nil && other.Rate != nil || resource.Rate != nil && other.Rate == nil {
		return false
	}

	if resource.Rate != nil {
		if !resource.Rate.Equals(*other.Rate) {
			return false
		}
	}
	if resource.TopMetrics == nil && other.TopMetrics != nil || resource.TopMetrics != nil && other.TopMetrics == nil {
		return false
	}

	if resource.TopMetrics != nil {
		if !resource.TopMetrics.Equals(*other.TopMetrics) {
			return false
		}
	}

	return true
}
