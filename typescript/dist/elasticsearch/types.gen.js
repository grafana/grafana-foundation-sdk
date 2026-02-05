"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.defaultMetricAggregationWithInlineScript = exports.defaultMetricAggregationWithMissingSupport = exports.defaultMetricAggregationWithField = exports.defaultBaseMetricAggregation = exports.defaultGeoHashGridSettings = exports.defaultFiltersSettings = exports.defaultTermsSettings = exports.defaultHistogramSettings = exports.defaultDateHistogramSettings = exports.defaultBucketAggregationWithField = exports.defaultBaseBucketAggregation = exports.defaultTopMetrics = exports.defaultRate = exports.defaultLogs = exports.defaultMovingFunction = exports.defaultAverage = exports.defaultSum = exports.defaultMax = exports.defaultMin = exports.defaultExtendedStats = exports.defaultPercentiles = exports.defaultUniqueCount = exports.defaultRawDocument = exports.defaultRawData = exports.defaultSerialDiff = exports.defaultMetricAggregationWithSettings = exports.defaultInlineScript = exports.defaultPipelineVariable = exports.defaultBucketScript = exports.defaultCumulativeSum = exports.defaultDerivative = exports.defaultMovingAverage = exports.defaultPipelineMetricAggregation = exports.defaultPipelineMetricAggregationType = exports.PipelineMetricAggregationType = exports.defaultMetricAggregationType = exports.defaultCount = exports.defaultMetricAggregation = exports.defaultNested = exports.defaultGeoHashGrid = exports.defaultFilter = exports.defaultFilters = exports.defaultTermsOrder = exports.TermsOrder = exports.defaultTerms = exports.defaultHistogram = exports.defaultBucketAggregationType = exports.BucketAggregationType = exports.defaultDateHistogram = exports.defaultBucketAggregation = void 0;
exports.defaultDataquery = exports.defaultMovingAverageHoltWintersModelSettings = exports.defaultMovingAverageHoltModelSettings = exports.defaultMovingAverageEWMAModelSettings = exports.defaultMovingAverageLinearModelSettings = exports.defaultMovingAverageSimpleModelSettings = exports.defaultBaseMovingAverageModelSettings = exports.defaultMovingAverageModelOption = exports.defaultMovingAverageModel = exports.MovingAverageModel = exports.defaultPipelineMetricAggregationWithMultipleBucketPaths = exports.defaultBasePipelineMetricAggregation = exports.defaultExtendedStat = exports.defaultExtendedStatMetaType = exports.ExtendedStatMetaType = void 0;
const defaultBucketAggregation = () => ((0, exports.defaultDateHistogram)());
exports.defaultBucketAggregation = defaultBucketAggregation;
const defaultDateHistogram = () => ({
    id: "",
    type: BucketAggregationType.DateHistogram,
});
exports.defaultDateHistogram = defaultDateHistogram;
var BucketAggregationType;
(function (BucketAggregationType) {
    BucketAggregationType["Terms"] = "terms";
    BucketAggregationType["Filters"] = "filters";
    BucketAggregationType["GeohashGrid"] = "geohash_grid";
    BucketAggregationType["DateHistogram"] = "date_histogram";
    BucketAggregationType["Histogram"] = "histogram";
    BucketAggregationType["Nested"] = "nested";
})(BucketAggregationType || (exports.BucketAggregationType = BucketAggregationType = {}));
const defaultBucketAggregationType = () => (BucketAggregationType.Terms);
exports.defaultBucketAggregationType = defaultBucketAggregationType;
const defaultHistogram = () => ({
    id: "",
    type: BucketAggregationType.Histogram,
});
exports.defaultHistogram = defaultHistogram;
const defaultTerms = () => ({
    id: "",
    type: BucketAggregationType.Terms,
});
exports.defaultTerms = defaultTerms;
var TermsOrder;
(function (TermsOrder) {
    TermsOrder["Desc"] = "desc";
    TermsOrder["Asc"] = "asc";
})(TermsOrder || (exports.TermsOrder = TermsOrder = {}));
const defaultTermsOrder = () => (TermsOrder.Desc);
exports.defaultTermsOrder = defaultTermsOrder;
const defaultFilters = () => ({
    id: "",
    type: BucketAggregationType.Filters,
});
exports.defaultFilters = defaultFilters;
const defaultFilter = () => ({
    query: "",
    label: "",
});
exports.defaultFilter = defaultFilter;
const defaultGeoHashGrid = () => ({
    id: "",
    type: BucketAggregationType.GeohashGrid,
});
exports.defaultGeoHashGrid = defaultGeoHashGrid;
const defaultNested = () => ({
    id: "",
    type: BucketAggregationType.Nested,
});
exports.defaultNested = defaultNested;
const defaultMetricAggregation = () => ((0, exports.defaultCount)());
exports.defaultMetricAggregation = defaultMetricAggregation;
const defaultCount = () => ({
    type: "unknown",
    id: "",
});
exports.defaultCount = defaultCount;
const defaultMetricAggregationType = () => ("count");
exports.defaultMetricAggregationType = defaultMetricAggregationType;
var PipelineMetricAggregationType;
(function (PipelineMetricAggregationType) {
    PipelineMetricAggregationType["MovingAvg"] = "moving_avg";
    PipelineMetricAggregationType["MovingFn"] = "moving_fn";
    PipelineMetricAggregationType["Derivative"] = "derivative";
    PipelineMetricAggregationType["SerialDiff"] = "serial_diff";
    PipelineMetricAggregationType["CumulativeSum"] = "cumulative_sum";
    PipelineMetricAggregationType["BucketScript"] = "bucket_script";
})(PipelineMetricAggregationType || (exports.PipelineMetricAggregationType = PipelineMetricAggregationType = {}));
const defaultPipelineMetricAggregationType = () => (PipelineMetricAggregationType.MovingAvg);
exports.defaultPipelineMetricAggregationType = defaultPipelineMetricAggregationType;
const defaultPipelineMetricAggregation = () => ((0, exports.defaultMovingAverage)());
exports.defaultPipelineMetricAggregation = defaultPipelineMetricAggregation;
const defaultMovingAverage = () => ({
    type: "unknown",
    id: "",
});
exports.defaultMovingAverage = defaultMovingAverage;
const defaultDerivative = () => ({
    type: "unknown",
    id: "",
});
exports.defaultDerivative = defaultDerivative;
const defaultCumulativeSum = () => ({
    type: "unknown",
    id: "",
});
exports.defaultCumulativeSum = defaultCumulativeSum;
const defaultBucketScript = () => ({
    type: "unknown",
    id: "",
});
exports.defaultBucketScript = defaultBucketScript;
const defaultPipelineVariable = () => ({
    name: "",
    pipelineAgg: "",
});
exports.defaultPipelineVariable = defaultPipelineVariable;
const defaultInlineScript = () => ("");
exports.defaultInlineScript = defaultInlineScript;
const defaultMetricAggregationWithSettings = () => ((0, exports.defaultBucketScript)());
exports.defaultMetricAggregationWithSettings = defaultMetricAggregationWithSettings;
const defaultSerialDiff = () => ({
    type: "unknown",
    id: "",
});
exports.defaultSerialDiff = defaultSerialDiff;
const defaultRawData = () => ({
    type: "unknown",
    id: "",
});
exports.defaultRawData = defaultRawData;
const defaultRawDocument = () => ({
    type: "unknown",
    id: "",
});
exports.defaultRawDocument = defaultRawDocument;
const defaultUniqueCount = () => ({
    type: "unknown",
    id: "",
});
exports.defaultUniqueCount = defaultUniqueCount;
const defaultPercentiles = () => ({
    type: "unknown",
    id: "",
});
exports.defaultPercentiles = defaultPercentiles;
const defaultExtendedStats = () => ({
    type: "unknown",
    id: "",
});
exports.defaultExtendedStats = defaultExtendedStats;
const defaultMin = () => ({
    type: "unknown",
    id: "",
});
exports.defaultMin = defaultMin;
const defaultMax = () => ({
    type: "unknown",
    id: "",
});
exports.defaultMax = defaultMax;
const defaultSum = () => ({
    type: "unknown",
    id: "",
});
exports.defaultSum = defaultSum;
const defaultAverage = () => ({
    type: "unknown",
    id: "",
});
exports.defaultAverage = defaultAverage;
const defaultMovingFunction = () => ({
    type: "unknown",
    id: "",
});
exports.defaultMovingFunction = defaultMovingFunction;
const defaultLogs = () => ({
    type: "unknown",
    id: "",
});
exports.defaultLogs = defaultLogs;
const defaultRate = () => ({
    type: "unknown",
    id: "",
});
exports.defaultRate = defaultRate;
const defaultTopMetrics = () => ({
    type: "unknown",
    id: "",
});
exports.defaultTopMetrics = defaultTopMetrics;
const defaultBaseBucketAggregation = () => ({
    id: "",
    type: BucketAggregationType.Terms,
});
exports.defaultBaseBucketAggregation = defaultBaseBucketAggregation;
const defaultBucketAggregationWithField = () => ({
    id: "",
    type: BucketAggregationType.Terms,
});
exports.defaultBucketAggregationWithField = defaultBucketAggregationWithField;
const defaultDateHistogramSettings = () => ({});
exports.defaultDateHistogramSettings = defaultDateHistogramSettings;
const defaultHistogramSettings = () => ({});
exports.defaultHistogramSettings = defaultHistogramSettings;
const defaultTermsSettings = () => ({});
exports.defaultTermsSettings = defaultTermsSettings;
const defaultFiltersSettings = () => ({});
exports.defaultFiltersSettings = defaultFiltersSettings;
const defaultGeoHashGridSettings = () => ({});
exports.defaultGeoHashGridSettings = defaultGeoHashGridSettings;
const defaultBaseMetricAggregation = () => ({
    type: (0, exports.defaultMetricAggregationType)(),
    id: "",
});
exports.defaultBaseMetricAggregation = defaultBaseMetricAggregation;
const defaultMetricAggregationWithField = () => ({
    type: (0, exports.defaultMetricAggregationType)(),
    id: "",
});
exports.defaultMetricAggregationWithField = defaultMetricAggregationWithField;
const defaultMetricAggregationWithMissingSupport = () => ({
    type: (0, exports.defaultMetricAggregationType)(),
    id: "",
});
exports.defaultMetricAggregationWithMissingSupport = defaultMetricAggregationWithMissingSupport;
const defaultMetricAggregationWithInlineScript = () => ({
    type: (0, exports.defaultMetricAggregationType)(),
    id: "",
});
exports.defaultMetricAggregationWithInlineScript = defaultMetricAggregationWithInlineScript;
var ExtendedStatMetaType;
(function (ExtendedStatMetaType) {
    ExtendedStatMetaType["Avg"] = "avg";
    ExtendedStatMetaType["Min"] = "min";
    ExtendedStatMetaType["Max"] = "max";
    ExtendedStatMetaType["Sum"] = "sum";
    ExtendedStatMetaType["Count"] = "count";
    ExtendedStatMetaType["StdDeviation"] = "std_deviation";
    ExtendedStatMetaType["StdDeviationBoundsUpper"] = "std_deviation_bounds_upper";
    ExtendedStatMetaType["StdDeviationBoundsLower"] = "std_deviation_bounds_lower";
})(ExtendedStatMetaType || (exports.ExtendedStatMetaType = ExtendedStatMetaType = {}));
const defaultExtendedStatMetaType = () => (ExtendedStatMetaType.Avg);
exports.defaultExtendedStatMetaType = defaultExtendedStatMetaType;
const defaultExtendedStat = () => ({
    label: "",
    value: ExtendedStatMetaType.Avg,
});
exports.defaultExtendedStat = defaultExtendedStat;
const defaultBasePipelineMetricAggregation = () => ({
    type: "",
    id: "",
});
exports.defaultBasePipelineMetricAggregation = defaultBasePipelineMetricAggregation;
const defaultPipelineMetricAggregationWithMultipleBucketPaths = () => ({
    type: (0, exports.defaultMetricAggregationType)(),
    id: "",
});
exports.defaultPipelineMetricAggregationWithMultipleBucketPaths = defaultPipelineMetricAggregationWithMultipleBucketPaths;
var MovingAverageModel;
(function (MovingAverageModel) {
    MovingAverageModel["Simple"] = "simple";
    MovingAverageModel["Linear"] = "linear";
    MovingAverageModel["Ewma"] = "ewma";
    MovingAverageModel["Holt"] = "holt";
    MovingAverageModel["HoltWinters"] = "holt_winters";
})(MovingAverageModel || (exports.MovingAverageModel = MovingAverageModel = {}));
const defaultMovingAverageModel = () => (MovingAverageModel.Simple);
exports.defaultMovingAverageModel = defaultMovingAverageModel;
const defaultMovingAverageModelOption = () => ({
    label: "",
    value: MovingAverageModel.Simple,
});
exports.defaultMovingAverageModelOption = defaultMovingAverageModelOption;
const defaultBaseMovingAverageModelSettings = () => ({
    model: MovingAverageModel.Simple,
    window: "",
    predict: "",
});
exports.defaultBaseMovingAverageModelSettings = defaultBaseMovingAverageModelSettings;
const defaultMovingAverageSimpleModelSettings = () => ({
    model: MovingAverageModel.Simple,
    window: "",
    predict: "",
});
exports.defaultMovingAverageSimpleModelSettings = defaultMovingAverageSimpleModelSettings;
const defaultMovingAverageLinearModelSettings = () => ({
    model: MovingAverageModel.Linear,
    window: "",
    predict: "",
});
exports.defaultMovingAverageLinearModelSettings = defaultMovingAverageLinearModelSettings;
const defaultMovingAverageEWMAModelSettings = () => ({
    model: MovingAverageModel.Ewma,
    window: "",
    minimize: false,
    predict: "",
});
exports.defaultMovingAverageEWMAModelSettings = defaultMovingAverageEWMAModelSettings;
const defaultMovingAverageHoltModelSettings = () => ({
    model: MovingAverageModel.Holt,
    settings: {},
    window: "",
    minimize: false,
    predict: "",
});
exports.defaultMovingAverageHoltModelSettings = defaultMovingAverageHoltModelSettings;
const defaultMovingAverageHoltWintersModelSettings = () => ({
    model: MovingAverageModel.HoltWinters,
    settings: {},
    window: "",
    minimize: false,
    predict: "",
});
exports.defaultMovingAverageHoltWintersModelSettings = defaultMovingAverageHoltWintersModelSettings;
const defaultDataquery = () => ({
    _implementsDataqueryVariant: () => { },
});
exports.defaultDataquery = defaultDataquery;
//# sourceMappingURL=types.gen.js.map