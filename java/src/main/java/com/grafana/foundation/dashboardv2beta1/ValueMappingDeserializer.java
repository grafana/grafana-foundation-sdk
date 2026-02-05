// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class ValueMappingDeserializer extends JsonDeserializer<ValueMapping> {

    @Override
    public ValueMapping deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        ValueMapping valueMapping = new ValueMapping();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for ValueMapping");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "range":
            valueMapping.rangeMap = mapper.convertValue(root, RangeMap.class);
            break;
        case "regex":
            valueMapping.regexMap = mapper.convertValue(root, RegexMap.class);
            break;
        case "special":
            valueMapping.specialValueMap = mapper.convertValue(root, SpecialValueMap.class);
            break;
        case "value":
            valueMapping.valueMap = mapper.convertValue(root, ValueMap.class);
            break;
        }
        
        return valueMapping;
    }
}
