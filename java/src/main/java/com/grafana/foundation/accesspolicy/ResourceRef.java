// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ResourceRef {
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("name")
    public String name;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ResourceRef> {
        protected final ResourceRef internal;
        
        public Builder() {
            this.internal = new ResourceRef();
        }
    public Builder kind(String kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    public ResourceRef build() {
            return this.internal;
        }
    }
}
