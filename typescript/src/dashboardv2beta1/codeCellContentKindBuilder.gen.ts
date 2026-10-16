// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class CodeCellContentKindBuilder implements cog.Builder<dashboardv2beta1.CodeCellContentKind> {
    protected readonly internal: dashboardv2beta1.CodeCellContentKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultCodeCellContentKind();
        this.internal.kind = "Code";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CodeCellContentKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.CodeCellContentSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

