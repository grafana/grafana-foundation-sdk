import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class AnnotationPermissionBuilder implements cog.Builder<dashboard.AnnotationPermission> {
    protected readonly internal: dashboard.AnnotationPermission;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationPermission;
    dashboard(dashboard: cog.Builder<dashboard.AnnotationActions>): this;
    organization(organization: cog.Builder<dashboard.AnnotationActions>): this;
}
