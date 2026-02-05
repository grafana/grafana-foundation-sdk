import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class QueryVariableBuilder implements cog.Builder<dashboardv2beta1.QueryVariableKind> {
    protected readonly internal: dashboardv2beta1.QueryVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.QueryVariableKind;
    spec(spec: dashboardv2beta1.QueryVariableSpec): this;
    name(name: string): this;
    current(current: dashboardv2beta1.VariableOption): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    refresh(refresh: dashboardv2beta1.VariableRefresh): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this;
    regex(regex: string): this;
    sort(sort: dashboardv2beta1.VariableSort): this;
    definition(definition: string): this;
    options(options: dashboardv2beta1.VariableOption[]): this;
    multi(multi: boolean): this;
    includeAll(includeAll: boolean): this;
    allValue(allValue: string): this;
    placeholder(placeholder: string): this;
    allowCustomValue(allowCustomValue: boolean): this;
    staticOptions(staticOptions: dashboardv2beta1.VariableOption[]): this;
    staticOptionsOrder(staticOptionsOrder: "before" | "after" | "sorted"): this;
}
