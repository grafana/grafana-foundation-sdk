// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class QueryEditorExpressionDeserializer extends JsonDeserializer<QueryEditorExpression> {

    @Override
    public QueryEditorExpression deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        QueryEditorExpression queryEditorExpression = new QueryEditorExpression();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for QueryEditorExpression");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "and":
            queryEditorExpression.queryEditorArrayExpression = mapper.convertValue(root, QueryEditorArrayExpression.class);
            break;
        case "function":
            queryEditorExpression.queryEditorFunctionExpression = mapper.convertValue(root, QueryEditorFunctionExpression.class);
            break;
        case "functionParameter":
            queryEditorExpression.queryEditorFunctionParameterExpression = mapper.convertValue(root, QueryEditorFunctionParameterExpression.class);
            break;
        case "groupBy":
            queryEditorExpression.queryEditorGroupByExpression = mapper.convertValue(root, QueryEditorGroupByExpression.class);
            break;
        case "operator":
            queryEditorExpression.queryEditorOperatorExpression = mapper.convertValue(root, QueryEditorOperatorExpression.class);
            break;
        case "or":
            queryEditorExpression.queryEditorArrayExpression = mapper.convertValue(root, QueryEditorArrayExpression.class);
            break;
        case "property":
            queryEditorExpression.queryEditorPropertyExpression = mapper.convertValue(root, QueryEditorPropertyExpression.class);
            break;
        }
        
        return queryEditorExpression;
    }
}
