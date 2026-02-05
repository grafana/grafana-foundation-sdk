"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AppInsightsGroupByQueryBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class AppInsightsGroupByQueryBuilder {
    constructor() {
        this.internal = azuremonitor.defaultAppInsightsGroupByQuery();
        this.internal.kind = "AppInsightsGroupByQuery";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    rawQuery(rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    metricName(metricName) {
        this.internal.metricName = metricName;
        return this;
    }
}
exports.AppInsightsGroupByQueryBuilder = AppInsightsGroupByQueryBuilder;
//# sourceMappingURL=appInsightsGroupByQueryBuilder.gen.js.map