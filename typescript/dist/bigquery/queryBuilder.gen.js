"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
const bigquery = tslib_1.__importStar(require("../bigquery"));
class QueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-bigquery-datasource";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    version(version) {
        this.internal.version = version;
        return this;
    }
    // New type for datasource reference
    // Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource(ref) {
        this.internal.datasource = ref;
        return this;
    }
    dataset(dataset) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.dataset = dataset;
        return this;
    }
    table(table) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.table = table;
        return this;
    }
    project(project) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.project = project;
        return this;
    }
    format(format) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.format = format;
        return this;
    }
    rawQuery(rawQuery) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.rawQuery = rawQuery;
        return this;
    }
    rawSql(rawSql) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.rawSql = rawSql;
        return this;
    }
    location(location) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.location = location;
        return this;
    }
    partitioned(partitioned) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.partitioned = partitioned;
        return this;
    }
    partitionedField(partitionedField) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.partitionedField = partitionedField;
        return this;
    }
    convertToUTC(convertToUTC) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.convertToUTC = convertToUTC;
        return this;
    }
    sharded(sharded) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.sharded = sharded;
        return this;
    }
    queryPriority(queryPriority) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.queryPriority = queryPriority;
        return this;
    }
    timeShift(timeShift) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.timeShift = timeShift;
        return this;
    }
    editorMode(editorMode) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.editorMode = editorMode;
        return this;
    }
    sql(sql) {
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
    refId(refId) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        if (!this.internal.spec) {
            this.internal.spec = bigquery.defaultDataquery();
        }
        this.internal.spec.queryType = queryType;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map