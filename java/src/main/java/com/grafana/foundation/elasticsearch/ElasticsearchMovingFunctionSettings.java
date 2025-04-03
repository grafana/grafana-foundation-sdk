// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchMovingFunctionSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("window")
    public String window;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("script")
    public InlineScript script;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("shift")
    public String shift;
    public ElasticsearchMovingFunctionSettings() {
    }
    public ElasticsearchMovingFunctionSettings(String window,InlineScript script,String shift) {
        this.window = window;
        this.script = script;
        this.shift = shift;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
