// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.gauge;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.BarGaugeSizing;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonProperty("showThresholdLabels")
    public Boolean showThresholdLabels;
    @JsonProperty("showThresholdMarkers")
    public Boolean showThresholdMarkers;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sizing")
    public BarGaugeSizing sizing;
    @JsonProperty("minVizWidth")
    public Integer minVizWidth;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("minVizHeight")
    public Integer minVizHeight;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    public Options() {
        this.showThresholdLabels = false;
        this.showThresholdMarkers = true;
        this.sizing = BarGaugeSizing.AUTO;
        this.minVizWidth = 75;
        this.reduceOptions = new com.grafana.foundation.common.ReduceDataOptions();
        this.minVizHeight = 75;
        this.orientation = VizOrientation.AUTO;
    }
    public Options(Boolean showThresholdLabels,Boolean showThresholdMarkers,BarGaugeSizing sizing,Integer minVizWidth,ReduceDataOptions reduceOptions,VizTextDisplayOptions text,Integer minVizHeight,VizOrientation orientation) {
        this.showThresholdLabels = showThresholdLabels;
        this.showThresholdMarkers = showThresholdMarkers;
        this.sizing = sizing;
        this.minVizWidth = minVizWidth;
        this.reduceOptions = reduceOptions;
        this.text = text;
        this.minVizHeight = minVizHeight;
        this.orientation = orientation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
