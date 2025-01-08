// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.LegendDisplayMode;
import com.grafana.foundation.common.LegendPlacement;

public class PieChartLegendOptions {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("values")
    public List<PieChartLegendValues> values;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("displayMode")
    public LegendDisplayMode displayMode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("placement")
    public LegendPlacement placement;
    @JsonProperty("showLegend")
    public Boolean showLegend;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("asTable")
    public Boolean asTable;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isVisible")
    public Boolean isVisible;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sortBy")
    public String sortBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sortDesc")
    public Boolean sortDesc;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("width")
    public Double width;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("calcs")
    public List<String> calcs;
    public PieChartLegendOptions() {
    }
    
    public PieChartLegendOptions(List<PieChartLegendValues> values,LegendDisplayMode displayMode,LegendPlacement placement,Boolean showLegend,Boolean asTable,Boolean isVisible,String sortBy,Boolean sortDesc,Double width,List<String> calcs) {
        this.values = values;
        this.displayMode = displayMode;
        this.placement = placement;
        this.showLegend = showLegend;
        this.asTable = asTable;
        this.isVisible = isVisible;
        this.sortBy = sortBy;
        this.sortDesc = sortDesc;
        this.width = width;
        this.calcs = calcs;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
