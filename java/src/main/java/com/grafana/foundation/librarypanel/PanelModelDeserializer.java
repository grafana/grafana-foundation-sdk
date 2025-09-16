// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

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
import java.util.Map;
import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;


import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;
import java.util.Map;

public class PanelModelDeserializer extends JsonDeserializer<PanelModel> {
    
    @Override
    public PanelModel deserialize(JsonParser jp, DeserializationContext cxt) throws IOException, JsonProcessingException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        // Deserialise all the fields
        PanelModel panelModel = new PanelModel();
        if (root.has("type")) {
            panelModel.type = mapper.convertValue(root.get("type"), new TypeReference<>() {});
        }
        if (root.has("pluginVersion")) {
            panelModel.pluginVersion = mapper.convertValue(root.get("pluginVersion"), new TypeReference<>() {});
        }
        if (root.has("tags")) {
            panelModel.tags = mapper.convertValue(root.get("tags"), new TypeReference<>() {});
        }
        if (root.has("targets")) {
            panelModel.targets = mapper.convertValue(root.get("targets"), new TypeReference<>() {});
        }
        if (root.has("title")) {
            panelModel.title = mapper.convertValue(root.get("title"), new TypeReference<>() {});
        }
        if (root.has("description")) {
            panelModel.description = mapper.convertValue(root.get("description"), new TypeReference<>() {});
        }
        if (root.has("transparent")) {
            panelModel.transparent = mapper.convertValue(root.get("transparent"), new TypeReference<>() {});
        }
        if (root.has("datasource")) {
            panelModel.datasource = mapper.convertValue(root.get("datasource"), new TypeReference<>() {});
        }
        if (root.has("links")) {
            panelModel.links = mapper.convertValue(root.get("links"), new TypeReference<>() {});
        }
        if (root.has("repeat")) {
            panelModel.repeat = mapper.convertValue(root.get("repeat"), new TypeReference<>() {});
        }
        if (root.has("repeatDirection")) {
            panelModel.repeatDirection = mapper.convertValue(root.get("repeatDirection"), new TypeReference<>() {});
        }
        if (root.has("maxPerRow")) {
            panelModel.maxPerRow = mapper.convertValue(root.get("maxPerRow"), new TypeReference<>() {});
        }
        if (root.has("maxDataPoints")) {
            panelModel.maxDataPoints = mapper.convertValue(root.get("maxDataPoints"), new TypeReference<>() {});
        }
        if (root.has("transformations")) {
            panelModel.transformations = mapper.convertValue(root.get("transformations"), new TypeReference<>() {});
        }
        if (root.has("interval")) {
            panelModel.interval = mapper.convertValue(root.get("interval"), new TypeReference<>() {});
        }
        if (root.has("timeFrom")) {
            panelModel.timeFrom = mapper.convertValue(root.get("timeFrom"), new TypeReference<>() {});
        }
        if (root.has("timeShift")) {
            panelModel.timeShift = mapper.convertValue(root.get("timeShift"), new TypeReference<>() {});
        }
        if (root.has("hideTimeOverride")) {
            panelModel.hideTimeOverride = mapper.convertValue(root.get("hideTimeOverride"), new TypeReference<>() {});
        }
        if (root.has("options")) {
            panelModel.options = mapper.convertValue(root.get("options"), new TypeReference<>() {});
        }
        if (root.has("fieldConfig")) {
            panelModel.fieldConfig = mapper.convertValue(root.get("fieldConfig"), new TypeReference<>() {});
        }
       
       
       // Dataquery stuff
       String datasourceType = "";
       panelModel.datasource = mapper.treeToValue(root.get("datasource"), DataSourceRef.class);
       if (panelModel.datasource != null) {
            datasourceType = panelModel.datasource.type;
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
      panelModel.targets = targets;
        
       return panelModel;
    }
}
