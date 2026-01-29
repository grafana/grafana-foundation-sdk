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
import com.grafana.foundation.cog.variants.PanelConfig;
import java.util.Map;

public class VariableValueSingleDeserializer extends JsonDeserializer<VariableValueSingle> {

    @Override
    public VariableValueSingle deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        VariableValueSingle variableValueSingle = new VariableValueSingle();
        if (root.isTextual()) {
            variableValueSingle.string = mapper.convertValue(root, String.class);
        }
        else if (root.isBoolean()) {
            variableValueSingle.bool = mapper.convertValue(root, Boolean.class);
        }
        else if (root.isDouble()) {
            variableValueSingle.float64 = mapper.convertValue(root, Double.class);
        }
        else if (root.isObject() && couldBe(mapper, root, CustomVariableValue.class)) {
            variableValueSingle.customVariableValue = mapper.convertValue(root, CustomVariableValue.class);
        }
        
        return variableValueSingle;
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
