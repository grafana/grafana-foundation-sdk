// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class QueryEditorOperatorValueTypeDeserializer extends JsonDeserializer<QueryEditorOperatorValueType> {

    @Override
    public QueryEditorOperatorValueType deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        QueryEditorOperatorValueType queryEditorOperatorValueType = new QueryEditorOperatorValueType();
        if (root.isTextual()) {
            queryEditorOperatorValueType.string = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isBoolean()) {
            queryEditorOperatorValueType.bool = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isObject()) {
            queryEditorOperatorValueType.int64 = mapper.convertValue(root, new TypeReference<>() {});
        }
        else if (root.isArray()) {
            queryEditorOperatorValueType.arrayOfQueryEditorOperatorType = mapper.convertValue(root, new TypeReference<>() {});
        }
        
        return queryEditorOperatorValueType;
    }
}
