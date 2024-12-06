// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class MetricNamespaceQueryBuilder implements cog.Builder<azuremonitor.MetricNamespaceQuery> {
    protected readonly internal: azuremonitor.MetricNamespaceQuery;

    constructor() {
        this.internal = azuremonitor.defaultMetricNamespaceQuery();
        this.internal.kind = "MetricNamespaceQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.MetricNamespaceQuery {
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

    metricNamespace(metricNamespace: string): this {
        this.internal.metricNamespace = metricNamespace;
        return this;
    }

    resourceName(resourceName: string): this {
        this.internal.resourceName = resourceName;
        return this;
    }
}
