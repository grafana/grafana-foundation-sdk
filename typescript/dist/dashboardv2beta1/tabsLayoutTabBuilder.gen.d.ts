import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TabsLayoutTabBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutTabKind> {
    protected readonly internal: dashboardv2beta1.TabsLayoutTabKind;
    constructor(title: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutTabKind;
    title(title: string): this;
    gridLayout(gridLayoutKind: cog.Builder<dashboardv2beta1.GridLayoutKind>): this;
    rowsLayout(rowsLayoutKind: cog.Builder<dashboardv2beta1.RowsLayoutKind>): this;
    autoGridLayout(autoGridLayoutKind: cog.Builder<dashboardv2beta1.AutoGridLayoutKind>): this;
    tabsLayout(tabsLayoutKind: cog.Builder<dashboardv2beta1.TabsLayoutKind>): this;
    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.TabRepeatOptions>): this;
}
