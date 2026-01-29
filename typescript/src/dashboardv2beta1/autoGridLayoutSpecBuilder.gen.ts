// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class AutoGridLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutSpec> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutSpec {
        return this.internal;
    }

    maxColumnCount(maxColumnCount: number): this {
        this.internal.maxColumnCount = maxColumnCount;
        return this;
    }

    columnWidthMode(columnWidthMode: "narrow" | "standard" | "wide" | "custom"): this {
        this.internal.columnWidthMode = columnWidthMode;
        return this;
    }

    columnWidth(columnWidth: number): this {
        this.internal.columnWidth = columnWidth;
        return this;
    }

    rowHeightMode(rowHeightMode: "short" | "standard" | "tall" | "custom"): this {
        this.internal.rowHeightMode = rowHeightMode;
        return this;
    }

    rowHeight(rowHeight: number): this {
        this.internal.rowHeight = rowHeight;
        return this;
    }

    fillScreen(fillScreen: boolean): this {
        this.internal.fillScreen = fillScreen;
        return this;
    }

    items(items: cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind>[]): this {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }
}

