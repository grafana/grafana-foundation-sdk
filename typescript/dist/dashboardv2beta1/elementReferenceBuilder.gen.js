"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ElementReferenceBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class ElementReferenceBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultElementReference();
        this.internal.kind = "ElementReference";
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
}
exports.ElementReferenceBuilder = ElementReferenceBuilder;
//# sourceMappingURL=elementReferenceBuilder.gen.js.map