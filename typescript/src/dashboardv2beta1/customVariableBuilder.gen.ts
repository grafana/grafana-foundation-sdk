// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Custom variable kind
export class CustomVariableBuilder implements cog.Builder<dashboardv2beta1.CustomVariableKind> {
    protected readonly internal: dashboardv2beta1.CustomVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultCustomVariableKind();
        this.internal.kind = "CustomVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CustomVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    options(options: dashboardv2beta1.VariableOption[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.options = options;
        return this;
    }

    multi(multi: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.multi = multi;
        return this;
    }

    includeAll(includeAll: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.includeAll = includeAll;
        return this;
    }

    allValue(allValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.allValue = allValue;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    allowCustomValue(allowCustomValue: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultCustomVariableSpec();
        }
        this.internal.spec.allowCustomValue = allowCustomValue;
        return this;
    }
}

