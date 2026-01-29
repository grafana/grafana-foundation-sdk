// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingGroupSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingGroupSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingGroupSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingGroupSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingGroupSpec {
        return this.internal;
    }

    visibility(visibility: "show" | "hide"): this {
        this.internal.visibility = visibility;
        return this;
    }

    condition(condition: "and" | "or"): this {
        this.internal.condition = condition;
        return this;
    }

    items(items: (cog.Builder<dashboardv2beta1.ConditionalRenderingVariableKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> | cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind>)[]): this {
        const itemsResources = items.map(builder1 => builder1.build());
        this.internal.items = itemsResources;
        return this;
    }
}

