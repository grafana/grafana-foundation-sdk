// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class TableCellOptionsDeserializer extends JsonDeserializer<TableCellOptions> {

    @Override
    public TableCellOptions deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        TableCellOptions tableCellOptions = new TableCellOptions();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for TableCellOptions");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "actions":
            tableCellOptions.tableActionsCellOptions = mapper.convertValue(root, TableActionsCellOptions.class);
            break;
        case "auto":
            tableCellOptions.tableAutoCellOptions = mapper.convertValue(root, TableAutoCellOptions.class);
            break;
        case "color-background":
            tableCellOptions.tableColoredBackgroundCellOptions = mapper.convertValue(root, TableColoredBackgroundCellOptions.class);
            break;
        case "color-text":
            tableCellOptions.tableColorTextCellOptions = mapper.convertValue(root, TableColorTextCellOptions.class);
            break;
        case "data-links":
            tableCellOptions.tableDataLinksCellOptions = mapper.convertValue(root, TableDataLinksCellOptions.class);
            break;
        case "gauge":
            tableCellOptions.tableBarGaugeCellOptions = mapper.convertValue(root, TableBarGaugeCellOptions.class);
            break;
        case "image":
            tableCellOptions.tableImageCellOptions = mapper.convertValue(root, TableImageCellOptions.class);
            break;
        case "json-view":
            tableCellOptions.tableJsonViewCellOptions = mapper.convertValue(root, TableJsonViewCellOptions.class);
            break;
        case "sparkline":
            tableCellOptions.tableSparklineCellOptions = mapper.convertValue(root, TableSparklineCellOptions.class);
            break;
        }
        
        return tableCellOptions;
    }
}
