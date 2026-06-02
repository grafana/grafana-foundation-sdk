// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class ResourceGraphQueryBuilder implements cog.Builder<azuremonitor.ResourceGraphQuery> {
    protected readonly internal: azuremonitor.ResourceGraphQuery;

    constructor() {
        this.internal = azuremonitor.defaultResourceGraphQuery();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.ResourceGraphQuery {
        return this.internal;
    }

    // Azure Resource Graph KQL query to be executed.
    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    // Specifies the format results should be returned as. Defaults to table.
    resultFormat(resultFormat: string): this {
        this.internal.resultFormat = resultFormat;
        return this;
    }
}

