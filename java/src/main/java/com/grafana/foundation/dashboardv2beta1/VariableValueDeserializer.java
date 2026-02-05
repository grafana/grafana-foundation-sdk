// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.common.DataSourceRef;
import java.util.List;

public class VariableValueDeserializer extends JsonDeserializer<VariableValue> {

    @Override
    public VariableValue deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        VariableValue variableValue = new VariableValue();
        if (root.isTextual()) {
            variableValue.string = mapper.convertValue(root, String.class);
        }
        else if (root.isBoolean()) {
            variableValue.bool = mapper.convertValue(root, Boolean.class);
        }
        else if (root.isDouble()) {
            variableValue.float64 = mapper.convertValue(root, Double.class);
        }
        else if (root.isObject() && couldBe(mapper, root, CustomVariableValue.class)) {
            variableValue.customVariableValue = mapper.convertValue(root, CustomVariableValue.class);
        }
        else if (root.isArray()) {
            variableValue.arrayOfVariableValueSingle = mapper.convertValue(root, new TypeReference<List<VariableValueSingle>>() {});
        }
        
        return variableValue;
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
