// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VisibilityMode;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.HeatmapCalculationOptions;

public class Options {
    // Controls if the heatmap should be calculated from data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("calculate")
    public Boolean calculate;
    // Calculation options for the heatmap
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("calculation")
    public HeatmapCalculationOptions calculation;
    // Controls the color options
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("color")
    public HeatmapColorOptions color;
    // Filters values between a given range
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filterValues")
    public FilterValueRange filterValues;
    // Controls tick alignment and value name when not calculating from data
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rowsFrame")
    public RowsHeatmapOptions rowsFrame;
    // | *{
    // 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls the display of the value in the cell
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls gap between cells
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cellGap")
    public Integer cellGap;
    // Controls cell radius
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cellRadius")
    public Float cellRadius;
    // Controls cell value unit
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cellValues")
    public CellValues cellValues;
    // Controls yAxis placement
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("yAxis")
    public YAxisConfig yAxis;
    // | *{
    // 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    // }
    // Controls legend options
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public HeatmapLegend legend;
    // Controls tooltip options
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public HeatmapTooltip tooltip;
    // Controls exemplar options
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("exemplars")
    public ExemplarConfig exemplars;
    public Options() {
        this.calculate = false;
        this.color = new HeatmapColorOptions(HeatmapColorMode.OPACITY, "Oranges", "dark-orange", HeatmapColorScale.LINEAR, 0.5f, 64L, false, 0.0f, 0.0f);
        this.filterValues = new FilterValueRange(0.0f, 0.0f);
        this.showValue = VisibilityMode.AUTO;
        this.cellGap = 1;
        this.legend = new HeatmapLegend(true);
        this.exemplars = new ExemplarConfig("rgba(255,0,255,0.7)");
    }
    public Options(Boolean calculate,HeatmapCalculationOptions calculation,HeatmapColorOptions color,FilterValueRange filterValues,RowsHeatmapOptions rowsFrame,VisibilityMode showValue,Integer cellGap,Float cellRadius,CellValues cellValues,YAxisConfig yAxis,HeatmapLegend legend,HeatmapTooltip tooltip,ExemplarConfig exemplars) {
        this.calculate = calculate;
        this.calculation = calculation;
        this.color = color;
        this.filterValues = filterValues;
        this.rowsFrame = rowsFrame;
        this.showValue = showValue;
        this.cellGap = cellGap;
        this.cellRadius = cellRadius;
        this.cellValues = cellValues;
        this.yAxis = yAxis;
        this.legend = legend;
        this.tooltip = tooltip;
        this.exemplars = exemplars;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
