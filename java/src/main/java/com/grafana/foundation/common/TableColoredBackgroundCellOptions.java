// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Colored background cell options
public class TableColoredBackgroundCellOptions {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mode")
    public TableCellBackgroundDisplayMode mode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("applyToRow")
    public Boolean applyToRow;
    public TableColoredBackgroundCellOptions() {
    }
    
    public TableColoredBackgroundCellOptions(String type,TableCellBackgroundDisplayMode mode,Boolean applyToRow) {
        this.type = type;
        this.mode = mode;
        this.applyToRow = applyToRow;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
