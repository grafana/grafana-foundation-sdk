// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class CellSpecBuilder implements cog.Builder<dashboardv2beta1.CellSpec> {
    protected readonly internal: dashboardv2beta1.CellSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultCellSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CellSpec {
        return this.internal;
    }

    content(content: cog.Builder<dashboardv2beta1.CellContentKind>): this {
        const contentResource = content.build();
        this.internal.content = contentResource;
        return this;
    }
}

