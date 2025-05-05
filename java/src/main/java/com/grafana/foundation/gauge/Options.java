// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.gauge;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonProperty("showThresholdLabels")
    public Boolean showThresholdLabels;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("showThresholdMarkers")
    public Boolean showThresholdMarkers;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    public Options() {
        this.showThresholdLabels = false;
        this.reduceOptions = new com.grafana.foundation.common.ReduceDataOptions();
        this.showThresholdMarkers = true;
        this.orientation = VizOrientation.AUTO;
    }
    public Options(Boolean showThresholdLabels,ReduceDataOptions reduceOptions,VizTextDisplayOptions text,Boolean showThresholdMarkers,VizOrientation orientation) {
        this.showThresholdLabels = showThresholdLabels;
        this.reduceOptions = reduceOptions;
        this.text = text;
        this.showThresholdMarkers = showThresholdMarkers;
        this.orientation = orientation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
