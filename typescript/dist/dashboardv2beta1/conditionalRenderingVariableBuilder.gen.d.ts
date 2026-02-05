import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingVariableBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingVariableKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingVariableKind;
    spec(spec: cog.Builder<dashboardv2beta1.ConditionalRenderingVariableSpec>): this;
    variable(variable: string): this;
    operator(operator: "equals" | "notEquals" | "matches" | "notMatches"): this;
    value(value: string): this;
}
