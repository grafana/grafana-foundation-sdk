// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AppInsightsMetricNameQueryBuilder implements cog.OptionsBuilder<azuremonitor.AppInsightsMetricNameQuery> {
    private readonly internal: azuremonitor.AppInsightsMetricNameQuery;

    constructor() {
        this.internal = azuremonitor.defaultAppInsightsMetricNameQuery();
        this.internal.kind = "AppInsightsMetricNameQuery";
    }

    build(): azuremonitor.AppInsightsMetricNameQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }
}
