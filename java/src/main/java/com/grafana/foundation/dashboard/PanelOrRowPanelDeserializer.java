// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class PanelOrRowPanelDeserializer extends JsonDeserializer<PanelOrRowPanel> {

    @Override
    public PanelOrRowPanel deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        PanelOrRowPanel panelOrRowPanel = new PanelOrRowPanel();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for PanelOrRowPanel");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        default:
            panelOrRowPanel.panel = mapper.convertValue(root, Panel.class);
            break;
        case "row":
            panelOrRowPanel.rowPanel = mapper.convertValue(root, RowPanel.class);
            break;
        }
        
        return panelOrRowPanel;
    }
}
