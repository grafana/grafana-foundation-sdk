// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export interface MetricStat {
	// AWS region to query for the metric
	region: string;
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	namespace: string;
	// Name of the metric
	metricName?: string;
	// The dimensions of the metric
	dimensions?: Dimensions;
	// Only show metrics that exactly match all defined dimension names.
	matchExact?: boolean;
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	period?: string;
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	accountId?: string;
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	statistic?: string;
	// @deprecated use statistic
	statistics?: string[];
}

export const defaultMetricStat = (): MetricStat => ({
	region: "",
	namespace: "",
});

// A name/value pair that is part of the identity of a metric. For example, you can get statistics for a specific EC2 instance by specifying the InstanceId dimension when you search for metrics.
export type Dimensions = Record<string, string | string[]>;

export const defaultDimensions = (): Dimensions => ({});

// Shape of a CloudWatch Metrics query
export interface CloudWatchMetricsQuery {
	// Whether a query is a Metrics, Logs, or Annotations query
	queryMode: CloudWatchQueryMode;
	// Whether to use a metric search or metric insights query
	metricQueryType?: MetricQueryType;
	// Whether to use the query builder or code editor to create the query
	metricEditorMode?: MetricEditorMode;
	// ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
	id: string;
	// Deprecated: use label
	// @deprecated use label
	alias?: string;
	// Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
	label?: string;
	// Math expression query
	expression?: string;
	// When the metric query type is set to `Insights`, this field is used to specify the query string.
	sqlExpression?: string;
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// AWS region to query for the metric
	region: string;
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	namespace: string;
	// Name of the metric
	metricName?: string;
	// The dimensions of the metric
	dimensions?: Dimensions;
	// Only show metrics that exactly match all defined dimension names.
	matchExact?: boolean;
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	period?: string;
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	accountId?: string;
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	statistic?: string;
	// When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
	sql?: SQLExpression;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// @deprecated use statistic
	statistics?: string[];
	_implementsDataqueryVariant(): void;
}

export const defaultCloudWatchMetricsQuery = (): CloudWatchMetricsQuery => ({
	queryMode: CloudWatchQueryMode.Metrics,
	id: "",
	refId: "",
	region: "",
	namespace: "",
	_implementsDataqueryVariant: () => {},
});

export enum CloudWatchQueryMode {
	Metrics = "Metrics",
	Logs = "Logs",
	Annotations = "Annotations",
}

export const defaultCloudWatchQueryMode = (): CloudWatchQueryMode => (CloudWatchQueryMode.Metrics);

export enum MetricQueryType {
	Search = 0,
	Insights = 1,
}

export const defaultMetricQueryType = (): MetricQueryType => (MetricQueryType.Search);

export enum MetricEditorMode {
	Builder = 0,
	Code = 1,
}

export const defaultMetricEditorMode = (): MetricEditorMode => (MetricEditorMode.Builder);

export interface SQLExpression {
	// SELECT part of the SQL expression
	select?: QueryEditorFunctionExpression;
	// FROM part of the SQL expression
	from?: QueryEditorPropertyExpression | QueryEditorFunctionExpression;
	// WHERE part of the SQL expression
	where?: QueryEditorArrayExpression;
	// GROUP BY part of the SQL expression
	groupBy?: QueryEditorArrayExpression;
	// ORDER BY part of the SQL expression
	orderBy?: QueryEditorFunctionExpression;
	// The sort order of the SQL expression, `ASC` or `DESC`
	orderByDirection?: string;
	// LIMIT part of the SQL expression
	limit?: number;
}

export const defaultSQLExpression = (): SQLExpression => ({
});

export interface QueryEditorFunctionExpression {
	type: QueryEditorExpressionType.Function;
	name?: string;
	parameters?: QueryEditorFunctionParameterExpression[];
}

export const defaultQueryEditorFunctionExpression = (): QueryEditorFunctionExpression => ({
	type: QueryEditorExpressionType.Function,
});

export enum QueryEditorExpressionType {
	Property = "property",
	Operator = "operator",
	Or = "or",
	And = "and",
	GroupBy = "groupBy",
	Function = "function",
	FunctionParameter = "functionParameter",
}

export const defaultQueryEditorExpressionType = (): QueryEditorExpressionType => (QueryEditorExpressionType.Property);

export interface QueryEditorFunctionParameterExpression {
	type: QueryEditorExpressionType.FunctionParameter;
	name?: string;
}

export const defaultQueryEditorFunctionParameterExpression = (): QueryEditorFunctionParameterExpression => ({
	type: QueryEditorExpressionType.FunctionParameter,
});

export interface QueryEditorPropertyExpression {
	type: QueryEditorExpressionType.Property;
	property: QueryEditorProperty;
}

export const defaultQueryEditorPropertyExpression = (): QueryEditorPropertyExpression => ({
	type: QueryEditorExpressionType.Property,
	property: defaultQueryEditorProperty(),
});

export interface QueryEditorProperty {
	type: QueryEditorPropertyType.String;
	name?: string;
}

