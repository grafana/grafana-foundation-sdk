"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConstantVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Constant variable kind
class ConstantVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultConstantVariableKind();
        this.internal.kind = "ConstantVariable";
        this.internal.spec.name = name;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    spec(spec) {
        this.internal.spec = spec;
        return this;
    }
    name(name) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}
exports.ConstantVariableBuilder = ConstantVariableBuilder;
//# sourceMappingURL=constantVariableBuilder.gen.js.map