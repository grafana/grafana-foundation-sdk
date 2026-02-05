// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Datasource variable kind
export class DatasourceVariableBuilder implements cog.Builder<dashboardv2beta1.DatasourceVariableKind> {
    protected readonly internal: dashboardv2beta1.DatasourceVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultDatasourceVariableKind();
        this.internal.kind = "DatasourceVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DatasourceVariableKind {
        return this.internal;
    }

    spec(spec: dashboardv2beta1.DatasourceVariableSpec): this {
        this.internal.spec = spec;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    pluginId(pluginId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.pluginId = pluginId;
        return this;
    }

    refresh(refresh: dashboardv2beta1.VariableRefresh): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }

    regex(regex: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.regex = regex;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2beta1.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    includeAll(includeAll: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }

    allValue(allValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    allowCustomValue(allowCustomValue: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultDatasourceVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
}

