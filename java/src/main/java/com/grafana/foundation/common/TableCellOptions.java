// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

// Table cell options. Each cell has a display mode
// and other potential options for that display.
@JsonDeserialize(using = TableCellOptionsDeserializer.class)
public class TableCellOptions { 
    @JsonUnwrapped
    public TableAutoCellOptions tableAutoCellOptions; 
    @JsonUnwrapped
    public TableSparklineCellOptions tableSparklineCellOptions; 
    @JsonUnwrapped
    public TableBarGaugeCellOptions tableBarGaugeCellOptions; 
    @JsonUnwrapped
    public TableColoredBackgroundCellOptions tableColoredBackgroundCellOptions; 
    @JsonUnwrapped
    public TableColorTextCellOptions tableColorTextCellOptions; 
    @JsonUnwrapped
    public TableImageCellOptions tableImageCellOptions; 
    @JsonUnwrapped
    public TableDataLinksCellOptions tableDataLinksCellOptions; 
    @JsonUnwrapped
    public TableJsonViewCellOptions tableJsonViewCellOptions;
    
    public String toJSON() throws JsonProcessingException {
        if (tableAutoCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableAutoCellOptions);
        }
        if (tableSparklineCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableSparklineCellOptions);
        }
        if (tableBarGaugeCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableBarGaugeCellOptions);
        }
        if (tableColoredBackgroundCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableColoredBackgroundCellOptions);
        }
        if (tableColorTextCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableColorTextCellOptions);
        }
        if (tableImageCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableImageCellOptions);
        }
        if (tableDataLinksCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableDataLinksCellOptions);
        }
        if (tableJsonViewCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableJsonViewCellOptions);
        }
        
        return null;
    }

}
