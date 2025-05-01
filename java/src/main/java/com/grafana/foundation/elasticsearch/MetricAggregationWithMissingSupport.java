// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MetricAggregationWithMissingSupport {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchMetricAggregationWithMissingSupportSettings settings;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public MetricAggregationType type;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    public MetricAggregationWithMissingSupport() {
        this.type = MetricAggregationType.COUNT;
        this.id = "";
    }
    public MetricAggregationWithMissingSupport(ElasticsearchMetricAggregationWithMissingSupportSettings settings,MetricAggregationType type,String id,Boolean hide) {
        this.settings = settings;
        this.type = type;
        this.id = id;
        this.hide = hide;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
