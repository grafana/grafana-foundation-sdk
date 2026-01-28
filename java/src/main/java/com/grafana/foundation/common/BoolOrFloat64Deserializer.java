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
import com.grafana.foundation.cog.variants.PanelConfig;
import com.grafana.foundation.common.DataSourceRef;
import java.util.Map;
import java.util.List;

public class BoolOrFloat64Deserializer extends JsonDeserializer<BoolOrFloat64> {

    @Override
    public BoolOrFloat64 deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        BoolOrFloat64 boolOrFloat64 = new BoolOrFloat64();
        if (root.isBoolean()) {
            boolOrFloat64.bool = mapper.convertValue(root, Boolean.class);
        }
        else if (root.isDouble()) {
            boolOrFloat64.float64 = mapper.convertValue(root, Double.class);
        }
        
        return boolOrFloat64;
    }
}
