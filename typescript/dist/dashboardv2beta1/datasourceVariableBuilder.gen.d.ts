import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class DatasourceVariableBuilder implements cog.Builder<dashboardv2beta1.DatasourceVariableKind> {
    protected readonly internal: dashboardv2beta1.DatasourceVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.DatasourceVariableKind;
    spec(spec: dashboardv2beta1.DatasourceVariableSpec): this;
    name(name: string): this;
    pluginId(pluginId: string): this;
    refresh(refresh: dashboardv2beta1.VariableRefresh): this;
    regex(regex: string): this;
    current(current: dashboardv2beta1.VariableOption): this;
    options(options: dashboardv2beta1.VariableOption[]): this;
    multi(multi: boolean): this;
    includeAll(includeAll: boolean): this;
    allValue(allValue: string): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
    allowCustomValue(allowCustomValue: boolean): this;
}
