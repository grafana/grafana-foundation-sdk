// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
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

type BucketAggregationWithField struct {
	Field    *string               `json:"field,omitempty"`
	Id       string                `json:"id"`
	Type     BucketAggregationType `json:"type"`
	Settings any                   `json:"settings,omitempty"`
}

type DateHistogram struct {
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Settings *struct {
		Interval    *string `json:"interval,omitempty"`
		MinDocCount *string `json:"min_doc_count,omitempty"`
		TrimEdges   *string `json:"trimEdges,omitempty"`
		Offset      *string `json:"offset,omitempty"`
		TimeZone    *string `json:"timeZone,omitempty"`
	} `json:"settings,omitempty"`
}

type DateHistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
	TrimEdges   *string `json:"trimEdges,omitempty"`
	Offset      *string `json:"offset,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}

type Histogram struct {
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Settings *struct {
		Interval    *string `json:"interval,omitempty"`
		MinDocCount *string `json:"min_doc_count,omitempty"`
	} `json:"settings,omitempty"`
}

type HistogramSettings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
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

type Terms struct {
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Settings *struct {
		Order       *TermsOrder `json:"order,omitempty"`
		Size        *string     `json:"size,omitempty"`
		MinDocCount *string     `json:"min_doc_count,omitempty"`
		OrderBy     *string     `json:"orderBy,omitempty"`
		Missing     *string     `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
}

type TermsSettings struct {
	Order       *TermsOrder `json:"order,omitempty"`
	Size        *string     `json:"size,omitempty"`
	MinDocCount *string     `json:"min_doc_count,omitempty"`
	OrderBy     *string     `json:"orderBy,omitempty"`
	Missing     *string     `json:"missing,omitempty"`
}

type Filters struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Settings *struct {
		Filters []Filter `json:"filters,omitempty"`
	} `json:"settings,omitempty"`
}

type Filter struct {
	Query string `json:"query"`
	Label string `json:"label"`
}

type FiltersSettings struct {
	Filters []Filter `json:"filters,omitempty"`
}

type GeoHashGrid struct {
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Type     string  `json:"type"`
	Settings *struct {
		Precision *string `json:"precision,omitempty"`
	} `json:"settings,omitempty"`
}

