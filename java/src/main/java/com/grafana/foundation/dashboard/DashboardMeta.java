// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class DashboardMeta { 
    @JsonProperty("annotationsPermissions")
    public AnnotationPermission annotationsPermissions; 
    @JsonProperty("canAdmin")
    public Boolean canAdmin; 
    @JsonProperty("canDelete")
    public Boolean canDelete; 
    @JsonProperty("canEdit")
    public Boolean canEdit; 
    @JsonProperty("canSave")
    public Boolean canSave; 
    @JsonProperty("canStar")
    public Boolean canStar; 
    @JsonProperty("created")
    public String created; 
    @JsonProperty("createdBy")
    public String createdBy; 
    @JsonProperty("expires")
    public String expires; 
    @JsonProperty("folderId")
    public Long folderId; 
    @JsonProperty("folderTitle")
    public String folderTitle; 
    @JsonProperty("folderUid")
    public String folderUid; 
    @JsonProperty("folderUrl")
    public String folderUrl; 
    @JsonProperty("hasAcl")
    public Boolean hasAcl; 
    @JsonProperty("isFolder")
    public Boolean isFolder; 
    @JsonProperty("isSnapshot")
    public Boolean isSnapshot; 
    @JsonProperty("isStarred")
    public Boolean isStarred; 
    @JsonProperty("provisioned")
    public Boolean provisioned; 
    @JsonProperty("provisionedExternalId")
    public String provisionedExternalId; 
    @JsonProperty("publicDashboardAccessToken")
    public String publicDashboardAccessToken; 
    @JsonProperty("publicDashboardEnabled")
    public Boolean publicDashboardEnabled; 
    @JsonProperty("publicDashboardUid")
    public String publicDashboardUid; 
    @JsonProperty("slug")
    public String slug; 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("updated")
    public String updated; 
    @JsonProperty("updatedBy")
    public String updatedBy; 
    @JsonProperty("url")
    public String url; 
    @JsonProperty("version")
    public Long version;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<DashboardMeta> {
        private final DashboardMeta internal;
        
        public Builder() {
            this.internal = new DashboardMeta();
        }
    public Builder annotationsPermissions(com.grafana.foundation.cog.Builder<AnnotationPermission> annotationsPermissions) {
    this.internal.annotationsPermissions = annotationsPermissions.build();
        return this;
    }
    
    public Builder canAdmin(Boolean canAdmin) {
    this.internal.canAdmin = canAdmin;
        return this;
    }
    
    public Builder canDelete(Boolean canDelete) {
    this.internal.canDelete = canDelete;
        return this;
    }
    
    public Builder canEdit(Boolean canEdit) {
    this.internal.canEdit = canEdit;
        return this;
    }
    
    public Builder canSave(Boolean canSave) {
    this.internal.canSave = canSave;
        return this;
    }
    
    public Builder canStar(Boolean canStar) {
    this.internal.canStar = canStar;
        return this;
    }
    
    public Builder created(String created) {
    this.internal.created = created;
        return this;
    }
    
    public Builder createdBy(String createdBy) {
    this.internal.createdBy = createdBy;
        return this;
    }
    
    public Builder expires(String expires) {
    this.internal.expires = expires;
        return this;
    }
    
    public Builder folderId(Long folderId) {
    this.internal.folderId = folderId;
        return this;
    }
    
    public Builder folderTitle(String folderTitle) {
    this.internal.folderTitle = folderTitle;
        return this;
    }
    
    public Builder folderUid(String folderUid) {
    this.internal.folderUid = folderUid;
        return this;
    }
    
    public Builder folderUrl(String folderUrl) {
    this.internal.folderUrl = folderUrl;
        return this;
    }
    
    public Builder hasAcl(Boolean hasAcl) {
    this.internal.hasAcl = hasAcl;
        return this;
    }
    
    public Builder isFolder(Boolean isFolder) {
    this.internal.isFolder = isFolder;
        return this;
    }
    
    public Builder isSnapshot(Boolean isSnapshot) {
    this.internal.isSnapshot = isSnapshot;
        return this;
    }
    
    public Builder isStarred(Boolean isStarred) {
    this.internal.isStarred = isStarred;
        return this;
    }
    
    public Builder provisioned(Boolean provisioned) {
    this.internal.provisioned = provisioned;
        return this;
    }
    
    public Builder provisionedExternalId(String provisionedExternalId) {
    this.internal.provisionedExternalId = provisionedExternalId;
        return this;
    }
    
    public Builder publicDashboardAccessToken(String publicDashboardAccessToken) {
    this.internal.publicDashboardAccessToken = publicDashboardAccessToken;
        return this;
    }
    
    public Builder publicDashboardEnabled(Boolean publicDashboardEnabled) {
    this.internal.publicDashboardEnabled = publicDashboardEnabled;
        return this;
    }
    
    public Builder publicDashboardUid(String publicDashboardUid) {
    this.internal.publicDashboardUid = publicDashboardUid;
        return this;
    }
    
    public Builder slug(String slug) {
    this.internal.slug = slug;
        return this;
    }
    
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    
    public Builder updated(String updated) {
    this.internal.updated = updated;
        return this;
    }
    
    public Builder updatedBy(String updatedBy) {
    this.internal.updatedBy = updatedBy;
        return this;
    }
    
    public Builder url(String url) {
    this.internal.url = url;
        return this;
    }
    
    public Builder version(Long version) {
    this.internal.version = version;
        return this;
    }
    public DashboardMeta build() {
            return this.internal;
        }
    }
}
