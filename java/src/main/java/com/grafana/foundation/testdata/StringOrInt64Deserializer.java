// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

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
import com.grafana.foundation.common.DataSourceRef;

public class StringOrInt64Deserializer extends JsonDeserializer<StringOrInt64> {

    @Override
    public StringOrInt64 deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        StringOrInt64 stringOrInt64 = new StringOrInt64();
        if (root.isTextual()) {
            stringOrInt64.string = mapper.convertValue(root, String.class);
        }
        else if (root.isIntegralNumber()) {
            stringOrInt64.int64 = mapper.convertValue(root, Long.class);
        }
        
        return stringOrInt64;
    }
}
