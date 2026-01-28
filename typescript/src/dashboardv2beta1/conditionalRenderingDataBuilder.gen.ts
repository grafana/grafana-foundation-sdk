// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class ConditionalRenderingDataBuilder implements cog.Builder<dashboardv2beta1.ConditionalRenderingDataKind> {
    protected readonly internal: dashboardv2beta1.ConditionalRenderingDataKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultConditionalRenderingDataKind();
        this.internal.kind = "ConditionalRenderingData";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ConditionalRenderingDataKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.ConditionalRenderingDataSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }

    value(value: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultConditionalRenderingDataSpec();
        }
        this.internal.spec.value = value;
        return this;
    }
}

