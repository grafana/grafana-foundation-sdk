// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"encoding/json"
	"errors"
	"fmt"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type MetricStat struct {
	// AWS region to query for the metric
	Region string `json:"region"`
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	Namespace string `json:"namespace"`
	// Name of the metric
	MetricName *string `json:"metricName,omitempty"`
	// The dimensions of the metric
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Only show metrics that exactly match all defined dimension names.
	MatchExact *bool `json:"matchExact,omitempty"`
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	Period *string `json:"period,omitempty"`
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	AccountId *string `json:"accountId,omitempty"`
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	Statistic *string `json:"statistic,omitempty"`
	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
type Dimensions map[string]StringOrArrayOfString

// Shape of a CloudWatch Metrics query
type CloudWatchMetricsQuery struct {
	// Whether a query is a Metrics, Logs, or Annotations query
	QueryMode CloudWatchQueryMode `json:"queryMode"`
	// Whether to use a metric search or metric insights query
	MetricQueryType *MetricQueryType `json:"metricQueryType,omitempty"`
	// Whether to use the query builder or code editor to create the query
	MetricEditorMode *MetricEditorMode `json:"metricEditorMode,omitempty"`
	// ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
	Id string `json:"id"`
	// Deprecated: use label
	// @deprecated use label
	Alias *string `json:"alias,omitempty"`
	// Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
	Label *string `json:"label,omitempty"`
	// Math expression query
	Expression *string `json:"expression,omitempty"`
	// When the metric query type is set to `Insights`, this field is used to specify the query string.
	SqlExpression *string `json:"sqlExpression,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// AWS region to query for the metric
	Region string `json:"region"`
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	Namespace string `json:"namespace"`
	// Name of the metric
	MetricName *string `json:"metricName,omitempty"`
	// The dimensions of the metric
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Only show metrics that exactly match all defined dimension names.
	MatchExact *bool `json:"matchExact,omitempty"`
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	Period *string `json:"period,omitempty"`
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	AccountId *string `json:"accountId,omitempty"`
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	Statistic *string `json:"statistic,omitempty"`
	// When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
	Sql *SQLExpression `json:"sql,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

func (resource CloudWatchMetricsQuery) ImplementsDataqueryVariant() {}

type CloudWatchQueryMode string

const (
	CloudWatchQueryModeMetrics     CloudWatchQueryMode = "Metrics"
	CloudWatchQueryModeLogs        CloudWatchQueryMode = "Logs"
	CloudWatchQueryModeAnnotations CloudWatchQueryMode = "Annotations"
)

type MetricQueryType int64

const (
	MetricQueryTypeSearch   MetricQueryType = 0
	MetricQueryTypeInsights MetricQueryType = 1
)

type MetricEditorMode int64

const (
	MetricEditorModeBuilder MetricEditorMode = 0
	MetricEditorModeCode    MetricEditorMode = 1
)

type SQLExpression struct {
	// SELECT part of the SQL expression
	Select *QueryEditorFunctionExpression `json:"select,omitempty"`
	// FROM part of the SQL expression
	From *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression `json:"from,omitempty"`
	// WHERE part of the SQL expression
	Where *QueryEditorArrayExpression `json:"where,omitempty"`
	// GROUP BY part of the SQL expression
	GroupBy *QueryEditorArrayExpression `json:"groupBy,omitempty"`
	// ORDER BY part of the SQL expression
	OrderBy *QueryEditorFunctionExpression `json:"orderBy,omitempty"`
	// The sort order of the SQL expression, `ASC` or `DESC`
	OrderByDirection *string `json:"orderByDirection,omitempty"`
	// LIMIT part of the SQL expression
	Limit *int64 `json:"limit,omitempty"`
}

type QueryEditorFunctionExpression struct {
	Type       string                                   `json:"type"`
	Name       *string                                  `json:"name,omitempty"`
	Parameters []QueryEditorFunctionParameterExpression `json:"parameters,omitempty"`
}

type QueryEditorExpressionType string

const (
	QueryEditorExpressionTypeProperty          QueryEditorExpressionType = "property"
	QueryEditorExpressionTypeOperator          QueryEditorExpressionType = "operator"
	QueryEditorExpressionTypeOr                QueryEditorExpressionType = "or"
	QueryEditorExpressionTypeAnd               QueryEditorExpressionType = "and"
	QueryEditorExpressionTypeGroupBy           QueryEditorExpressionType = "groupBy"
	QueryEditorExpressionTypeFunction          QueryEditorExpressionType = "function"
	QueryEditorExpressionTypeFunctionParameter QueryEditorExpressionType = "functionParameter"
)

type QueryEditorFunctionParameterExpression struct {
	Type string  `json:"type"`
	Name *string `json:"name,omitempty"`
}

type QueryEditorPropertyExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
}

type QueryEditorGroupByExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
}

type QueryEditorOperatorExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
	// TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
	Operator QueryEditorOperator `json:"operator"`
}

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
type QueryEditorOperator struct {
	Name  *string                                              `json:"name,omitempty"`
	Value *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType `json:"value,omitempty"`
}

type QueryEditorOperatorValueType = StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType

type QueryEditorOperatorType = StringOrBoolOrInt64

type QueryEditorProperty struct {
	Type QueryEditorPropertyType `json:"type"`
	Name *string                 `json:"name,omitempty"`
}

type QueryEditorPropertyType string

const (
	QueryEditorPropertyTypeString QueryEditorPropertyType = "string"
)

type QueryEditorArrayExpression struct {
	Type        QueryEditorArrayExpressionType `json:"type"`
	Expressions []QueryEditorExpression        `json:"expressions"`
}

type QueryEditorExpression = QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression

