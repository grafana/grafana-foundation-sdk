// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class ElasticsearchTopMetricsSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("order")
    public String order;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orderBy")
    public String orderBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("metrics")
    public List<String> metrics;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchTopMetricsSettings> {
        protected final ElasticsearchTopMetricsSettings internal;
        
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
