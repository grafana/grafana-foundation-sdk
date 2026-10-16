// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

/**
 * A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
 * Panel cells are not represented here — they reuse V2PanelKind.
 */
export class CellKindBuilder implements cog.Builder<dashboardv2beta1.CellKind> {
    protected readonly internal: dashboardv2beta1.CellKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultCellKind();
        this.internal.kind = "Cell";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CellKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.CellSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

