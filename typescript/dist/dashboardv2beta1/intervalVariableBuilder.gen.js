"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.IntervalVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Interval variable kind
class IntervalVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultIntervalVariableKind();
        this.internal.kind = "IntervalVariable";
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
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    query(query) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    options(options) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }
    auto(auto) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto = auto;
        return this;
    }
    autoMin(autoMin) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto_min = autoMin;
        return this;
    }
    autoCount(autoCount) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.auto_count = autoCount;
        return this;
    }
    refresh(refresh) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultIntervalVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}
exports.IntervalVariableBuilder = IntervalVariableBuilder;
//# sourceMappingURL=intervalVariableBuilder.gen.js.map