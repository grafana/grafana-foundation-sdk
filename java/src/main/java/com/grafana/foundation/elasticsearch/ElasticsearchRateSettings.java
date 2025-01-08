// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchRateSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("unit")
    public String unit;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public String mode;
    public ElasticsearchRateSettings() {
    }
    
    public ElasticsearchRateSettings(String unit,String mode) {
        this.unit = unit;
        this.mode = mode;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
