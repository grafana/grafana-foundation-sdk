// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CustomRoleRef {
    @JsonProperty("kind")
    public String kind;
    @JsonProperty("name")
    public String name;
    public CustomRoleRef() {
        this.kind = "";
        this.name = "";
    }
    public CustomRoleRef(String kind,String name) {
        this.kind = kind;
        this.name = name;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
