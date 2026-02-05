"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataqueryBuilder = void 0;
const tslib_1 = require("tslib");
const bigquery = tslib_1.__importStar(require("../bigquery"));
// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
class DataqueryBuilder {
    constructor() {
        this.internal = bigquery.defaultDataquery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    dataset(dataset) {
        this.internal.dataset = dataset;
        return this;
    }
    table(table) {
        this.internal.table = table;
        return this;
    }
    project(project) {
        this.internal.project = project;
        return this;
    }
    format(format) {
        this.internal.format = format;
        return this;
    }
    rawQuery(rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    rawSql(rawSql) {
        this.internal.rawSql = rawSql;
        return this;
    }
    location(location) {
        this.internal.location = location;
        return this;
    }
    partitioned(partitioned) {
        this.internal.partitioned = partitioned;
        return this;
    }
    partitionedField(partitionedField) {
        this.internal.partitionedField = partitionedField;
        return this;
    }
    convertToUTC(convertToUTC) {
        this.internal.convertToUTC = convertToUTC;
        return this;
    }
    sharded(sharded) {
        this.internal.sharded = sharded;
        return this;
    }
    queryPriority(queryPriority) {
        this.internal.queryPriority = queryPriority;
        return this;
    }
    timeShift(timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    editorMode(editorMode) {
        this.internal.editorMode = editorMode;
        return this;
    }
    sql(sql) {
        const sqlResource = sql.build();
        this.internal.sql = sqlResource;
        return this;
    }
    // A unique identifier for the query within the list of targets.
    // In server side expressions, the refId is used as a variable name to identify results.
    // By default, the UI will assign A->Z; however setting meaningful names may be useful.
    refId(refId) {
        this.internal.refId = refId;
        return this;
    }
    // If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Specify the query flavor
    // TODO make this required and give it a default
    queryType(queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    // For mixed data sources the selected datasource is on the query level.
    // For non mixed scenarios this is undefined.
    // TODO find a better way to do this ^ that's friendly to schema
    // TODO this shouldn't be unknown but DataSourceRef | null
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
}
exports.DataqueryBuilder = DataqueryBuilder;
//# sourceMappingURL=dataqueryBuilder.gen.js.map