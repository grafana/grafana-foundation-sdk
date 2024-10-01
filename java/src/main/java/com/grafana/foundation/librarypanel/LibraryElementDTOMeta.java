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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LibraryElementDTOMeta> {
        protected final LibraryElementDTOMeta internal;
        
        public Builder() {
            this.internal = new LibraryElementDTOMeta();
        }
    public Builder folderName(String folderName) {
    this.internal.folderName = folderName;
        return this;
    }
    
    public Builder folderUid(String folderUid) {
    this.internal.folderUid = folderUid;
        return this;
    }
    
    public Builder connectedDashboards(Long connectedDashboards) {
    this.internal.connectedDashboards = connectedDashboards;
        return this;
    }
    
    public Builder created(String created) {
    this.internal.created = created;
        return this;
    }
    
    public Builder updated(String updated) {
    this.internal.updated = updated;
        return this;
    }
    
    public Builder createdBy(com.grafana.foundation.cog.Builder<LibraryElementDTOMetaUser> createdBy) {
    this.internal.createdBy = createdBy.build();
        return this;
    }
    
    public Builder updatedBy(com.grafana.foundation.cog.Builder<LibraryElementDTOMetaUser> updatedBy) {
    this.internal.updatedBy = updatedBy.build();
        return this;
    }
    public LibraryElementDTOMeta build() {
            return this.internal;
        }
    }
}
