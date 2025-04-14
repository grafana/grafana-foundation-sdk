// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class XYSeriesConfig {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public XychartXYSeriesConfigName name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("frame")
    public XychartXYSeriesConfigFrame frame;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("x")
    public XychartXYSeriesConfigX x;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("y")
    public XychartXYSeriesConfigY y;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("color")
    public XychartXYSeriesConfigColor color;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("size")
    public XychartXYSeriesConfigSize size;
    public XYSeriesConfig() {
    }
    public XYSeriesConfig(XychartXYSeriesConfigName name,XychartXYSeriesConfigFrame frame,XychartXYSeriesConfigX x,XychartXYSeriesConfigY y,XychartXYSeriesConfigColor color,XychartXYSeriesConfigSize size) {
        this.name = name;
        this.frame = frame;
        this.x = x;
        this.y = y;
        this.color = color;
        this.size = size;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
