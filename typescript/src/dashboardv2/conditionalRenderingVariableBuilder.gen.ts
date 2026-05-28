// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ConditionalRenderingVariableBuilder implements cog.Builder<dashboardv2.ConditionalRenderingVariableKind> {
    protected readonly internal: dashboardv2.ConditionalRenderingVariableKind;

    constructor() {
        this.internal = dashboardv2.defaultConditionalRenderingVariableKind();
        this.internal.kind = "ConditionalRenderingVariable";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ConditionalRenderingVariableKind {
        return this.internal;
    }

    variable(variable: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingVariableSpec();
        }
        this.internal.spec.variable = variable;
        return this;
    }

    operator(operator: "equals" | "notEquals" | "matches" | "notMatches"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingVariableSpec();
        }
        this.internal.spec.operator = operator;
        return this;
    }

    value(value: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingVariableSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}

