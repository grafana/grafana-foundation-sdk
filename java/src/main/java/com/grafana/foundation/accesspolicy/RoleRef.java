// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RoleRef {
    // Policies can apply to roles, teams, or users
    // Applying policies to individual users is supported, but discouraged
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("kind")
    public RoleRefKind kind;
    @JsonProperty("name")
    public String name;
    @JsonProperty("xname")
    public String xname;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RoleRef> {
        private final RoleRef internal;
        
        public Builder() {
            this.internal = new RoleRef();
        }
    public Builder kind(RoleRefKind kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder xname(String xname) {
    this.internal.xname = xname;
        return this;
    }
    public RoleRef build() {
            return this.internal;
        }
    }
}
