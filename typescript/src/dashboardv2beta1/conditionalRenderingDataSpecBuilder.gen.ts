// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingDataSpecBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingDataSpec> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingDataSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingDataSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingDataSpec {
        return this.internal;
    }

    value(value: boolean): this {
        this.internal.value = value;
        return this;
    }
}

