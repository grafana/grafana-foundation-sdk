import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TextVariableBuilder implements cog.Builder<dashboardv2beta1.TextVariableKind> {
    protected readonly internal: dashboardv2beta1.TextVariableKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TextVariableKind;
    spec(spec: dashboardv2beta1.TextVariableSpec): this;
    name(name: string): this;
    current(current: dashboardv2beta1.VariableOption): this;
    query(query: string): this;
    label(label: string): this;
    hide(hide: dashboardv2beta1.VariableHide): this;
    skipUrlSync(skipUrlSync: boolean): this;
    description(description: string): this;
}
