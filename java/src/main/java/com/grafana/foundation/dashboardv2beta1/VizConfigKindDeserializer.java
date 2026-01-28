// Code generated - EDITING IS FUTILE. DO NOT EDIT.


package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import com.grafana.foundation.cog.variants.PanelConfig;
import com.grafana.foundation.cog.variants.Registry;

import java.io.IOException;
import java.util.Iterator;
import java.util.Map;

public class VizConfigKindDeserializer extends JsonDeserializer<VizConfigKind> {

    @Override
    public VizConfigKind deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        
        VizConfigKind vizConfigKind = new VizConfigKind();
        if (root.has("kind")) {
            vizConfigKind.kind = mapper.convertValue(root.get("kind"), new TypeReference<>() {});
        }
        else if (root.has("group")) {
            vizConfigKind.group = mapper.convertValue(root.get("group"), new TypeReference<>() {});
        }
        else if (root.has("version")) {
            vizConfigKind.version = mapper.convertValue(root.get("version"), new TypeReference<>() {});
        }
            
        VizConfigSpec spec = new VizConfigSpec();
        JsonNode specNode = root.get("spec");
        if (specNode != null && !specNode.isNull()) {
            PanelConfig panelConfig = Registry.getPanel(vizConfigKind.group);
            if (panelConfig == null) {
                throw new IllegalArgumentException("Unknown panel type: " + vizConfigKind.group);
            }
            if (specNode.has("options")) {
                    spec.options = mapper.treeToValue(root.get("options"), panelConfig.getOptionsClass());
            }
            else if (specNode.has("fieldConfig")) {
                    FieldConfigSource fieldConfig = mapper.treeToValue(specNode.get("fieldConfig"), FieldConfigSource.class);
                    if (fieldConfig != null && fieldConfig.defaults != null) {
                        JsonNode customNode = specNode.get("fieldConfig").get("defaults").get("custom");
                        Class<?> customClass = panelConfig.getFieldConfigClass();
                        if (customNode != null && customClass != null) {
                            fieldConfig.defaults.custom = mapper.treeToValue(customNode, panelConfig.getFieldConfigClass());
                        }
                        spec.fieldConfig = fieldConfig;
                    }
            }
        }
        vizConfigKind.spec = spec;
        
        return vizConfigKind;
    }
}