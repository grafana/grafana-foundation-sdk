// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

// Azure Monitor Logs sub-query properties
export class AzureLogsQueryBuilder implements cog.Builder<azuremonitor.AzureLogsQuery> {
    protected readonly internal: azuremonitor.AzureLogsQuery;

    constructor() {
        this.internal = azuremonitor.defaultAzureLogsQuery();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureLogsQuery {
        return this.internal;
    }

    // KQL query to be executed.
    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    // Specifies the format results should be returned as.
    resultFormat(resultFormat: azuremonitor.ResultFormat): this {
        this.internal.resultFormat = resultFormat;
        return this;
    }

    // Array of resource URIs to be queried.
    resources(resources: string[]): this {
        this.internal.resources = resources;
        return this;
    }

    // If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
    dashboardTime(dashboardTime: boolean): this {
        this.internal.dashboardTime = dashboardTime;
        return this;
    }

    // If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
    timeColumn(timeColumn: string): this {
        this.internal.timeColumn = timeColumn;
        return this;
    }

    // Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
    workspace(workspace: string): this {
        this.internal.workspace = workspace;
        return this;
    }

    // @deprecated Use resources instead
    resource(resource: string): this {
        this.internal.resource = resource;
        return this;
    }

    // @deprecated Use dashboardTime instead
    intersectTime(intersectTime: boolean): this {
        this.internal.intersectTime = intersectTime;
        return this;
    }
}
