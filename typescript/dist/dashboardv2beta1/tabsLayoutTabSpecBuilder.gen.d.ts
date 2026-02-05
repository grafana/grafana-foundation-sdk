import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TabsLayoutTabSpecBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutTabSpec> {
    protected readonly internal: dashboardv2beta1.TabsLayoutTabSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutTabSpec;
    title(title: string): this;
    layout(layout: cog.Builder<dashboardv2beta1.GridLayoutKind> | cog.Builder<dashboardv2beta1.RowsLayoutKind> | cog.Builder<dashboardv2beta1.AutoGridLayoutKind> | cog.Builder<dashboardv2beta1.TabsLayoutKind>): this;
    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.TabRepeatOptions>): this;
}
