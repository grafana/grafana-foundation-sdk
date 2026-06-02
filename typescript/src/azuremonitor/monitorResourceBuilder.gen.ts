// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class MonitorResourceBuilder implements cog.Builder<azuremonitor.MonitorResource> {
    protected readonly internal: azuremonitor.MonitorResource;

    constructor() {
        this.internal = azuremonitor.defaultMonitorResource();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.MonitorResource {
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

