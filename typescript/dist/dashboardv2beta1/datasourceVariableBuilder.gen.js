"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DatasourceVariableBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
// Datasource variable kind
class DatasourceVariableBuilder {
    constructor(name) {
        this.internal = dashboardv2beta1.defaultDatasourceVariableKind();
        this.internal.kind = "DatasourceVariable";
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
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }
    pluginId(pluginId) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.pluginId = pluginId;
        return this;
    }
    refresh(refresh) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }
    regex(regex) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.regex = regex;
        return this;
    }
    current(current) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }
    options(options) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }
    multi(multi) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }
    includeAll(includeAll) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }
    allValue(allValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }
    label(label) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }
    hide(hide) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }
    skipUrlSync(skipUrlSync) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }
    description(description) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
    allowCustomValue(allowCustomValue) {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
}
exports.DatasourceVariableBuilder = DatasourceVariableBuilder;
//# sourceMappingURL=datasourceVariableBuilder.gen.js.map