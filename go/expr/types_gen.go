// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"encoding/json"
	"errors"
	"fmt"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type Expr = TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql

func (resource Expr) ImplementsDataqueryVariant() {}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "__expr__",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := Expr{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

type TypeMath struct {
	// The datasource
	Datasource *struct {
		// The apiserver version
		ApiVersion *string `json:"apiVersion,omitempty"`
		// The datasource plugin type
		Type string `json:"type"`
		// Datasource UID (NOTE: name in k8s)
		Uid *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	// General math expression
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *struct {
		// Maximum frame count
		MaxFrames *int64 `json:"maxFrames,omitempty"`
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""`
		//  - `"timeseries-wide"`
		//  - `"timeseries-long"`
		//  - `"timeseries-many"`
		//  - `"timeseries-multi"`
		//  - `"directory-listing"`
		//  - `"table"`
		//  - `"numeric-wide"`
		//  - `"numeric-multi"`
		//  - `"numeric-long"`
		//  - `"log-lines"`
		Type *TypeMathType `json:"type,omitempty"`
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		TypeVersion []int64 `json:"typeVersion"`
	} `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *struct {
		// From is the start time of the query.
		From string `json:"from"`
		// To is the end time of the query.
		To string `json:"to"`
	} `json:"timeRange,omitempty"`
	Type string `json:"type"`
}

func (resource TypeMath) ImplementsDataqueryVariant() {}

type TypeReduce struct {
	// The datasource
	Datasource *struct {
		// The apiserver version
		ApiVersion *string `json:"apiVersion,omitempty"`
		// The datasource plugin type
		Type string `json:"type"`
		// Datasource UID (NOTE: name in k8s)
		Uid *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	// Reference to single query result
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// The reducer
	// Possible enum values:
	//  - `"sum"`
	//  - `"mean"`
	//  - `"min"`
	//  - `"max"`
	//  - `"count"`
	//  - `"last"`
	//  - `"median"`
	Reducer TypeReduceReducer `json:"reducer"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *struct {
		// Maximum frame count
		MaxFrames *int64 `json:"maxFrames,omitempty"`
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""`
		//  - `"timeseries-wide"`
		//  - `"timeseries-long"`
		//  - `"timeseries-many"`
		//  - `"timeseries-multi"`
		//  - `"directory-listing"`
		//  - `"table"`
		//  - `"numeric-wide"`
		//  - `"numeric-multi"`
		//  - `"numeric-long"`
		//  - `"log-lines"`
		Type *TypeReduceType `json:"type,omitempty"`
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		TypeVersion []int64 `json:"typeVersion"`
	} `json:"resultAssertions,omitempty"`
	// Reducer Options
	Settings *struct {
		// Non-number reduce behavior
		// Possible enum values:
		//  - `"dropNN"` Drop non-numbers
		//  - `"replaceNN"` Replace non-numbers
		Mode TypeReduceMode `json:"mode"`
		// Only valid when mode is replace
		ReplaceWithValue *float64 `json:"replaceWithValue,omitempty"`
	} `json:"settings,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *struct {
		// From is the start time of the query.
		From string `json:"from"`
		// To is the end time of the query.
		To string `json:"to"`
	} `json:"timeRange,omitempty"`
	Type string `json:"type"`
}

func (resource TypeReduce) ImplementsDataqueryVariant() {}

type TypeResample struct {
	// The datasource
	Datasource *struct {
		// The apiserver version
		ApiVersion *string `json:"apiVersion,omitempty"`
		// The datasource plugin type
		Type string `json:"type"`
		// Datasource UID (NOTE: name in k8s)
		Uid *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	// The downsample function
	// Possible enum values:
	//  - `"sum"`
	//  - `"mean"`
	//  - `"min"`
	//  - `"max"`
	//  - `"count"`
	//  - `"last"`
	//  - `"median"`
	Downsampler TypeResampleDownsampler `json:"downsampler"`
	// The math expression
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *struct {
		// Maximum frame count
		MaxFrames *int64 `json:"maxFrames,omitempty"`
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""`
		//  - `"timeseries-wide"`
		//  - `"timeseries-long"`
		//  - `"timeseries-many"`
		//  - `"timeseries-multi"`
		//  - `"directory-listing"`
		//  - `"table"`
		//  - `"numeric-wide"`
		//  - `"numeric-multi"`
		//  - `"numeric-long"`
		//  - `"log-lines"`
		Type *TypeResampleType `json:"type,omitempty"`
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		TypeVersion []int64 `json:"typeVersion"`
	} `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *struct {
		// From is the start time of the query.
		From string `json:"from"`
		// To is the end time of the query.
		To string `json:"to"`
	} `json:"timeRange,omitempty"`
	Type string `json:"type"`
	// The upsample function
	// Possible enum values:
	//  - `"pad"` Use the last seen value
	//  - `"backfilling"` backfill
	//  - `"fillna"` Do not fill values (nill)
	Upsampler TypeResampleUpsampler `json:"upsampler"`
	// The time duration
	Window string `json:"window"`
}

func (resource TypeResample) ImplementsDataqueryVariant() {}

type TypeClassicConditions struct {
	Conditions []struct {
		Evaluator struct {
			Params []float64 `json:"params"`
			// e.g. "gt"
			Type string `json:"type"`
		} `json:"evaluator"`
		Operator struct {
			Type TypeClassicConditionsType `json:"type"`
		} `json:"operator"`
		Query struct {
			Params []string `json:"params"`
		} `json:"query"`
		Reducer struct {
			Type string `json:"type"`
		} `json:"reducer"`
	} `json:"conditions"`
	// The datasource
	Datasource *struct {
		// The apiserver version
		ApiVersion *string `json:"apiVersion,omitempty"`
		// The datasource plugin type
		Type string `json:"type"`
		// Datasource UID (NOTE: name in k8s)
		Uid *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *struct {
		// Maximum frame count
		MaxFrames *int64 `json:"maxFrames,omitempty"`
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""`
		//  - `"timeseries-wide"`
		//  - `"timeseries-long"`
		//  - `"timeseries-many"`
		//  - `"timeseries-multi"`
		//  - `"directory-listing"`
		//  - `"table"`
		//  - `"numeric-wide"`
		//  - `"numeric-multi"`
		//  - `"numeric-long"`
		//  - `"log-lines"`
		Type *TypeClassicConditionsType `json:"type,omitempty"`
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		TypeVersion []int64 `json:"typeVersion"`
	} `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *struct {
		// From is the start time of the query.
		From string `json:"from"`
		// To is the end time of the query.
		To string `json:"to"`
	} `json:"timeRange,omitempty"`
	Type string `json:"type"`
}

func (resource TypeClassicConditions) ImplementsDataqueryVariant() {}

type TypeThreshold struct {
	// Threshold Conditions
	Conditions []struct {
		Evaluator struct {
			Params []float64 `json:"params"`
			// e.g. "gt"
			Type TypeThresholdType `json:"type"`
		} `json:"evaluator"`
		LoadedDimensions any `json:"loadedDimensions,omitempty"`
		UnloadEvaluator  *struct {
			Params []float64 `json:"params"`
			// e.g. "gt"
			Type TypeThresholdType `json:"type"`
		} `json:"unloadEvaluator,omitempty"`
	} `json:"conditions"`
	// The datasource
	Datasource *struct {
		// The apiserver version
		ApiVersion *string `json:"apiVersion,omitempty"`
		// The datasource plugin type
		Type string `json:"type"`
		// Datasource UID (NOTE: name in k8s)
		Uid *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	// Reference to single query result
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *struct {
		// Maximum frame count
		MaxFrames *int64 `json:"maxFrames,omitempty"`
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""`
		//  - `"timeseries-wide"`
		//  - `"timeseries-long"`
		//  - `"timeseries-many"`
		//  - `"timeseries-multi"`
		//  - `"directory-listing"`
		//  - `"table"`
		//  - `"numeric-wide"`
		//  - `"numeric-multi"`
		//  - `"numeric-long"`
		//  - `"log-lines"`
		Type *TypeThresholdType `json:"type,omitempty"`
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		TypeVersion []int64 `json:"typeVersion"`
	} `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *struct {
		// From is the start time of the query.
		From string `json:"from"`
		// To is the end time of the query.
		To string `json:"to"`
	} `json:"timeRange,omitempty"`
	Type string `json:"type"`
}

func (resource TypeThreshold) ImplementsDataqueryVariant() {}

type TypeSql struct {
	// The datasource
	Datasource *struct {
		// The apiserver version
		ApiVersion *string `json:"apiVersion,omitempty"`
		// The datasource plugin type
		Type string `json:"type"`
		// Datasource UID (NOTE: name in k8s)
		Uid *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *struct {
		// Maximum frame count
		MaxFrames *int64 `json:"maxFrames,omitempty"`
		// Type asserts that the frame matches a known type structure.
		// Possible enum values:
		//  - `""`
		//  - `"timeseries-wide"`
		//  - `"timeseries-long"`
		//  - `"timeseries-many"`
		//  - `"timeseries-multi"`
		//  - `"directory-listing"`
		//  - `"table"`
		//  - `"numeric-wide"`
		//  - `"numeric-multi"`
		//  - `"numeric-long"`
		//  - `"log-lines"`
		Type *TypeSqlType `json:"type,omitempty"`
		// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
		// contract documentation https://grafana.github.io/dataplane/contract/.
		TypeVersion []int64 `json:"typeVersion"`
	} `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *struct {
		// From is the start time of the query.
		From string `json:"from"`
		// To is the end time of the query.
		To string `json:"to"`
	} `json:"timeRange,omitempty"`
	Type string `json:"type"`
}

