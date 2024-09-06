// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BuiltinRoleRef {
    @JsonProperty("kind")
    public String kind;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("name")
    public BuiltinRoleRefName name;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BuiltinRoleRef> {
        private final BuiltinRoleRef internal;
        
        public Builder() {
            this.internal = new BuiltinRoleRef();
    this.internal.kind = "BuiltinRole";
        }
    public Builder name(BuiltinRoleRefName name) {
    this.internal.name = name;
        return this;
    }
    public BuiltinRoleRef build() {
            return this.internal;
        }
    }
}
