// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.folder;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO:
// common metadata will soon support setting the parent folder in the metadata
public class Folder {
    // Unique folder id. (will be k8s name)
    @JsonProperty("uid")
    public String uid;
    // Folder title
    @JsonProperty("title")
    public String title;
    // Description of the folder.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    public Folder() {
    }
    public Folder(String uid,String title,String description) {
        this.uid = uid;
        this.title = title;
        this.description = description;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
