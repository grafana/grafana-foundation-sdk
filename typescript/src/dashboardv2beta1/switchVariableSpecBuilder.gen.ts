// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class SwitchVariableSpecBuilder implements cog.Builder<dashboardv2beta1.SwitchVariableSpec> {
    protected readonly internal: dashboardv2beta1.SwitchVariableSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultSwitchVariableSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.SwitchVariableSpec {
        return this.internal;
    }

    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    current(current: string): this {
        this.internal.current = current;
        return this;
    }

    enabledValue(enabledValue: string): this {
        this.internal.enabledValue = enabledValue;
        return this;
    }

    disabledValue(disabledValue: string): this {
        this.internal.disabledValue = disabledValue;
        return this;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    hide(hide: dashboardv2beta1.VariableHide): this {
        this.internal.hide = hide;
        return this;
    }

    skipUrlSync(skipUrlSync: boolean): this {
        this.internal.skipUrlSync = skipUrlSync;
        return this;
    }

    description(description: string): this {
        this.internal.description = description;
        return this;
    }
}

