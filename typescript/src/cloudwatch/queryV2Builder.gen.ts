// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';
import * as cloudwatch from '../cloudwatch';

export class QueryV2Builder implements cog.Builder<dashboardv2.DataQueryKind> {
    protected readonly internal: dashboardv2.DataQueryKind;

    constructor() {
        this.internal = dashboardv2.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloudwatch";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    labels(labels: Record<string, string>): this {
        this.internal.labels = labels;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    cloudWatchMetricsQuery(cloudWatchMetricsQuery: cog.Builder<cloudwatch.CloudWatchMetricsQuery>): this {
        const cloudWatchMetricsQueryResource = cloudWatchMetricsQuery.build();
        this.internal.spec = cloudWatchMetricsQueryResource;
        return this;
    }

    cloudWatchLogsQuery(cloudWatchLogsQuery: cog.Builder<cloudwatch.CloudWatchLogsQuery>): this {
        const cloudWatchLogsQueryResource = cloudWatchLogsQuery.build();
        this.internal.spec = cloudWatchLogsQueryResource;
        return this;
    }

    cloudWatchAnnotationQuery(cloudWatchAnnotationQuery: cog.Builder<cloudwatch.CloudWatchAnnotationQuery>): this {
        const cloudWatchAnnotationQueryResource = cloudWatchAnnotationQuery.build();
        this.internal.spec = cloudWatchAnnotationQueryResource;
        return this;
    }
}

