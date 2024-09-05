// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class MetricAggregationType {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("String")
    public String string;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("PipelineMetricAggregationType")
    public PipelineMetricAggregationType pipelineMetricAggregationType;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MetricAggregationType> {
        private final MetricAggregationType internal;
        
        public Builder() {
            this.internal = new MetricAggregationType();
    this.internal.string = "count";
        }
    public Builder pipelineMetricAggregationType(PipelineMetricAggregationType pipelineMetricAggregationType) {
    this.internal.pipelineMetricAggregationType = pipelineMetricAggregationType;
        return this;
    }
    public MetricAggregationType build() {
            return this.internal;
        }
    }
}
