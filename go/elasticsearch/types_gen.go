// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

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
	PipelineAgg *string `json:"pipelineAgg,omitempty"`
	Field       *string `json:"field,omitempty"`
	Type        string  `json:"type"`
	Id          string  `json:"id"`
	Settings    any     `json:"settings,omitempty"`
	Hide        *bool   `json:"hide,omitempty"`
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
	Alias      *string             `json:"alias,omitempty"`
	Query      *string             `json:"query,omitempty"`
	TimeField  *string             `json:"timeField,omitempty"`
	BucketAggs []BucketAggregation `json:"bucketAggs,omitempty"`
	Metrics    []MetricAggregation `json:"metrics,omitempty"`
	RefId      *string             `json:"refId,omitempty"`
	Hide       *bool               `json:"hide,omitempty"`
	QueryType  *string             `json:"queryType,omitempty"`
	Datasource any                 `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

type DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested struct {
	DateHistogram *DateHistogram `json:"DateHistogram,omitempty"`
	Histogram     *Histogram     `json:"Histogram,omitempty"`
	Terms         *Terms         `json:"Terms,omitempty"`
	Filters       *Filters       `json:"Filters,omitempty"`
	GeoHashGrid   *GeoHashGrid   `json:"GeoHashGrid,omitempty"`
	Nested        *Nested        `json:"Nested,omitempty"`
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

type StringOrPipelineMetricAggregationType struct {
	String                        *string                        `json:"String,omitempty"`
	PipelineMetricAggregationType *PipelineMetricAggregationType `json:"PipelineMetricAggregationType,omitempty"`
}

type StringOrStruct struct {
	String *string `json:"String,omitempty"`
	Struct *struct {
		Inline string `json:"inline,omitempty"`
	} `json:"Struct,omitempty"`
}

type MovingAverageOrDerivativeOrCumulativeSumOrBucketScript struct {
	MovingAverage *MovingAverage `json:"MovingAverage,omitempty"`
	Derivative    *Derivative    `json:"Derivative,omitempty"`
	CumulativeSum *CumulativeSum `json:"CumulativeSum,omitempty"`
	BucketScript  *BucketScript  `json:"BucketScript,omitempty"`
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
