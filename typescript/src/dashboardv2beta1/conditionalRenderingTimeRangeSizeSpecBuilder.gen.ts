// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingTimeRangeSizeSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingTimeRangeSizeSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec {
        return this.internal;
    }

    value(value: string): this {
        this.internal.value = value;
        return this;
    }
}

