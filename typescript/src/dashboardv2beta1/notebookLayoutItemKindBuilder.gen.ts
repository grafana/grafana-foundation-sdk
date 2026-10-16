// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class NotebookLayoutItemKindBuilder implements cog.Builder<dashboardv2beta1.NotebookLayoutItemKind> {
    protected readonly internal: dashboardv2beta1.NotebookLayoutItemKind;

    constructor() {
        this.internal = dashboardv2beta1.defaultNotebookLayoutItemKind();
        this.internal.kind = "NotebookLayoutItem";
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.NotebookLayoutItemKind {
        return this.internal;
    }

    spec(spec: cog.Builder<dashboardv2beta1.NotebookLayoutItemSpec>): this {
        const specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
}

