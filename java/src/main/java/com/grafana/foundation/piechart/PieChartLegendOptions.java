// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import java.util.List;
import com.grafana.foundation.common.LegendDisplayMode;
import com.grafana.foundation.common.LegendPlacement;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PieChartLegendOptions { 
    @JsonProperty("values")
    public List<PieChartLegendValues> values; 
    @JsonProperty("displayMode")
    public LegendDisplayMode displayMode; 
    @JsonProperty("placement")
    public LegendPlacement placement; 
    @JsonProperty("showLegend")
    public Boolean showLegend; 
    @JsonProperty("asTable")
    public Boolean asTable; 
    @JsonProperty("isVisible")
    public Boolean isVisible; 
    @JsonProperty("sortBy")
    public String sortBy; 
    @JsonProperty("sortDesc")
    public Boolean sortDesc; 
    @JsonProperty("width")
    public Double width; 
    @JsonProperty("calcs")
    public List<String> calcs;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
