// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

export class MetricStatBuilder implements cog.Builder<cloudwatch.MetricStat> {
    protected readonly internal: cloudwatch.MetricStat;

    constructor() {
        this.internal = cloudwatch.defaultMetricStat();
    }

    /**
     * Builds the object.
     */
    build(): cloudwatch.MetricStat {
        return this.internal;
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

    // @deprecated use statistic
    statistics(statistics: string[]): this {
        this.internal.statistics = statistics;
        return this;
    }
}
