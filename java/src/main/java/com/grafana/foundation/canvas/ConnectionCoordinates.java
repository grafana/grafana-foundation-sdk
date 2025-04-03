// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ConnectionCoordinates {
    @JsonProperty("x")
    public Double x;
    @JsonProperty("y")
    public Double y;
    public ConnectionCoordinates() {
    }
    public ConnectionCoordinates(Double x,Double y) {
        this.x = x;
        this.y = y;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
