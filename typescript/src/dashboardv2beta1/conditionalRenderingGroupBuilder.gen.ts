// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingGroupBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingGroupKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingGroupKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingGroupKind();
        this.internal.kind = "ConditionalRenderingGroup";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingGroupKind {
        return this.internal;
    }

    visibility(visibility: "show" | "hide"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        this.internal.spec.visibility = visibility;
        return this;
    }

    condition(condition: "and" | "or"): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        this.internal.spec.condition = condition;
        return this;
    }

    items(items: (cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind>)[]): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.spec.items = itemsResources;
        return this;
    }

    item(item: cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
        }
        if (!this.internal.spec.items) {
            this.internal.spec.items = [];
        }
        const itemResource = item.build();
        this.internal.spec.items.push(itemResource);
        return this;
    }
}

