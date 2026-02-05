import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class GroupByVariableBuilder implements cog.Builder<dashboardv2beta1.GroupByVariableKind> {
    protected readonly internal: dashboardv2beta1.GroupByVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.GroupByVariableKind;
    group(group: string): this;
    datasource(datasource: {
        name?: string;
    }): this;
    spec(spec: dashboardv2beta1.GroupByVariableSpec): this;
    name(name: string): this;
    defaultValue(defaultValue: dashboardv2beta1.VariableOption): this;
    current(current: dashboardv2beta1.VariableOption): this;
    options(options: dashboardv2beta1.VariableOption[]): this;
    multi(multi: boolean): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
}
