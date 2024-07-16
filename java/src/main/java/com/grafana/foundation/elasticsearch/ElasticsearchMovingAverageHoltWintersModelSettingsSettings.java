// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ElasticsearchMovingAverageHoltWintersModelSettingsSettings { 
    @JsonProperty("alpha")
    public String alpha; 
    @JsonProperty("beta")
    public String beta; 
    @JsonProperty("gamma")
    public String gamma; 
    @JsonProperty("period")
    public String period; 
    @JsonProperty("pad")
    public Boolean pad;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltWintersModelSettingsSettings> {
        private final ElasticsearchMovingAverageHoltWintersModelSettingsSettings internal;
        
        public Builder() {
            this.internal = new ElasticsearchMovingAverageHoltWintersModelSettingsSettings();
        }
    public Builder alpha(String alpha) {
    this.internal.alpha = alpha;
        return this;
    }
    
    public Builder beta(String beta) {
    this.internal.beta = beta;
        return this;
    }
    
    public Builder gamma(String gamma) {
    this.internal.gamma = gamma;
        return this;
    }
    
    public Builder period(String period) {
    this.internal.period = period;
        return this;
    }
    
    public Builder pad(Boolean pad) {
    this.internal.pad = pad;
        return this;
    }
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettings build() {
            return this.internal;
        }
    }
}