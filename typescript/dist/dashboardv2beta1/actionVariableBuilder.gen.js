"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ActionVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ActionVariableBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultActionVariable();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    key(key) {
        this.internal.key = key;
        return this;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
}
exports.ActionVariableBuilder = ActionVariableBuilder;
//# sourceMappingURL=actionVariableBuilder.gen.js.map