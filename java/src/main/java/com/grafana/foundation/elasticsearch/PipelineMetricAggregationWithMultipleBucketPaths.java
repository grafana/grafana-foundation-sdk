// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class PipelineMetricAggregationWithMultipleBucketPaths {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pipelineVariables")
    public List<PipelineVariable> pipelineVariables;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public MetricAggregationType type;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PipelineMetricAggregationWithMultipleBucketPaths> {
        protected final PipelineMetricAggregationWithMultipleBucketPaths internal;
        
        public Builder() {
            this.internal = new PipelineMetricAggregationWithMultipleBucketPaths();
        }
    public Builder pipelineVariables(com.grafana.foundation.cog.Builder<List<PipelineVariable>> pipelineVariables) {
    this.internal.pipelineVariables = pipelineVariables.build();
        return this;
    }
    
    public Builder type(com.grafana.foundation.cog.Builder<MetricAggregationType> type) {
    this.internal.type = type.build();
        return this;
    }
    
    public Builder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public PipelineMetricAggregationWithMultipleBucketPaths build() {
            return this.internal;
        }
    }
}