// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class RowBuilder implements cog.Builder<dashboardv2.RowsLayoutRowKind> {
    protected readonly internal: dashboardv2.RowsLayoutRowKind;

    constructor() {
        this.internal = dashboardv2.defaultRowsLayoutRowKind();
        this.internal.kind = "RowsLayoutRow";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.RowsLayoutRowKind {
        return this.internal;
    }

    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    collapse(collapse: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.collapse = collapse;
        return this;
    }

    hideHeader(hideHeader: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.hideHeader = hideHeader;
        return this;
    }

    fillScreen(fillScreen: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2.ConditionalRenderingGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2.RowRepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    layout(layout: cog.Builder<dashboardv2.GridLayoutKind> | cog.Builder<dashboardv2.AutoGridLayoutKind> | cog.Builder<dashboardv2.TabsLayoutKind> | cog.Builder<dashboardv2.RowsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        const layoutResource = layout.build();
        this.internal.spec.layout = layoutResource;
        return this;
    }

    variables(variables: cog.Builder<dashboardv2.VariableKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultRowsLayoutRowSpec();
        }
        const variablesResources = variables.map(builder1 => builder1.build());
        this.internal.spec.variables = variablesResources;
        return this;
    }
}

export function row(title: string): RowBuilder {
	const builder = new RowBuilder();
	builder.title(title);

	return builder;
}

