// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TabBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutTabKind> {
    protected readonly internal: dashboardv2beta1.TabsLayoutTabKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutTabKind();
        this.internal.kind = "TabsLayoutTab";
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

    layout(layout: cog.Builder<dashboardv2beta1.GridLayoutKind> | cog.Builder<dashboardv2beta1.RowsLayoutKind> | cog.Builder<dashboardv2beta1.AutoGridLayoutKind> | cog.Builder<dashboardv2beta1.TabsLayoutKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultTabsLayoutTabSpec();
        }
        const layoutResource = layout.build();
        this.internal.spec.layout = layoutResource;
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

export function tab(title: string): TabBuilder {
	const builder = new TabBuilder();
	builder.title(title);

	return builder;
}

