// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.dashboard.DataSourceRef;

public class BoolOrUint32Deserializer extends JsonDeserializer<BoolOrUint32> {

    @Override
    public BoolOrUint32 deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        BoolOrUint32 boolOrUint32 = new BoolOrUint32();
        if (root.isBoolean()) {
            boolOrUint32.bool = mapper.convertValue(root, Boolean.class);
        }
        else if (root.isInt()) {
            boolOrUint32.uint32 = mapper.convertValue(root, Integer.class);
        }
        
        return boolOrUint32;
    }
}
