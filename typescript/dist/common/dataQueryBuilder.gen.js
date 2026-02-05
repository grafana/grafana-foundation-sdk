"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataQueryBuilder = void 0;
const tslib_1 = require("tslib");
const common = tslib_1.__importStar(require("../common"));
// These are the common properties available to all queries in all datasources.
// Specific implementations will *extend* this interface, adding the required
// properties for the given context.
class DataQueryBuilder {
    constructor() {
        this.internal = common.defaultDataQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
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
exports.DataQueryBuilder = DataQueryBuilder;
//# sourceMappingURL=dataQueryBuilder.gen.js.map