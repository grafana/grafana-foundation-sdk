// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class CellContentKindDeserializer extends JsonDeserializer<CellContentKind> {

    @Override
    public CellContentKind deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        CellContentKind cellContentKind = new CellContentKind();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for CellContentKind");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "Code":
            cellContentKind.codeCellContentKind = mapper.convertValue(root, CodeCellContentKind.class);
            break;
        case "Markdown":
            cellContentKind.markdownCellContentKind = mapper.convertValue(root, MarkdownCellContentKind.class);
            break;
        }
        
        return cellContentKind;
    }
}
