import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class CustomVariableBuilder implements cog.Builder<dashboard.VariableModel> {
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
    values(query: string | Record<string, any>): this;
    current(current: dashboard.VariableOption): this;
    multi(multi: boolean): this;
    allowCustomValue(allowCustomValue: boolean): this;
    options(options: dashboard.VariableOption[]): this;
    includeAll(includeAll: boolean): this;
    allValue(allValue: string): this;
}
