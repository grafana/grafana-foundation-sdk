// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PipelineVariable {
    @JsonProperty("name")
    public String name;
    @JsonProperty("pipelineAgg")
    public String pipelineAgg;
    public PipelineVariable() {
    }
    public PipelineVariable(String name,String pipelineAgg) {
        this.name = name;
        this.pipelineAgg = pipelineAgg;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
