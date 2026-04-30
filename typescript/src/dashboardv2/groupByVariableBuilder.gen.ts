// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Group variable kind
export class GroupByVariableBuilder implements cog.Builder<dashboardv2.GroupByVariableKind> {
    protected readonly internal: dashboardv2.GroupByVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2.defaultGroupByVariableKind();
        this.internal.kind = "GroupByVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.GroupByVariableKind {
        return this.internal;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    labels(labels: Record<string, string>): this {
        this.internal.labels = labels;
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
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    defaultValue(defaultValue: dashboardv2.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.defaultValue = defaultValue;
        return this;
    }

    current(current: dashboardv2.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    origin(origin: cog.Builder<dashboardv2.ControlSourceRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultGroupByVariableSpec();
        }
        const originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
}

