// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
import * as dashboard from '../dashboard';

// Shape of a CloudWatch Metrics query
export class CloudWatchMetricsQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: cloudwatch.CloudWatchMetricsQuery;

    constructor() {
        this.internal = cloudwatch.defaultCloudWatchMetricsQuery();
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.CloudWatchMetricsQuery {
        return this.internal;
    }

    // Whether a query is a Metrics, Logs, or Annotations query
    queryMode(queryMode: cloudwatch.CloudWatchQueryMode): this {
        this.internal.queryMode = queryMode;
        return this;
    }

    // Whether to use a metric search or metric query. Metric query is referred to as "Metrics Insights" in the AWS console.
    metricQueryType(metricQueryType: cloudwatch.MetricQueryType): this {
        this.internal.metricQueryType = metricQueryType;
        return this;
    }

    // Whether to use the query builder or code editor to create the query
    metricEditorMode(metricEditorMode: cloudwatch.MetricEditorMode): this {
        this.internal.metricEditorMode = metricEditorMode;
        return this;
    }

    // ID can be used to reference other queries in math expressions. The ID can include numbers, letters, and underscore, and must start with a lowercase letter.
    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    // Deprecated: use label
    // @deprecated use label
    alias(alias: string): this {
        this.internal.alias = alias;
        return this;
    }

    // Change the time series legend names using dynamic labels. See https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph-dynamic-labels.html for more details.
    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    // Math expression query
    expression(expression: string): this {
        this.internal.expression = expression;
        return this;
    }

    // When the metric query type is `metricQueryType` is set to `Query`, this field is used to specify the query string.
    sqlExpression(sqlExpression: string): this {
        this.internal.sqlExpression = sqlExpression;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // AWS region to query for the metric
    region(region: string): this {
        this.internal.region = region;
        return this;
    }

    // A namespace is a container for CloudWatch metrics. Metrics in different namespaces are isolated from each other, so that metrics from different applications are not mistakenly aggregated into the same statistics. For example, Amazon EC2 uses the AWS/EC2 namespace.
    namespace(namespace: string): this {
        this.internal.namespace = namespace;
        return this;
    }

    // Name of the metric
    metricName(metricName: string): this {
        this.internal.metricName = metricName;
        return this;
    }

    // The dimensions of the metric
    dimensions(dimensions: cloudwatch.Dimensions): this {
        this.internal.dimensions = dimensions;
        return this;
    }

    // Only show metrics that exactly match all defined dimension names.
    matchExact(matchExact: boolean): this {
        this.internal.matchExact = matchExact;
        return this;
    }

    // The length of time associated with a specific Amazon CloudWatch statistic. Can be specified by a number of seconds, 'auto', or as a duration string e.g. '15m' being 15 minutes
    period(period: string): this {
        this.internal.period = period;
        return this;
    }

    // The ID of the AWS account to query for the metric, specifying `all` will query all accounts that the monitoring account is permitted to query.
    accountId(accountId: string): this {
        this.internal.accountId = accountId;
        return this;
    }

    // Metric data aggregations over specified periods of time. For detailed definitions of the statistics supported by CloudWatch, see https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html.
    statistic(statistic: string): this {
        this.internal.statistic = statistic;
        return this;
    }

    // When the metric query type is `metricQueryType` is set to `Query` and the `metricEditorMode` is set to `Builder`, this field is used to build up an object representation of a SQL query.
    sql(sql: cog.Builder<cloudwatch.SQLExpression>): this {
        const sqlResource = sql.build();
        this.internal.sql = sqlResource;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: dashboard.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }

    // @deprecated use statistic
    statistics(statistics: string[]): this {
        this.internal.statistics = statistics;
        return this;
    }
}
