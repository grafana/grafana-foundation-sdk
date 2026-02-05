import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class RowsLayoutRowSpecBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutRowSpec> {
    protected readonly internal: dashboardv2beta1.RowsLayoutRowSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutRowSpec;
    title(title: string): this;
    collapse(collapse: boolean): this;
    hideHeader(hideHeader: boolean): this;
    fillScreen(fillScreen: boolean): this;
    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.RowRepeatOptions>): this;
    layout(layout: cog.Builder<dashboardv2beta1.GridLayoutKind> | cog.Builder<dashboardv2beta1.AutoGridLayoutKind> | cog.Builder<dashboardv2beta1.TabsLayoutKind> | cog.Builder<dashboardv2beta1.RowsLayoutKind>): this;
}
