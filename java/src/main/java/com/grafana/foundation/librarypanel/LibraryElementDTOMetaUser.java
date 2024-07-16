// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class LibraryElementDTOMetaUser { 
    @JsonProperty("id")
    public Long id; 
    @JsonProperty("name")
    public String name; 
    @JsonProperty("avatarUrl")
    public String avatarUrl;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<LibraryElementDTOMetaUser> {
        private final LibraryElementDTOMetaUser internal;
        
        public Builder() {
            this.internal = new LibraryElementDTOMetaUser();
        }
    public Builder id(Long id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder avatarUrl(String avatarUrl) {
    this.internal.avatarUrl = avatarUrl;
        return this;
    }
    public LibraryElementDTOMetaUser build() {
            return this.internal;
        }
    }
}
