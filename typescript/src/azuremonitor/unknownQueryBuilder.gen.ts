// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class UnknownQueryBuilder implements cog.Builder<azuremonitor.UnknownQuery> {
    protected readonly internal: azuremonitor.UnknownQuery;

    constructor() {
        this.internal = azuremonitor.defaultUnknownQuery();
        this.internal.kind = "UnknownQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.UnknownQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }
}
