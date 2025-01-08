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
    public MovingAverageHoltModelSettings() {
    }
    
    public MovingAverageHoltModelSettings(String model,ElasticsearchMovingAverageHoltModelSettingsSettings settings,String window,Boolean minimize,String predict) {
        this.model = model;
        this.settings = settings;
        this.window = window;
        this.minimize = minimize;
        this.predict = predict;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
