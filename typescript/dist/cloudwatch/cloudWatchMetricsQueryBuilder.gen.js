"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CloudWatchMetricsQueryBuilder = void 0;
const tslib_1 = require("tslib");
const cloudwatch = tslib_1.__importStar(require("../cloudwatch"));
// Shape of a CloudWatch Metrics query
class CloudWatchMetricsQueryBuilder {
    constructor() {
        this.internal = cloudwatch.defaultCloudWatchMetricsQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Whether a query is a Metrics, Logs, or Annotations query
    queryMode(queryMode) {
        this.internal.queryMode = queryMode;
        return this;
    }
    // Whether to use a metric search or metric insights query
    metricQueryType(metricQueryType) {
        this.internal.metricQueryType = metricQueryType;
        return this;
    }
    // Whether to use the query builder or code editor to create the query
    metricEditorMode(metricEditorMode) {
        this.internal.metricEditorMode = metricEditorMode;
        return this;
    }
    // ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
    id(id) {
        this.internal.id = id;
        return this;
    }
    // Deprecated: use label
    // @deprecated use label
    alias(alias) {
        this.internal.alias = alias;
        return this;
    }
    // Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
    label(label) {
        this.internal.label = label;
        return this;
    }
    // Math expression query
    expression(expression) {
        this.internal.expression = expression;
        return this;
    }
    // When the metric query type is set to `Insights`, this field is used to specify the query string.
    sqlExpression(sqlExpression) {
        this.internal.sqlExpression = sqlExpression;
        return this;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        this.internal.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    // AWS region to query for the metric
    region(region) {
        this.internal.region = region;
        return this;
    }
    // A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    namespace(namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    // Name of the metric
    metricName(metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    // The dimensions of the metric
    dimensions(dimensions) {
        this.internal.dimensions = dimensions;
        return this;
    }
    // Only show metrics that exactly match all defined dimension names.
    matchExact(matchExact) {
        this.internal.matchExact = matchExact;
        return this;
    }
    // The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    period(period) {
        this.internal.period = period;
        return this;
    }
    // The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    accountId(accountId) {
        this.internal.accountId = accountId;
        return this;
    }
    // Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    statistic(statistic) {
        this.internal.statistic = statistic;
        return this;
    }
    // When the metric query type is set to `Insights` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
    sql(sql) {
        const sqlResource = sql.build();
        this.internal.sql = sqlResource;
        return this;
    }
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    // @deprecated use statistic
    statistics(statistics) {
        this.internal.statistics = statistics;
        return this;
    }
}
exports.CloudWatchMetricsQueryBuilder = CloudWatchMetricsQueryBuilder;
//# sourceMappingURL=cloudWatchMetricsQueryBuilder.gen.js.map