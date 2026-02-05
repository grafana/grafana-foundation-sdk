"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SwitchVariableKindBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class SwitchVariableKindBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultSwitchVariableKind();
        this.internal.kind = "SwitchVariable";
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
}
exports.SwitchVariableKindBuilder = SwitchVariableKindBuilder;
//# sourceMappingURL=switchVariableKindBuilder.gen.js.map