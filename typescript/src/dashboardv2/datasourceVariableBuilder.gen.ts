// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Datasource variable kind
export class DatasourceVariableBuilder implements cog.Builder<dashboardv2.DatasourceVariableKind> {
    protected readonly internal: dashboardv2.DatasourceVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2.defaultDatasourceVariableKind();
        this.internal.kind = "DatasourceVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.DatasourceVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    pluginId(pluginId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.pluginId = pluginId;
        return this;
    }

    refresh(refresh: dashboardv2.VariableRefresh): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.refresh = refresh;
        return this;
    }

    regex(regex: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.regex = regex;
        return this;
    }

    current(current: dashboardv2.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    includeAll(includeAll: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }

    allValue(allValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    allowCustomValue(allowCustomValue: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }

    origin(origin: cog.Builder<dashboardv2.ControlSourceRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultDatasourceVariableSpec();
        }
        const originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
}

