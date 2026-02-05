import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AutoGridLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutSpec> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutSpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutSpec;
    maxColumnCount(maxColumnCount: number): this;
    columnWidthMode(columnWidthMode: "narrow" | "standard" | "wide" | "custom"): this;
    columnWidth(columnWidth: number): this;
    rowHeightMode(rowHeightMode: "short" | "standard" | "tall" | "custom"): this;
    rowHeight(rowHeight: number): this;
    fillScreen(fillScreen: boolean): this;
    items(items: cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind>[]): this;
}
