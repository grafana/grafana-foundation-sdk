// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricStat` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricStat) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is null"))...)

		}
		delete(fields, "region")
	} else {
		errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is missing from input"))...)
	}
	// Field "namespace"
	if fields["namespace"] != nil {
		if string(fields["namespace"]) != "null" {
			if err := json.Unmarshal(fields["namespace"], &resource.Namespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("namespace", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("namespace", errors.New("required field is null"))...)

		}
		delete(fields, "namespace")
	} else {
		errs = append(errs, cog.MakeBuildErrors("namespace", errors.New("required field is missing from input"))...)
	}
	// Field "metricName"
	if fields["metricName"] != nil {
		if string(fields["metricName"]) != "null" {
			if err := json.Unmarshal(fields["metricName"], &resource.MetricName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricName", err)...)
			}

		}
		delete(fields, "metricName")

	}
	// Field "dimensions"
	if fields["dimensions"] != nil {
		if string(fields["dimensions"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["dimensions"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]StringOrArrayOfString, len(partialMap))
			for key1 := range partialMap {
				var result1 StringOrArrayOfString

				result1 = StringOrArrayOfString{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("dimensions["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Dimensions = cog.ToPtr(Dimensions(parsedMap1))

		}
		delete(fields, "dimensions")

	}
	// Field "matchExact"
	if fields["matchExact"] != nil {
		if string(fields["matchExact"]) != "null" {
			if err := json.Unmarshal(fields["matchExact"], &resource.MatchExact); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matchExact", err)...)
			}

		}
		delete(fields, "matchExact")

	}
	// Field "period"
	if fields["period"] != nil {
		if string(fields["period"]) != "null" {
			if err := json.Unmarshal(fields["period"], &resource.Period); err != nil {
				errs = append(errs, cog.MakeBuildErrors("period", err)...)
			}

		}
		delete(fields, "period")

	}
	// Field "accountId"
	if fields["accountId"] != nil {
		if string(fields["accountId"]) != "null" {
			if err := json.Unmarshal(fields["accountId"], &resource.AccountId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("accountId", err)...)
			}

		}
		delete(fields, "accountId")

	}
	// Field "statistic"
	if fields["statistic"] != nil {
		if string(fields["statistic"]) != "null" {
			if err := json.Unmarshal(fields["statistic"], &resource.Statistic); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statistic", err)...)
			}

		}
		delete(fields, "statistic")

	}
	// Field "statistics"
	if fields["statistics"] != nil {
		if string(fields["statistics"]) != "null" {

			if err := json.Unmarshal(fields["statistics"], &resource.Statistics); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statistics", err)...)
			}

		}
		delete(fields, "statistics")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricStat", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricStat` objects.
func (resource MetricStat) Equals(other MetricStat) bool {
	if resource.Region != other.Region {
		return false
	}
	if resource.Namespace != other.Namespace {
		return false
	}
	if resource.MetricName == nil && other.MetricName != nil || resource.MetricName != nil && other.MetricName == nil {
		return false
	}

	if resource.MetricName != nil {
		if *resource.MetricName != *other.MetricName {
			return false
		}
	}
	if resource.Dimensions == nil && other.Dimensions != nil || resource.Dimensions != nil && other.Dimensions == nil {
		return false
	}

	if resource.Dimensions != nil {

		if len(*resource.Dimensions) != len(*other.Dimensions) {
			return false
		}

		for key1 := range *resource.Dimensions {
			if !(*resource.Dimensions)[key1].Equals((*other.Dimensions)[key1]) {
				return false
			}
		}
	}
	if resource.MatchExact == nil && other.MatchExact != nil || resource.MatchExact != nil && other.MatchExact == nil {
		return false
	}

	if resource.MatchExact != nil {
		if *resource.MatchExact != *other.MatchExact {
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
	if resource.AccountId == nil && other.AccountId != nil || resource.AccountId != nil && other.AccountId == nil {
		return false
	}

	if resource.AccountId != nil {
		if *resource.AccountId != *other.AccountId {
			return false
		}
	}
	if resource.Statistic == nil && other.Statistic != nil || resource.Statistic != nil && other.Statistic == nil {
		return false
	}

	if resource.Statistic != nil {
		if *resource.Statistic != *other.Statistic {
			return false
		}
	}

	if len(resource.Statistics) != len(other.Statistics) {
		return false
	}

	for i1 := range resource.Statistics {
		if resource.Statistics[i1] != other.Statistics[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `MetricStat` fields for violations and returns them.
func (resource MetricStat) Validate() error {
	return nil
}

// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
type Dimensions map[string]StringOrArrayOfString

// Shape of a CloudWatch Metrics query
type CloudWatchMetricsQuery struct {
	// Whether a query is a Metrics, Logs, or Annotations query
	QueryMode CloudWatchQueryMode `json:"queryMode"`
	// Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
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
	// When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
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
	// When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
	Sql *SQLExpression `json:"sql,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

func (resource CloudWatchMetricsQuery) ImplementsDataqueryVariant() {}

func (resource CloudWatchMetricsQuery) DataqueryType() string {
	return "cloudwatch"
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchMetricsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CloudWatchMetricsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "queryMode"
	if fields["queryMode"] != nil {
		if string(fields["queryMode"]) != "null" {
			if err := json.Unmarshal(fields["queryMode"], &resource.QueryMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("queryMode", errors.New("required field is null"))...)

		}
		delete(fields, "queryMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("queryMode", errors.New("required field is missing from input"))...)
	}
	// Field "metricQueryType"
	if fields["metricQueryType"] != nil {
		if string(fields["metricQueryType"]) != "null" {
			if err := json.Unmarshal(fields["metricQueryType"], &resource.MetricQueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricQueryType", err)...)
			}

		}
		delete(fields, "metricQueryType")

	}
	// Field "metricEditorMode"
	if fields["metricEditorMode"] != nil {
		if string(fields["metricEditorMode"]) != "null" {
			if err := json.Unmarshal(fields["metricEditorMode"], &resource.MetricEditorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricEditorMode", err)...)
			}

		}
		delete(fields, "metricEditorMode")

	}
	// Field "id"
	if fields["id"] != nil {
		if string(fields["id"]) != "null" {
			if err := json.Unmarshal(fields["id"], &resource.Id); err != nil {
				errs = append(errs, cog.MakeBuildErrors("id", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is null"))...)

		}
		delete(fields, "id")
	} else {
		errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is missing from input"))...)
	}
	// Field "alias"
	if fields["alias"] != nil {
		if string(fields["alias"]) != "null" {
			if err := json.Unmarshal(fields["alias"], &resource.Alias); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alias", err)...)
			}

		}
		delete(fields, "alias")

	}
	// Field "label"
	if fields["label"] != nil {
		if string(fields["label"]) != "null" {
			if err := json.Unmarshal(fields["label"], &resource.Label); err != nil {
				errs = append(errs, cog.MakeBuildErrors("label", err)...)
			}

		}
		delete(fields, "label")

	}
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}

		}
		delete(fields, "expression")

	}
	// Field "sqlExpression"
	if fields["sqlExpression"] != nil {
		if string(fields["sqlExpression"]) != "null" {
			if err := json.Unmarshal(fields["sqlExpression"], &resource.SqlExpression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sqlExpression", err)...)
			}

		}
		delete(fields, "sqlExpression")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is null"))...)

		}
		delete(fields, "region")
	} else {
		errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is missing from input"))...)
	}
	// Field "namespace"
	if fields["namespace"] != nil {
		if string(fields["namespace"]) != "null" {
			if err := json.Unmarshal(fields["namespace"], &resource.Namespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("namespace", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("namespace", errors.New("required field is null"))...)

		}
		delete(fields, "namespace")
	} else {
		errs = append(errs, cog.MakeBuildErrors("namespace", errors.New("required field is missing from input"))...)
	}
	// Field "metricName"
	if fields["metricName"] != nil {
		if string(fields["metricName"]) != "null" {
			if err := json.Unmarshal(fields["metricName"], &resource.MetricName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricName", err)...)
			}

		}
		delete(fields, "metricName")

	}
	// Field "dimensions"
	if fields["dimensions"] != nil {
		if string(fields["dimensions"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["dimensions"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]StringOrArrayOfString, len(partialMap))
			for key1 := range partialMap {
				var result1 StringOrArrayOfString

				result1 = StringOrArrayOfString{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("dimensions["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Dimensions = cog.ToPtr(Dimensions(parsedMap1))

		}
		delete(fields, "dimensions")

	}
	// Field "matchExact"
	if fields["matchExact"] != nil {
		if string(fields["matchExact"]) != "null" {
			if err := json.Unmarshal(fields["matchExact"], &resource.MatchExact); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matchExact", err)...)
			}

		}
		delete(fields, "matchExact")

	}
	// Field "period"
	if fields["period"] != nil {
		if string(fields["period"]) != "null" {
			if err := json.Unmarshal(fields["period"], &resource.Period); err != nil {
				errs = append(errs, cog.MakeBuildErrors("period", err)...)
			}

		}
		delete(fields, "period")

	}
	// Field "accountId"
	if fields["accountId"] != nil {
		if string(fields["accountId"]) != "null" {
			if err := json.Unmarshal(fields["accountId"], &resource.AccountId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("accountId", err)...)
			}

		}
		delete(fields, "accountId")

	}
	// Field "statistic"
	if fields["statistic"] != nil {
		if string(fields["statistic"]) != "null" {
			if err := json.Unmarshal(fields["statistic"], &resource.Statistic); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statistic", err)...)
			}

		}
		delete(fields, "statistic")

	}
	// Field "sql"
	if fields["sql"] != nil {
		if string(fields["sql"]) != "null" {

			resource.Sql = &SQLExpression{}
			if err := resource.Sql.UnmarshalJSONStrict(fields["sql"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sql", err)...)
			}

		}
		delete(fields, "sql")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "statistics"
	if fields["statistics"] != nil {
		if string(fields["statistics"]) != "null" {

			if err := json.Unmarshal(fields["statistics"], &resource.Statistics); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statistics", err)...)
			}

		}
		delete(fields, "statistics")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CloudWatchMetricsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
func (resource CloudWatchMetricsQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(CloudWatchMetricsQuery)
	if !ok {
		return false
	}
	if resource.QueryMode != other.QueryMode {
		return false
	}
	if resource.MetricQueryType == nil && other.MetricQueryType != nil || resource.MetricQueryType != nil && other.MetricQueryType == nil {
		return false
	}

	if resource.MetricQueryType != nil {
		if *resource.MetricQueryType != *other.MetricQueryType {
			return false
		}
	}
	if resource.MetricEditorMode == nil && other.MetricEditorMode != nil || resource.MetricEditorMode != nil && other.MetricEditorMode == nil {
		return false
	}

	if resource.MetricEditorMode != nil {
		if *resource.MetricEditorMode != *other.MetricEditorMode {
			return false
		}
	}
	if resource.Id != other.Id {
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
	if resource.Label == nil && other.Label != nil || resource.Label != nil && other.Label == nil {
		return false
	}

	if resource.Label != nil {
		if *resource.Label != *other.Label {
			return false
		}
	}
	if resource.Expression == nil && other.Expression != nil || resource.Expression != nil && other.Expression == nil {
		return false
	}

	if resource.Expression != nil {
		if *resource.Expression != *other.Expression {
			return false
		}
	}
	if resource.SqlExpression == nil && other.SqlExpression != nil || resource.SqlExpression != nil && other.SqlExpression == nil {
		return false
	}

	if resource.SqlExpression != nil {
		if *resource.SqlExpression != *other.SqlExpression {
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
	if resource.Region != other.Region {
		return false
	}
	if resource.Namespace != other.Namespace {
		return false
	}
	if resource.MetricName == nil && other.MetricName != nil || resource.MetricName != nil && other.MetricName == nil {
		return false
	}

	if resource.MetricName != nil {
		if *resource.MetricName != *other.MetricName {
			return false
		}
	}
	if resource.Dimensions == nil && other.Dimensions != nil || resource.Dimensions != nil && other.Dimensions == nil {
		return false
	}

	if resource.Dimensions != nil {

		if len(*resource.Dimensions) != len(*other.Dimensions) {
			return false
		}

		for key1 := range *resource.Dimensions {
			if !(*resource.Dimensions)[key1].Equals((*other.Dimensions)[key1]) {
				return false
			}
		}
	}
	if resource.MatchExact == nil && other.MatchExact != nil || resource.MatchExact != nil && other.MatchExact == nil {
		return false
	}

	if resource.MatchExact != nil {
		if *resource.MatchExact != *other.MatchExact {
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
	if resource.AccountId == nil && other.AccountId != nil || resource.AccountId != nil && other.AccountId == nil {
		return false
	}

	if resource.AccountId != nil {
		if *resource.AccountId != *other.AccountId {
			return false
		}
	}
	if resource.Statistic == nil && other.Statistic != nil || resource.Statistic != nil && other.Statistic == nil {
		return false
	}

	if resource.Statistic != nil {
		if *resource.Statistic != *other.Statistic {
			return false
		}
	}
	if resource.Sql == nil && other.Sql != nil || resource.Sql != nil && other.Sql == nil {
		return false
	}

	if resource.Sql != nil {
		if !resource.Sql.Equals(*other.Sql) {
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

	if len(resource.Statistics) != len(other.Statistics) {
		return false
	}

	for i1 := range resource.Statistics {
		if resource.Statistics[i1] != other.Statistics[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CloudWatchMetricsQuery` fields for violations and returns them.
func (resource CloudWatchMetricsQuery) Validate() error {
	var errs cog.BuildErrors
	if resource.Sql != nil {
		if err := resource.Sql.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("sql", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CloudWatchQueryMode string

const (
	CloudWatchQueryModeMetrics     CloudWatchQueryMode = "Metrics"
	CloudWatchQueryModeLogs        CloudWatchQueryMode = "Logs"
	CloudWatchQueryModeAnnotations CloudWatchQueryMode = "Annotations"
)

type MetricQueryType int64

const (
	MetricQueryTypeSearch MetricQueryType = 0
	MetricQueryTypeQuery  MetricQueryType = 1
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SQLExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SQLExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "select"
	if fields["select"] != nil {
		if string(fields["select"]) != "null" {

			resource.Select = &QueryEditorFunctionExpression{}
			if err := resource.Select.UnmarshalJSONStrict(fields["select"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("select", err)...)
			}

		}
		delete(fields, "select")

	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {

			resource.From = &QueryEditorPropertyExpressionOrQueryEditorFunctionExpression{}
			if err := resource.From.UnmarshalJSONStrict(fields["from"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}

		}
		delete(fields, "from")

	}
	// Field "where"
	if fields["where"] != nil {
		if string(fields["where"]) != "null" {

			resource.Where = &QueryEditorArrayExpression{}
			if err := resource.Where.UnmarshalJSONStrict(fields["where"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("where", err)...)
			}

		}
		delete(fields, "where")

	}
	// Field "groupBy"
	if fields["groupBy"] != nil {
		if string(fields["groupBy"]) != "null" {

			resource.GroupBy = &QueryEditorArrayExpression{}
			if err := resource.GroupBy.UnmarshalJSONStrict(fields["groupBy"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupBy", err)...)
			}

		}
		delete(fields, "groupBy")

	}
	// Field "orderBy"
	if fields["orderBy"] != nil {
		if string(fields["orderBy"]) != "null" {

			resource.OrderBy = &QueryEditorFunctionExpression{}
			if err := resource.OrderBy.UnmarshalJSONStrict(fields["orderBy"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
			}

		}
		delete(fields, "orderBy")

	}
	// Field "orderByDirection"
	if fields["orderByDirection"] != nil {
		if string(fields["orderByDirection"]) != "null" {
			if err := json.Unmarshal(fields["orderByDirection"], &resource.OrderByDirection); err != nil {
				errs = append(errs, cog.MakeBuildErrors("orderByDirection", err)...)
			}

		}
		delete(fields, "orderByDirection")

	}
	// Field "limit"
	if fields["limit"] != nil {
		if string(fields["limit"]) != "null" {
			if err := json.Unmarshal(fields["limit"], &resource.Limit); err != nil {
				errs = append(errs, cog.MakeBuildErrors("limit", err)...)
			}

		}
		delete(fields, "limit")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SQLExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SQLExpression` objects.
func (resource SQLExpression) Equals(other SQLExpression) bool {
	if resource.Select == nil && other.Select != nil || resource.Select != nil && other.Select == nil {
		return false
	}

	if resource.Select != nil {
		if !resource.Select.Equals(*other.Select) {
			return false
		}
	}
	if resource.From == nil && other.From != nil || resource.From != nil && other.From == nil {
		return false
	}

	if resource.From != nil {
		if !resource.From.Equals(*other.From) {
			return false
		}
	}
	if resource.Where == nil && other.Where != nil || resource.Where != nil && other.Where == nil {
		return false
	}

	if resource.Where != nil {
		if !resource.Where.Equals(*other.Where) {
			return false
		}
	}
	if resource.GroupBy == nil && other.GroupBy != nil || resource.GroupBy != nil && other.GroupBy == nil {
		return false
	}

	if resource.GroupBy != nil {
		if !resource.GroupBy.Equals(*other.GroupBy) {
			return false
		}
	}
	if resource.OrderBy == nil && other.OrderBy != nil || resource.OrderBy != nil && other.OrderBy == nil {
		return false
	}

	if resource.OrderBy != nil {
		if !resource.OrderBy.Equals(*other.OrderBy) {
			return false
		}
	}
	if resource.OrderByDirection == nil && other.OrderByDirection != nil || resource.OrderByDirection != nil && other.OrderByDirection == nil {
		return false
	}

	if resource.OrderByDirection != nil {
		if *resource.OrderByDirection != *other.OrderByDirection {
			return false
		}
	}
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

// Validate checks all the validation constraints that may be defined on `SQLExpression` fields for violations and returns them.
func (resource SQLExpression) Validate() error {
	var errs cog.BuildErrors
	if resource.Select != nil {
		if err := resource.Select.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("select", err)...)
		}
	}
	if resource.From != nil {
		if err := resource.From.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("from", err)...)
		}
	}
	if resource.Where != nil {
		if err := resource.Where.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("where", err)...)
		}
	}
	if resource.GroupBy != nil {
		if err := resource.GroupBy.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("groupBy", err)...)
		}
	}
	if resource.OrderBy != nil {
		if err := resource.OrderBy.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("orderBy", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorFunctionExpression struct {
	Type       string                                   `json:"type"`
	Name       *string                                  `json:"name,omitempty"`
	Parameters []QueryEditorFunctionParameterExpression `json:"parameters,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorFunctionExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorFunctionExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "parameters"
	if fields["parameters"] != nil {
		if string(fields["parameters"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["parameters"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 QueryEditorFunctionParameterExpression

				result1 = QueryEditorFunctionParameterExpression{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("parameters["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Parameters = append(resource.Parameters, result1)
			}

		}
		delete(fields, "parameters")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorFunctionExpression` objects.
func (resource QueryEditorFunctionExpression) Equals(other QueryEditorFunctionExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	if len(resource.Parameters) != len(other.Parameters) {
		return false
	}

	for i1 := range resource.Parameters {
		if !resource.Parameters[i1].Equals(other.Parameters[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorFunctionExpression` fields for violations and returns them.
func (resource QueryEditorFunctionExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "function") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == function"),
		)...)
	}

	for i1 := range resource.Parameters {
		if err := resource.Parameters[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("parameters["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorFunctionParameterExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorFunctionParameterExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionParameterExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorFunctionParameterExpression` objects.
func (resource QueryEditorFunctionParameterExpression) Equals(other QueryEditorFunctionParameterExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorFunctionParameterExpression` fields for violations and returns them.
func (resource QueryEditorFunctionParameterExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "functionParameter") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == functionParameter"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorPropertyExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorPropertyExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "property"
	if fields["property"] != nil {
		if string(fields["property"]) != "null" {

			resource.Property = QueryEditorProperty{}
			if err := resource.Property.UnmarshalJSONStrict(fields["property"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("property", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is null"))...)

		}
		delete(fields, "property")
	} else {
		errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorPropertyExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorPropertyExpression` objects.
func (resource QueryEditorPropertyExpression) Equals(other QueryEditorPropertyExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Property.Equals(other.Property) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorPropertyExpression` fields for violations and returns them.
func (resource QueryEditorPropertyExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "property") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == property"),
		)...)
	}
	if err := resource.Property.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("property", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorGroupByExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorGroupByExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorGroupByExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "property"
	if fields["property"] != nil {
		if string(fields["property"]) != "null" {

			resource.Property = QueryEditorProperty{}
			if err := resource.Property.UnmarshalJSONStrict(fields["property"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("property", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is null"))...)

		}
		delete(fields, "property")
	} else {
		errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorGroupByExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorGroupByExpression` objects.
func (resource QueryEditorGroupByExpression) Equals(other QueryEditorGroupByExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Property.Equals(other.Property) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorGroupByExpression` fields for violations and returns them.
func (resource QueryEditorGroupByExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "groupBy") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == groupBy"),
		)...)
	}
	if err := resource.Property.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("property", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorOperatorExpression struct {
	Type     string              `json:"type"`
	Property QueryEditorProperty `json:"property"`
	// TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
	Operator QueryEditorOperator `json:"operator"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorOperatorExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorOperatorExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "property"
	if fields["property"] != nil {
		if string(fields["property"]) != "null" {

			resource.Property = QueryEditorProperty{}
			if err := resource.Property.UnmarshalJSONStrict(fields["property"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("property", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is null"))...)

		}
		delete(fields, "property")
	} else {
		errs = append(errs, cog.MakeBuildErrors("property", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {

			resource.Operator = QueryEditorOperator{}
			if err := resource.Operator.UnmarshalJSONStrict(fields["operator"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorOperatorExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorOperatorExpression` objects.
func (resource QueryEditorOperatorExpression) Equals(other QueryEditorOperatorExpression) bool {
	if resource.Type != other.Type {
		return false
	}
	if !resource.Property.Equals(other.Property) {
		return false
	}
	if !resource.Operator.Equals(other.Operator) {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorOperatorExpression` fields for violations and returns them.
func (resource QueryEditorOperatorExpression) Validate() error {
	var errs cog.BuildErrors
	if !(resource.Type == "operator") {
		errs = append(errs, cog.MakeBuildErrors(
			"type",
			errors.New("must be == operator"),
		)...)
	}
	if err := resource.Property.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("property", err)...)
	}
	if err := resource.Operator.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("operator", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
type QueryEditorOperator struct {
	Name  *string                                              `json:"name,omitempty"`
	Value *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType `json:"value,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorOperator` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorOperator) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {

			resource.Value = &StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType{}
			if err := resource.Value.UnmarshalJSONStrict(fields["value"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}

		}
		delete(fields, "value")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorOperator", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorOperator` objects.
func (resource QueryEditorOperator) Equals(other QueryEditorOperator) bool {
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}
	if resource.Value == nil && other.Value != nil || resource.Value != nil && other.Value == nil {
		return false
	}

	if resource.Value != nil {
		if !resource.Value.Equals(*other.Value) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorOperator` fields for violations and returns them.
func (resource QueryEditorOperator) Validate() error {
	var errs cog.BuildErrors
	if resource.Value != nil {
		if err := resource.Value.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("value", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type QueryEditorOperatorValueType = StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType

type QueryEditorOperatorType = StringOrBoolOrInt64

type QueryEditorProperty struct {
	Type QueryEditorPropertyType `json:"type"`
	Name *string                 `json:"name,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorProperty` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorProperty) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}

		}
		delete(fields, "name")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorProperty", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorProperty` objects.
func (resource QueryEditorProperty) Equals(other QueryEditorProperty) bool {
	if resource.Type != other.Type {
		return false
	}
	if resource.Name == nil && other.Name != nil || resource.Name != nil && other.Name == nil {
		return false
	}

	if resource.Name != nil {
		if *resource.Name != *other.Name {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorProperty` fields for violations and returns them.
func (resource QueryEditorProperty) Validate() error {
	return nil
}

type QueryEditorPropertyType string

const (
	QueryEditorPropertyTypeString QueryEditorPropertyType = "string"
)

type QueryEditorArrayExpression struct {
	Type        QueryEditorArrayExpressionType `json:"type"`
	Expressions []QueryEditorExpression        `json:"expressions"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorArrayExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorArrayExpression) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "expressions"
	if fields["expressions"] != nil {
		if string(fields["expressions"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["expressions"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 QueryEditorExpression

				result1 = QueryEditorExpression{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("expressions["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Expressions = append(resource.Expressions, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is null"))...)

		}
		delete(fields, "expressions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expressions", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorArrayExpression", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `QueryEditorArrayExpression` objects.
func (resource QueryEditorArrayExpression) Equals(other QueryEditorArrayExpression) bool {
	if resource.Type != other.Type {
		return false
	}

	if len(resource.Expressions) != len(other.Expressions) {
		return false
	}

	for i1 := range resource.Expressions {
		if !resource.Expressions[i1].Equals(other.Expressions[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorArrayExpression` fields for violations and returns them.
func (resource QueryEditorArrayExpression) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Expressions {
		if err := resource.Expressions[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("expressions["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource CloudWatchLogsQuery) ImplementsDataqueryVariant() {}

func (resource CloudWatchLogsQuery) DataqueryType() string {
	return "cloudwatch"
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchLogsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CloudWatchLogsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "queryMode"
	if fields["queryMode"] != nil {
		if string(fields["queryMode"]) != "null" {
			if err := json.Unmarshal(fields["queryMode"], &resource.QueryMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("queryMode", errors.New("required field is null"))...)

		}
		delete(fields, "queryMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("queryMode", errors.New("required field is missing from input"))...)
	}
	// Field "id"
	if fields["id"] != nil {
		if string(fields["id"]) != "null" {
			if err := json.Unmarshal(fields["id"], &resource.Id); err != nil {
				errs = append(errs, cog.MakeBuildErrors("id", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is null"))...)

		}
		delete(fields, "id")
	} else {
		errs = append(errs, cog.MakeBuildErrors("id", errors.New("required field is missing from input"))...)
	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is null"))...)

		}
		delete(fields, "region")
	} else {
		errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is missing from input"))...)
	}
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}

		}
		delete(fields, "expression")

	}
	// Field "statsGroups"
	if fields["statsGroups"] != nil {
		if string(fields["statsGroups"]) != "null" {

			if err := json.Unmarshal(fields["statsGroups"], &resource.StatsGroups); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statsGroups", err)...)
			}

		}
		delete(fields, "statsGroups")

	}
	// Field "logGroups"
	if fields["logGroups"] != nil {
		if string(fields["logGroups"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["logGroups"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 LogGroup

				result1 = LogGroup{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("logGroups["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.LogGroups = append(resource.LogGroups, result1)
			}

		}
		delete(fields, "logGroups")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "logGroupNames"
	if fields["logGroupNames"] != nil {
		if string(fields["logGroupNames"]) != "null" {

			if err := json.Unmarshal(fields["logGroupNames"], &resource.LogGroupNames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("logGroupNames", err)...)
			}

		}
		delete(fields, "logGroupNames")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CloudWatchLogsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
func (resource CloudWatchLogsQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(CloudWatchLogsQuery)
	if !ok {
		return false
	}
	if resource.QueryMode != other.QueryMode {
		return false
	}
	if resource.Id != other.Id {
		return false
	}
	if resource.Region != other.Region {
		return false
	}
	if resource.Expression == nil && other.Expression != nil || resource.Expression != nil && other.Expression == nil {
		return false
	}

	if resource.Expression != nil {
		if *resource.Expression != *other.Expression {
			return false
		}
	}

	if len(resource.StatsGroups) != len(other.StatsGroups) {
		return false
	}

	for i1 := range resource.StatsGroups {
		if resource.StatsGroups[i1] != other.StatsGroups[i1] {
			return false
		}
	}

	if len(resource.LogGroups) != len(other.LogGroups) {
		return false
	}

	for i1 := range resource.LogGroups {
		if !resource.LogGroups[i1].Equals(other.LogGroups[i1]) {
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

	if len(resource.LogGroupNames) != len(other.LogGroupNames) {
		return false
	}

	for i1 := range resource.LogGroupNames {
		if resource.LogGroupNames[i1] != other.LogGroupNames[i1] {
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

// Validate checks all the validation constraints that may be defined on `CloudWatchLogsQuery` fields for violations and returns them.
func (resource CloudWatchLogsQuery) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.LogGroups {
		if err := resource.LogGroups[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("logGroups["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LogGroup` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LogGroup) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "arn"
	if fields["arn"] != nil {
		if string(fields["arn"]) != "null" {
			if err := json.Unmarshal(fields["arn"], &resource.Arn); err != nil {
				errs = append(errs, cog.MakeBuildErrors("arn", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("arn", errors.New("required field is null"))...)

		}
		delete(fields, "arn")
	} else {
		errs = append(errs, cog.MakeBuildErrors("arn", errors.New("required field is missing from input"))...)
	}
	// Field "name"
	if fields["name"] != nil {
		if string(fields["name"]) != "null" {
			if err := json.Unmarshal(fields["name"], &resource.Name); err != nil {
				errs = append(errs, cog.MakeBuildErrors("name", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is null"))...)

		}
		delete(fields, "name")
	} else {
		errs = append(errs, cog.MakeBuildErrors("name", errors.New("required field is missing from input"))...)
	}
	// Field "accountId"
	if fields["accountId"] != nil {
		if string(fields["accountId"]) != "null" {
			if err := json.Unmarshal(fields["accountId"], &resource.AccountId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("accountId", err)...)
			}

		}
		delete(fields, "accountId")

	}
	// Field "accountLabel"
	if fields["accountLabel"] != nil {
		if string(fields["accountLabel"]) != "null" {
			if err := json.Unmarshal(fields["accountLabel"], &resource.AccountLabel); err != nil {
				errs = append(errs, cog.MakeBuildErrors("accountLabel", err)...)
			}

		}
		delete(fields, "accountLabel")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LogGroup", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LogGroup` objects.
func (resource LogGroup) Equals(other LogGroup) bool {
	if resource.Arn != other.Arn {
		return false
	}
	if resource.Name != other.Name {
		return false
	}
	if resource.AccountId == nil && other.AccountId != nil || resource.AccountId != nil && other.AccountId == nil {
		return false
	}

	if resource.AccountId != nil {
		if *resource.AccountId != *other.AccountId {
			return false
		}
	}
	if resource.AccountLabel == nil && other.AccountLabel != nil || resource.AccountLabel != nil && other.AccountLabel == nil {
		return false
	}

	if resource.AccountLabel != nil {
		if *resource.AccountLabel != *other.AccountLabel {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `LogGroup` fields for violations and returns them.
func (resource LogGroup) Validate() error {
	return nil
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
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// @deprecated use statistic
	Statistics []string `json:"statistics,omitempty"`
}

func (resource CloudWatchAnnotationQuery) ImplementsDataqueryVariant() {}

func (resource CloudWatchAnnotationQuery) DataqueryType() string {
	return "cloudwatch"
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchAnnotationQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CloudWatchAnnotationQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "queryMode"
	if fields["queryMode"] != nil {
		if string(fields["queryMode"]) != "null" {
			if err := json.Unmarshal(fields["queryMode"], &resource.QueryMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("queryMode", errors.New("required field is null"))...)

		}
		delete(fields, "queryMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("queryMode", errors.New("required field is missing from input"))...)
	}
	// Field "prefixMatching"
	if fields["prefixMatching"] != nil {
		if string(fields["prefixMatching"]) != "null" {
			if err := json.Unmarshal(fields["prefixMatching"], &resource.PrefixMatching); err != nil {
				errs = append(errs, cog.MakeBuildErrors("prefixMatching", err)...)
			}

		}
		delete(fields, "prefixMatching")

	}
	// Field "actionPrefix"
	if fields["actionPrefix"] != nil {
		if string(fields["actionPrefix"]) != "null" {
			if err := json.Unmarshal(fields["actionPrefix"], &resource.ActionPrefix); err != nil {
				errs = append(errs, cog.MakeBuildErrors("actionPrefix", err)...)
			}

		}
		delete(fields, "actionPrefix")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "region"
	if fields["region"] != nil {
		if string(fields["region"]) != "null" {
			if err := json.Unmarshal(fields["region"], &resource.Region); err != nil {
				errs = append(errs, cog.MakeBuildErrors("region", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is null"))...)

		}
		delete(fields, "region")
	} else {
		errs = append(errs, cog.MakeBuildErrors("region", errors.New("required field is missing from input"))...)
	}
	// Field "namespace"
	if fields["namespace"] != nil {
		if string(fields["namespace"]) != "null" {
			if err := json.Unmarshal(fields["namespace"], &resource.Namespace); err != nil {
				errs = append(errs, cog.MakeBuildErrors("namespace", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("namespace", errors.New("required field is null"))...)

		}
		delete(fields, "namespace")
	} else {
		errs = append(errs, cog.MakeBuildErrors("namespace", errors.New("required field is missing from input"))...)
	}
	// Field "metricName"
	if fields["metricName"] != nil {
		if string(fields["metricName"]) != "null" {
			if err := json.Unmarshal(fields["metricName"], &resource.MetricName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricName", err)...)
			}

		}
		delete(fields, "metricName")

	}
	// Field "dimensions"
	if fields["dimensions"] != nil {
		if string(fields["dimensions"]) != "null" {

			partialMap := make(map[string]json.RawMessage)
			if err := json.Unmarshal(fields["dimensions"], &partialMap); err != nil {
				return err
			}
			parsedMap1 := make(map[string]StringOrArrayOfString, len(partialMap))
			for key1 := range partialMap {
				var result1 StringOrArrayOfString

				result1 = StringOrArrayOfString{}
				if err := result1.UnmarshalJSONStrict(partialMap[key1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("dimensions["+key1+"]", err)...)
				}
				parsedMap1[key1] = result1
			}
			resource.Dimensions = cog.ToPtr(Dimensions(parsedMap1))

		}
		delete(fields, "dimensions")

	}
	// Field "matchExact"
	if fields["matchExact"] != nil {
		if string(fields["matchExact"]) != "null" {
			if err := json.Unmarshal(fields["matchExact"], &resource.MatchExact); err != nil {
				errs = append(errs, cog.MakeBuildErrors("matchExact", err)...)
			}

		}
		delete(fields, "matchExact")

	}
	// Field "period"
	if fields["period"] != nil {
		if string(fields["period"]) != "null" {
			if err := json.Unmarshal(fields["period"], &resource.Period); err != nil {
				errs = append(errs, cog.MakeBuildErrors("period", err)...)
			}

		}
		delete(fields, "period")

	}
	// Field "accountId"
	if fields["accountId"] != nil {
		if string(fields["accountId"]) != "null" {
			if err := json.Unmarshal(fields["accountId"], &resource.AccountId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("accountId", err)...)
			}

		}
		delete(fields, "accountId")

	}
	// Field "statistic"
	if fields["statistic"] != nil {
		if string(fields["statistic"]) != "null" {
			if err := json.Unmarshal(fields["statistic"], &resource.Statistic); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statistic", err)...)
			}

		}
		delete(fields, "statistic")

	}
	// Field "alarmNamePrefix"
	if fields["alarmNamePrefix"] != nil {
		if string(fields["alarmNamePrefix"]) != "null" {
			if err := json.Unmarshal(fields["alarmNamePrefix"], &resource.AlarmNamePrefix); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alarmNamePrefix", err)...)
			}

		}
		delete(fields, "alarmNamePrefix")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "statistics"
	if fields["statistics"] != nil {
		if string(fields["statistics"]) != "null" {

			if err := json.Unmarshal(fields["statistics"], &resource.Statistics); err != nil {
				errs = append(errs, cog.MakeBuildErrors("statistics", err)...)
			}

		}
		delete(fields, "statistics")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CloudWatchAnnotationQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
func (resource CloudWatchAnnotationQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(CloudWatchAnnotationQuery)
	if !ok {
		return false
	}
	if resource.QueryMode != other.QueryMode {
		return false
	}
	if resource.PrefixMatching == nil && other.PrefixMatching != nil || resource.PrefixMatching != nil && other.PrefixMatching == nil {
		return false
	}

	if resource.PrefixMatching != nil {
		if *resource.PrefixMatching != *other.PrefixMatching {
			return false
		}
	}
	if resource.ActionPrefix == nil && other.ActionPrefix != nil || resource.ActionPrefix != nil && other.ActionPrefix == nil {
		return false
	}

	if resource.ActionPrefix != nil {
		if *resource.ActionPrefix != *other.ActionPrefix {
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
	if resource.Region != other.Region {
		return false
	}
	if resource.Namespace != other.Namespace {
		return false
	}
	if resource.MetricName == nil && other.MetricName != nil || resource.MetricName != nil && other.MetricName == nil {
		return false
	}

	if resource.MetricName != nil {
		if *resource.MetricName != *other.MetricName {
			return false
		}
	}
	if resource.Dimensions == nil && other.Dimensions != nil || resource.Dimensions != nil && other.Dimensions == nil {
		return false
	}

	if resource.Dimensions != nil {

		if len(*resource.Dimensions) != len(*other.Dimensions) {
			return false
		}

		for key1 := range *resource.Dimensions {
			if !(*resource.Dimensions)[key1].Equals((*other.Dimensions)[key1]) {
				return false
			}
		}
	}
	if resource.MatchExact == nil && other.MatchExact != nil || resource.MatchExact != nil && other.MatchExact == nil {
		return false
	}

	if resource.MatchExact != nil {
		if *resource.MatchExact != *other.MatchExact {
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
	if resource.AccountId == nil && other.AccountId != nil || resource.AccountId != nil && other.AccountId == nil {
		return false
	}

	if resource.AccountId != nil {
		if *resource.AccountId != *other.AccountId {
			return false
		}
	}
	if resource.Statistic == nil && other.Statistic != nil || resource.Statistic != nil && other.Statistic == nil {
		return false
	}

	if resource.Statistic != nil {
		if *resource.Statistic != *other.Statistic {
			return false
		}
	}
	if resource.AlarmNamePrefix == nil && other.AlarmNamePrefix != nil || resource.AlarmNamePrefix != nil && other.AlarmNamePrefix == nil {
		return false
	}

	if resource.AlarmNamePrefix != nil {
		if *resource.AlarmNamePrefix != *other.AlarmNamePrefix {
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

	if len(resource.Statistics) != len(other.Statistics) {
		return false
	}

	for i1 := range resource.Statistics {
		if resource.Statistics[i1] != other.Statistics[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CloudWatchAnnotationQuery` fields for violations and returns them.
func (resource CloudWatchAnnotationQuery) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CloudWatchQuery = CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery

// VariantConfig returns the configuration related to cloudwatch dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "cloudwatch",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &CloudWatchQuery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &CloudWatchQuery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery CloudWatchQuery
			if cast, ok := input.(*CloudWatchQuery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(CloudWatchQuery)
			}

			if dataquery.CloudWatchMetricsQuery != nil {
				return CloudWatchMetricsQueryConverter(*dataquery.CloudWatchMetricsQuery)
			}
			if dataquery.CloudWatchLogsQuery != nil {
				return CloudWatchLogsQueryConverter(*dataquery.CloudWatchLogsQuery)
			}
			if dataquery.CloudWatchAnnotationQuery != nil {
				return CloudWatchAnnotationQueryConverter(*dataquery.CloudWatchAnnotationQuery)
			}

			return ""
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

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrArrayOfString` as JSON.
func (resource StringOrArrayOfString) MarshalJSON() ([]byte, error) {
	if resource.String != nil {
		return json.Marshal(resource.String)
	}

	if resource.ArrayOfString != nil {
		return json.Marshal(resource.ArrayOfString)
	}

	return nil, fmt.Errorf("no value for disjunction of scalars")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrArrayOfString` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrArrayOfString) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string

	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// ArrayOfString
	var ArrayOfString []string

	if err := json.Unmarshal(raw, &ArrayOfString); err != nil {
		errList = append(errList, err)
	} else {
		resource.ArrayOfString = ArrayOfString
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrArrayOfString", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrArrayOfString` objects.
func (resource StringOrArrayOfString) Equals(other StringOrArrayOfString) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}

	if len(resource.ArrayOfString) != len(other.ArrayOfString) {
		return false
	}

	for i1 := range resource.ArrayOfString {
		if resource.ArrayOfString[i1] != other.ArrayOfString[i1] {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrArrayOfString` fields for violations and returns them.
func (resource StringOrArrayOfString) Validate() error {
	return nil
}

type QueryEditorPropertyExpressionOrQueryEditorFunctionExpression struct {
	QueryEditorPropertyExpression *QueryEditorPropertyExpression `json:"QueryEditorPropertyExpression,omitempty"`
	QueryEditorFunctionExpression *QueryEditorFunctionExpression `json:"QueryEditorFunctionExpression,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` as JSON.
func (resource QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) MarshalJSON() ([]byte, error) {
	if resource.QueryEditorPropertyExpression != nil {
		return json.Marshal(resource.QueryEditorPropertyExpression)
	}
	if resource.QueryEditorFunctionExpression != nil {
		return json.Marshal(resource.QueryEditorFunctionExpression)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) UnmarshalJSONStrict(raw []byte) error {
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
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "function":
		queryEditorFunctionExpression := &QueryEditorFunctionExpression{}
		if err := queryEditorFunctionExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorFunctionExpression = queryEditorFunctionExpression
		return nil
	case "property":
		queryEditorPropertyExpression := &QueryEditorPropertyExpression{}
		if err := queryEditorPropertyExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorPropertyExpression = queryEditorPropertyExpression
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` objects.
func (resource QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) Equals(other QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) bool {
	if resource.QueryEditorPropertyExpression == nil && other.QueryEditorPropertyExpression != nil || resource.QueryEditorPropertyExpression != nil && other.QueryEditorPropertyExpression == nil {
		return false
	}

	if resource.QueryEditorPropertyExpression != nil {
		if !resource.QueryEditorPropertyExpression.Equals(*other.QueryEditorPropertyExpression) {
			return false
		}
	}
	if resource.QueryEditorFunctionExpression == nil && other.QueryEditorFunctionExpression != nil || resource.QueryEditorFunctionExpression != nil && other.QueryEditorFunctionExpression == nil {
		return false
	}

	if resource.QueryEditorFunctionExpression != nil {
		if !resource.QueryEditorFunctionExpression.Equals(*other.QueryEditorFunctionExpression) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorPropertyExpressionOrQueryEditorFunctionExpression` fields for violations and returns them.
func (resource QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) Validate() error {
	var errs cog.BuildErrors
	if resource.QueryEditorPropertyExpression != nil {
		if err := resource.QueryEditorPropertyExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorPropertyExpression", err)...)
		}
	}
	if resource.QueryEditorFunctionExpression != nil {
		if err := resource.QueryEditorFunctionExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionExpression", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType struct {
	String                         *string                   `json:"String,omitempty"`
	Bool                           *bool                     `json:"Bool,omitempty"`
	Int64                          *int64                    `json:"Int64,omitempty"`
	ArrayOfQueryEditorOperatorType []QueryEditorOperatorType `json:"ArrayOfQueryEditorOperatorType,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` as JSON.
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

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string

	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool

	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Int64
	var Int64 int64

	if err := json.Unmarshal(raw, &Int64); err != nil {
		errList = append(errList, err)
	} else {
		resource.Int64 = &Int64
		return nil
	}

	// ArrayOfQueryEditorOperatorType
	var ArrayOfQueryEditorOperatorType []QueryEditorOperatorType

	if err := json.Unmarshal(raw, &ArrayOfQueryEditorOperatorType); err != nil {
		errList = append(errList, err)
	} else {
		resource.ArrayOfQueryEditorOperatorType = ArrayOfQueryEditorOperatorType
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` objects.
func (resource StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) Equals(other StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.Bool == nil && other.Bool != nil || resource.Bool != nil && other.Bool == nil {
		return false
	}

	if resource.Bool != nil {
		if *resource.Bool != *other.Bool {
			return false
		}
	}
	if resource.Int64 == nil && other.Int64 != nil || resource.Int64 != nil && other.Int64 == nil {
		return false
	}

	if resource.Int64 != nil {
		if *resource.Int64 != *other.Int64 {
			return false
		}
	}

	if len(resource.ArrayOfQueryEditorOperatorType) != len(other.ArrayOfQueryEditorOperatorType) {
		return false
	}

	for i1 := range resource.ArrayOfQueryEditorOperatorType {
		if !resource.ArrayOfQueryEditorOperatorType[i1].Equals(other.ArrayOfQueryEditorOperatorType[i1]) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType` fields for violations and returns them.
func (resource StringOrBoolOrInt64OrArrayOfQueryEditorOperatorType) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.ArrayOfQueryEditorOperatorType {
		if err := resource.ArrayOfQueryEditorOperatorType[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("ArrayOfQueryEditorOperatorType["+strconv.Itoa(i1)+"]", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type StringOrBoolOrInt64 struct {
	String *string `json:"String,omitempty"`
	Bool   *bool   `json:"Bool,omitempty"`
	Int64  *int64  `json:"Int64,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `StringOrBoolOrInt64` as JSON.
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

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StringOrBoolOrInt64` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *StringOrBoolOrInt64) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors
	var errList []error

	// String
	var String string

	if err := json.Unmarshal(raw, &String); err != nil {
		errList = append(errList, err)
	} else {
		resource.String = &String
		return nil
	}

	// Bool
	var Bool bool

	if err := json.Unmarshal(raw, &Bool); err != nil {
		errList = append(errList, err)
	} else {
		resource.Bool = &Bool
		return nil
	}

	// Int64
	var Int64 int64

	if err := json.Unmarshal(raw, &Int64); err != nil {
		errList = append(errList, err)
	} else {
		resource.Int64 = &Int64
		return nil
	}

	if len(errList) != 0 {
		errs = append(errs, cog.MakeBuildErrors("StringOrBoolOrInt64", errors.Join(errList...))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `StringOrBoolOrInt64` objects.
func (resource StringOrBoolOrInt64) Equals(other StringOrBoolOrInt64) bool {
	if resource.String == nil && other.String != nil || resource.String != nil && other.String == nil {
		return false
	}

	if resource.String != nil {
		if *resource.String != *other.String {
			return false
		}
	}
	if resource.Bool == nil && other.Bool != nil || resource.Bool != nil && other.Bool == nil {
		return false
	}

	if resource.Bool != nil {
		if *resource.Bool != *other.Bool {
			return false
		}
	}
	if resource.Int64 == nil && other.Int64 != nil || resource.Int64 != nil && other.Int64 == nil {
		return false
	}

	if resource.Int64 != nil {
		if *resource.Int64 != *other.Int64 {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `StringOrBoolOrInt64` fields for violations and returns them.
func (resource StringOrBoolOrInt64) Validate() error {
	return nil
}

type QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression struct {
	QueryEditorArrayExpression             *QueryEditorArrayExpression             `json:"QueryEditorArrayExpression,omitempty"`
	QueryEditorPropertyExpression          *QueryEditorPropertyExpression          `json:"QueryEditorPropertyExpression,omitempty"`
	QueryEditorGroupByExpression           *QueryEditorGroupByExpression           `json:"QueryEditorGroupByExpression,omitempty"`
	QueryEditorFunctionExpression          *QueryEditorFunctionExpression          `json:"QueryEditorFunctionExpression,omitempty"`
	QueryEditorFunctionParameterExpression *QueryEditorFunctionParameterExpression `json:"QueryEditorFunctionParameterExpression,omitempty"`
	QueryEditorOperatorExpression          *QueryEditorOperatorExpression          `json:"QueryEditorOperatorExpression,omitempty"`
}

// MarshalJSON implements a custom JSON marshalling logic to encode `QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression` as JSON.
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

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression) UnmarshalJSONStrict(raw []byte) error {
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
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "and":
		queryEditorArrayExpression := &QueryEditorArrayExpression{}
		if err := queryEditorArrayExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorArrayExpression = queryEditorArrayExpression
		return nil
	case "function":
		queryEditorFunctionExpression := &QueryEditorFunctionExpression{}
		if err := queryEditorFunctionExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorFunctionExpression = queryEditorFunctionExpression
		return nil
	case "functionParameter":
		queryEditorFunctionParameterExpression := &QueryEditorFunctionParameterExpression{}
		if err := queryEditorFunctionParameterExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorFunctionParameterExpression = queryEditorFunctionParameterExpression
		return nil
	case "groupBy":
		queryEditorGroupByExpression := &QueryEditorGroupByExpression{}
		if err := queryEditorGroupByExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorGroupByExpression = queryEditorGroupByExpression
		return nil
	case "operator":
		queryEditorOperatorExpression := &QueryEditorOperatorExpression{}
		if err := queryEditorOperatorExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorOperatorExpression = queryEditorOperatorExpression
		return nil
	case "or":
		queryEditorArrayExpression := &QueryEditorArrayExpression{}
		if err := queryEditorArrayExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorArrayExpression = queryEditorArrayExpression
		return nil
	case "property":
		queryEditorPropertyExpression := &QueryEditorPropertyExpression{}
		if err := queryEditorPropertyExpression.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.QueryEditorPropertyExpression = queryEditorPropertyExpression
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two `QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression` objects.
func (resource QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression) Equals(other QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression) bool {
	if resource.QueryEditorArrayExpression == nil && other.QueryEditorArrayExpression != nil || resource.QueryEditorArrayExpression != nil && other.QueryEditorArrayExpression == nil {
		return false
	}

	if resource.QueryEditorArrayExpression != nil {
		if !resource.QueryEditorArrayExpression.Equals(*other.QueryEditorArrayExpression) {
			return false
		}
	}
	if resource.QueryEditorPropertyExpression == nil && other.QueryEditorPropertyExpression != nil || resource.QueryEditorPropertyExpression != nil && other.QueryEditorPropertyExpression == nil {
		return false
	}

	if resource.QueryEditorPropertyExpression != nil {
		if !resource.QueryEditorPropertyExpression.Equals(*other.QueryEditorPropertyExpression) {
			return false
		}
	}
	if resource.QueryEditorGroupByExpression == nil && other.QueryEditorGroupByExpression != nil || resource.QueryEditorGroupByExpression != nil && other.QueryEditorGroupByExpression == nil {
		return false
	}

	if resource.QueryEditorGroupByExpression != nil {
		if !resource.QueryEditorGroupByExpression.Equals(*other.QueryEditorGroupByExpression) {
			return false
		}
	}
	if resource.QueryEditorFunctionExpression == nil && other.QueryEditorFunctionExpression != nil || resource.QueryEditorFunctionExpression != nil && other.QueryEditorFunctionExpression == nil {
		return false
	}

	if resource.QueryEditorFunctionExpression != nil {
		if !resource.QueryEditorFunctionExpression.Equals(*other.QueryEditorFunctionExpression) {
			return false
		}
	}
	if resource.QueryEditorFunctionParameterExpression == nil && other.QueryEditorFunctionParameterExpression != nil || resource.QueryEditorFunctionParameterExpression != nil && other.QueryEditorFunctionParameterExpression == nil {
		return false
	}

	if resource.QueryEditorFunctionParameterExpression != nil {
		if !resource.QueryEditorFunctionParameterExpression.Equals(*other.QueryEditorFunctionParameterExpression) {
			return false
		}
	}
	if resource.QueryEditorOperatorExpression == nil && other.QueryEditorOperatorExpression != nil || resource.QueryEditorOperatorExpression != nil && other.QueryEditorOperatorExpression == nil {
		return false
	}

	if resource.QueryEditorOperatorExpression != nil {
		if !resource.QueryEditorOperatorExpression.Equals(*other.QueryEditorOperatorExpression) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression` fields for violations and returns them.
func (resource QueryEditorArrayExpressionOrQueryEditorPropertyExpressionOrQueryEditorGroupByExpressionOrQueryEditorFunctionExpressionOrQueryEditorFunctionParameterExpressionOrQueryEditorOperatorExpression) Validate() error {
	var errs cog.BuildErrors
	if resource.QueryEditorArrayExpression != nil {
		if err := resource.QueryEditorArrayExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorArrayExpression", err)...)
		}
	}
	if resource.QueryEditorPropertyExpression != nil {
		if err := resource.QueryEditorPropertyExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorPropertyExpression", err)...)
		}
	}
	if resource.QueryEditorGroupByExpression != nil {
		if err := resource.QueryEditorGroupByExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorGroupByExpression", err)...)
		}
	}
	if resource.QueryEditorFunctionExpression != nil {
		if err := resource.QueryEditorFunctionExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionExpression", err)...)
		}
	}
	if resource.QueryEditorFunctionParameterExpression != nil {
		if err := resource.QueryEditorFunctionParameterExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorFunctionParameterExpression", err)...)
		}
	}
	if resource.QueryEditorOperatorExpression != nil {
		if err := resource.QueryEditorOperatorExpression.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("QueryEditorOperatorExpression", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

type CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery struct {
	CloudWatchMetricsQuery    *CloudWatchMetricsQuery    `json:"CloudWatchMetricsQuery,omitempty"`
	CloudWatchLogsQuery       *CloudWatchLogsQuery       `json:"CloudWatchLogsQuery,omitempty"`
	CloudWatchAnnotationQuery *CloudWatchAnnotationQuery `json:"CloudWatchAnnotationQuery,omitempty"`
}

func (resource CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) ImplementsDataqueryVariant() {
}

func (resource CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) DataqueryType() string {
	return "cloudwatch"
}

// MarshalJSON implements a custom JSON marshalling logic to encode `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` as JSON.
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

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) UnmarshalJSONStrict(raw []byte) error {
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
		return fmt.Errorf("discriminator field 'queryMode' not found in payload")
	}

	switch discriminator {
	case "Annotations":
		cloudWatchAnnotationQuery := &CloudWatchAnnotationQuery{}
		if err := cloudWatchAnnotationQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CloudWatchAnnotationQuery = cloudWatchAnnotationQuery
		return nil
	case "Logs":
		cloudWatchLogsQuery := &CloudWatchLogsQuery{}
		if err := cloudWatchLogsQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CloudWatchLogsQuery = cloudWatchLogsQuery
		return nil
	case "Metrics":
		cloudWatchMetricsQuery := &CloudWatchMetricsQuery{}
		if err := cloudWatchMetricsQuery.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.CloudWatchMetricsQuery = cloudWatchMetricsQuery
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `queryMode = %v`", discriminator)
}

// Equals tests the equality of two dataqueries.
func (resource CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery)
	if !ok {
		return false
	}
	if resource.CloudWatchMetricsQuery == nil && other.CloudWatchMetricsQuery != nil || resource.CloudWatchMetricsQuery != nil && other.CloudWatchMetricsQuery == nil {
		return false
	}

	if resource.CloudWatchMetricsQuery != nil {
		if !resource.CloudWatchMetricsQuery.Equals(*other.CloudWatchMetricsQuery) {
			return false
		}
	}
	if resource.CloudWatchLogsQuery == nil && other.CloudWatchLogsQuery != nil || resource.CloudWatchLogsQuery != nil && other.CloudWatchLogsQuery == nil {
		return false
	}

	if resource.CloudWatchLogsQuery != nil {
		if !resource.CloudWatchLogsQuery.Equals(*other.CloudWatchLogsQuery) {
			return false
		}
	}
	if resource.CloudWatchAnnotationQuery == nil && other.CloudWatchAnnotationQuery != nil || resource.CloudWatchAnnotationQuery != nil && other.CloudWatchAnnotationQuery == nil {
		return false
	}

	if resource.CloudWatchAnnotationQuery != nil {
		if !resource.CloudWatchAnnotationQuery.Equals(*other.CloudWatchAnnotationQuery) {
			return false
		}
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery` fields for violations and returns them.
func (resource CloudWatchMetricsQueryOrCloudWatchLogsQueryOrCloudWatchAnnotationQuery) Validate() error {
	var errs cog.BuildErrors
	if resource.CloudWatchMetricsQuery != nil {
		if err := resource.CloudWatchMetricsQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CloudWatchMetricsQuery", err)...)
		}
	}
	if resource.CloudWatchLogsQuery != nil {
		if err := resource.CloudWatchLogsQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CloudWatchLogsQuery", err)...)
		}
	}
	if resource.CloudWatchAnnotationQuery != nil {
		if err := resource.CloudWatchAnnotationQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("CloudWatchAnnotationQuery", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
