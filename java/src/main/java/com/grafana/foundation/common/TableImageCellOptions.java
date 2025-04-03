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
    public TableCellDisplayMode type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alt")
    public String alt;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("title")
    public String title;
    public TableImageCellOptions() {
        this.type = TableCellDisplayMode.IMAGE;
    }
    public TableImageCellOptions(String alt,String title) {
        this.type = TableCellDisplayMode.IMAGE;
        this.alt = alt;
        this.title = title;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
