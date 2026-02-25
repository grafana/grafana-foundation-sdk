// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class AutoGridBuilder implements cog.Builder<dashboardv2beta1.AutoGridLayoutKind> {
    protected readonly internal: dashboardv2beta1.AutoGridLayoutKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultAutoGridLayoutKind();
        this.internal.kind = "AutoGridLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AutoGridLayoutKind {
        return this.internal;
    }

    maxColumnCount(maxColumnCount: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.maxColumnCount = maxColumnCount;
        return this;
    }

    columnWidthMode(columnWidthMode: "narrow" | "standard" | "wide" | "custom"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.columnWidthMode = columnWidthMode;
        return this;
    }

    columnWidth(columnWidth: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.columnWidth = columnWidth;
        return this;
    }

    rowHeightMode(rowHeightMode: "short" | "standard" | "tall" | "custom"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.rowHeightMode = rowHeightMode;
        return this;
    }

    rowHeight(rowHeight: number): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.rowHeight = rowHeight;
        return this;
    }

    fillScreen(fillScreen: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }

    items(items: cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }

    item(item: cog.Builder<dashboardv2beta1.AutoGridLayoutItemKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }

    withItem(name: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultAutoGridLayoutSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        this.internal.spec.items.push({
        kind: "AutoGridLayoutItem",
        spec: {
        element: {
        kind: "ElementReference",
        name: name,
    },
    },
    });
        return this;
    }
}

export function autoGrid(): AutoGridBuilder {
	const builder = new AutoGridBuilder();

	return builder;
}

