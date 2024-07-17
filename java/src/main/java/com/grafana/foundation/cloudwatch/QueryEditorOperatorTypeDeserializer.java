// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class QueryEditorOperatorTypeDeserializer extends JsonDeserializer<QueryEditorOperatorType> {

    @Override
    public QueryEditorOperatorType deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        QueryEditorOperatorType queryEditorOperatorType = new QueryEditorOperatorType();
        if (root.isTextual()) {
            queryEditorOperatorType.string = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isBoolean()) {
            queryEditorOperatorType.bool = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isObject()) {
            queryEditorOperatorType.int64 = mapper.convertValue(root, new TypeReference<>() {});
        }
        
        return queryEditorOperatorType;
    }
}
