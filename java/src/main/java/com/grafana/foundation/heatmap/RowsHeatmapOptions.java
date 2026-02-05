// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.HeatmapCellLayout;

// Controls frame rows options
public class RowsHeatmapOptions {
    // Sets the name of the cell when not calculating from data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public String value;
    // Controls tick alignment when not calculating from data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("layout")
    public HeatmapCellLayout layout;
    public RowsHeatmapOptions() {
    }
    public RowsHeatmapOptions(String value,HeatmapCellLayout layout) {
        this.value = value;
        this.layout = layout;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
