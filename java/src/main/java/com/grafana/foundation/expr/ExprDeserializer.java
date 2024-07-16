// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class ExprDeserializer extends JsonDeserializer<Expr> {

    @Override
    public Expr deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        Expr expr = new Expr();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for Expr");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "classic_conditions":
            expr.typeClassicConditions = mapper.convertValue(root, TypeClassicConditions.class);
            break;
        case "math":
            expr.typeMath = mapper.convertValue(root, TypeMath.class);
            break;
        case "reduce":
            expr.typeReduce = mapper.convertValue(root, TypeReduce.class);
            break;
        case "resample":
            expr.typeResample = mapper.convertValue(root, TypeResample.class);
            break;
        case "sql":
            expr.typeSql = mapper.convertValue(root, TypeSql.class);
            break;
        case "threshold":
            expr.typeThreshold = mapper.convertValue(root, TypeThreshold.class);
            break;
        }
        
        return expr;
    }
}
