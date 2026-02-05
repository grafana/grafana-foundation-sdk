import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class VariableValueOptionBuilder implements cog.Builder<dashboardv2beta1.VariableValueOption> {
    protected readonly internal: dashboardv2beta1.VariableValueOption;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.VariableValueOption;
    label(label: string): this;
    value(value: dashboardv2beta1.VariableValueSingle): this;
    group(group: string): this;
}
