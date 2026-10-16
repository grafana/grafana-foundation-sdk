// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';

/**
 * One ordered item in a notebook layout. `element` references either a CellKind
 * (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
 * records who authored the cell; `collapsed` hides the body in the UI.
 */
export class NotebookLayoutItemSpecBuilder implements cog.Builder<dashboardv2beta1.NotebookLayoutItemSpec> {
    protected readonly internal: dashboardv2beta1.NotebookLayoutItemSpec;

    constructor() {
        this.internal = dashboardv2beta1.defaultNotebookLayoutItemSpec();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.NotebookLayoutItemSpec {
        return this.internal;
    }

    element(element: cog.Builder<dashboardv2beta1.ElementReference>): this {
        const elementResource = element.build();
        this.internal.element = elementResource;
        return this;
    }

    source(source: "assistant" | "user"): this {
        this.internal.source = source;
        return this;
    }

    collapsed(collapsed: boolean): this {
        this.internal.collapsed = collapsed;
        return this;
    }
}

