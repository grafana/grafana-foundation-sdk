// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

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

func (resource *BucketScriptOrCumulativeSumOrDerivativeOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error) {
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

	return nil, nil
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

func (resource *CountOrMovingAverageOrDerivativeOrCumulativeSumOrBucketScriptOrSerialDiffOrRawDataOrRawDocumentOrUniqueCountOrPercentilesOrExtendedStatsOrMinOrMaxOrSumOrAverageOrMovingFunctionOrLogsOrRateOrTopMetrics) MarshalJSON() ([]byte, error) {
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

	return nil, nil
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

func (resource *DateHistogramOrHistogramOrTermsOrFiltersOrGeoHashGridOrNested) MarshalJSON() ([]byte, error) {
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

	return nil, nil
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

func (resource *MovingAverageOrDerivativeOrCumulativeSumOrBucketScript) MarshalJSON() ([]byte, error) {
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

	return nil, nil
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
