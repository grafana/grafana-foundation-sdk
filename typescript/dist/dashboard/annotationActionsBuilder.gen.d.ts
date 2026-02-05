import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class AnnotationActionsBuilder implements cog.Builder<dashboard.AnnotationActions> {
    protected readonly internal: dashboard.AnnotationActions;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationActions;
    canAdd(canAdd: boolean): this;
    canDelete(canDelete: boolean): this;
    canEdit(canEdit: boolean): this;
}
