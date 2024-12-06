// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AppInsightsMetricNameQueryBuilder implements cog.Builder<azuremonitor.AppInsightsMetricNameQuery> {
    protected readonly internal: azuremonitor.AppInsightsMetricNameQuery;

    constructor() {
        this.internal = azuremonitor.defaultAppInsightsMetricNameQuery();
        this.internal.kind = "AppInsightsMetricNameQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AppInsightsMetricNameQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }
}
