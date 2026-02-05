"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.DashboardMetaBuilder = void 0;
const tslib_1 = require("tslib");
const dashboard = tslib_1.__importStar(require("../dashboard"));
class DashboardMetaBuilder {
    constructor() {
        this.internal = dashboard.defaultDashboardMeta();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    annotationsPermissions(annotationsPermissions) {
        const annotationsPermissionsResource = annotationsPermissions.build();
        this.internal.annotationsPermissions = annotationsPermissionsResource;
        return this;
    }
    apiVersion(apiVersion) {
        this.internal.apiVersion = apiVersion;
        return this;
    }
    canAdmin(canAdmin) {
        this.internal.canAdmin = canAdmin;
        return this;
    }
    canDelete(canDelete) {
        this.internal.canDelete = canDelete;
        return this;
    }
    canEdit(canEdit) {
        this.internal.canEdit = canEdit;
        return this;
    }
    canSave(canSave) {
        this.internal.canSave = canSave;
        return this;
    }
    canStar(canStar) {
        this.internal.canStar = canStar;
        return this;
    }
    created(created) {
        this.internal.created = created;
        return this;
    }
    createdBy(createdBy) {
        this.internal.createdBy = createdBy;
        return this;
    }
    expires(expires) {
        this.internal.expires = expires;
        return this;
    }
    // Deprecated: use FolderUID instead
    folderId(folderId) {
        this.internal.folderId = folderId;
        return this;
    }
    folderTitle(folderTitle) {
        this.internal.folderTitle = folderTitle;
        return this;
    }
    folderUid(folderUid) {
        this.internal.folderUid = folderUid;
        return this;
    }
    folderUrl(folderUrl) {
        this.internal.folderUrl = folderUrl;
        return this;
    }
    hasAcl(hasAcl) {
        this.internal.hasAcl = hasAcl;
        return this;
    }
    isFolder(isFolder) {
        this.internal.isFolder = isFolder;
        return this;
    }
    isSnapshot(isSnapshot) {
        this.internal.isSnapshot = isSnapshot;
        return this;
    }
    isStarred(isStarred) {
        this.internal.isStarred = isStarred;
        return this;
    }
    provisioned(provisioned) {
        this.internal.provisioned = provisioned;
        return this;
    }
    provisionedExternalId(provisionedExternalId) {
        this.internal.provisionedExternalId = provisionedExternalId;
        return this;
    }
    publicDashboardEnabled(publicDashboardEnabled) {
        this.internal.publicDashboardEnabled = publicDashboardEnabled;
        return this;
    }
    slug(slug) {
        this.internal.slug = slug;
        return this;
    }
    type(type) {
        this.internal.type = type;
        return this;
    }
    updated(updated) {
        this.internal.updated = updated;
        return this;
    }
    updatedBy(updatedBy) {
        this.internal.updatedBy = updatedBy;
        return this;
    }
    url(url) {
        this.internal.url = url;
        return this;
    }
    version(version) {
        this.internal.version = version;
        return this;
    }
}
exports.DashboardMetaBuilder = DashboardMetaBuilder;
//# sourceMappingURL=dashboardMetaBuilder.gen.js.map