// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class ResourceNamesQueryBuilder implements cog.Builder<azuremonitor.ResourceNamesQuery> {
    protected readonly internal: azuremonitor.ResourceNamesQuery;

    constructor() {
        this.internal = azuremonitor.defaultResourceNamesQuery();
        this.internal.kind = "ResourceNamesQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.ResourceNamesQuery {
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
}
