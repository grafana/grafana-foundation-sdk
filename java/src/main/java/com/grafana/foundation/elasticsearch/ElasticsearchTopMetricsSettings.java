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
    public ElasticsearchTopMetricsSettings() {
    }
    public ElasticsearchTopMetricsSettings(String order,String orderBy,List<String> metrics) {
        this.order = order;
        this.orderBy = orderBy;
        this.metrics = metrics;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
