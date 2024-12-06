// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AppInsightsGroupByQueryBuilder implements cog.Builder<azuremonitor.AppInsightsGroupByQuery> {
    protected readonly internal: azuremonitor.AppInsightsGroupByQuery;

    constructor() {
        this.internal = azuremonitor.defaultAppInsightsGroupByQuery();
        this.internal.kind = "AppInsightsGroupByQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AppInsightsGroupByQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }

    metricName(metricName: string): this {
        this.internal.metricName = metricName;
        return this;
    }
}
