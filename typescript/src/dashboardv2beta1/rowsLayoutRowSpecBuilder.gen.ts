// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class RowsLayoutRowSpecBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutRowSpec> {
    protected readonly internal: dashboardv2beta1.RowsLayoutRowSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutRowSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.RowsLayoutRowSpec {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    collapse(collapse: boolean): this {
        this.internal.collapse = collapse;
        return this;
    }

    hideHeader(hideHeader: boolean): this {
        this.internal.hideHeader = hideHeader;
        return this;
    }

    fillScreen(fillScreen: boolean): this {
        this.internal.fillScreen = fillScreen;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this {
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.RowRepeatOptions>): this {
        const repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }

    layout(layout: cog.Builder<dashboardv2beta1.GridLayoutKind> | cog.Builder<dashboardv2beta1.AutoGridLayoutKind> | cog.Builder<dashboardv2beta1.TabsLayoutKind> | cog.Builder<dashboardv2beta1.RowsLayoutKind>): this {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }
}

