// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class NotebookLayoutKindBuilder implements cog.Builder<dashboardv2beta1.NotebookLayoutKind> {
    protected readonly internal: dashboardv2beta1.NotebookLayoutKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultNotebookLayoutKind();
        this.internal.kind = "NotebookLayout";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.NotebookLayoutKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.NotebookLayoutSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

