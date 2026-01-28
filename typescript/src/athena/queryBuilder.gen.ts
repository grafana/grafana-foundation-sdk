// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as athena from '../athena';
import * as common from '../common';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-athena-datasource";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DataQueryKind {
        return this.internal;
    }

    version(version: string): this {
        this.internal.version = version;
        return this;
    }

    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref: {
	name?: string;
}): this {
        this.internal.datasource = ref;
        return this;
    }

    format(format: athena.FormatOptions): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.format = format;
        return this;
    }

    connectionArgs(connectionArgs: cog.Builder<athena.ConnectionArgs>): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        const connectionArgsResource = connectionArgs.build();
        this.internal.spec.connectionArgs = connectionArgsResource;
        return this;
    }

    table(table: string): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.table = table;
        return this;
    }

    column(column: string): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.column = column;
        return this;
    }

    queryID(queryID: string): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.queryID = queryID;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    rawSQL(rawSQL: string): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.rawSQL = rawSQL;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    oldDatasource(datasource: common.DataSourceRef): this {
        if (!this.internal.spec) {
            this.internal.spec = athena.defaultDataquery();
        }
        this.internal.spec.datasource = datasource;
        return this;
    }
}

