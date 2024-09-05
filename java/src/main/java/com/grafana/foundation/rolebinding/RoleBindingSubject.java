// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class RoleBindingSubject {
    @JsonProperty("kind")
    public RoleBindingSubjectKind kind;
    // The team/user identifier name
    @JsonProperty("name")
    public String name;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RoleBindingSubject> {
        private final RoleBindingSubject internal;
        
        public Builder() {
            this.internal = new RoleBindingSubject();
        }
    public Builder kind(RoleBindingSubjectKind kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    public RoleBindingSubject build() {
            return this.internal;
        }
    }
}
