// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Gauge cell options
public class TableBarGaugeCellOptions { 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("mode")
    public BarGaugeDisplayMode mode; 
    @JsonProperty("valueDisplayMode")
    public BarGaugeValueMode valueDisplayMode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableBarGaugeCellOptions> {
        private final TableBarGaugeCellOptions internal;
        
        public Builder() {
            this.internal = new TableBarGaugeCellOptions();
    this.internal.type = "gauge";
        }
    public Builder mode(BarGaugeDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder valueDisplayMode(BarGaugeValueMode valueDisplayMode) {
    this.internal.valueDisplayMode = valueDisplayMode;
        return this;
    }
    public TableBarGaugeCellOptions build() {
            return this.internal;
        }
    }
}
