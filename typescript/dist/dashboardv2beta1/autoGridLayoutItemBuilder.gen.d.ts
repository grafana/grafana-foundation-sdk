import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AutoGridLayoutItemBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutItemKind;
    constructor(name: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutItemKind;
    element(element: cog.Builder<dashboardv2beta1.ElementReference>): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.AutoGridRepeatOptions>): this;
    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this;
    name(name: string): this;
}
