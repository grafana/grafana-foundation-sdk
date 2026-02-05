// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class RowsLayoutRowBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutRowKind> {
    protected readonly internal: dashboardv2beta1.RowsLayoutRowKind;

    constructor(title: string) {
        this.internal = dashboardv2beta1.defaultRowsLayoutRowKind();
        this.internal.kind = "RowsLayoutRow";
        this.internal.spec.title = title;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutRowKind {
        return this.internal;
    }

    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    collapse(collapse: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.collapse = collapse;
        return this;
    }

    hideHeader(hideHeader: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.hideHeader = hideHeader;
        return this;
    }

    fillScreen(fillScreen: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.RowRepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    gridLayout(gridLayoutKind: cog.Builder<dashboardv2beta1.GridLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const gridLayoutKindResource = gridLayoutKind.build();
        this.internal.spec.layout = gridLayoutKindResource;
        return this;
    }

    autoGridLayout(autoGridLayoutKind: cog.Builder<dashboardv2beta1.AutoGridLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const autoGridLayoutKindResource = autoGridLayoutKind.build();
        this.internal.spec.layout = autoGridLayoutKindResource;
        return this;
    }

    tabsLayout(tabsLayoutKind: cog.Builder<dashboardv2beta1.TabsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const tabsLayoutKindResource = tabsLayoutKind.build();
        this.internal.spec.layout = tabsLayoutKindResource;
        return this;
    }

    rowsLayout(rowsLayoutKind: cog.Builder<dashboardv2beta1.RowsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const rowsLayoutKindResource = rowsLayoutKind.build();
        this.internal.spec.layout = rowsLayoutKindResource;
        return this;
    }
}

