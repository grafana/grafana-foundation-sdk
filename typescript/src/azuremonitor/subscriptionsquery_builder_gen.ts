// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class SubscriptionsQueryBuilder implements cog.Builder<azuremonitor.SubscriptionsQuery> {
    private readonly internal: azuremonitor.SubscriptionsQuery;

    constructor() {
        this.internal = azuremonitor.defaultSubscriptionsQuery();
        this.internal.kind = "SubscriptionsQuery";
    }

    build(): azuremonitor.SubscriptionsQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }
}
