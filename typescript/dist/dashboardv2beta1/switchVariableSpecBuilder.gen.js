"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SwitchVariableSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class SwitchVariableSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultSwitchVariableSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    current(current) {
        this.internal.current = current;
        return this;
    }
    enabledValue(enabledValue) {
        this.internal.enabledValue = enabledValue;
        return this;
    }
    disabledValue(disabledValue) {
        this.internal.disabledValue = disabledValue;
        return this;
    }
    label(label) {
        this.internal.label = label;
        return this;
    }
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        this.internal.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        this.internal.description = description;
        return this;
    }
}
exports.SwitchVariableSpecBuilder = SwitchVariableSpecBuilder;
//# sourceMappingURL=switchVariableSpecBuilder.gen.js.map