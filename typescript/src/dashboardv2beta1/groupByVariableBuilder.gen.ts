// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Group variable kind
export class GroupByVariableBuilder implements cog.Builder<dashboardv2beta1.GroupByVariableKind> {
    protected readonly internal: dashboardv2beta1.GroupByVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultGroupByVariableKind();
        this.internal.kind = "GroupByVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.GroupByVariableKind {
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

    spec(spec: dashboardv2beta1.GroupByVariableSpec): this {
        this.internal.spec = spec;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    defaultValue(defaultValue: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.defaultValue = defaultValue;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2beta1.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultGroupByVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}

