// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class CodeCellContentSpecBuilder implements cog.Builder<dashboardv2beta1.CodeCellContentSpec> {
    protected readonly internal: dashboardv2beta1.CodeCellContentSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultCodeCellContentSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.CodeCellContentSpec {
        return this.internal;
    }

    language(language: string): this {
        this.internal.language = language;
        return this;
    }

    code(code: string): this {
        this.internal.code = code;
        return this;
    }

    highlight(highlight: number[]): this {
        this.internal.highlight = highlight;
        return this;
    }

    annotation(annotation: string): this {
        this.internal.annotation = annotation;
        return this;
    }
}

