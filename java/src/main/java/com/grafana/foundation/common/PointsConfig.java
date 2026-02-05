// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class PointsConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showPoints")
    public VisibilityMode showPoints;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSize")
    public Double pointSize;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointColor")
    public String pointColor;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("pointSymbol")
    public String pointSymbol;
    public PointsConfig() {
    }
    public PointsConfig(VisibilityMode showPoints,Double pointSize,String pointColor,String pointSymbol) {
        this.showPoints = showPoints;
        this.pointSize = pointSize;
        this.pointColor = pointColor;
        this.pointSymbol = pointSymbol;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
