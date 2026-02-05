"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TransformationBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class TransformationBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultTransformationKind();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // The kind of a TransformationKind is the transformation ID
    kind(kind) {
        this.internal.kind = kind;
        return this;
    }
    // Unique identifier of transformer
    id(id) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.id = id;
        return this;
    }
    // Disabled transformations are skipped
    disabled(disabled) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.disabled = disabled;
        return this;
    }
    // Optional frame matcher. When missing it will be applied to all results
    filter(filter) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.filter = filter;
        return this;
    }
    // Where to pull DataFrames from as input to transformation
    topic(topic) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.topic = topic;
        return this;
    }
    // Options to be passed to the transformer
    // Valid options depend on the transformer id
    options(options) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDataTransformerConfig();
        }
        this.internal.spec.options = options;
        return this;
    }
}
exports.TransformationBuilder = TransformationBuilder;
//# sourceMappingURL=transformationBuilder.gen.js.map