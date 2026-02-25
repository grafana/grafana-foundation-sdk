// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class SwitchVariableBuilder implements cog.Builder<dashboardv2beta1.SwitchVariableKind> {
    protected readonly internal: dashboardv2beta1.SwitchVariableKind;

    constructor(name: string) {
        this.internal = dashboardv2beta1.defaultSwitchVariableKind();
        this.internal.kind = "SwitchVariable";
        this.internal.spec.name = name;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.SwitchVariableKind {
        return this.internal;
    }

    name(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.name = name;
        return this;
    }

    current(current: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.current = current;
        return this;
    }

    enabledValue(enabledValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.enabledValue = enabledValue;
        return this;
    }

    disabledValue(disabledValue: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.disabledValue = disabledValue;
        return this;
    }

    label(label: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultSwitchVariableSpec();
        }
        this.internal.spec.description = description;
        return this;
    }
}

