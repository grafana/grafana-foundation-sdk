// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MovingAverageEWMAModelSettings {
    @JsonProperty("model")
    public MovingAverageModel model;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("settings")
    public ElasticsearchMovingAverageEWMAModelSettingsSettings settings;
    @JsonProperty("window")
    public String window;
    @JsonProperty("minimize")
    public Boolean minimize;
    @JsonProperty("predict")
    public String predict;
    public MovingAverageEWMAModelSettings() {
        this.model = MovingAverageModel.EWMA;
        this.window = "";
        this.minimize = false;
        this.predict = "";
    }
    public MovingAverageEWMAModelSettings(ElasticsearchMovingAverageEWMAModelSettingsSettings settings,String window,Boolean minimize,String predict) {
        this.model = MovingAverageModel.EWMA;
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
