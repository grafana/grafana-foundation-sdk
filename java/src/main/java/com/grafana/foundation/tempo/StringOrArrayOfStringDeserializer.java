// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class StringOrArrayOfStringDeserializer extends JsonDeserializer<StringOrArrayOfString> {

    @Override
    public StringOrArrayOfString deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        StringOrArrayOfString stringOrArrayOfString = new StringOrArrayOfString();
        if (root.isTextual()) {
            stringOrArrayOfString.string = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isArray()) {
            stringOrArrayOfString.arrayOfString = mapper.convertValue(root, new TypeReference<>() {});
        }
        
        return stringOrArrayOfString;
    }
}
