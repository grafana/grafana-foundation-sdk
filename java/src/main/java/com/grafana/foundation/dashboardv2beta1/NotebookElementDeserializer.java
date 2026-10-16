// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class NotebookElementDeserializer extends JsonDeserializer<NotebookElement> {

    @Override
    public NotebookElement deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        NotebookElement notebookElement = new NotebookElement();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for NotebookElement");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "Cell":
            notebookElement.cellKind = mapper.convertValue(root, CellKind.class);
            break;
        case "LibraryPanel":
            notebookElement.libraryPanelKind = mapper.convertValue(root, LibraryPanelKind.class);
            break;
        case "Panel":
            notebookElement.v2PanelKind = mapper.convertValue(root, V2PanelKind.class);
            break;
        }
        
        return notebookElement;
    }
}
