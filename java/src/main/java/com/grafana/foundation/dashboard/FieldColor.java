// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Map a field to a color.
public class FieldColor {
    // The main color scheme mode.
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public FieldColorModeId mode;
    // The fixed color value for fixed or shades color modes.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixedColor")
    public String fixedColor;
    // Some visualizations need to know how to assign a series color from by value color schemes.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("seriesBy")
    public FieldColorSeriesByMode seriesBy;
    public FieldColor() {
        this.mode = FieldColorModeId.THRESHOLDS;
    }
    public FieldColor(FieldColorModeId mode,String fixedColor,FieldColorSeriesByMode seriesBy) {
        this.mode = mode;
        this.fixedColor = fixedColor;
        this.seriesBy = seriesBy;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
