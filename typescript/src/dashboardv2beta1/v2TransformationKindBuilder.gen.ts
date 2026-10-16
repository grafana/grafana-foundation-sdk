// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

/**
 * Dashboard v2 shape: the transformation ID moved from `kind` to `group`.
 */
export class V2TransformationKindBuilder implements cog.Builder<dashboardv2beta1.V2TransformationKind> {
    protected readonly internal: dashboardv2beta1.V2TransformationKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultV2TransformationKind();
        this.internal.kind = "Transformation";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.V2TransformationKind {
        return this.internal;
    }

    // The group is the transformation ID
    group(group: string): this {
        this.internal.group = group;
        return this;
    }

    spec(spec: cog.Builder<dashboardv2beta1.V2TransformationSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

