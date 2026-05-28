// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// FIXME: should we introduce this? --- Variable value option
export class VariableValueOptionBuilder implements cog.Builder<dashboardv2.VariableValueOption> {
    protected readonly internal: dashboardv2.VariableValueOption;

    constructor() {
        this.internal = dashboardv2.defaultVariableValueOption();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.VariableValueOption {
        return this.internal;
    }

    label(label: string): this {
        this.internal.label = label;
        return this;
    }

    value(value: dashboardv2.VariableValueSingle): this {
        this.internal.value = value;
        return this;
    }

    group(group: string): this {
        this.internal.group = group;
        return this;
    }
}

