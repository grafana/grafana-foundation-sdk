"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.GroupByVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Group variable kind
class GroupByVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultGroupByVariableKind();
        this.internal.kind = "GroupByVariable";
        this.internal.spec.name = name;
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    group(group) {
        this.internal.group = group;
        return this;
    }
    datasource(datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    spec(spec) {
        this.internal.spec = spec;
        return this;
    }
    name(name) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    defaultValue(defaultValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.defaultValue = defaultValue;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    options(options) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }
    multi(multi) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}
exports.GroupByVariableBuilder = GroupByVariableBuilder;
//# sourceMappingURL=groupByVariableBuilder.gen.js.map