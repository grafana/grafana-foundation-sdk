// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.role;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Role {
    // The role identifier `managed:builtins:editor:permissions`
    @JsonProperty("name")
    public String name;
    // Optional display
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("displayName")
    public String displayName;
    // Name of the team.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("groupName")
    public String groupName;
    // Role description
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("description")
    public String description;
    // Do not show this role
    @JsonProperty("hidden")
    public Boolean hidden;
    public Role() {
        this.name = "";
        this.hidden = false;
    }
    public Role(String name,String displayName,String groupName,String description,Boolean hidden) {
        this.name = name;
        this.displayName = displayName;
        this.groupName = groupName;
        this.description = description;
        this.hidden = hidden;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
