// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class MarkdownCellContentSpecBuilder implements cog.Builder<dashboardv2beta1.MarkdownCellContentSpec> {
    protected readonly internal: dashboardv2beta1.MarkdownCellContentSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultMarkdownCellContentSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.MarkdownCellContentSpec {
        return this.internal;
    }

    text(text: string): this {
        this.internal.text = text;
        return this;
    }
}

