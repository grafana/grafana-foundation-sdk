// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class RowBuilder implements cog.Builder<dashboardv2beta1.RowsLayoutRowKind> {
    protected readonly internal: dashboardv2beta1.RowsLayoutRowKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultRowsLayoutRowKind();
        this.internal.kind = "RowsLayoutRow";
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

    layout(layout: cog.Builder<dashboardv2beta1.GridLayoutKind> | cog.Builder<dashboardv2beta1.AutoGridLayoutKind> | cog.Builder<dashboardv2beta1.TabsLayoutKind> | cog.Builder<dashboardv2beta1.RowsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultRowsLayoutRowSpec();
        }
        const layoutResource = layout.build();
        this.internal.spec.layout = layoutResource;
        return this;
    }
}

export function row(title: string): RowBuilder {
	const builder = new RowBuilder();
	builder.title(title);

	return builder;
}

