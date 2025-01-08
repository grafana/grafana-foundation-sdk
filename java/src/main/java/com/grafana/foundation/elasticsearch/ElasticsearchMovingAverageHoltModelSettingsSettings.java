// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ElasticsearchMovingAverageHoltModelSettingsSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alpha")
    public String alpha;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("beta")
    public String beta;
    public ElasticsearchMovingAverageHoltModelSettingsSettings() {
    }
    
    public ElasticsearchMovingAverageHoltModelSettingsSettings(String alpha,String beta) {
        this.alpha = alpha;
        this.beta = beta;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
