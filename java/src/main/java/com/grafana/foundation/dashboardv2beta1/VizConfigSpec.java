// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// --- Kinds ---
public class VizConfigSpec {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("options")
    public Object options;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("fieldConfig")
    public FieldConfigSource fieldConfig;
    public VizConfigSpec() {
        this.fieldConfig = new com.grafana.foundation.dashboardv2beta1.FieldConfigSource();
    }
    public VizConfigSpec(Object options,FieldConfigSource fieldConfig) {
        this.options = options;
        this.fieldConfig = fieldConfig;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
