// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class TargetBuilder implements cog.Builder<dashboardv2.PanelQueryKind> {
    protected readonly internal: dashboardv2.PanelQueryKind;

    constructor() {
        this.internal = dashboardv2.defaultPanelQueryKind();
        this.internal.kind = "PanelQuery";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.PanelQueryKind {
        return this.internal;
    }

    query(query: cog.Builder<dashboardv2.DataQueryKind>): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelQuerySpec();
        }
        const queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }

    refId(refId: string): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelQuerySpec();
        }
        this.internal.spec.refId = refId;
        return this;
    }

    hidden(hidden: boolean): this {
        if (!this.internal.spec) {
            this.internal.spec = dashboardv2.defaultPanelQuerySpec();
        }
        this.internal.spec.hidden = hidden;
        return this;
    }
}

