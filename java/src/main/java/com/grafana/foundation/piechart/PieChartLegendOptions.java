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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PieChartLegendOptions> {
        protected final PieChartLegendOptions internal;
        
        public Builder() {
            this.internal = new PieChartLegendOptions();
        }
    public Builder values(List<PieChartLegendValues> values) {
    this.internal.values = values;
        return this;
    }
    
    public Builder displayMode(LegendDisplayMode displayMode) {
    this.internal.displayMode = displayMode;
        return this;
    }
    
    public Builder placement(LegendPlacement placement) {
    this.internal.placement = placement;
        return this;
    }
    
    public Builder showLegend(Boolean showLegend) {
    this.internal.showLegend = showLegend;
        return this;
    }
    
    public Builder asTable(Boolean asTable) {
    this.internal.asTable = asTable;
        return this;
    }
    
    public Builder isVisible(Boolean isVisible) {
    this.internal.isVisible = isVisible;
        return this;
    }
    
    public Builder sortBy(String sortBy) {
    this.internal.sortBy = sortBy;
        return this;
    }
    
    public Builder sortDesc(Boolean sortDesc) {
    this.internal.sortDesc = sortDesc;
        return this;
    }
    
    public Builder width(Double width) {
    this.internal.width = width;
        return this;
    }
    
    public Builder calcs(List<String> calcs) {
    this.internal.calcs = calcs;
        return this;
    }
    public PieChartLegendOptions build() {
            return this.internal;
        }
    }
}
