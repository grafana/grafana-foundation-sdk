// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

// --- Common types ---
export class KindBuilder implements cog.Builder<dashboardv2beta1.Kind> {
    protected readonly internal: dashboardv2beta1.Kind;

    constructor() {
        this.internal = dashboardv2beta1.defaultKind();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.Kind {
        return this.internal;
    }

    kind(kind: string): this {
        this.internal.kind = kind;
        return this;
    }

    spec(spec: any): this {
        this.internal.spec = spec;
        return this;
    }

    metadata(metadata: any): this {
        this.internal.metadata = metadata;
        return this;
    }
}

