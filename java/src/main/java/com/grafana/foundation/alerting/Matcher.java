// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class Matcher {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Type")
    public MatchType type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("Value")
    public String value;
    public Matcher() {
    }
    public Matcher(String name,MatchType type,String value) {
        this.name = name;
        this.type = type;
        this.value = value;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
