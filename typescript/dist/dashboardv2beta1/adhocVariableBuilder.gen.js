"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdhocVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Adhoc variable kind
class AdhocVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultAdhocVariableKind();
        this.internal.kind = "AdhocVariable";
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
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    baseFilters(baseFilters) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        const baseFiltersResources = baseFilters.map(builder1 => builder1.build());
        this.internal.spec.baseFilters = baseFiltersResources;
        return this;
    }
    filters(filters) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.spec.filters = filtersResources;
        return this;
    }
    defaultKeys(defaultKeys) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        const defaultKeysResources = defaultKeys.map(builder1 => builder1.build());
        this.internal.spec.defaultKeys = defaultKeysResources;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
    allowCustomValue(allowCustomValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
}
exports.AdhocVariableBuilder = AdhocVariableBuilder;
//# sourceMappingURL=adhocVariableBuilder.gen.js.map