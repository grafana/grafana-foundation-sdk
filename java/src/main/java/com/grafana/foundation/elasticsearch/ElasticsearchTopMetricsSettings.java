// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchTopMetricsSettings { 
    @JsonProperty("order")
    public String order; 
    @JsonProperty("orderBy")
    public String orderBy; 
    @JsonProperty("metrics")
    public List<String> metrics;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchTopMetricsSettings> {
        private final ElasticsearchTopMetricsSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchTopMetricsSettings();
        }
    public Builder order(String order) {
    this.internal.order = order;
        return this;
    }
    
    public Builder orderBy(String orderBy) {
    this.internal.orderBy = orderBy;
        return this;
    }
    
    public Builder metrics(List<String> metrics) {
    this.internal.metrics = metrics;
        return this;
    }
    public ElasticsearchTopMetricsSettings build() {
            return this.internal;
        }
    }
}
