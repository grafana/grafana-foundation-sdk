// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.team;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Team {
    // Name of the team.
    @JsonProperty("name")
    public String name;
    // Email of the team.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("email")
    public String email;
    public Team() {
    }
    public Team(String name,String email) {
        this.name = name;
        this.email = email;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
