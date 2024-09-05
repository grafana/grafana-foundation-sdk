// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.gauge;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.ReduceDataOptions;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonProperty("showThresholdLabels")
    public Boolean showThresholdLabels;
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("showThresholdMarkers")
    public Boolean showThresholdMarkers;
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
