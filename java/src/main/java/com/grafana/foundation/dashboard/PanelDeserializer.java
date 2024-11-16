// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.grafana.foundation.cog.variants.UnknownDataquery;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.cog.variants.PanelConfig;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;
import java.util.Map;

public class PanelDeserializer extends JsonDeserializer<Panel> {
    
    @Override
    public Panel deserialize(JsonParser jp, DeserializationContext cxt) throws IOException, JsonProcessingException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        // Deserialise all the fields
        Panel panel = new Panel();
        if (root.has("type")) {
            panel.type = mapper.convertValue(root.get("type"), new TypeReference<>() {});
        }
        if (root.has("id")) {
            panel.id = mapper.convertValue(root.get("id"), new TypeReference<>() {});
        }
        if (root.has("pluginVersion")) {
            panel.pluginVersion = mapper.convertValue(root.get("pluginVersion"), new TypeReference<>() {});
        }
        if (root.has("tags")) {
            panel.tags = mapper.convertValue(root.get("tags"), new TypeReference<>() {});
        }
        if (root.has("title")) {
            panel.title = mapper.convertValue(root.get("title"), new TypeReference<>() {});
        }
        if (root.has("description")) {
            panel.description = mapper.convertValue(root.get("description"), new TypeReference<>() {});
        }
        if (root.has("transparent")) {
            panel.transparent = mapper.convertValue(root.get("transparent"), new TypeReference<>() {});
        }
        if (root.has("datasource")) {
            panel.datasource = mapper.convertValue(root.get("datasource"), new TypeReference<>() {});
        }
        if (root.has("gridPos")) {
            panel.gridPos = mapper.convertValue(root.get("gridPos"), new TypeReference<>() {});
        }
        if (root.has("links")) {
            panel.links = mapper.convertValue(root.get("links"), new TypeReference<>() {});
        }
        if (root.has("repeat")) {
            panel.repeat = mapper.convertValue(root.get("repeat"), new TypeReference<>() {});
        }
        if (root.has("repeatDirection")) {
            panel.repeatDirection = mapper.convertValue(root.get("repeatDirection"), new TypeReference<>() {});
        }
        if (root.has("repeatPanelId")) {
            panel.repeatPanelId = mapper.convertValue(root.get("repeatPanelId"), new TypeReference<>() {});
        }
        if (root.has("maxDataPoints")) {
            panel.maxDataPoints = mapper.convertValue(root.get("maxDataPoints"), new TypeReference<>() {});
        }
        if (root.has("transformations")) {
            panel.transformations = mapper.convertValue(root.get("transformations"), new TypeReference<>() {});
        }
        if (root.has("interval")) {
            panel.interval = mapper.convertValue(root.get("interval"), new TypeReference<>() {});
        }
        if (root.has("timeFrom")) {
            panel.timeFrom = mapper.convertValue(root.get("timeFrom"), new TypeReference<>() {});
        }
        if (root.has("timeShift")) {
            panel.timeShift = mapper.convertValue(root.get("timeShift"), new TypeReference<>() {});
        }
        if (root.has("libraryPanel")) {
            panel.libraryPanel = mapper.convertValue(root.get("libraryPanel"), new TypeReference<>() {});
        }
        if (root.has("options")) {
            panel.options = mapper.convertValue(root.get("options"), new TypeReference<>() {});
        }
        if (root.has("fieldConfig")) {
            panel.fieldConfig = mapper.convertValue(root.get("fieldConfig"), new TypeReference<>() {});
        }
        // Deserialise panels
        PanelConfig config = Registry.getPanel(panel.type);
        if (config != null) {
            panel.options = mapper.treeToValue(root.get("options"), config.getOptionsClass());
            
            FieldConfigSource fieldConfigSource = mapper.treeToValue(root.get("fieldConfig"), FieldConfigSource.class);
            if (fieldConfigSource != null) {
                FieldConfig fieldConfig = fieldConfigSource.defaults;
                if (fieldConfig != null) {
                    JsonNode customNode = root.get("fieldConfig").get("defaults").get("custom");
                    if (customNode != null) {
                        Class<?> customClass = config.getFieldConfigClass();
                        if (customClass != null) {
                            fieldConfig.custom = mapper.treeToValue(customNode, customClass);
                        }
                    }
                    panel.fieldConfig = fieldConfigSource;
                }
            }
        } else {
            throw new IllegalArgumentException("Unknown panel type: " + panel.type);
        }
       
       
       // Dataquery stuff
       String datasourceType = "";
       panel.datasource = mapper.treeToValue(root.get("datasource"), DataSourceRef.class);
       if (panel.datasource != null) {
            datasourceType = panel.datasource.type;
       }
       List<Dataquery> targets = new ArrayList<>();
       if (root.has("targets")) {
           for (JsonNode node : root.get("targets")) {
                Class<? extends Dataquery> clazz = Registry.getDataquery(datasourceType);
                if (clazz != null) {
                    Dataquery dataquery = mapper.treeToValue(node, clazz);
                    targets.add(dataquery);
                } else {
                  UnknownDataquery unknownDataquery = new UnknownDataquery();
                  Iterator<Map.Entry<String, JsonNode>> fieldsIterator = node.fields();
                  while (fieldsIterator.hasNext()) {
                      Map.Entry<String, JsonNode> field = fieldsIterator.next();
                      unknownDataquery.genericFields.put(field.getKey(), mapper.treeToValue(field.getValue(), Object.class));
                  }
                  targets.add(unknownDataquery);
                }
          }
      }
      panel.targets = targets;
        
       return panel;
    }
}
