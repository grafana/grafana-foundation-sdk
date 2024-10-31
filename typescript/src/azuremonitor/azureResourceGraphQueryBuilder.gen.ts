// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class AzureResourceGraphQueryBuilder implements cog.Builder<azuremonitor.AzureResourceGraphQuery> {
    protected readonly internal: azuremonitor.AzureResourceGraphQuery;

    constructor() {
        this.internal = azuremonitor.defaultAzureResourceGraphQuery();
    }

    /**
     * Builds the object.
     */
    build(): azuremonitor.AzureResourceGraphQuery {
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
