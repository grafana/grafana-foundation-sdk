import * as cog from '../cog';
import * as dashboard from '../dashboard';
import * as common from '../common';
export declare class AdHocVariableBuilder implements cog.Builder<dashboard.VariableModel> {
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
    datasource(datasource: common.DataSourceRef): this;
    allowCustomValue(allowCustomValue: boolean): this;
}
