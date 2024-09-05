// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.role;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class Role {
    // The role identifier `managed:builtins:editor:permissions`
    @JsonProperty("name")
    public String name;
    // Optional display
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("displayName")
    public String displayName;
    // Name of the team.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("groupName")
    public String groupName;
    // Role description
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("description")
    public String description;
    // Do not show this role
    @JsonProperty("hidden")
    public Boolean hidden;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Role> {
        private final Role internal;
        
        public Builder() {
            this.internal = new Role();
        this.hidden(false);
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder displayName(String displayName) {
    this.internal.displayName = displayName;
        return this;
    }
    
    public Builder groupName(String groupName) {
    this.internal.groupName = groupName;
        return this;
    }
    
    public Builder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public Builder hidden(Boolean hidden) {
    this.internal.hidden = hidden;
        return this;
    }
    public Role build() {
            return this.internal;
        }
    }
}
