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
    public RoleRef() {
        this.kind = RoleRefKind.ROLE;
        this.name = "";
        this.xname = "";
    }
    public RoleRef(RoleRefKind kind,String name,String xname) {
        this.kind = kind;
        this.name = name;
        this.xname = xname;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
