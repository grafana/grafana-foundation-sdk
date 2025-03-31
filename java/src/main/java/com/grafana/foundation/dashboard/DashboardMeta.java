// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class DashboardMeta {
    // +k8s:deepcopy-gen=true
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotationsPermissions")
    public AnnotationPermission annotationsPermissions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("apiVersion")
    public String apiVersion;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canAdmin")
    public Boolean canAdmin;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canDelete")
    public Boolean canDelete;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canEdit")
    public Boolean canEdit;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canSave")
    public Boolean canSave;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("canStar")
    public Boolean canStar;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("created")
    public String created;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("createdBy")
    public String createdBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("expires")
    public String expires;
    // Deprecated: use FolderUID instead
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderId")
    public Long folderId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderTitle")
    public String folderTitle;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderUid")
    public String folderUid;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderUrl")
    public String folderUrl;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hasAcl")
    public Boolean hasAcl;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isFolder")
    public Boolean isFolder;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isSnapshot")
    public Boolean isSnapshot;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isStarred")
    public Boolean isStarred;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provisioned")
    public Boolean provisioned;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provisionedExternalId")
    public String provisionedExternalId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("publicDashboardEnabled")
    public Boolean publicDashboardEnabled;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("slug")
    public String slug;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("updated")
    public String updated;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("updatedBy")
    public String updatedBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("url")
    public String url;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("version")
    public Long version;
    public DashboardMeta() {
    }
    
    public DashboardMeta(AnnotationPermission annotationsPermissions,String apiVersion,Boolean canAdmin,Boolean canDelete,Boolean canEdit,Boolean canSave,Boolean canStar,String created,String createdBy,String expires,Long folderId,String folderTitle,String folderUid,String folderUrl,Boolean hasAcl,Boolean isFolder,Boolean isSnapshot,Boolean isStarred,Boolean provisioned,String provisionedExternalId,Boolean publicDashboardEnabled,String slug,String type,String updated,String updatedBy,String url,Long version) {
        this.annotationsPermissions = annotationsPermissions;
        this.apiVersion = apiVersion;
        this.canAdmin = canAdmin;
        this.canDelete = canDelete;
        this.canEdit = canEdit;
        this.canSave = canSave;
        this.canStar = canStar;
        this.created = created;
        this.createdBy = createdBy;
        this.expires = expires;
        this.folderId = folderId;
        this.folderTitle = folderTitle;
        this.folderUid = folderUid;
        this.folderUrl = folderUrl;
        this.hasAcl = hasAcl;
        this.isFolder = isFolder;
        this.isSnapshot = isSnapshot;
        this.isStarred = isStarred;
        this.provisioned = provisioned;
        this.provisionedExternalId = provisionedExternalId;
        this.publicDashboardEnabled = publicDashboardEnabled;
        this.slug = slug;
        this.type = type;
        this.updated = updated;
        this.updatedBy = updatedBy;
        this.url = url;
        this.version = version;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