type GeoHashGridSettings struct {
	Precision *string `json:"precision,omitempty"`
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

type PipelineVariable struct {
	Name        string `json:"name"`
	PipelineAgg string `json:"pipelineAgg"`
}

type MetricAggregationWithField struct {
	Field *string               `json:"field,omitempty"`
	Type  MetricAggregationType `json:"type"`
	Id    string                `json:"id"`
	Hide  *bool                 `json:"hide,omitempty"`
}

type MetricAggregationWithMissingSupport struct {
	Settings *struct {
		Missing *string `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
	Type MetricAggregationType `json:"type"`
	Id   string                `json:"id"`
	Hide *bool                 `json:"hide,omitempty"`
}

type InlineScript = StringOrStruct

type MetricAggregationWithInlineScript struct {
	Settings *struct {
		Script *InlineScript `json:"script,omitempty"`
	} `json:"settings,omitempty"`
	Type MetricAggregationType `json:"type"`
	Id   string                `json:"id"`
	Hide *bool                 `json:"hide,omitempty"`
}

type Count struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Hide *bool  `json:"hide,omitempty"`
}

type Average struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		Script  *InlineScript `json:"script,omitempty"`
		Missing *string       `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type Sum struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		Script  *InlineScript `json:"script,omitempty"`
		Missing *string       `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type Max struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		Script  *InlineScript `json:"script,omitempty"`
		Missing *string       `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type Min struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		Script  *InlineScript `json:"script,omitempty"`
		Missing *string       `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
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

type ExtendedStats struct {
	Type     string `json:"type"`
	Settings *struct {
		Script  *InlineScript `json:"script,omitempty"`
		Missing *string       `json:"missing,omitempty"`
		Sigma   *string       `json:"sigma,omitempty"`
	} `json:"settings,omitempty"`
	Field *string `json:"field,omitempty"`
	Id    string  `json:"id"`
	Meta  any     `json:"meta,omitempty"`
	Hide  *bool   `json:"hide,omitempty"`
}

type Percentiles struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		Script   *InlineScript `json:"script,omitempty"`
		Missing  *string       `json:"missing,omitempty"`
		Percents []string      `json:"percents,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type UniqueCount struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		PrecisionThreshold *string `json:"precision_threshold,omitempty"`
		Missing            *string `json:"missing,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type RawDocument struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	Settings *struct {
		Size *string `json:"size,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type RawData struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	Settings *struct {
		Size *string `json:"size,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type Logs struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	Settings *struct {
		Limit *string `json:"limit,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type Rate struct {
	Type     string  `json:"type"`
	Field    *string `json:"field,omitempty"`
	Id       string  `json:"id"`
	Settings *struct {
		Unit *string `json:"unit,omitempty"`
		Mode *string `json:"mode,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type BasePipelineMetricAggregation struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Hide        *bool   `json:"hide,omitempty"`
}

type PipelineMetricAggregationWithMultipleBucketPaths struct {
	PipelineVariables []PipelineVariable    `json:"pipelineVariables,omitempty"`
	Type              MetricAggregationType `json:"type"`
	Id                string                `json:"id"`
	Hide              *bool                 `json:"hide,omitempty"`
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

type BaseMovingAverageModelSettings struct {
	Model   MovingAverageModel `json:"model"`
	Window  string             `json:"window"`
	Predict string             `json:"predict"`
}

type MovingAverageSimpleModelSettings struct {
	Model   string `json:"model"`
	Window  string `json:"window"`
	Predict string `json:"predict"`
}

type MovingAverageLinearModelSettings struct {
	Model   string `json:"model"`
	Window  string `json:"window"`
	Predict string `json:"predict"`
}

type MovingAverageEWMAModelSettings struct {
	Model    string `json:"model"`
	Settings *struct {
		Alpha *string `json:"alpha,omitempty"`
	} `json:"settings,omitempty"`
	Window   string `json:"window"`
	Minimize bool   `json:"minimize"`
	Predict  string `json:"predict"`
}

type MovingAverageHoltModelSettings struct {
	Model    string `json:"model"`
	Settings struct {
		Alpha *string `json:"alpha,omitempty"`
		Beta  *string `json:"beta,omitempty"`
	} `json:"settings"`
	Window   string `json:"window"`
	Minimize bool   `json:"minimize"`
	Predict  string `json:"predict"`
}

type MovingAverageHoltWintersModelSettings struct {
	Model    string `json:"model"`
	Settings struct {
		Alpha  *string `json:"alpha,omitempty"`
		Beta   *string `json:"beta,omitempty"`
		Gamma  *string `json:"gamma,omitempty"`
		Period *string `json:"period,omitempty"`
		Pad    *bool   `json:"pad,omitempty"`
	} `json:"settings"`
	Window   string `json:"window"`
	Minimize bool   `json:"minimize"`
	Predict  string `json:"predict"`
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

type MovingFunction struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Settings    *struct {
		Window *string       `json:"window,omitempty"`
		Script *InlineScript `json:"script,omitempty"`
		Shift  *string       `json:"shift,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type Derivative struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Settings    *struct {
		Unit *string `json:"unit,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type SerialDiff struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Settings    *struct {
		Lag *string `json:"lag,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type CumulativeSum struct {
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Settings    *struct {
		Format *string `json:"format,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type BucketScript struct {
	Type              string             `json:"type"`
	PipelineVariables []PipelineVariable `json:"pipelineVariables,omitempty"`
	Id                string             `json:"id"`
	Settings          *struct {
		Script *InlineScript `json:"script,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
}

type TopMetrics struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	Settings *struct {
		Order   *string  `json:"order,omitempty"`
		OrderBy *string  `json:"orderBy,omitempty"`
		Metrics []string `json:"metrics,omitempty"`
	} `json:"settings,omitempty"`
	Hide *bool `json:"hide,omitempty"`
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
	Datasource any `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "elasticsearch",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := Dataquery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
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

type StringOrPipelineMetricAggregationType struct {
	String                        *string                        `json:"String,omitempty"`
	PipelineMetricAggregationType *PipelineMetricAggregationType `json:"PipelineMetricAggregationType,omitempty"`
}

type StringOrStruct struct {
	String *string `json:"String,omitempty"`
	Struct *struct {
		Inline *string `json:"inline,omitempty"`
	} `json:"Struct,omitempty"`
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
