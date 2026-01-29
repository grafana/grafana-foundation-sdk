// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingVariableSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingVariableSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingVariableSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingVariableSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingVariableSpec {
        return this.internal;
    }

    variable(variable: string): this {
        this.internal.variable = variable;
        return this;
    }

    operator(operator: "equals" | "notEquals" | "matches" | "notMatches"): this {
        this.internal.operator = operator;
        return this;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

