// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class TopMetrics {
    @JsonProperty("type")
    public String type;
    @JsonProperty("id")
    public String id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchTopMetricsSettings settings;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hide")
    public Boolean hide;
    public TopMetrics() {
    }
    
    public TopMetrics(String type,String id,ElasticsearchTopMetricsSettings settings,Boolean hide) {
        this.type = type;
        this.id = id;
        this.settings = settings;
        this.hide = hide;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
