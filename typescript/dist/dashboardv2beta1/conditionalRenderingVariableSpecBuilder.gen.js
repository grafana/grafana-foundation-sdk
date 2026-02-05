"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConditionalRenderingVariableSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ConditionalRenderingVariableSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingVariableSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    variable(variable) {
        this.internal.variable = variable;
        return this;
    }
    operator(operator) {
        this.internal.operator = operator;
        return this;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
}
exports.ConditionalRenderingVariableSpecBuilder = ConditionalRenderingVariableSpecBuilder;
//# sourceMappingURL=conditionalRenderingVariableSpecBuilder.gen.js.map