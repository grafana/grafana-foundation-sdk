import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConditionalRenderingVariableSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingVariableSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingVariableSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingVariableSpec;
    variable(variable: string): this;
    operator(operator: "equals" | "notEquals" | "matches" | "notMatches"): this;
    value(value: string): this;
}
