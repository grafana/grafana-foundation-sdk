// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ElasticsearchMovingAverageHoltModelSettingsSettings {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("alpha")
    public String alpha;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("beta")
    public String beta;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltModelSettingsSettings> {
        private final ElasticsearchMovingAverageHoltModelSettingsSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchMovingAverageHoltModelSettingsSettings();
        }
    public Builder alpha(String alpha) {
    this.internal.alpha = alpha;
        return this;
    }
    
    public Builder beta(String beta) {
    this.internal.beta = beta;
        return this;
    }
    public ElasticsearchMovingAverageHoltModelSettingsSettings build() {
            return this.internal;
        }
    }
}
