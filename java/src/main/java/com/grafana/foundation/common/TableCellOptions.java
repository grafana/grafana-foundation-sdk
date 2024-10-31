// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

// Table cell options. Each cell has a display mode
// and other potential options for that display.
@JsonDeserialize(using = TableCellOptionsDeserializer.class)
public class TableCellOptions {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableAutoCellOptions tableAutoCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableSparklineCellOptions tableSparklineCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableBarGaugeCellOptions tableBarGaugeCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableColoredBackgroundCellOptions tableColoredBackgroundCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableColorTextCellOptions tableColorTextCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableImageCellOptions tableImageCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableDataLinksCellOptions tableDataLinksCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableActionsCellOptions tableActionsCellOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TableJsonViewCellOptions tableJsonViewCellOptions;
    protected TableCellOptions() {}
    public static TableCellOptions createTableAutoCellOptions(com.grafana.foundation.cog.Builder<TableAutoCellOptions> tableAutoCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableAutoCellOptions = tableAutoCellOptions.build();
        return tableCellOptions;
    }
    public static TableCellOptions createTableSparklineCellOptions(com.grafana.foundation.cog.Builder<TableSparklineCellOptions> tableSparklineCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableSparklineCellOptions = tableSparklineCellOptions.build();
        return tableCellOptions;
    }
    public static TableCellOptions createTableBarGaugeCellOptions(com.grafana.foundation.cog.Builder<TableBarGaugeCellOptions> tableBarGaugeCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableBarGaugeCellOptions = tableBarGaugeCellOptions.build();
        return tableCellOptions;
    }
    public static TableCellOptions createTableColoredBackgroundCellOptions(com.grafana.foundation.cog.Builder<TableColoredBackgroundCellOptions> tableColoredBackgroundCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableColoredBackgroundCellOptions = tableColoredBackgroundCellOptions.build();
        return tableCellOptions;
    }
    public static TableCellOptions createTableColorTextCellOptions(com.grafana.foundation.cog.Builder<TableColorTextCellOptions> tableColorTextCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableColorTextCellOptions = tableColorTextCellOptions.build();
        return tableCellOptions;
    }
    public static TableCellOptions createTableImageCellOptions(com.grafana.foundation.cog.Builder<TableImageCellOptions> tableImageCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableImageCellOptions = tableImageCellOptions.build();
        return tableCellOptions;
    }
    public static TableCellOptions createTableDataLinksCellOptions(TableDataLinksCellOptions tableDataLinksCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableDataLinksCellOptions = tableDataLinksCellOptions;
        return tableCellOptions;
    }
    public static TableCellOptions createTableActionsCellOptions(TableActionsCellOptions tableActionsCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableActionsCellOptions = tableActionsCellOptions;
        return tableCellOptions;
    }
    public static TableCellOptions createTableJsonViewCellOptions(TableJsonViewCellOptions tableJsonViewCellOptions) {
        TableCellOptions tableCellOptions = new TableCellOptions();
        tableCellOptions.tableJsonViewCellOptions = tableJsonViewCellOptions;
        return tableCellOptions;
    }
    
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
        if (tableActionsCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableActionsCellOptions);
        }
        if (tableJsonViewCellOptions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(tableJsonViewCellOptions);
        }
        
        return null;
    }

}
