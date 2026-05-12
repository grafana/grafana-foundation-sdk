// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

export class AnnotationPanelFilterBuilder implements cog.Builder<dashboardv2.AnnotationPanelFilter> {
    protected readonly internal: dashboardv2.AnnotationPanelFilter;

    constructor() {
        this.internal = dashboardv2.defaultAnnotationPanelFilter();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.AnnotationPanelFilter {
        return this.internal;
    }

    // Should the specified panels be included or excluded
    exclude(exclude: boolean): this {
        this.internal.exclude = exclude;
        return this;
    }

    // Panel IDs that should be included or excluded
    ids(ids: number[]): this {
        this.internal.ids = ids;
        return this;
    }
}

