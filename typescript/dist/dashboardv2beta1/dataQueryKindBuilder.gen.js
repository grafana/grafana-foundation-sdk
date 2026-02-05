"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataQueryKindBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class DataQueryKindBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    group(group) {
        this.internal.group = group;
        return this;
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
    spec(spec) {
        this.internal.spec = spec;
        return this;
    }
}
exports.DataQueryKindBuilder = DataQueryKindBuilder;
//# sourceMappingURL=dataQueryKindBuilder.gen.js.map