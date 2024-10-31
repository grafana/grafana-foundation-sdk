// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class WorkspacesQueryBuilder implements cog.Builder<azuremonitor.WorkspacesQuery> {
    protected readonly internal: azuremonitor.WorkspacesQuery;

    constructor() {
        this.internal = azuremonitor.defaultWorkspacesQuery();
        this.internal.kind = "WorkspacesQuery";
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.WorkspacesQuery {
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
