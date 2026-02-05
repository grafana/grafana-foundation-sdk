"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConditionalRenderingVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ConditionalRenderingVariableBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingVariableKind();
        this.internal.kind = "ConditionalRenderingVariable";
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
    variable(variable) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingVariableSpec();
        }
        this.internal.spec.variable = variable;
        return this;
    }
    operator(operator) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingVariableSpec();
        }
        this.internal.spec.operator = operator;
        return this;
    }
    value(value) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingVariableSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}
exports.ConditionalRenderingVariableBuilder = ConditionalRenderingVariableBuilder;
//# sourceMappingURL=conditionalRenderingVariableBuilder.gen.js.map