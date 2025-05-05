// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class MovingAverageHoltWintersModelSettings {
    @JsonProperty("model")
    public MovingAverageModel model;
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
        this.model = MovingAverageModel.HOLT_WINTERS;
        this.settings = new com.grafana.foundation.elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings();
        this.window = "";
        this.minimize = false;
        this.predict = "";
    }
    public MovingAverageHoltWintersModelSettings(ElasticsearchMovingAverageHoltWintersModelSettingsSettings settings,String window,Boolean minimize,String predict) {
        this.model = MovingAverageModel.HOLT_WINTERS;
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
