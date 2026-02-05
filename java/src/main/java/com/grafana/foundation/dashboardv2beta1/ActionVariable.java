// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ActionVariable {
    @JsonProperty("key")
    public String key;
    @JsonProperty("name")
    public String name;
    @JsonProperty("type")
    public String type;
    public ActionVariable() {
        this.key = "";
        this.name = "";
        this.type = Constants.ActionVariableType;
    }
    public ActionVariable(String key,String name) {
        this.key = key;
        this.name = name;
        this.type = Constants.ActionVariableType;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
