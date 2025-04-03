// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class StringOrMapDeserializer extends JsonDeserializer<StringOrMap> {

    @Override
    public StringOrMap deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        StringOrMap stringOrMap = new StringOrMap();
        if (root.isTextual()) {
            stringOrMap.string = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isObject()) {
            stringOrMap.map = mapper.convertValue(root, new TypeReference<>() {});
        }
        
        return stringOrMap;
    }
}
