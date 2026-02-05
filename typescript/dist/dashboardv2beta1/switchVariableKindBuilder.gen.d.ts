import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class SwitchVariableKindBuilder implements cog.Builder<dashboardv2beta1.SwitchVariableKind> {
    protected readonly internal: dashboardv2beta1.SwitchVariableKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.SwitchVariableKind;
    spec(spec: cog.Builder<dashboardv2beta1.SwitchVariableSpec>): this;
}
