import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AutoGridLayoutItemSpecBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutItemSpec> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutItemSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutItemSpec;
    element(element: cog.Builder<dashboardv2beta1.ElementReference>): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.AutoGridRepeatOptions>): this;
    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this;
}
