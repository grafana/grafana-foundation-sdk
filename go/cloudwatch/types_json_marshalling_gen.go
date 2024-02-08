// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "cloudwatch",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := CloudWatchQuery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

func (resource CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) MarshalJSON() ([]byte, error) {
	if resource.CloudWatchMetricsQuery != nil {
		return json.Marshal(resource.CloudWatchMetricsQuery)
	}
	if resource.CloudWatchLogsQuery != nil {
		return json.Marshal(resource.CloudWatchLogsQuery)
	}
	if resource.CloudWatchAnnotationQuery != nil {
		return json.Marshal(resource.CloudWatchAnnotationQuery)
	}

	return nil, nil
}

func (resource *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["queryMode"]
	if !found {
		return errors.New("discriminator field 'queryMode' not found in payload")
	}

	switch discriminator {
	case "Annotations":
		var cloudWatchAnnotationQuery CloudWatchAnnotationQuery
		if err := json.Unmarshal(raw, &cloudWatchAnnotationQuery); err != nil {
			return err
		}

		resource.CloudWatchAnnotationQuery = &cloudWatchAnnotationQuery
		return nil
	case "Logs":
		var cloudWatchLogsQuery CloudWatchLogsQuery
		if err := json.Unmarshal(raw, &cloudWatchLogsQuery); err != nil {
			return err
		}

		resource.CloudWatchLogsQuery = &cloudWatchLogsQuery
		return nil
	case "Metrics":
		var cloudWatchMetricsQuery CloudWatchMetricsQuery
		if err := json.Unmarshal(raw, &cloudWatchMetricsQuery); err != nil {
			return err
		}

		resource.CloudWatchMetricsQuery = &cloudWatchMetricsQuery
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `queryMode = %v`", discriminator)
}

func (resource QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression) MarshalJSON() ([]byte, error) {
	if resource.QueryEditorArrayExpression != nil {
		return json.Marshal(resource.QueryEditorArrayExpression)
	}
	if resource.QueryEditorPropertyExpression != nil {
		return json.Marshal(resource.QueryEditorPropertyExpression)
	}
	if resource.QueryEditorGroupByExpression != nil {
		return json.Marshal(resource.QueryEditorGroupByExpression)
	}
	if resource.QueryEditorFunctionExpression != nil {
		return json.Marshal(resource.QueryEditorFunctionExpression)
	}
	if resource.QueryEditorFunctionParameterExpression != nil {
		return json.Marshal(resource.QueryEditorFunctionParameterExpression)
	}
	if resource.QueryEditorOperatorExpression != nil {
		return json.Marshal(resource.QueryEditorOperatorExpression)
	}

	return nil, nil
}

func (resource *QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression) UnmarshalJSON(raw []byte) error {
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
	case "and":
		var queryEditorArrayExpression QueryEditorArrayExpression
		if err := json.Unmarshal(raw, &queryEditorArrayExpression); err != nil {
			return err
		}

		resource.QueryEditorArrayExpression = &queryEditorArrayExpression
		return nil
	case "function":
		var queryEditorFunctionExpression QueryEditorFunctionExpression
		if err := json.Unmarshal(raw, &queryEditorFunctionExpression); err != nil {
			return err
		}

		resource.QueryEditorFunctionExpression = &queryEditorFunctionExpression
		return nil
	case "functionParameter":
		var queryEditorFunctionParameterExpression QueryEditorFunctionParameterExpression
		if err := json.Unmarshal(raw, &queryEditorFunctionParameterExpression); err != nil {
			return err
		}

		resource.QueryEditorFunctionParameterExpression = &queryEditorFunctionParameterExpression
		return nil
	case "groupBy":
		var queryEditorGroupByExpression QueryEditorGroupByExpression
		if err := json.Unmarshal(raw, &queryEditorGroupByExpression); err != nil {
			return err
		}

		resource.QueryEditorGroupByExpression = &queryEditorGroupByExpression
		return nil
	case "operator":
		var queryEditorOperatorExpression QueryEditorOperatorExpression
		if err := json.Unmarshal(raw, &queryEditorOperatorExpression); err != nil {
			return err
		}

		resource.QueryEditorOperatorExpression = &queryEditorOperatorExpression
		return nil
	case "or":
		var queryEditorArrayExpression QueryEditorArrayExpression
		if err := json.Unmarshal(raw, &queryEditorArrayExpression); err != nil {
			return err
		}

		resource.QueryEditorArrayExpression = &queryEditorArrayExpression
		return nil
	case "property":
		var queryEditorPropertyExpression QueryEditorPropertyExpression
		if err := json.Unmarshal(raw, &queryEditorPropertyExpression); err != nil {
			return err
		}

		resource.QueryEditorPropertyExpression = &queryEditorPropertyExpression
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) MarshalJSON() ([]byte, error) {
	if resource.QueryEditorPropertyExpression != nil {
		return json.Marshal(resource.QueryEditorPropertyExpression)
	}
	if resource.QueryEditorFunctionExpression != nil {
		return json.Marshal(resource.QueryEditorFunctionExpression)
	}

	return nil, nil
}

func (resource *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) UnmarshalJSON(raw []byte) error {
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
	case "function":
		var queryEditorFunctionExpression QueryEditorFunctionExpression
		if err := json.Unmarshal(raw, &queryEditorFunctionExpression); err != nil {
			return err
		}

		resource.QueryEditorFunctionExpression = &queryEditorFunctionExpression
		return nil
	case "property":
		var queryEditorPropertyExpression QueryEditorPropertyExpression
		if err := json.Unmarshal(raw, &queryEditorPropertyExpression); err != nil {
			return err
		}

		resource.QueryEditorPropertyExpression = &queryEditorPropertyExpression
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return nil, nil
}

func (resource *StringOrArrayOfString) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// ArrayOfString
	var ArrayOfString []string
	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
		resource.ArrayOfString = nil
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	return errors.Join(errList...)
}

func (resource StringOrBoolOrInt64) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Int64 != nil {
		return json.Marshal(resource.Int64)
	}

	return nil, nil
}

func (resource *StringOrBoolOrInt64) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Int64
	var Int64 int64
	if err := json.Unmarshal(raw, &Int64); err != nil {
		errList = append(errList, err)
		resource.Int64 = nil
	} else {
		resource.Int64 = &Int64
		return nil
	}

	return errors.Join(errList...)
}

func (resource StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.Bool != nil {
		return json.Marshal(resource.Bool)
	}

	if resource.Int64 != nil {
		return json.Marshal(resource.Int64)
	}

	if resource.ArrayOfQueryEditorOperatorType != nil {
		return json.Marshal(resource.ArrayOfQueryEditorOperatorType)
	}

	return nil, nil
}

func (resource *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	var errList []error

	// String
	var String string
	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
		resource.String = nil
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool
	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
		resource.Bool = nil
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Int64
	var Int64 int64
	if err := json.Unmarshal(raw, &Int64); err != nil {
		errList = append(errList, err)
		resource.Int64 = nil
	} else {
		resource.Int64 = &Int64
		return nil
	}

	// ArrayOfQueryEditorOperatorType
	var ArrayOfQueryEditorOperatorType []QueryEditorOperatorType
	if err := json.Unmarshal(raw, &ArrayOfQueryEditorOperatorType); err != nil {
		errList = append(errList, err)
		resource.ArrayOfQueryEditorOperatorType = nil
	} else {
		resource.ArrayOfQueryEditorOperatorType = ArrayOfQueryEditorOperatorType
		return nil
	}

	return errors.Join(errList...)
}
