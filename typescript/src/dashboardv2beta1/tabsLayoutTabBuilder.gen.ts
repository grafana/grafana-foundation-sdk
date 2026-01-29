// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TabsLayoutTabBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutTabKind> {
    protected readonly internal: dashboardv2beta1.TabsLayoutTabKind;

    constructor(title: string) {
        this.internal = dashboardv2beta1.defaultTabsLayoutTabKind();
        this.internal.kind = "TabsLayoutTab";
        this.internal.spec.title = title;
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutTabKind {
        return this.internal;
    }

    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    gridLayout(gridLayoutKind: cog.Builder<dashboardv2beta1.GridLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const gridLayoutKindResource = gridLayoutKind.build();
        this.internal.spec.layout = gridLayoutKindResource;
        return this;
    }

    rowsLayout(rowsLayoutKind: cog.Builder<dashboardv2beta1.RowsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const rowsLayoutKindResource = rowsLayoutKind.build();
        this.internal.spec.layout = rowsLayoutKindResource;
        return this;
    }

    autoGridLayout(autoGridLayoutKind: cog.Builder<dashboardv2beta1.AutoGridLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const autoGridLayoutKindResource = autoGridLayoutKind.build();
        this.internal.spec.layout = autoGridLayoutKindResource;
        return this;
    }

    tabsLayout(tabsLayoutKind: cog.Builder<dashboardv2beta1.TabsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const tabsLayoutKindResource = tabsLayoutKind.build();
        this.internal.spec.layout = tabsLayoutKindResource;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.TabRepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
}

