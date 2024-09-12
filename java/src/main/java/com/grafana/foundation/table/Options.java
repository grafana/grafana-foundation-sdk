// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.table;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.common.TableSortByFieldState;
import com.grafana.foundation.common.TableFooterOptions;
import com.grafana.foundation.common.TableCellHeight;

public class Options {
    // Represents the index of the selected frame
    @JsonProperty("frameIndex")
    public Double frameIndex;
    // Controls whether the panel should show the header
    @JsonProperty("showHeader")
    public Boolean showHeader;
    // Controls whether the header should show icons for the column types
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("showTypeIcons")
    public Boolean showTypeIcons;
    // Used to control row sorting
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("sortBy")
    public List<TableSortByFieldState> sortBy;
    // Controls footer options
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("footer")
    public TableFooterOptions footer;
    // Controls the height of the rows
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("cellHeight")
    public TableCellHeight cellHeight;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
