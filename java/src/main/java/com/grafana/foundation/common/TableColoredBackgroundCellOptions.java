// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Colored background cell options
public class TableColoredBackgroundCellOptions {
    @JsonProperty("type")
    public String type;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("mode")
    public TableCellBackgroundDisplayMode mode;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("applyToRow")
    public Boolean applyToRow;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("wrapText")
    public Boolean wrapText;
    
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
    
    public Builder applyToRow(Boolean applyToRow) {
    this.internal.applyToRow = applyToRow;
        return this;
    }
    
    public Builder wrapText(Boolean wrapText) {
    this.internal.wrapText = wrapText;
        return this;
    }
    public TableColoredBackgroundCellOptions build() {
            return this.internal;
        }
    }
}
