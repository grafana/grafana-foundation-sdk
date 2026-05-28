// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class SwitchVariableBuilder implements cog.Builder<dashboardv2.SwitchVariableKind> {
    protected readonly internal: dashboardv2.SwitchVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2.defaultSwitchVariableKind();
        this.internal.kind = "SwitchVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.SwitchVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    current(current: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    enabledValue(enabledValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.enabledValue = enabledValue;
        return this;
    }

    disabledValue(disabledValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.disabledValue = disabledValue;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }

    origin(origin: cog.Builder<dashboardv2.ControlSourceRef>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultSwitchVariableSpec();
        }
        const originResource = origin.build();
        this.internal.spec.origin = originResource;
        return this;
    }
}

