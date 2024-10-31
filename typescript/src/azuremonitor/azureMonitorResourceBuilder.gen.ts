// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AzureMonitorResourceBuilder implements cog.Builder<azuremonitor.AzureMonitorResource> {
    protected readonly internal: azuremonitor.AzureMonitorResource;

    constructor() {
        this.internal = azuremonitor.defaultAzureMonitorResource();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureMonitorResource {
        return this.internal;
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

    region(region: string): this {
        this.internal.region = region;
        return this;
    }
}
