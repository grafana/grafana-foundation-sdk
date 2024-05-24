// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
export class DataQueryBuilder implements cog.Builder<common.DataQuery> {
    protected readonly internal: common.DataQuery;

    constructor() {
        this.internal = common.defaultDataQuery();
    }

    build(): common.DataQuery {
        return this.internal;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // true if query is disabled (ie should not be returned to the dashboard)
    // Note this does not always imply that the query should not be executed since
    // the results from a hidden query may be used as the input to other queries (SSE etc)
    hide(hide: boolean): this {
        this.internal.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        this.internal.queryType = queryType;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: any): this {
        this.internal.datasource = datasource;
        return this;
    }
}
