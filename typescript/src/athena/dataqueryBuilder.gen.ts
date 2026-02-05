// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as athena from '../athena';
import * as common from '../common';

// Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: athena.Dataquery;

    constructor() {
        this.internal = athena.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): athena.Dataquery {
        return this.internal;
    }

    format(format: athena.FormatOptions): this {
        this.internal.format = format;
        return this;
    }

    connectionArgs(connectionArgs: cog.Builder<athena.ConnectionArgs>): this {
        const connectionArgsResource = connectionArgs.build();
        this.internal.connectionArgs = connectionArgsResource;
        return this;
    }

    table(table: string): this {
        this.internal.table = table;
        return this;
    }

    column(column: string): this {
        this.internal.column = column;
        return this;
    }

    queryID(queryID: string): this {
        this.internal.queryID = queryID;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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

    rawSQL(rawSQL: string): this {
        this.internal.rawSQL = rawSQL;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource: common.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }
}

