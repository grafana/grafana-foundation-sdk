"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AnnotationPermissionBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
class AnnotationPermissionBuilder {
    constructor() {
        this.internal = dashboard.defaultAnnotationPermission();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    dashboard(dashboard) {
        const dashboardResource = dashboard.build();
        this.internal.dashboard = dashboardResource;
        return this;
    }
    organization(organization) {
        const organizationResource = organization.build();
        this.internal.organization = organizationResource;
        return this;
    }
}
exports.AnnotationPermissionBuilder = AnnotationPermissionBuilder;
//# sourceMappingURL=annotationPermissionBuilder.gen.js.map