// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import java.util.List;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options { 
    @JsonProperty("pieType")
    public PieChartType pieType; 
    @JsonProperty("displayLabels")
    public List<PieChartLabels> displayLabels; 
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip; 
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions; 
    @JsonProperty("text")
    public VizTextDisplayOptions text; 
    @JsonProperty("legend")
    public PieChartLegendOptions legend; 
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
