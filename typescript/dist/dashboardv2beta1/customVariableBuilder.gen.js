"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CustomVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Custom variable kind
class CustomVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultCustomVariableKind();
        this.internal.kind = "CustomVariable";
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
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    options(options) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }
    multi(multi) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }
    includeAll(includeAll) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }
    allValue(allValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
    allowCustomValue(allowCustomValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
}
exports.CustomVariableBuilder = CustomVariableBuilder;
//# sourceMappingURL=customVariableBuilder.gen.js.map