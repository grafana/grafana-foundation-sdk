import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class ConstantVariableBuilder implements cog.Builder<dashboard.VariableModel> {
    protected readonly internal: dashboard.VariableModel;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboard.VariableModel;
    name(name: string): this;
    label(label: string): this;
    description(description: string): this;
    value(query: string | Record<string, any>): this;
    allowCustomValue(allowCustomValue: boolean): this;
}
