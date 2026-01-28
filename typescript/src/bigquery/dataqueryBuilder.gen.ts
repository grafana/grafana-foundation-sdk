// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as bigquery from '../bigquery';
import * as common from '../common';

// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
export class DataqueryBuilder implements cog.Builder<cog.Dataquery> {
    protected readonly internal: bigquery.Dataquery;

    constructor() {
        this.internal = bigquery.defaultDataquery();
    }

    /**
     * Builds the object.
     */
    build(): bigquery.Dataquery {
        return this.internal;
    }

    dataset(dataset: string): this {
        this.internal.dataset = dataset;
        return this;
    }

    table(table: string): this {
        this.internal.table = table;
        return this;
    }

    project(project: string): this {
        this.internal.project = project;
        return this;
    }

    format(format: bigquery.QueryFormat): this {
        this.internal.format = format;
        return this;
    }

    rawQuery(rawQuery: boolean): this {
        this.internal.rawQuery = rawQuery;
        return this;
    }

    rawSql(rawSql: string): this {
        this.internal.rawSql = rawSql;
        return this;
    }

    location(location: string): this {
        this.internal.location = location;
        return this;
    }

    partitioned(partitioned: boolean): this {
        this.internal.partitioned = partitioned;
        return this;
    }

    partitionedField(partitionedField: string): this {
        this.internal.partitionedField = partitionedField;
        return this;
    }

    convertToUTC(convertToUTC: boolean): this {
        this.internal.convertToUTC = convertToUTC;
        return this;
    }

    sharded(sharded: boolean): this {
        this.internal.sharded = sharded;
        return this;
    }

    queryPriority(queryPriority: bigquery.QueryPriority): this {
        this.internal.queryPriority = queryPriority;
        return this;
    }

    timeShift(timeShift: string): this {
        this.internal.timeShift = timeShift;
        return this;
    }

    editorMode(editorMode: bigquery.EditorMode): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    sql(sql: cog.Builder<bigquery.SQLExpression>): this {
        const sqlResource = sql.build();
        this.internal.sql = sqlResource;
        return this;
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
    datasource(datasource: common.DataSourceRef): this {
        this.internal.datasource = datasource;
        return this;
    }
}

