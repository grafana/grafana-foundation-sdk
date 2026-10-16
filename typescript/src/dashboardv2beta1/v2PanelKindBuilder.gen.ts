// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

/**
 * The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
 * which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
 * to be forked rather than shared because PanelKind reaches TransformationKind through
 * QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
 * change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
 */
export class V2PanelKindBuilder implements cog.Builder<dashboardv2beta1.V2PanelKind> {
    protected readonly internal: dashboardv2beta1.V2PanelKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultV2PanelKind();
        this.internal.kind = "Panel";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.V2PanelKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.V2PanelSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

