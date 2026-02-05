import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AutoGridLayoutBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutKind> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutKind;
    maxColumnCount(maxColumnCount: number): this;
    columnWidthMode(columnWidthMode: "narrow" | "standard" | "wide" | "custom"): this;
    columnWidth(columnWidth: number): this;
    rowHeightMode(rowHeightMode: "short" | "standard" | "tall" | "custom"): this;
    rowHeight(rowHeight: number): this;
    fillScreen(fillScreen: boolean): this;
    items(items: cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind>[]): this;
    item(item: cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind>): this;
}
