// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as cloudwatch from '../cloudwatch';

// Shape of a CloudWatch Logs query
export class CloudWatchLogsQueryBuilder implements cog.Builder<cog.Dataquery> {
    private readonly internal: cloudwatch.CloudWatchLogsQuery;

    constructor() {
        this.internal = cloudwatch.defaultCloudWatchLogsQuery();
    }

    build(): cloudwatch.CloudWatchLogsQuery {
        return this.internal;
    }

    // Whether a query is a Metrics, Logs, or Annotations query
    queryMode(queryMode: cloudwatch.CloudWatchQueryMode): this {
        this.internal.queryMode = queryMode;
        return this;
    }

    id(id: string): this {
        this.internal.id = id;
        return this;
    }

    // AWS region to query for the logs
    region(region: string): this {
        this.internal.region = region;
        return this;
    }

    // The CloudWatch Logs Insights query to execute
    expression(expression: string): this {
        this.internal.expression = expression;
        return this;
    }

    // Fields to group the results by, this field is automatically populated whenever the query is updated
    statsGroups(statsGroups: string[]): this {
        this.internal.statsGroups = statsGroups;
        return this;
    }

    // Log groups to query
    logGroups(logGroups: cog.Builder<cloudwatch.LogGroup>[]): this {
        const logGroupsResources = logGroups.map(builder1 => builder1.build());
        this.internal.logGroups = logGroupsResources;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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

    // @deprecated use logGroups
    logGroupNames(logGroupNames: string[]): this {
        this.internal.logGroupNames = logGroupNames;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: any): this {
        this.internal.datasource = datasource;
        return this;
    }
}
