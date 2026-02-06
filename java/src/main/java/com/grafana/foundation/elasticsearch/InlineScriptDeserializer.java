// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.List;

public class InlineScriptDeserializer extends JsonDeserializer<InlineScript> {

    @Override
    public InlineScript deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        InlineScript inlineScript = new InlineScript();
        if (root.isTextual()) {
            inlineScript.string = mapper.convertValue(root, String.class);
        }
        else if (root.isObject() && couldBe(mapper, root, ElasticsearchInlineScript.class)) {
            inlineScript.elasticsearchInlineScript = mapper.convertValue(root, ElasticsearchInlineScript.class);
        }
        
        return inlineScript;
    }
    
    private <T> boolean couldBe(ObjectMapper mapper, JsonNode root, Class<T> clazz) {
        try {
            mapper.convertValue(root, clazz);
        } catch (Exception e) {
           return false;
        }
        
        return true;
    }
}
