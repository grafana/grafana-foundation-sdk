// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.grafana.foundation.common.HeatmapCellLayout;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Controls frame rows options
public class RowsHeatmapOptions {
    // Sets the name of the cell when not calculating from data 
    @JsonProperty("value")
    public String value;
    // Controls tick alignment when not calculating from data 
    @JsonProperty("layout")
    public HeatmapCellLayout layout;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
