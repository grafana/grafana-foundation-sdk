// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.debug;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class UpdateConfig {
    @JsonProperty("render")
    public Boolean render;
    @JsonProperty("dataChanged")
    public Boolean dataChanged;
    @JsonProperty("schemaChanged")
    public Boolean schemaChanged;
    public UpdateConfig() {
    }
    public UpdateConfig(Boolean render,Boolean dataChanged,Boolean schemaChanged) {
        this.render = render;
        this.dataChanged = dataChanged;
        this.schemaChanged = schemaChanged;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
