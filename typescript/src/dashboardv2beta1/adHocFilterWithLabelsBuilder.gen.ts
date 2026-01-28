// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// Define the AdHocFilterWithLabels type
export class AdHocFilterWithLabelsBuilder implements cog.Builder<dashboardv2beta1.AdHocFilterWithLabels> {
    protected readonly internal: dashboardv2beta1.AdHocFilterWithLabels;

    constructor() {
        this.internal = dashboardv2beta1.defaultAdHocFilterWithLabels();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AdHocFilterWithLabels {
        return this.internal;
    }

    key(key: string): this {
        this.internal.key = key;
        return this;
    }

    operator(operator: string): this {
        this.internal.operator = operator;
        return this;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }

    values(values: string[]): this {
        this.internal.values = values;
        return this;
    }

    keyLabel(keyLabel: string): this {
        this.internal.keyLabel = keyLabel;
        return this;
    }

    valueLabels(valueLabels: string[]): this {
        this.internal.valueLabels = valueLabels;
        return this;
    }

    forceEdit(forceEdit: boolean): this {
        this.internal.forceEdit = forceEdit;
        return this;
    }

    origin(origin: "dashboard"): this {
        this.internal.origin = origin;
        return this;
    }

    // @deprecated
    condition(condition: string): this {
        this.internal.condition = condition;
        return this;
    }
}

