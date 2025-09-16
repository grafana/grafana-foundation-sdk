// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.cog.variants.PanelConfig;
import java.util.Map;

public class BuilderQueryEditorOperatorTypeDeserializer extends JsonDeserializer<BuilderQueryEditorOperatorType> {

    @Override
    public BuilderQueryEditorOperatorType deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        BuilderQueryEditorOperatorType builderQueryEditorOperatorType = new BuilderQueryEditorOperatorType();
        if (root.isTextual()) {
            builderQueryEditorOperatorType.string = mapper.convertValue(root, String.class);
        }
        else if (root.isBoolean()) {
            builderQueryEditorOperatorType.bool = mapper.convertValue(root, Boolean.class);
        }
        else if (root.isDouble()) {
            builderQueryEditorOperatorType.float64 = mapper.convertValue(root, Double.class);
        }
        else if (root.isObject() && couldBe(mapper, root, SelectableValue.class)) {
            builderQueryEditorOperatorType.selectableValue = mapper.convertValue(root, SelectableValue.class);
        }
        
        return builderQueryEditorOperatorType;
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
