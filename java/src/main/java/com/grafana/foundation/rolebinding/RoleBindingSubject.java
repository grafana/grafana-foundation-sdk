// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RoleBindingSubject {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("kind")
    public RoleBindingSubjectKind kind;
    // The team/user identifier name
    @JsonProperty("name")
    public String name;
    public RoleBindingSubject() {
    }
    public RoleBindingSubject(RoleBindingSubjectKind kind,String name) {
        this.kind = kind;
        this.name = name;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
