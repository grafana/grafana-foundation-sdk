import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class CustomFormatterVariableBuilder implements cog.Builder<dashboardv2beta1.CustomFormatterVariable> {
    protected readonly internal: dashboardv2beta1.CustomFormatterVariable;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CustomFormatterVariable;
    name(name: string): this;
    type(type: dashboardv2beta1.VariableType): this;
    multi(multi: boolean): this;
    includeAll(includeAll: boolean): this;
}
