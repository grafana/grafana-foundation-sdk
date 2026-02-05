// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class ElementDeserializer extends JsonDeserializer<Element> {

    @Override
    public Element deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        Element element = new Element();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for Element");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "LibraryPanel":
            element.libraryPanelKind = mapper.convertValue(root, LibraryPanelKind.class);
            break;
        case "Panel":
            element.panelKind = mapper.convertValue(root, PanelKind.class);
            break;
        }
        
        return element;
    }
}
