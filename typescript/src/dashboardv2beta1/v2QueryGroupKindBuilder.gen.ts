// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class V2QueryGroupKindBuilder implements cog.Builder<dashboardv2beta1.V2QueryGroupKind> {
    protected readonly internal: dashboardv2beta1.V2QueryGroupKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultV2QueryGroupKind();
        this.internal.kind = "QueryGroup";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.V2QueryGroupKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.V2QueryGroupSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

