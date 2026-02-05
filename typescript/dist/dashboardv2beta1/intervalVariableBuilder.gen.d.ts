import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class IntervalVariableBuilder implements cog.Builder<dashboardv2beta1.IntervalVariableKind> {
    protected readonly internal: dashboardv2beta1.IntervalVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.IntervalVariableKind;
    spec(spec: dashboardv2beta1.IntervalVariableSpec): this;
    name(name: string): this;
    query(query: string): this;
    current(current: dashboardv2beta1.VariableOption): this;
    options(options: dashboardv2beta1.VariableOption[]): this;
    auto(auto: boolean): this;
    autoMin(autoMin: string): this;
    autoCount(autoCount: number): this;
    refresh(refresh: dashboardv2beta1.VariableRefresh): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
}
