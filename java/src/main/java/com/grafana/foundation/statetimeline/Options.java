// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;
import com.grafana.foundation.common.TimelineValueAlignment;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options {
    // Show timeline values on chart 
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls the row height 
    @JsonProperty("rowHeight")
    public Double rowHeight;
    // Merge equal consecutive values 
    @JsonProperty("mergeValues")
    public Boolean mergeValues; 
    @JsonProperty("legend")
    public VizLegendOptions legend; 
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip; 
    @JsonProperty("timezone")
    public List<String> timezone;
    // Controls value alignment on the timelines 
    @JsonProperty("alignValue")
    public TimelineValueAlignment alignValue;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
