// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class MarkdownCellContentKindBuilder implements cog.Builder<dashboardv2beta1.MarkdownCellContentKind> {
    protected readonly internal: dashboardv2beta1.MarkdownCellContentKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultMarkdownCellContentKind();
        this.internal.kind = "Markdown";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.MarkdownCellContentKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.MarkdownCellContentSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

