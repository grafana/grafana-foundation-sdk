// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.grafana.foundation.common.BarGaugeNamePlacement;
import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options { 
    @JsonProperty("displayMode")
    public BarGaugeDisplayMode displayMode; 
    @JsonProperty("valueMode")
    public BarGaugeValueMode valueMode; 
    @JsonProperty("namePlacement")
    public BarGaugeNamePlacement namePlacement; 
    @JsonProperty("showUnfilled")
    public Boolean showUnfilled; 
    @JsonProperty("sizing")
    public BarGaugeSizing sizing; 
    @JsonProperty("minVizWidth")
    public Integer minVizWidth; 
    @JsonProperty("minVizHeight")
    public Integer minVizHeight; 
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions; 
    @JsonProperty("text")
    public VizTextDisplayOptions text; 
    @JsonProperty("maxVizHeight")
    public Integer maxVizHeight; 
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
