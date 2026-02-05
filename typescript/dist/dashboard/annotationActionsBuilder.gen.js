"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AnnotationActionsBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
class AnnotationActionsBuilder {
    constructor() {
        this.internal = dashboard.defaultAnnotationActions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    canAdd(canAdd) {
        this.internal.canAdd = canAdd;
        return this;
    }
    canDelete(canDelete) {
        this.internal.canDelete = canDelete;
        return this;
    }
    canEdit(canEdit) {
        this.internal.canEdit = canEdit;
        return this;
    }
}
exports.AnnotationActionsBuilder = AnnotationActionsBuilder;
//# sourceMappingURL=annotationActionsBuilder.gen.js.map