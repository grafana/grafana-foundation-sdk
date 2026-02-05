import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class SwitchVariableSpecBuilder implements cog.Builder<dashboardv2beta1.SwitchVariableSpec> {
    protected readonly internal: dashboardv2beta1.SwitchVariableSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.SwitchVariableSpec;
    name(name: string): this;
    current(current: string): this;
    enabledValue(enabledValue: string): this;
    disabledValue(disabledValue: string): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
}
