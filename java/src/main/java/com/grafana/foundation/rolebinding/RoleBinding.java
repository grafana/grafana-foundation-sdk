// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RoleBinding {
    // The role we are discussing
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("role")
    public BuiltinRoleRefOrCustomRoleRef role;
    // The team or user that has the specified role
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("subject")
    public RoleBindingSubject subject;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RoleBinding> {
        private final RoleBinding internal;
        
        public Builder() {
            this.internal = new RoleBinding();
        }
    public Builder role(BuiltinRoleRefOrCustomRoleRef role) {
    this.internal.role = role;
        return this;
    }
    
    public Builder subject(com.grafana.foundation.cog.Builder<RoleBindingSubject> subject) {
    this.internal.subject = subject.build();
        return this;
    }
    public RoleBinding build() {
            return this.internal;
        }
    }
}
