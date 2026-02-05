"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AnnotationPanelFilterBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
class AnnotationPanelFilterBuilder {
    constructor() {
        this.internal = dashboard.defaultAnnotationPanelFilter();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Should the specified panels be included or excluded
    exclude(exclude) {
        this.internal.exclude = exclude;
        return this;
    }
    // Panel IDs that should be included or excluded
    ids(ids) {
        this.internal.ids = ids;
        return this;
    }
}
exports.AnnotationPanelFilterBuilder = AnnotationPanelFilterBuilder;
//# sourceMappingURL=annotationPanelFilterBuilder.gen.js.map