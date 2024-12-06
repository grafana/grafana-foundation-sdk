// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

// @deprecated Use MetricNamespaceQuery instead
export class MetricDefinitionsQueryBuilder implements cog.Builder<azuremonitor.MetricDefinitionsQuery> {
    protected readonly internal: azuremonitor.MetricDefinitionsQuery;

    constructor() {
        this.internal = azuremonitor.defaultMetricDefinitionsQuery();
        this.internal.kind = "MetricDefinitionsQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.MetricDefinitionsQuery {
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
