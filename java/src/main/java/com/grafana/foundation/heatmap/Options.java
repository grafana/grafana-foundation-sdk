// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.HeatmapCalculationOptions;
import com.grafana.foundation.common.VisibilityMode;

public class Options {
    // Controls if the heatmap should be calculated from data
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("calculate")
    public Boolean calculate;
    // Calculation options for the heatmap
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("calculation")
    public HeatmapCalculationOptions calculation;
    // Controls the color options
    @JsonProperty("color")
    public HeatmapColorOptions color;
    // Filters values between a given range
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("filterValues")
    public FilterValueRange filterValues;
    // Controls tick alignment and value name when not calculating from data
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("rowsFrame")
    public RowsHeatmapOptions rowsFrame;
    // | *{
    // 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls the display of the value in the cell
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls gap between cells
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("cellGap")
    public Integer cellGap;
    // Controls cell radius
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("cellRadius")
    public Float cellRadius;
    // Controls cell value unit
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("cellValues")
    public CellValues cellValues;
    // Controls yAxis placement
    @JsonProperty("yAxis")
    public YAxisConfig yAxis;
    // | *{
    // 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls legend options
    @JsonProperty("legend")
    public HeatmapLegend legend;
    // Controls tooltip options
    @JsonProperty("tooltip")
    public HeatmapTooltip tooltip;
    // Controls exemplar options
    @JsonProperty("exemplars")
    public ExemplarConfig exemplars;
    // Controls which axis to allow selection on
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("selectionMode")
    public HeatmapSelectionMode selectionMode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
