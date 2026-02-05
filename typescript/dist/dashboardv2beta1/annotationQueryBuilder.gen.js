"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AnnotationQueryBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class AnnotationQueryBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultAnnotationQueryKind();
        this.internal.kind = "AnnotationQuery";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    spec(spec) {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }
    enable(enable) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.enable = enable;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    iconColor(iconColor) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.iconColor = iconColor;
        return this;
    }
    name(name) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    builtIn(builtIn) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.builtIn = builtIn;
        return this;
    }
    filter(filter) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        const filterResource = filter.build();
        this.internal.spec.filter = filterResource;
        return this;
    }
    // Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    placement(placement) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.placement = placement;
        return this;
    }
    // Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    legacyOptions(legacyOptions) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAnnotationQuerySpec();
        }
        this.internal.spec.legacyOptions = legacyOptions;
        return this;
    }
}
exports.AnnotationQueryBuilder = AnnotationQueryBuilder;
//# sourceMappingURL=annotationQueryBuilder.gen.js.map