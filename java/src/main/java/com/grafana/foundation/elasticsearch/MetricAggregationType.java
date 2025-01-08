// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MetricAggregationType {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("String")
    public String string;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("PipelineMetricAggregationType")
    public PipelineMetricAggregationType pipelineMetricAggregationType;
    public MetricAggregationType() {
    }
    
    public MetricAggregationType(String string,PipelineMetricAggregationType pipelineMetricAggregationType) {
        this.string = string;
        this.pipelineMetricAggregationType = pipelineMetricAggregationType;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
