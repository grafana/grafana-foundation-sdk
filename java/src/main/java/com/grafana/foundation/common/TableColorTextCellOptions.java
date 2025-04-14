// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Colored text cell options
public class TableColorTextCellOptions {
    @JsonProperty("type")
    public TableCellDisplayMode type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("wrapText")
    public Boolean wrapText;
    public TableColorTextCellOptions() {
        this.type = TableCellDisplayMode.COLOR_TEXT;
    }
    public TableColorTextCellOptions(Boolean wrapText) {
        this.type = TableCellDisplayMode.COLOR_TEXT;
        this.wrapText = wrapText;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
