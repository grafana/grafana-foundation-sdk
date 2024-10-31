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
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("wrapText")
    public Boolean wrapText;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableColorTextCellOptions> {
        protected final TableColorTextCellOptions internal;
        
        public Builder() {
            this.internal = new TableColorTextCellOptions();
    this.internal.type = "color-text";
        }
    public Builder wrapText(Boolean wrapText) {
    this.internal.wrapText = wrapText;
        return this;
    }
    public TableColorTextCellOptions build() {
            return this.internal;
        }
    }
}
