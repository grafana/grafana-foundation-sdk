// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;

// #MovingAverage's settings are overridden in types.ts
public class MovingAverage {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pipelineAgg")
    public String pipelineAgg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonProperty("type")
    public MetricAggregationType type;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public Map<String, Object> settings;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    public MovingAverage() {
        this.type = MetricAggregationType.MOVING_AVG;
    }
    public MovingAverage(String pipelineAgg,String field,String id,Map<String, Object> settings,Boolean hide) {
        this.pipelineAgg = pipelineAgg;
        this.field = field;
        this.type = MetricAggregationType.MOVING_AVG;
        this.id = id;
        this.settings = settings;
        this.hide = hide;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
