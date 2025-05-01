// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DashboardMetaBuilder implements com.grafana.foundation.cog.Builder<DashboardMeta> {
    protected final DashboardMeta internal;
    
    public DashboardMetaBuilder() {
        this.internal = new DashboardMeta();
    }
    public DashboardMetaBuilder annotationsPermissions(com.grafana.foundation.cog.Builder<AnnotationPermission> annotationsPermissions) {
    AnnotationPermission annotationsPermissionsResource = annotationsPermissions.build();
        this.internal.annotationsPermissions = annotationsPermissionsResource;
        return this;
    }
    
    public DashboardMetaBuilder apiVersion(String apiVersion) {
        this.internal.apiVersion = apiVersion;
        return this;
    }
    
    public DashboardMetaBuilder canAdmin(Boolean canAdmin) {
        this.internal.canAdmin = canAdmin;
        return this;
    }
    
    public DashboardMetaBuilder canDelete(Boolean canDelete) {
        this.internal.canDelete = canDelete;
        return this;
    }
    
    public DashboardMetaBuilder canEdit(Boolean canEdit) {
        this.internal.canEdit = canEdit;
        return this;
    }
    
    public DashboardMetaBuilder canSave(Boolean canSave) {
        this.internal.canSave = canSave;
        return this;
    }
    
    public DashboardMetaBuilder canStar(Boolean canStar) {
        this.internal.canStar = canStar;
        return this;
    }
    
    public DashboardMetaBuilder created(String created) {
        this.internal.created = created;
        return this;
    }
    
    public DashboardMetaBuilder createdBy(String createdBy) {
        this.internal.createdBy = createdBy;
        return this;
    }
    
    public DashboardMetaBuilder expires(String expires) {
        this.internal.expires = expires;
        return this;
    }
    
    public DashboardMetaBuilder folderId(Long folderId) {
        this.internal.folderId = folderId;
        return this;
    }
    
    public DashboardMetaBuilder folderTitle(String folderTitle) {
        this.internal.folderTitle = folderTitle;
        return this;
    }
    
    public DashboardMetaBuilder folderUid(String folderUid) {
        this.internal.folderUid = folderUid;
        return this;
    }
    
    public DashboardMetaBuilder folderUrl(String folderUrl) {
        this.internal.folderUrl = folderUrl;
        return this;
    }
    
    public DashboardMetaBuilder hasAcl(Boolean hasAcl) {
        this.internal.hasAcl = hasAcl;
        return this;
    }
    
    public DashboardMetaBuilder isFolder(Boolean isFolder) {
        this.internal.isFolder = isFolder;
        return this;
    }
    
    public DashboardMetaBuilder isSnapshot(Boolean isSnapshot) {
        this.internal.isSnapshot = isSnapshot;
        return this;
    }
    
    public DashboardMetaBuilder isStarred(Boolean isStarred) {
        this.internal.isStarred = isStarred;
        return this;
    }
    
    public DashboardMetaBuilder provisioned(Boolean provisioned) {
        this.internal.provisioned = provisioned;
        return this;
    }
    
    public DashboardMetaBuilder provisionedExternalId(String provisionedExternalId) {
        this.internal.provisionedExternalId = provisionedExternalId;
        return this;
    }
    
    public DashboardMetaBuilder publicDashboardEnabled(Boolean publicDashboardEnabled) {
        this.internal.publicDashboardEnabled = publicDashboardEnabled;
        return this;
    }
    
    public DashboardMetaBuilder slug(String slug) {
        this.internal.slug = slug;
        return this;
    }
    
    public DashboardMetaBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public DashboardMetaBuilder updated(String updated) {
        this.internal.updated = updated;
        return this;
    }
    
    public DashboardMetaBuilder updatedBy(String updatedBy) {
        this.internal.updatedBy = updatedBy;
        return this;
    }
    
    public DashboardMetaBuilder url(String url) {
        this.internal.url = url;
        return this;
    }
    
    public DashboardMetaBuilder version(Long version) {
        this.internal.version = version;
        return this;
    }
    public DashboardMeta build() {
        return this.internal;
    }
}
