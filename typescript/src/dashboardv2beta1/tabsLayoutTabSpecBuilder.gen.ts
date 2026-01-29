// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TabsLayoutTabSpecBuilder implements cog.Builder<dashboardv2beta1.TabsLayoutTabSpec> {
    protected readonly internal: dashboardv2beta1.TabsLayoutTabSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultTabsLayoutTabSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.TabsLayoutTabSpec {
        return this.internal;
    }

    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    layout(layout: cog.Builder<dashboardv2beta1.GridLayoutKind> | cog.Builder<dashboardv2beta1.RowsLayoutKind> | cog.Builder<dashboardv2beta1.AutoGridLayoutKind> | cog.Builder<dashboardv2beta1.TabsLayoutKind>): this {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }

    conditionalRendering(conditionalRendering: cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind>): this {
        const conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }

    repeat(repeat: cog.Builder<dashboardv2beta1.TabRepeatOptions>): this {
        const repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }
}

