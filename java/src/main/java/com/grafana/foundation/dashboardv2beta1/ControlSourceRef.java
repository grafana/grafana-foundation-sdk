// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ControlSourceRef {
    @JsonProperty("type")
    public String type;
    // The plugin type-id
    @JsonProperty("group")
    public String group;
    public ControlSourceRef() {
        this.type = "";
        this.group = "";
    }
    public ControlSourceRef(String type,String group) {
        this.type = type;
        this.group = group;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
