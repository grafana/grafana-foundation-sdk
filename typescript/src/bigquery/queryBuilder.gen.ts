// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
import * as bigquery from '../bigquery';
import * as common from '../common';

export class QueryBuilder implements cog.Builder<dashboardv2beta1.DataQueryKind> {
    protected readonly internal: dashboardv2beta1.DataQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-bigquery-datasource";
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

    dataset(dataset: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.dataset = dataset;
        return this;
    }

    table(table: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.table = table;
        return this;
    }

    project(project: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.project = project;
        return this;
    }

    format(format: bigquery.QueryFormat): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.format = format;
        return this;
    }

    rawQuery(rawQuery: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.rawQuery = rawQuery;
        return this;
    }

    rawSql(rawSql: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.rawSql = rawSql;
        return this;
    }

    location(location: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.location = location;
        return this;
    }

    partitioned(partitioned: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.partitioned = partitioned;
        return this;
    }

    partitionedField(partitionedField: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.partitionedField = partitionedField;
        return this;
    }

    convertToUTC(convertToUTC: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.convertToUTC = convertToUTC;
        return this;
    }

    sharded(sharded: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.sharded = sharded;
        return this;
    }

    queryPriority(queryPriority: bigquery.QueryPriority): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.queryPriority = queryPriority;
        return this;
    }

    timeShift(timeShift: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.timeShift = timeShift;
        return this;
    }

    editorMode(editorMode: bigquery.EditorMode): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.editorMode = editorMode;
        return this;
    }

    sql(sql: cog.Builder<bigquery.SQLExpression>): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        const sqlResource = sql.build();
        this.internal.spec.sql = sqlResource;
        return this;
    }

    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType: string): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }

    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    oldDatasource(datasource: common.DataSourceRef): this {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.datasource = datasource;
        return this;
    }
}

