// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

// TODO docs
public class VizLegendOptions {
    @JsonProperty("displayMode")
    public LegendDisplayMode displayMode;
    @JsonProperty("placement")
    public LegendPlacement placement;
    @JsonProperty("showLegend")
    public Boolean showLegend;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("asTable")
    public Boolean asTable;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("isVisible")
    public Boolean isVisible;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("sortBy")
    public String sortBy;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("sortDesc")
    public Boolean sortDesc;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("width")
    public Double width;
    @JsonProperty("calcs")
    public List<String> calcs;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<VizLegendOptions> {
        private final VizLegendOptions internal;
        
        public Builder() {
            this.internal = new VizLegendOptions();
        this.displayMode(LegendDisplayMode.LIST);
        this.placement(LegendPlacement.BOTTOM);
        this.calcs(List.of());
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
    public VizLegendOptions build() {
            return this.internal;
        }
    }
}
