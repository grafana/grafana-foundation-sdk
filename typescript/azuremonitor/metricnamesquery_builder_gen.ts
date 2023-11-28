// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class MetricNamesQueryBuilder implements cog.OptionsBuilder<azuremonitor.MetricNamesQuery> {
    private readonly internal: azuremonitor.MetricNamesQuery;

    constructor() {
        this.internal = azuremonitor.defaultMetricNamesQuery();
        this.internal.kind = "MetricNamesQuery";
    }

    build(): azuremonitor.MetricNamesQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }

    subscription(subscription: string): this {
        this.internal.subscription = subscription;
        return this;
    }

    resourceGroup(resourceGroup: string): this {
        this.internal.resourceGroup = resourceGroup;
        return this;
    }

    resourceName(resourceName: string): this {
        this.internal.resourceName = resourceName;
        return this;
    }

    metricNamespace(metricNamespace: string): this {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }
}
