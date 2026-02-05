// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class TargetBuilder implements cog.Builder<dashboardv2beta1.PanelQueryKind> {
    protected readonly internal: dashboardv2beta1.PanelQueryKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultPanelQueryKind();
        this.internal.kind = "PanelQuery";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.PanelQueryKind {
        return this.internal;
    }

    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelQuerySpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }

    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelQuerySpec();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    hidden(hidden: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2beta1.defaultPanelQuerySpec();
        }
        this.internal.spec.hidden = hidden;
        return this;
    }
}

