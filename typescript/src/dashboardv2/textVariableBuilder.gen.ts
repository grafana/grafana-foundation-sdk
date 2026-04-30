// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Text variable kind
export class TextVariableBuilder implements cog.Builder<dashboardv2.TextVariableKind> {
    protected readonly internal: dashboardv2.TextVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2.defaultTextVariableKind();
        this.internal.kind = "TextVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.TextVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    current(current: dashboardv2.VariableOption): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    query(query: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.query = query;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    origin(origin: cog.Builder<dashboardv2.ControlSourceRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTextVariableSpec();
        }
        const originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
}

