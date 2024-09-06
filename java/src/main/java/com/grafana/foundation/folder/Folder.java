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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Folder> {
        private final Folder internal;
        
        public Builder() {
            this.internal = new Folder();
        }
    public Builder uid(String uid) {
    this.internal.uid = uid;
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    
    public Builder description(String description) {
    this.internal.description = description;
        return this;
    }
    public Folder build() {
            return this.internal;
        }
    }
}