// Shape of a CloudWatch Logs query
type CloudWatchLogsQuery struct {
	// Whether a query is a Metrics, Logs, or Annotations query
	QueryMode CloudWatchQueryMode `json:"queryMode"`
	Id        string              `json:"id"`
	// AWS region to query for the logs
	Region string `json:"region"`
	// The CloudWatch Logs Insights query to execute
	Expression *string `json:"expression,omitempty"`
	// Fields to group the results by, this field is automatically populated whenever the query is updated
	StatsGroups []string `json:"statsGroups,omitempty"`
	// Log groups to query
	LogGroups []LogGroup `json:"logGroups,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// @deprecated use logGroups
	LogGroupNames []string `json:"logGroupNames,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
}

func (resource CloudWatchLogsQuery) ImplementsDataqueryVariant() {}

type LogGroup struct {
	// ARN of the log group
	Arn string `json:"arn"`
	// Name of the log group
	Name string `json:"name"`
	// AccountId of the log group
	AccountId *string `json:"accountId,omitempty"`
	// Label of the log group
	AccountLabel *string `json:"accountLabel,omitempty"`
}

// Shape of a CloudWatch Annotation query
// TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
// #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
type CloudWatchAnnotationQuery struct {
	// Whether a query is a Metrics, Logs, or Annotations query
	QueryMode CloudWatchQueryMode `json:"queryMode"`
	// Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
	PrefixMatching *bool `json:"prefixMatching,omitempty"`
	// Use this parameter to filter the results of the operation to only those alarms
	// that use a certain alarm action. For example, you could specify the ARN of
	// an SNS topic to find all alarms that send notifications to that topic.
	// e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
	// but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
	ActionPrefix *string `json:"actionPrefix,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// AWS region to query for the metric
	Region string `json:"region"`
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	Namespace string `json:"namespace"`
	// Name of the metric
	MetricName *string `json:"metricName,omitempty"`
	// The dimensions of the metric
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Only show metrics that exactly match all defined dimension names.
	MatchExact *bool `json:"matchExact,omitempty"`
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	Period *string `json:"period,omitempty"`
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	AccountId *string `json:"accountId,omitempty"`
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	Statistic *string `json:"statistic,omitempty"`
	// An alarm name prefix. If you specify this parameter, you receive information
	// about all alarms that have names that start with this prefix.
	// e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
	AlarmNamePrefix *string `json:"alarmNamePrefix,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

func (resource CloudWatchAnnotationQuery) ImplementsDataqueryVariant() {}

type CloudWatchQuery = CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery

func (resource CloudWatchQuery) ImplementsDataqueryVariant() {}

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

type QueryEditorArrayExpressionType string

const (
	QueryEditorArrayExpressionTypeAnd QueryEditorArrayExpressionType = "and"
	QueryEditorArrayExpressionTypeOr  QueryEditorArrayExpressionType = "or"
)

type StringOrArrayOfString struct {
	String        *string  `json:"String,omitempty"`
	ArrayOfString []string `json:"ArrayOfString,omitempty"`
}

func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
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

type QueryEditorPropertyExpressionOrQueryEditorFunctionExpression struct {
	QueryEditorPropertyExpression *QueryEditorPropertyExpression `json:"QueryEditorPropertyExpression,omitempty"`
	QueryEditorFunctionExpression *QueryEditorFunctionExpression `json:"QueryEditorFunctionExpression,omitempty"`
}

func (resource QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) MarshalJSON() ([]byte, error) {
	if resource.QueryEditorPropertyExpression != nil {
		return json.Marshal(resource.QueryEditorPropertyExpression)
	}
	if resource.QueryEditorFunctionExpression != nil {
		return json.Marshal(resource.QueryEditorFunctionExpression)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
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

type StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType struct {
	String                         *string                   `json:"String,omitempty"`
	Bool                           *bool                     `json:"Bool,omitempty"`
	Int64                          *int64                    `json:"Int64,omitempty"`
	ArrayOfQueryEditorOperatorType []QueryEditorOperatorType `json:"ArrayOfQueryEditorOperatorType,omitempty"`
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

	return nil, fmt.Errorf("no value for disjunction of scalars")
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

type StringOrBoolOrInt64 struct {
	String *string `json:"String,omitempty"`
	Bool   *bool   `json:"Bool,omitempty"`
	Int64  *int64  `json:"Int64,omitempty"`
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

	return nil, fmt.Errorf("no value for disjunction of scalars")
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

type QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression struct {
	QueryEditorArrayExpression             *QueryEditorArrayExpression             `json:"QueryEditorArrayExpression,omitempty"`
	QueryEditorPropertyExpression          *QueryEditorPropertyExpression          `json:"QueryEditorPropertyExpression,omitempty"`
	QueryEditorGroupByExpression           *QueryEditorGroupByExpression           `json:"QueryEditorGroupByExpression,omitempty"`
	QueryEditorFunctionExpression          *QueryEditorFunctionExpression          `json:"QueryEditorFunctionExpression,omitempty"`
	QueryEditorFunctionParameterExpression *QueryEditorFunctionParameterExpression `json:"QueryEditorFunctionParameterExpression,omitempty"`
	QueryEditorOperatorExpression          *QueryEditorOperatorExpression          `json:"QueryEditorOperatorExpression,omitempty"`
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

	return nil, fmt.Errorf("no value for disjunction of refs")
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

type CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery struct {
	CloudWatchMetricsQuery    *CloudWatchMetricsQuery    `json:"CloudWatchMetricsQuery,omitempty"`
	CloudWatchLogsQuery       *CloudWatchLogsQuery       `json:"CloudWatchLogsQuery,omitempty"`
	CloudWatchAnnotationQuery *CloudWatchAnnotationQuery `json:"CloudWatchAnnotationQuery,omitempty"`
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

	return nil, fmt.Errorf("no value for disjunction of refs")
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
