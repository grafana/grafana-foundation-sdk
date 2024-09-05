// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

// Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
// Generally defines alignment, filtering capabilties, display options, etc.
public class TableFieldOptions {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("width")
    public Double width;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("minWidth")
    public Double minWidth;
    @JsonProperty("align")
    public FieldTextAlignment align;
    // This field is deprecated in favor of using cellOptions
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("displayMode")
    public TableCellDisplayMode displayMode;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("cellOptions")
    public TableCellOptions cellOptions;
    // ?? default is missing or false ??
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hidden")
    public Boolean hidden;
    @JsonProperty("inspect")
    public Boolean inspect;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("filterable")
    public Boolean filterable;
    // Hides any header for a column, usefull for columns that show some static content or buttons.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("hideHeader")
    public Boolean hideHeader;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TableFieldOptions> {
        private final TableFieldOptions internal;
        
        public Builder() {
            this.internal = new TableFieldOptions();
        this.align(FieldTextAlignment.AUTO);
        this.inspect(false);
        }
    public Builder width(Double width) {
    this.internal.width = width;
        return this;
    }
    
    public Builder minWidth(Double minWidth) {
    this.internal.minWidth = minWidth;
        return this;
    }
    
    public Builder align(FieldTextAlignment align) {
    this.internal.align = align;
        return this;
    }
    
    public Builder displayMode(TableCellDisplayMode displayMode) {
    this.internal.displayMode = displayMode;
        return this;
    }
    
    public Builder cellOptions(TableCellOptions cellOptions) {
    this.internal.cellOptions = cellOptions;
        return this;
    }
    
    public Builder hidden(Boolean hidden) {
    this.internal.hidden = hidden;
        return this;
    }
    
    public Builder inspect(Boolean inspect) {
    this.internal.inspect = inspect;
        return this;
    }
    
    public Builder filterable(Boolean filterable) {
    this.internal.filterable = filterable;
        return this;
    }
    
    public Builder hideHeader(Boolean hideHeader) {
    this.internal.hideHeader = hideHeader;
        return this;
    }
    public TableFieldOptions build() {
            return this.internal;
        }
    }
}
