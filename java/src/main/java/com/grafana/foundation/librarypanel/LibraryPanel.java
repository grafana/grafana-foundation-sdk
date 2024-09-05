// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class LibraryPanel {
    // Folder UID
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("folderUid")
    public String folderUid;
    // Library element UID
    @JsonProperty("uid")
    public String uid;
    // Panel name (also saved in the model)
    @JsonProperty("name")
    public String name;
    // Panel description
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("description")
    public String description;
    // The panel type (from inside the model)
    @JsonProperty("type")
    public String type;
    // Dashboard version when this was saved (zero if unknown)
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("schemaVersion")
    public Short schemaVersion;
    // panel version, incremented each time the dashboard is updated.
    @JsonProperty("version")
    public Long version;
    // TODO: should be the same panel schema defined in dashboard
    // Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
    @JsonProperty("model")
    public LibrarypanelLibraryPanelModel model;
    // Object storage metadata
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("meta")
    public LibraryElementDTOMeta meta;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LibraryPanel> {
        private final LibraryPanel internal;
        
        public Builder() {
            this.internal = new LibraryPanel();
        }
    public Builder folderUid(String folderUid) {
    this.internal.folderUid = folderUid;
        return this;
    }
    
    public Builder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    
    public Builder name(String name) {
        if (!(name.length() >= 1)) {
            throw new IllegalArgumentException("name.length() must be >= 1");
        }
    this.internal.name = name;
        return this;
    }
    
    public Builder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public Builder type(String type) {
        if (!(type.length() >= 1)) {
            throw new IllegalArgumentException("type.length() must be >= 1");
        }
    this.internal.type = type;
        return this;
    }
    
    public Builder schemaVersion(Short schemaVersion) {
    this.internal.schemaVersion = schemaVersion;
        return this;
    }
    
    public Builder version(Long version) {
    this.internal.version = version;
        return this;
    }
    
    public Builder model(com.grafana.foundation.cog.Builder<LibrarypanelLibraryPanelModel> model) {
    this.internal.model = model.build();
        return this;
    }
    
    public Builder meta(com.grafana.foundation.cog.Builder<LibraryElementDTOMeta> meta) {
    this.internal.meta = meta.build();
        return this;
    }
    public LibraryPanel build() {
            return this.internal;
        }
    }
}
