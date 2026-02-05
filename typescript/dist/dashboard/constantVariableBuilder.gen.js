"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConstantVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
class ConstantVariableBuilder {
    constructor(name) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Constant;
        this.internal.hide = dashboard.VariableHide.HideVariable;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // Name of variable
    name(name) {
        this.internal.name = name;
        return this;
    }
    // Optional display name
    label(label) {
        this.internal.label = label;
        return this;
    }
    // Description of variable. It can be defined but `null`.
    description(description) {
        this.internal.description = description;
        return this;
    }
    // Query used to fetch values for a variable
    value(query) {
        this.internal.query = query;
        return this;
    }
    // Allow custom values to be entered in the variable
    allowCustomValue(allowCustomValue) {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
}
exports.ConstantVariableBuilder = ConstantVariableBuilder;
//# sourceMappingURL=constantVariableBuilder.gen.js.map