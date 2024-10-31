// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class DashboardMetaBuilder implements cog.Builder<dashboard.DashboardMeta> {
    protected readonly internal: dashboard.DashboardMeta;

    constructor() {
        this.internal = dashboard.defaultDashboardMeta();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.DashboardMeta {
        return this.internal;
    }

    annotationsPermissions(annotationsPermissions: cog.Builder<dashboard.AnnotationPermission>): this {
        const annotationsPermissionsResource = annotationsPermissions.build();
        this.internal.annotationsPermissions = annotationsPermissionsResource;
        return this;
    }

    canAdmin(canAdmin: boolean): this {
        this.internal.canAdmin = canAdmin;
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

    canSave(canSave: boolean): this {
        this.internal.canSave = canSave;
        return this;
    }

    canStar(canStar: boolean): this {
        this.internal.canStar = canStar;
        return this;
    }

    created(created: string): this {
        this.internal.created = created;
        return this;
    }

    createdBy(createdBy: string): this {
        this.internal.createdBy = createdBy;
        return this;
    }

    expires(expires: string): this {
        this.internal.expires = expires;
        return this;
    }

    folderId(folderId: number): this {
        this.internal.folderId = folderId;
        return this;
    }

    folderTitle(folderTitle: string): this {
        this.internal.folderTitle = folderTitle;
        return this;
    }

    folderUid(folderUid: string): this {
        this.internal.folderUid = folderUid;
        return this;
    }

    folderUrl(folderUrl: string): this {
        this.internal.folderUrl = folderUrl;
        return this;
    }

    hasAcl(hasAcl: boolean): this {
        this.internal.hasAcl = hasAcl;
        return this;
    }

    isFolder(isFolder: boolean): this {
        this.internal.isFolder = isFolder;
        return this;
    }

    isSnapshot(isSnapshot: boolean): this {
        this.internal.isSnapshot = isSnapshot;
        return this;
    }

    isStarred(isStarred: boolean): this {
        this.internal.isStarred = isStarred;
        return this;
    }

    provisioned(provisioned: boolean): this {
        this.internal.provisioned = provisioned;
        return this;
    }

    provisionedExternalId(provisionedExternalId: string): this {
        this.internal.provisionedExternalId = provisionedExternalId;
        return this;
    }

    publicDashboardAccessToken(publicDashboardAccessToken: string): this {
        this.internal.publicDashboardAccessToken = publicDashboardAccessToken;
        return this;
    }

    publicDashboardEnabled(publicDashboardEnabled: boolean): this {
        this.internal.publicDashboardEnabled = publicDashboardEnabled;
        return this;
    }

    publicDashboardUid(publicDashboardUid: string): this {
        this.internal.publicDashboardUid = publicDashboardUid;
        return this;
    }

    slug(slug: string): this {
        this.internal.slug = slug;
        return this;
    }

    type(type: string): this {
        this.internal.type = type;
        return this;
    }

    updated(updated: string): this {
        this.internal.updated = updated;
        return this;
    }

    updatedBy(updatedBy: string): this {
        this.internal.updatedBy = updatedBy;
        return this;
    }

    url(url: string): this {
        this.internal.url = url;
        return this;
    }

    version(version: number): this {
        this.internal.version = version;
        return this;
    }
}
