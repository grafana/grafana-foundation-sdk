// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ConditionalRenderingGroupBuilder implements cog.Builder<dashboardv2.ConditionalRenderingGroupKind> {
    protected readonly internal: dashboardv2.ConditionalRenderingGroupKind;

    constructor() {
        this.internal = dashboardv2.defaultConditionalRenderingGroupKind();
        this.internal.kind = "ConditionalRenderingGroup";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ConditionalRenderingGroupKind {
        return this.internal;
    }

    visibility(visibility: "show" | "hide"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingGroupSpec();
        }
        this.internal.spec.visibility = visibility;
        return this;
    }

    condition(condition: "and" | "or"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingGroupSpec();
        }
        this.internal.spec.condition = condition;
        return this;
    }

    items(items: (cog.Builder<dashboardv2.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2.ConditionalRenderingDataKind> | cog.Builder<dashboardv2.ConditionalRenderingTimeRangeSizeKind>)[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingGroupSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }

    item(item: cog.Builder<dashboardv2.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2.ConditionalRenderingDataKind> | cog.Builder<dashboardv2.ConditionalRenderingTimeRangeSizeKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingGroupSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }
}

