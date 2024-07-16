// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// Colored background cell options
public class TableColoredBackgroundCellOptions { 
    @JsonProperty("type")
    public String type; 
    @JsonProperty("mode")
    public TableCellBackgroundDisplayMode mode;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableColoredBackgroundCellOptions> {
        private final TableColoredBackgroundCellOptions internal;
        
        public Builder() {
            this.internal = new TableColoredBackgroundCellOptions();
    this.internal.type = "color-background";
        }
    public Builder mode(TableCellBackgroundDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    public TableColoredBackgroundCellOptions build() {
            return this.internal;
        }
    }
}
