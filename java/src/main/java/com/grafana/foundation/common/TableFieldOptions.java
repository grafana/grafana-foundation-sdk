// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
// Generally defines alignment, filtering capabilties, display options, etc.
public class TableFieldOptions {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("width")
    public Double width;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("minWidth")
    public Double minWidth;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("align")
    public FieldTextAlignment align;
    // This field is deprecated in favor of using cellOptions
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("displayMode")
    public TableCellDisplayMode displayMode;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cellOptions")
    public TableCellOptions cellOptions;
    // ?? default is missing or false ??
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hidden")
    public Boolean hidden;
    @JsonProperty("inspect")
    public Boolean inspect;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("filterable")
    public Boolean filterable;
    // Hides any header for a column, usefull for columns that show some static content or buttons.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hideHeader")
    public Boolean hideHeader;
    public TableFieldOptions() {
        this.align = FieldTextAlignment.AUTO;
        this.inspect = false;
    }
    public TableFieldOptions(Double width,Double minWidth,FieldTextAlignment align,TableCellDisplayMode displayMode,TableCellOptions cellOptions,Boolean hidden,Boolean inspect,Boolean filterable,Boolean hideHeader) {
        this.width = width;
        this.minWidth = minWidth;
        this.align = align;
        this.displayMode = displayMode;
        this.cellOptions = cellOptions;
        this.hidden = hidden;
        this.inspect = inspect;
        this.filterable = filterable;
        this.hideHeader = hideHeader;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
