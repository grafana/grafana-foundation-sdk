// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';
import * as dashboard from '../dashboard';

// Shape of a CloudWatch Annotation query
// TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
// #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
export class CloudWatchAnnotationQueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: cloudwatch.CloudWatchAnnotationQuery;

    constructor() {
        this.internal = cloudwatch.defaultCloudWatchAnnotationQuery();
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.CloudWatchAnnotationQuery {
        return this.internal;
    }

    // Whether a query is a Metrics, Logs, or Annotations query
    queryMode(queryMode: cloudwatch.CloudWatchQueryMode): this {
        this.internal.queryMode = queryMode;
        return this;
    }

    // Enable matching on the prefix of the action name or alarm name, specify the prefixes with actionPrefix and/or alarmNamePrefix
    prefixMatching(prefixMatching: boolean): this {
        this.internal.prefixMatching = prefixMatching;
        return this;
    }

    // Use this parameter to filter the results of the operation to only those alarms
    // that use a certain alarm action. For example, you could specify the ARN of
    // an SNS topic to find all alarms that send notifications to that topic.
    // e.g. `arn:aws:sns:us-east-1:123456789012:my-app-` would match `arn:aws:sns:us-east-1:123456789012:my-app-action`
    // but not match `arn:aws:sns:us-east-1:123456789012:your-app-action`
    actionPrefix(actionPrefix: string): this {
        this.internal.actionPrefix = actionPrefix;
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

    // An alarm name prefix. If you specify this parameter, you receive information
    // about all alarms that have names that start with this prefix.
    // e.g. `my-team-service-` would match `my-team-service-high-cpu` but not match `your-team-service-high-cpu`
    alarmNamePrefix(alarmNamePrefix: string): this {
        this.internal.alarmNamePrefix = alarmNamePrefix;
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
