"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AnnotationQuerySpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class AnnotationQuerySpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAnnotationQuerySpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    query(query) {
        const queryResource = query.build();
        this.internal.query = queryResource;
        return this;
    }
    enable(enable) {
        this.internal.enable = enable;
        return this;
    }
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    iconColor(iconColor) {
        this.internal.iconColor = iconColor;
        return this;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    builtIn(builtIn) {
        this.internal.builtIn = builtIn;
        return this;
    }
    filter(filter) {
        const filterResource = filter.build();
        this.internal.filter = filterResource;
        return this;
    }
    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    placement(placement) {
        this.internal.placement = placement;
        return this;
    }
    // Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    legacyOptions(legacyOptions) {
        this.internal.legacyOptions = legacyOptions;
        return this;
    }
}
exports.AnnotationQuerySpecBuilder = AnnotationQuerySpecBuilder;
//# sourceMappingURL=annotationQuerySpecBuilder.gen.js.map