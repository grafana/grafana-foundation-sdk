---
title: <span class="badge object-type-struct"></span> CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics
---
# <span class="badge object-type-struct"></span> CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics

## Definition

```go
type CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics struct {
    Count *elasticsearch.Count `json:"Count,omitempty"`
    MovingAverage *elasticsearch.MovingAverage `json:"MovingAverage,omitempty"`
    Derivative *elasticsearch.Derivative `json:"Derivative,omitempty"`
    CumulativeSum *elasticsearch.CumulativeSum `json:"CumulativeSum,omitempty"`
    BucketScript *elasticsearch.BucketScript `json:"BucketScript,omitempty"`
    SerialDiff *elasticsearch.SerialDiff `json:"SerialDiff,omitempty"`
    RawData *elasticsearch.RawData `json:"RawData,omitempty"`
    RawDocument *elasticsearch.RawDocument `json:"RawDocument,omitempty"`
    UniqueCount *elasticsearch.UniqueCount `json:"UniqueCount,omitempty"`
    Percentiles *elasticsearch.Percentiles `json:"Percentiles,omitempty"`
    ExtendedStats *elasticsearch.ExtendedStats `json:"ExtendedStats,omitempty"`
    Min *elasticsearch.Min `json:"Min,omitempty"`
    Max *elasticsearch.Max `json:"Max,omitempty"`
    Sum *elasticsearch.Sum `json:"Sum,omitempty"`
    Average *elasticsearch.Average `json:"Average,omitempty"`
    MovingFunction *elasticsearch.MovingFunction `json:"MovingFunction,omitempty"`
    Logs *elasticsearch.Logs `json:"Logs,omitempty"`
    Rate *elasticsearch.Rate `json:"Rate,omitempty"`
    TopMetrics *elasticsearch.TopMetrics `json:"TopMetrics,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> MarshalJSON

MarshalJSON implements a custom JSON marshalling logic to encode `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` as JSON.

```go
func (countOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error)
```

### <span class="badge object-method"></span> UnmarshalJSON

UnmarshalJSON implements a custom JSON unmarshalling logic to decode `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` from JSON.

```go
func (countOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSON(raw []byte) error
```

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (countOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` objects.

```go
func (countOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Equals(other CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics` fields for violations and returns them.

```go
func (countOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) Validate() error
```

