// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class ConditionalRenderingDataBuilder implements cog.Builder<dashboardv2.ConditionalRenderingDataKind> {
    protected readonly internal: dashboardv2.ConditionalRenderingDataKind;

    constructor() {
        this.internal = dashboardv2.defaultConditionalRenderingDataKind();
        this.internal.kind = "ConditionalRenderingData";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.ConditionalRenderingDataKind {
        return this.internal;
    }

    value(value: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultConditionalRenderingDataSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}

