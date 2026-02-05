import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ConstantVariableBuilder implements cog.Builder<dashboardv2beta1.ConstantVariableKind> {
    protected readonly internal: dashboardv2beta1.ConstantVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConstantVariableKind;
    spec(spec: dashboardv2beta1.ConstantVariableSpec): this;
    name(name: string): this;
    query(query: string): this;
    current(current: dashboardv2beta1.VariableOption): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
}
