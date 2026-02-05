import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ActionVariableBuilder implements cog.Builder<dashboardv2beta1.ActionVariable> {
    protected readonly internal: dashboardv2beta1.ActionVariable;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ActionVariable;
    key(key: string): this;
    name(name: string): this;
}
