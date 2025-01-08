// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class MovingAverageLinearModelSettings {
    @JsonProperty("model")
    public String model;
    @JsonProperty("window")
    public String window;
    @JsonProperty("predict")
    public String predict;
    public MovingAverageLinearModelSettings() {
    }
    
    public MovingAverageLinearModelSettings(String model,String window,String predict) {
        this.model = model;
        this.window = window;
        this.predict = predict;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
