// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class TabBuilder implements cog.Builder<dashboardv2.TabsLayoutTabKind> {
    protected readonly internal: dashboardv2.TabsLayoutTabKind;

    constructor() {
        this.internal = dashboardv2.defaultTabsLayoutTabKind();
        this.internal.kind = "TabsLayoutTab";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.TabsLayoutTabKind {
        return this.internal;
    }

    title(title: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutTabSpec();
        }
        this.internal.spec.title = title;
        return this;
    }

    layout(layout: cog.Builder<dashboardv2.GridLayoutKind> | cog.Builder<dashboardv2.RowsLayoutKind> | cog.Builder<dashboardv2.AutoGridLayoutKind> | cog.Builder<dashboardv2.TabsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutTabSpec();
        }
        const layoutResource = layout.build();
        this.internal.spec.layout = layoutResource;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2.ConditionalRenderingGroupKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutTabSpec();
        }
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2.TabRepeatOptions>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutTabSpec();
        }
        const repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }

    variables(variables: cog.Builder<dashboardv2.VariableKind>[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultTabsLayoutTabSpec();
        }
        const variablesResources = variables.map(builder1 => builder1.build());
        this.internal.spec.variables = variablesResources;
        return this;
    }
}

export function tab(title: string): TabBuilder {
	const builder = new TabBuilder();
	builder.title(title);

	return builder;
}

