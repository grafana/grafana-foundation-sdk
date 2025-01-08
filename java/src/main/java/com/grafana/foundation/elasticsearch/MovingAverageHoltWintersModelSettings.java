// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MovingAverageHoltWintersModelSettings {
    @JsonProperty("model")
    public String model;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("settings")
    public ElasticsearchMovingAverageHoltWintersModelSettingsSettings settings;
    @JsonProperty("window")
    public String window;
    @JsonProperty("minimize")
    public Boolean minimize;
    @JsonProperty("predict")
    public String predict;
    public MovingAverageHoltWintersModelSettings() {
    }
    
    public MovingAverageHoltWintersModelSettings(String model,ElasticsearchMovingAverageHoltWintersModelSettingsSettings settings,String window,Boolean minimize,String predict) {
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
