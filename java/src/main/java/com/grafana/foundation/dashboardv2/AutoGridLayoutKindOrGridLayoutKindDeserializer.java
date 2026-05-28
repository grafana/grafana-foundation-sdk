// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class AutoGridLayoutKindOrGridLayoutKindDeserializer extends JsonDeserializer<AutoGridLayoutKindOrGridLayoutKind> {

    @Override
    public AutoGridLayoutKindOrGridLayoutKind deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        AutoGridLayoutKindOrGridLayoutKind autoGridLayoutKindOrGridLayoutKind = new AutoGridLayoutKindOrGridLayoutKind();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for AutoGridLayoutKindOrGridLayoutKind");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "AutoGridLayout":
            autoGridLayoutKindOrGridLayoutKind.autoGridLayoutKind = mapper.convertValue(root, AutoGridLayoutKind.class);
            break;
        case "GridLayout":
            autoGridLayoutKindOrGridLayoutKind.gridLayoutKind = mapper.convertValue(root, GridLayoutKind.class);
            break;
        }
        
        return autoGridLayoutKindOrGridLayoutKind;
    }
}
