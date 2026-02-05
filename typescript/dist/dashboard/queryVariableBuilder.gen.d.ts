import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as common from '../common';
export declare class QueryVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboard.VariableModel;
    name(name: string): this;
    label(label: string): this;
    hide(hide: dashboard.VariableHide): this;
    description(description: string): this;
    query(query: string | Record<string, any>): this;
    datasource(datasource: common.DataSourceRef): this;
    current(current: dashboard.VariableOption): this;
    multi(multi: boolean): this;
    allowCustomValue(allowCustomValue: boolean): this;
    options(options: dashboard.VariableOption[]): this;
    refresh(refresh: dashboard.VariableRefresh): this;
    sort(sort: dashboard.VariableSort): this;
    includeAll(includeAll: boolean): this;
    allValue(allValue: string): this;
    regex(regex: string): this;
}
