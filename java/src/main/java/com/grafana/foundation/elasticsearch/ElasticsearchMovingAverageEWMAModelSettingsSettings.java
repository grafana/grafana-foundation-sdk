// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchMovingAverageEWMAModelSettingsSettings { 
    @JsonProperty("alpha")
    public String alpha;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageEWMAModelSettingsSettings> {
        private final ElasticsearchMovingAverageEWMAModelSettingsSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchMovingAverageEWMAModelSettingsSettings();
        }
    public Builder alpha(String alpha) {
    this.internal.alpha = alpha;
        return this;
    }
    public ElasticsearchMovingAverageEWMAModelSettingsSettings build() {
            return this.internal;
        }
    }
}