func (resource TypeSql) ImplementsDataqueryVariant() {}

type TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql struct {
	TypeMath              *TypeMath              `json:"TypeMath,omitempty"`
	TypeReduce            *TypeReduce            `json:"TypeReduce,omitempty"`
	TypeResample          *TypeResample          `json:"TypeResample,omitempty"`
	TypeClassicConditions *TypeClassicConditions `json:"TypeClassicConditions,omitempty"`
	TypeThreshold         *TypeThreshold         `json:"TypeThreshold,omitempty"`
	TypeSql               *TypeSql               `json:"TypeSql,omitempty"`
}

func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) MarshalJSON() ([]byte, error) {
	if resource.TypeMath != nil {
		return json.Marshal(resource.TypeMath)
	}
	if resource.TypeReduce != nil {
		return json.Marshal(resource.TypeReduce)
	}
	if resource.TypeResample != nil {
		return json.Marshal(resource.TypeResample)
	}
	if resource.TypeClassicConditions != nil {
		return json.Marshal(resource.TypeClassicConditions)
	}
	if resource.TypeThreshold != nil {
		return json.Marshal(resource.TypeThreshold)
	}
	if resource.TypeSql != nil {
		return json.Marshal(resource.TypeSql)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) UnmarshalJSON(raw []byte) error {
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
	case "classic_conditions":
		var typeClassicConditions TypeClassicConditions
		if err := json.Unmarshal(raw, &typeClassicConditions); err != nil {
			return err
		}

		resource.TypeClassicConditions = &typeClassicConditions
		return nil
	case "math":
		var typeMath TypeMath
		if err := json.Unmarshal(raw, &typeMath); err != nil {
			return err
		}

		resource.TypeMath = &typeMath
		return nil
	case "reduce":
		var typeReduce TypeReduce
		if err := json.Unmarshal(raw, &typeReduce); err != nil {
			return err
		}

		resource.TypeReduce = &typeReduce
		return nil
	case "resample":
		var typeResample TypeResample
		if err := json.Unmarshal(raw, &typeResample); err != nil {
			return err
		}

		resource.TypeResample = &typeResample
		return nil
	case "sql":
		var typeSql TypeSql
		if err := json.Unmarshal(raw, &typeSql); err != nil {
			return err
		}

		resource.TypeSql = &typeSql
		return nil
	case "threshold":
		var typeThreshold TypeThreshold
		if err := json.Unmarshal(raw, &typeThreshold); err != nil {
			return err
		}

		resource.TypeThreshold = &typeThreshold
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

type TypeMathType string

const (
	TypeMathTypeNone             TypeMathType = ""
	TypeMathTypeTimeseriesWide   TypeMathType = "timeseries-wide"
	TypeMathTypeTimeseriesLong   TypeMathType = "timeseries-long"
	TypeMathTypeTimeseriesMany   TypeMathType = "timeseries-many"
	TypeMathTypeTimeseriesMulti  TypeMathType = "timeseries-multi"
	TypeMathTypeDirectoryListing TypeMathType = "directory-listing"
	TypeMathTypeTable            TypeMathType = "table"
	TypeMathTypeNumericWide      TypeMathType = "numeric-wide"
	TypeMathTypeNumericMulti     TypeMathType = "numeric-multi"
	TypeMathTypeNumericLong      TypeMathType = "numeric-long"
	TypeMathTypeLogLines         TypeMathType = "log-lines"
)

type TypeReduceReducer string

const (
	TypeReduceReducerSum    TypeReduceReducer = "sum"
	TypeReduceReducerMean   TypeReduceReducer = "mean"
	TypeReduceReducerMin    TypeReduceReducer = "min"
	TypeReduceReducerMax    TypeReduceReducer = "max"
	TypeReduceReducerCount  TypeReduceReducer = "count"
	TypeReduceReducerLast   TypeReduceReducer = "last"
	TypeReduceReducerMedian TypeReduceReducer = "median"
)

type TypeReduceType string

const (
	TypeReduceTypeNone             TypeReduceType = ""
	TypeReduceTypeTimeseriesWide   TypeReduceType = "timeseries-wide"
	TypeReduceTypeTimeseriesLong   TypeReduceType = "timeseries-long"
	TypeReduceTypeTimeseriesMany   TypeReduceType = "timeseries-many"
	TypeReduceTypeTimeseriesMulti  TypeReduceType = "timeseries-multi"
	TypeReduceTypeDirectoryListing TypeReduceType = "directory-listing"
	TypeReduceTypeTable            TypeReduceType = "table"
	TypeReduceTypeNumericWide      TypeReduceType = "numeric-wide"
	TypeReduceTypeNumericMulti     TypeReduceType = "numeric-multi"
	TypeReduceTypeNumericLong      TypeReduceType = "numeric-long"
	TypeReduceTypeLogLines         TypeReduceType = "log-lines"
)

type TypeReduceMode string

const (
	TypeReduceModeDropNN    TypeReduceMode = "dropNN"
	TypeReduceModeReplaceNN TypeReduceMode = "replaceNN"
)

type TypeResampleDownsampler string

const (
	TypeResampleDownsamplerSum    TypeResampleDownsampler = "sum"
	TypeResampleDownsamplerMean   TypeResampleDownsampler = "mean"
	TypeResampleDownsamplerMin    TypeResampleDownsampler = "min"
	TypeResampleDownsamplerMax    TypeResampleDownsampler = "max"
	TypeResampleDownsamplerCount  TypeResampleDownsampler = "count"
	TypeResampleDownsamplerLast   TypeResampleDownsampler = "last"
	TypeResampleDownsamplerMedian TypeResampleDownsampler = "median"
)

type TypeResampleType string

const (
	TypeResampleTypeNone             TypeResampleType = ""
	TypeResampleTypeTimeseriesWide   TypeResampleType = "timeseries-wide"
	TypeResampleTypeTimeseriesLong   TypeResampleType = "timeseries-long"
	TypeResampleTypeTimeseriesMany   TypeResampleType = "timeseries-many"
	TypeResampleTypeTimeseriesMulti  TypeResampleType = "timeseries-multi"
	TypeResampleTypeDirectoryListing TypeResampleType = "directory-listing"
	TypeResampleTypeTable            TypeResampleType = "table"
	TypeResampleTypeNumericWide      TypeResampleType = "numeric-wide"
	TypeResampleTypeNumericMulti     TypeResampleType = "numeric-multi"
	TypeResampleTypeNumericLong      TypeResampleType = "numeric-long"
	TypeResampleTypeLogLines         TypeResampleType = "log-lines"
)

type TypeResampleUpsampler string

const (
	TypeResampleUpsamplerPad         TypeResampleUpsampler = "pad"
	TypeResampleUpsamplerBackfilling TypeResampleUpsampler = "backfilling"
	TypeResampleUpsamplerFillna      TypeResampleUpsampler = "fillna"
)

type TypeClassicConditionsType string

const (
	TypeClassicConditionsTypeAnd TypeClassicConditionsType = "and"
	TypeClassicConditionsTypeOr  TypeClassicConditionsType = "or"
)

type TypeThresholdType string

const (
	TypeThresholdTypeGt           TypeThresholdType = "gt"
	TypeThresholdTypeLt           TypeThresholdType = "lt"
	TypeThresholdTypeWithinRange  TypeThresholdType = "within_range"
	TypeThresholdTypeOutsideRange TypeThresholdType = "outside_range"
)

type TypeSqlType string

const (
	TypeSqlTypeNone             TypeSqlType = ""
	TypeSqlTypeTimeseriesWide   TypeSqlType = "timeseries-wide"
	TypeSqlTypeTimeseriesLong   TypeSqlType = "timeseries-long"
	TypeSqlTypeTimeseriesMany   TypeSqlType = "timeseries-many"
	TypeSqlTypeTimeseriesMulti  TypeSqlType = "timeseries-multi"
	TypeSqlTypeDirectoryListing TypeSqlType = "directory-listing"
	TypeSqlTypeTable            TypeSqlType = "table"
	TypeSqlTypeNumericWide      TypeSqlType = "numeric-wide"
	TypeSqlTypeNumericMulti     TypeSqlType = "numeric-multi"
	TypeSqlTypeNumericLong      TypeSqlType = "numeric-long"
	TypeSqlTypeLogLines         TypeSqlType = "log-lines"
)
