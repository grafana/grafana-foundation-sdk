// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

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

public class StringOrBoolDeserializer extends JsonDeserializer<StringOrBool> {

    @Override
    public StringOrBool deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        StringOrBool stringOrBool = new StringOrBool();
        if (root.isTextual()) {
            stringOrBool.string = mapper.convertValue(root, String.class);
        }
        else if (root.isBoolean()) {
            stringOrBool.bool = mapper.convertValue(root, Boolean.class);
        }
        
        return stringOrBool;
    }
}
