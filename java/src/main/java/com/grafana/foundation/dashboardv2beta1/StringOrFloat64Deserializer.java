// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;

public class StringOrFloat64Deserializer extends JsonDeserializer<StringOrFloat64> {

    @Override
    public StringOrFloat64 deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        StringOrFloat64 stringOrFloat64 = new StringOrFloat64();
        if (root.isTextual()) {
            stringOrFloat64.string = mapper.convertValue(root, String.class);
        }
        else if (root.isDouble()) {
            stringOrFloat64.float64 = mapper.convertValue(root, Double.class);
        }
        
        return stringOrFloat64;
    }
}
