// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingTimeRangeSizeBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingTimeRangeSizeKind();
        this.internal.kind = "ConditionalRenderingTimeRangeSize";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind {
        return this.internal;
    }

    value(value: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingTimeRangeSizeSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}

