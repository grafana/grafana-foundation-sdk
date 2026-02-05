import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AdHocFilterWithLabelsBuilder implements cog.Builder<dashboardv2beta1.AdHocFilterWithLabels> {
    protected readonly internal: dashboardv2beta1.AdHocFilterWithLabels;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AdHocFilterWithLabels;
    key(key: string): this;
    operator(operator: string): this;
    value(value: string): this;
    values(values: string[]): this;
    keyLabel(keyLabel: string): this;
    valueLabels(valueLabels: string[]): this;
    forceEdit(forceEdit: boolean): this;
    origin(origin: "dashboard"): this;
    condition(condition: string): this;
}
