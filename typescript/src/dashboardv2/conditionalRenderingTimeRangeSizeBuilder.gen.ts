// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ConditionalRenderingTimeRangeSizeBuilder implements cog.Builder<dashboardv2.ConditionalRenderingTimeRangeSizeKind> {
    protected readonly internal: dashboardv2.ConditionalRenderingTimeRangeSizeKind;

    constructor() {
        this.internal = dashboardv2.defaultConditionalRenderingTimeRangeSizeKind();
        this.internal.kind = "ConditionalRenderingTimeRangeSize";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ConditionalRenderingTimeRangeSizeKind {
        return this.internal;
    }

    value(value: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingTimeRangeSizeSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}

