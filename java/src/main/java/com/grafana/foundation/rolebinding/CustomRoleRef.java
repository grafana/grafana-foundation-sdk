// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CustomRoleRef {
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("name")
    public String name;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CustomRoleRef> {
        private final CustomRoleRef internal;
        
        public Builder() {
            this.internal = new CustomRoleRef();
    this.internal.kind = "Role";
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    public CustomRoleRef build() {
            return this.internal;
        }
    }
}
