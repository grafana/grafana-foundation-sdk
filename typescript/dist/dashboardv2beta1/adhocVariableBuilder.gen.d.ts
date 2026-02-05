import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AdhocVariableBuilder implements cog.Builder<dashboardv2beta1.AdhocVariableKind> {
    protected readonly internal: dashboardv2beta1.AdhocVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AdhocVariableKind;
    group(group: string): this;
    datasource(datasource: {
        name?: string;
    }): this;
    spec(spec: dashboardv2beta1.AdhocVariableSpec): this;
    name(name: string): this;
    baseFilters(baseFilters: cog.Builder<dashboardv2beta1.AdHocFilterWithLabels>[]): this;
    filters(filters: cog.Builder<dashboardv2beta1.AdHocFilterWithLabels>[]): this;
    defaultKeys(defaultKeys: cog.Builder<dashboardv2beta1.MetricFindValue>[]): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
    allowCustomValue(allowCustomValue: boolean): this;
}
