// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// FIXME: should we introduce this? --- Variable value option
export class VariableValueOptionBuilder implements cog.Builder<dashboardv2beta1.VariableValueOption> {
    protected readonly internal: dashboardv2beta1.VariableValueOption;

    constructor() {
        this.internal = dashboardv2beta1.defaultVariableValueOption();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.VariableValueOption {
        return this.internal;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    value(value: dashboardv2beta1.VariableValueSingle): this {
        this.internal.value = value;
        return this;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }
}

