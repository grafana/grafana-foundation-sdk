// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class LibraryPanel {
    // Folder UID
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("folderUid")
    public String folderUid;
    // Library element UID
    @JsonProperty("uid")
    public String uid;
    // Panel name (also saved in the model)
    @JsonProperty("name")
    public String name;
    // Panel description
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // The panel type (from inside the model)
    @JsonProperty("type")
    public String type;
    // Dashboard version when this was saved (zero if unknown)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("schemaVersion")
    public Short schemaVersion;
    // panel version, incremented each time the dashboard is updated.
    @JsonProperty("version")
    public Long version;
    // TODO: should be the same panel schema defined in dashboard
    // Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("model")
    public PanelModel model;
    // Object storage metadata
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("meta")
    public LibraryElementDTOMeta meta;
    public LibraryPanel() {
    }
    public LibraryPanel(String folderUid,String uid,String name,String description,String type,Short schemaVersion,Long version,PanelModel model,LibraryElementDTOMeta meta) {
        this.folderUid = folderUid;
        this.uid = uid;
        this.name = name;
        this.description = description;
        this.type = type;
        this.schemaVersion = schemaVersion;
        this.version = version;
        this.model = model;
        this.meta = meta;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
