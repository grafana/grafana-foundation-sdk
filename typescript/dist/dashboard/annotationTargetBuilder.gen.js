"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AnnotationTargetBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most dashboards
class AnnotationTargetBuilder {
    constructor() {
        this.internal = dashboard.defaultAnnotationTarget();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    limit(limit) {
        this.internal.limit = limit;
        return this;
    }
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    matchAny(matchAny) {
        this.internal.matchAny = matchAny;
        return this;
    }
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    tags(tags) {
        this.internal.tags = tags;
        return this;
    }
    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    type(type) {
        this.internal.type = type;
        return this;
    }
}
exports.AnnotationTargetBuilder = AnnotationTargetBuilder;
//# sourceMappingURL=annotationTargetBuilder.gen.js.map