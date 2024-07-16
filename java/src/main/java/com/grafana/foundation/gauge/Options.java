// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.gauge;

import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options { 
    @JsonProperty("showThresholdLabels")
    public Boolean showThresholdLabels; 
    @JsonProperty("showThresholdMarkers")
    public Boolean showThresholdMarkers; 
    @JsonProperty("sizing")
    public BarGaugeSizing sizing; 
    @JsonProperty("minVizWidth")
    public Integer minVizWidth; 
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions; 
    @JsonProperty("text")
    public VizTextDisplayOptions text; 
    @JsonProperty("minVizHeight")
    public Integer minVizHeight; 
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
