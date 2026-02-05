"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TextBoxVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
class TextBoxVariableBuilder {
    constructor(name) {
        this.internal = dashboard.defaultVariableModel();
        this.internal.name = name;
        this.internal.type = dashboard.VariableType.Textbox;
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
    // Visibility configuration for the variable
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    // Description of variable. It can be defined but `null`.
    description(description) {
        this.internal.description = description;
        return this;
    }
    // Query used to fetch values for a variable
    defaultValue(query) {
        this.internal.query = query;
        return this;
    }
    // Shows current selected variable text/value on the dashboard
    current(current) {
        this.internal.current = current;
        return this;
    }
    // Allow custom values to be entered in the variable
    allowCustomValue(allowCustomValue) {
        this.internal.allowCustomValue = allowCustomValue;
        return this;
    }
    // Options that can be selected for a variable.
    options(options) {
        this.internal.options = options;
        return this;
    }
}
exports.TextBoxVariableBuilder = TextBoxVariableBuilder;
//# sourceMappingURL=textBoxVariableBuilder.gen.js.map