export const defaultQueryEditorProperty = (): QueryEditorProperty => ({
	type: QueryEditorPropertyType.String,
});

export interface QueryEditorArrayExpression {
	type: "and" | "or";
	expressions: QueryEditorExpression[];
}

export const defaultQueryEditorArrayExpression = (): QueryEditorArrayExpression => ({
	type: "and",
	expressions: [],
});

export type QueryEditorExpression = QueryEditorArrayExpression | QueryEditorPropertyExpression | QueryEditorGroupByExpression | QueryEditorFunctionExpression | QueryEditorFunctionParameterExpression | QueryEditorOperatorExpression;

export const defaultQueryEditorExpression = (): QueryEditorExpression => (defaultQueryEditorArrayExpression());

export interface QueryEditorGroupByExpression {
	type: QueryEditorExpressionType.GroupBy;
	property: QueryEditorProperty;
}

export const defaultQueryEditorGroupByExpression = (): QueryEditorGroupByExpression => ({
	type: QueryEditorExpressionType.GroupBy,
	property: defaultQueryEditorProperty(),
});

export interface QueryEditorOperatorExpression {
	type: QueryEditorExpressionType.Operator;
	property: QueryEditorProperty;
	// TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
	operator: QueryEditorOperator;
}

export const defaultQueryEditorOperatorExpression = (): QueryEditorOperatorExpression => ({
	type: QueryEditorExpressionType.Operator,
	property: defaultQueryEditorProperty(),
	operator: defaultQueryEditorOperator(),
});

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
export interface QueryEditorOperator {
	name?: string;
	value?: QueryEditorOperatorType | QueryEditorOperatorType[];
}

export const defaultQueryEditorOperator = (): QueryEditorOperator => ({
});

export type QueryEditorOperatorType = string | boolean | number;

export const defaultQueryEditorOperatorType = (): QueryEditorOperatorType => ("");

export type QueryEditorOperatorValueType = QueryEditorOperatorType | QueryEditorOperatorType[];

export const defaultQueryEditorOperatorValueType = (): QueryEditorOperatorValueType => (defaultQueryEditorOperatorType());

export enum QueryEditorPropertyType {
	String = "string",
}

export const defaultQueryEditorPropertyType = (): QueryEditorPropertyType => (QueryEditorPropertyType.String);

// Shape of a CloudWatch Logs query
export interface CloudWatchLogsQuery {
	// Whether a query is a Metrics, Logs, or Annotations query
	queryMode: CloudWatchQueryMode;
	id: string;
	// AWS region to query for the logs
	region: string;
	// The CloudWatch Logs Insights query to execute
	expression?: string;
	// Fields to group the results by, this field is automatically populated whenever the query is updated
	statsGroups?: string[];
	// Log groups to query
	logGroups?: LogGroup[];
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// @deprecated use logGroups
	logGroupNames?: string[];
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	_implementsDataqueryVariant(): void;
}

export const defaultCloudWatchLogsQuery = (): CloudWatchLogsQuery => ({
	queryMode: CloudWatchQueryMode.Logs,
	id: "",
	region: "",
	refId: "",
	_implementsDataqueryVariant: () => {},
});

export interface LogGroup {
	// ARN of the log group
	arn: string;
	// Name of the log group
	name: string;
	// AccountId of the log group
	accountId?: string;
	// Label of the log group
	accountLabel?: string;
}

export const defaultLogGroup = (): LogGroup => ({
	arn: "",
	name: "",
});

// Shape of a CloudWatch Annotation query
// TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
// #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
export interface CloudWatchAnnotationQuery {
	// Whether a query is a Metrics, Logs, or Annotations query
	queryMode: CloudWatchQueryMode;
	// Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
	prefixMatching?: boolean;
	// Use this parameter to filter the results of the operation to only those alarms
	// that use a certain alarm action. For example, you could specify the ARN of
	// an SNS topic to find all alarms that send notifications to that topic.
	// e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
	// but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
	actionPrefix?: string;
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// AWS region to query for the metric
	region: string;
	// A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
	namespace: string;
	// Name of the metric
	metricName?: string;
	// The dimensions of the metric
	dimensions?: Dimensions;
	// Only show metrics that exactly match all defined dimension names.
	matchExact?: boolean;
	// The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
	period?: string;
	// The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
	accountId?: string;
	// Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
	statistic?: string;
	// An alarm name prefix. If you specify this parameter, you receive information
	// about all alarms that have names that start with this prefix.
	// e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
	alarmNamePrefix?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	// @deprecated use statistic
	statistics?: string[];
	_implementsDataqueryVariant(): void;
}

export const defaultCloudWatchAnnotationQuery = (): CloudWatchAnnotationQuery => ({
	queryMode: CloudWatchQueryMode.Annotations,
	refId: "",
	region: "",
	namespace: "",
	_implementsDataqueryVariant: () => {},
});

export type CloudWatchQuery = CloudWatchMetricsQuery | CloudWatchLogsQuery | CloudWatchAnnotationQuery;

export const defaultCloudWatchQuery = (): CloudWatchQuery => (defaultCloudWatchMetricsQuery());

