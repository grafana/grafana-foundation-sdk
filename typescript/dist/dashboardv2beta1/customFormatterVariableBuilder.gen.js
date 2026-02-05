"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CustomFormatterVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Custom formatter variable
class CustomFormatterVariableBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultCustomFormatterVariable();
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
    type(type) {
        this.internal.type = type;
        return this;
    }
    multi(multi) {
        this.internal.multi = multi;
        return this;
    }
    includeAll(includeAll) {
        this.internal.includeAll = includeAll;
        return this;
    }
}
exports.CustomFormatterVariableBuilder = CustomFormatterVariableBuilder;
//# sourceMappingURL=customFormatterVariableBuilder.gen.js.map