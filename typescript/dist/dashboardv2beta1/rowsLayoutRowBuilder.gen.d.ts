import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class RowsLayoutRowBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutRowKind> {
    protected readonly internal: dashboardv2beta1.RowsLayoutRowKind;
    constructor(title: string);
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutRowKind;
    title(title: string): this;
    collapse(collapse: boolean): this;
    hideHeader(hideHeader: boolean): this;
    fillScreen(fillScreen: boolean): this;
    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this;
    repeat(repeat: cog.Builder<dashboardv2beta1.RowRepeatOptions>): this;
    gridLayout(gridLayoutKind: cog.Builder<dashboardv2beta1.GridLayoutKind>): this;
    autoGridLayout(autoGridLayoutKind: cog.Builder<dashboardv2beta1.AutoGridLayoutKind>): this;
    tabsLayout(tabsLayoutKind: cog.Builder<dashboardv2beta1.TabsLayoutKind>): this;
    rowsLayout(rowsLayoutKind: cog.Builder<dashboardv2beta1.RowsLayoutKind>): this;
}
