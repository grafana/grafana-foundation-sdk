// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class AnnotationPanelFilterBuilder implements cog.Builder<dashboard.AnnotationPanelFilter> {
    protected readonly internal: dashboard.AnnotationPanelFilter;

    constructor() {
        this.internal = dashboard.defaultAnnotationPanelFilter();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationPanelFilter {
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
