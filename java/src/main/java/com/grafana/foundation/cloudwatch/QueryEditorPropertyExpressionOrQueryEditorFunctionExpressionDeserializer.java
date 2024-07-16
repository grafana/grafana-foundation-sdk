// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class QueryEditorPropertyExpressionOrQueryEditorFunctionExpressionDeserializer extends JsonDeserializer<QueryEditorPropertyExpressionOrQueryEditorFunctionExpression> {

    @Override
    public QueryEditorPropertyExpressionOrQueryEditorFunctionExpression deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        QueryEditorPropertyExpressionOrQueryEditorFunctionExpression queryEditorPropertyExpressionOrQueryEditorFunctionExpression = new QueryEditorPropertyExpressionOrQueryEditorFunctionExpression();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for QueryEditorPropertyExpressionOrQueryEditorFunctionExpression");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "function":
            queryEditorPropertyExpressionOrQueryEditorFunctionExpression.queryEditorFunctionExpression = mapper.convertValue(root, QueryEditorFunctionExpression.class);
            break;
        case "property":
            queryEditorPropertyExpressionOrQueryEditorFunctionExpression.queryEditorPropertyExpression = mapper.convertValue(root, QueryEditorPropertyExpression.class);
            break;
        }
        
        return queryEditorPropertyExpressionOrQueryEditorFunctionExpression;
    }
}
