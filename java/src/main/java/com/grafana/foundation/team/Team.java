// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.team;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Team {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("email")
    public String email;
    @JsonProperty("name")
    public String name;
    public Team() {
    }
    
    public Team(String email,String name) {
        this.email = email;
        this.name = name;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
