"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.PromQLQueryBuilder = void 0;
const tslib_1 = require("tslib");
const googlecloudmonitoring = tslib_1.__importStar(require("../googlecloudmonitoring"));
// PromQL sub-query properties.
class PromQLQueryBuilder {
    constructor() {
        this.internal = googlecloudmonitoring.defaultPromQLQuery();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // GCP project to execute the query against.
    projectName(projectName) {
        this.internal.projectName = projectName;
        return this;
    }
    // PromQL expression/query to be executed.
    expr(expr) {
        this.internal.expr = expr;
        return this;
    }
    // PromQL min step
    step(step) {
        this.internal.step = step;
        return this;
    }
}
exports.PromQLQueryBuilder = PromQLQueryBuilder;
//# sourceMappingURL=promQLQueryBuilder.gen.js.map