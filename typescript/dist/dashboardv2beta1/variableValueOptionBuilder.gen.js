"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.VariableValueOptionBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// FIXME: should we introduce this? --- Variable value option
class VariableValueOptionBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultVariableValueOption();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    label(label) {
        this.internal.label = label;
        return this;
    }
    value(value) {
        this.internal.value = value;
        return this;
    }
    group(group) {
        this.internal.group = group;
        return this;
    }
}
exports.VariableValueOptionBuilder = VariableValueOptionBuilder;
//# sourceMappingURL=variableValueOptionBuilder.gen.js.map