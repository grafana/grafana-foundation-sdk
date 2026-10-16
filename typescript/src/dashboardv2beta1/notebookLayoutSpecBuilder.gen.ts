// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

export class NotebookLayoutSpecBuilder implements cog.Builder<dashboardv2beta1.NotebookLayoutSpec> {
    protected readonly internal: dashboardv2beta1.NotebookLayoutSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultNotebookLayoutSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.NotebookLayoutSpec {
        return this.internal;
    }

    cells(cells: cog.Builder<dashboardv2beta1.NotebookLayoutItemKind>[]): this {
        const cellsResources = cells.map(builder1 => builder1.build());
        this.internal.cells = cellsResources;
        return this;
    }
}

