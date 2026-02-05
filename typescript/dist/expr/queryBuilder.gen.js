"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class QueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultDataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "__expr__";
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
    typeMath(typeMath) {
        const typeMathResource = typeMath.build();
        this.internal.spec = typeMathResource;
        return this;
    }
    typeReduce(typeReduce) {
        const typeReduceResource = typeReduce.build();
        this.internal.spec = typeReduceResource;
        return this;
    }
    typeResample(typeResample) {
        const typeResampleResource = typeResample.build();
        this.internal.spec = typeResampleResource;
        return this;
    }
    typeClassicConditions(typeClassicConditions) {
        const typeClassicConditionsResource = typeClassicConditions.build();
        this.internal.spec = typeClassicConditionsResource;
        return this;
    }
    typeThreshold(typeThreshold) {
        const typeThresholdResource = typeThreshold.build();
        this.internal.spec = typeThresholdResource;
        return this;
    }
    typeSql(typeSql) {
        const typeSqlResource = typeSql.build();
        this.internal.spec = typeSqlResource;
        return this;
    }
}
exports.QueryBuilder = QueryBuilder;
//# sourceMappingURL=queryBuilder.gen.js.map