// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class LibraryElementDTOMeta {
    @JsonProperty("folderName")
    public String folderName;
    @JsonProperty("folderUid")
    public String folderUid;
    @JsonProperty("connectedDashboards")
    public Long connectedDashboards;
    @JsonProperty("created")
    public String created;
    @JsonProperty("updated")
    public String updated;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("createdBy")
    public LibraryElementDTOMetaUser createdBy;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("updatedBy")
    public LibraryElementDTOMetaUser updatedBy;
    public LibraryElementDTOMeta() {
    }
    public LibraryElementDTOMeta(String folderName,String folderUid,Long connectedDashboards,String created,String updated,LibraryElementDTOMetaUser createdBy,LibraryElementDTOMetaUser updatedBy) {
        this.folderName = folderName;
        this.folderUid = folderUid;
        this.connectedDashboards = connectedDashboards;
        this.created = created;
        this.updated = updated;
        this.createdBy = createdBy;
        this.updatedBy = updatedBy;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
