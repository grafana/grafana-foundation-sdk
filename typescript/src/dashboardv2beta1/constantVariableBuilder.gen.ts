// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Constant variable kind
export class ConstantVariableBuilder implements cog.Builder<dashboardv2beta1.ConstantVariableKind> {
    protected readonly internal: dashboardv2beta1.ConstantVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultConstantVariableKind();
        this.internal.kind = "ConstantVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConstantVariableKind {
        return this.internal;
    }

    spec(spec: dashboardv2beta1.ConstantVariableSpec): this {
        this.internal.spec = spec;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConstantVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}

