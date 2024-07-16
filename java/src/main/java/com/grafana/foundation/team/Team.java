// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.team;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Team {
    // Name of the team. 
    @JsonProperty("name")
    public String name;
    // Email of the team. 
    @JsonProperty("email")
    public String email;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Team> {
        private final Team internal;
        
        public Builder(String name) {
            this.internal = new Team();
    this.internal.name = name;
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder email(String email) {
    this.internal.email = email;
        return this;
    }
    public Team build() {
            return this.internal;
        }
    }
}
