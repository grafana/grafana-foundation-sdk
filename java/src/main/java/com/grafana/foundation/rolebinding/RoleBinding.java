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
    public RoleBinding() {
        this.role = new com.grafana.foundation.rolebinding.BuiltinRoleRefOrCustomRoleRef();
        this.subject = new com.grafana.foundation.rolebinding.RoleBindingSubject();
    }
    public RoleBinding(BuiltinRoleRefOrCustomRoleRef role,RoleBindingSubject subject) {
        this.role = role;
        this.subject = subject;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
