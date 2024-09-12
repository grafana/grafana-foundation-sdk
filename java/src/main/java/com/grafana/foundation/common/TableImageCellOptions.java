// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Json view cell options
public class TableImageCellOptions {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alt")
    public String alt;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableImageCellOptions> {
        private final TableImageCellOptions internal;
        
        public Builder() {
            this.internal = new TableImageCellOptions();
    this.internal.type = "image";
        }
    public Builder alt(String alt) {
    this.internal.alt = alt;
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    public TableImageCellOptions build() {
            return this.internal;
        }
    }
}
