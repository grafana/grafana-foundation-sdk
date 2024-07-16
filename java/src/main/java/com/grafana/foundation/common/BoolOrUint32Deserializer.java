// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class BoolOrUint32Deserializer extends JsonDeserializer<BoolOrUint32> {

    @Override
    public BoolOrUint32 deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        BoolOrUint32 boolOrUint32 = new BoolOrUint32();
        if (root.isBoolean()) {
            boolOrUint32.bool = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isInt()) {
            boolOrUint32.uint32 = mapper.convertValue(root, new TypeReference<>() {});
        }
        
        return boolOrUint32;
    }
}
