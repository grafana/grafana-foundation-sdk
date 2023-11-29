// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as azuremonitor from '../azuremonitor';

export class BaseGrafanaTemplateVariableQueryBuilder implements cog.Builder<azuremonitor.BaseGrafanaTemplateVariableQuery> {
    private readonly internal: azuremonitor.BaseGrafanaTemplateVariableQuery;

    constructor() {
        this.internal = azuremonitor.defaultBaseGrafanaTemplateVariableQuery();
    }

    build(): azuremonitor.BaseGrafanaTemplateVariableQuery {
        return this.internal;
    }

    rawQuery(rawQuery: string): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }
}
