// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Adhoc variable kind
export class AdhocVariableBuilder implements cog.Builder<dashboardv2beta1.AdhocVariableKind> {
    protected readonly internal: dashboardv2beta1.AdhocVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultAdhocVariableKind();
        this.internal.kind = "AdhocVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AdhocVariableKind {
        return this.internal;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    datasource(datasource: {
	name?: string;
}): this {
        this.internal.datasource = datasource;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    baseFilters(baseFilters: cog.Builder<dashboardv2beta1.AdHocFilterWithLabels>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        const baseFiltersResources = baseFilters.map(builder1 => builder1.build());
        this.internal.spec.baseFilters = baseFiltersResources;
        return this;
    }

    filters(filters: cog.Builder<dashboardv2beta1.AdHocFilterWithLabels>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        const filtersResources = filters.map(builder1 => builder1.build());
        this.internal.spec.filters = filtersResources;
        return this;
    }

    defaultKeys(defaultKeys: cog.Builder<dashboardv2beta1.MetricFindValue>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        const defaultKeysResources = defaultKeys.map(builder1 => builder1.build());
        this.internal.spec.defaultKeys = defaultKeysResources;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    allowCustomValue(allowCustomValue: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAdhocVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
}

