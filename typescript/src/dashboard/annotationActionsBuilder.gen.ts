// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class AnnotationActionsBuilder implements cog.Builder<dashboard.AnnotationActions> {
    protected readonly internal: dashboard.AnnotationActions;

    constructor() {
        this.internal = dashboard.defaultAnnotationActions();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationActions {
        return this.internal;
    }

    canAdd(canAdd: boolean): this {
        this.internal.canAdd = canAdd;
        return this;
    }

    canDelete(canDelete: boolean): this {
        this.internal.canDelete = canDelete;
        return this;
    }

    canEdit(canEdit: boolean): this {
        this.internal.canEdit = canEdit;
        return this;
    }
}
