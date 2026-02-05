import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class DashboardMetaBuilder implements cog.Builder<dashboard.DashboardMeta> {
    protected readonly internal: dashboard.DashboardMeta;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.DashboardMeta;
    annotationsPermissions(annotationsPermissions: cog.Builder<dashboard.AnnotationPermission>): this;
    apiVersion(apiVersion: string): this;
    canAdmin(canAdmin: boolean): this;
    canDelete(canDelete: boolean): this;
    canEdit(canEdit: boolean): this;
    canSave(canSave: boolean): this;
    canStar(canStar: boolean): this;
    created(created: string): this;
    createdBy(createdBy: string): this;
    expires(expires: string): this;
    folderId(folderId: number): this;
    folderTitle(folderTitle: string): this;
    folderUid(folderUid: string): this;
    folderUrl(folderUrl: string): this;
    hasAcl(hasAcl: boolean): this;
    isFolder(isFolder: boolean): this;
    isSnapshot(isSnapshot: boolean): this;
    isStarred(isStarred: boolean): this;
    provisioned(provisioned: boolean): this;
    provisionedExternalId(provisionedExternalId: string): this;
    publicDashboardEnabled(publicDashboardEnabled: boolean): this;
    slug(slug: string): this;
    type(type: string): this;
    updated(updated: string): this;
    updatedBy(updatedBy: string): this;
    url(url: string): this;
    version(version: number): this;
}
