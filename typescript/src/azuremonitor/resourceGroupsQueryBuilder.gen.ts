// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class ResourceGroupsQueryBuilder implements cog.Builder<azuremonitor.ResourceGroupsQuery> {
    protected readonly internal: azuremonitor.ResourceGroupsQuery;

    constructor() {
        this.internal = azuremonitor.defaultResourceGroupsQuery();
        this.internal.kind = "ResourceGroupsQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.ResourceGroupsQuery {
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
}
