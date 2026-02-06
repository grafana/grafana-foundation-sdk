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
import java.util.Map;
import java.util.List;

public class StringOrArrayOfStringDeserializer extends JsonDeserializer<StringOrArrayOfString> {

    @Override
    public StringOrArrayOfString deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        StringOrArrayOfString stringOrArrayOfString = new StringOrArrayOfString();
        if (root.isTextual()) {
            stringOrArrayOfString.string = mapper.convertValue(root, String.class);
        }
        else if (root.isArray()) {
            stringOrArrayOfString.arrayOfString = mapper.convertValue(root, new TypeReference<List<String>>() {});
        }
        
        return stringOrArrayOfString;
    }
}
