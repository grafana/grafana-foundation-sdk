// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Text variable kind
export class TextVariableBuilder implements cog.Builder<dashboardv2beta1.TextVariableKind> {
    protected readonly internal: dashboardv2beta1.TextVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultTextVariableKind();
        this.internal.kind = "TextVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TextVariableKind {
        return this.internal;
    }

    spec(spec: dashboardv2beta1.TextVariableSpec): this {
        this.internal.spec = spec;
        return this;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    current(current: dashboardv2beta1.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTextVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}

