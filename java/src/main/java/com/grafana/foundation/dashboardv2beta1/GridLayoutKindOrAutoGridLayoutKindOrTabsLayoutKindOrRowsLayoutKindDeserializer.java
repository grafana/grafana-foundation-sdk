// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKindDeserializer extends JsonDeserializer<GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind> {

    @Override
    public GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "AutoGridLayout":
            gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.autoGridLayoutKind = mapper.convertValue(root, AutoGridLayoutKind.class);
            break;
        case "GridLayout":
            gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.gridLayoutKind = mapper.convertValue(root, GridLayoutKind.class);
            break;
        case "RowsLayout":
            gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.rowsLayoutKind = mapper.convertValue(root, RowsLayoutKind.class);
            break;
        case "TabsLayout":
            gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.tabsLayoutKind = mapper.convertValue(root, TabsLayoutKind.class);
            break;
        }
        
        return gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
    }
}
