// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MovingAverageHoltModelSettings {
    @JsonProperty("model")
    public String model;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("settings")
    public ElasticsearchMovingAverageHoltModelSettingsSettings settings;
    @JsonProperty("window")
    public String window;
    @JsonProperty("minimize")
    public Boolean minimize;
    @JsonProperty("predict")
    public String predict;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MovingAverageHoltModelSettings> {
        private final MovingAverageHoltModelSettings internal;
        
        public Builder() {
            this.internal = new MovingAverageHoltModelSettings();
    this.internal.model = "holt";
        }
    public Builder settings(com.grafana.foundation.cog.Builder<ElasticsearchMovingAverageHoltModelSettingsSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public Builder window(String window) {
    this.internal.window = window;
        return this;
    }
    
    public Builder minimize(Boolean minimize) {
    this.internal.minimize = minimize;
        return this;
    }
    
    public Builder predict(String predict) {
    this.internal.predict = predict;
        return this;
    }
    public MovingAverageHoltModelSettings build() {
            return this.internal;
        }
    }
}
