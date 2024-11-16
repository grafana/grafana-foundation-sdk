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
import com.grafana.foundation.dashboard.DataSourceRef;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;
import java.util.Map;

public class LibrarypanelLibraryPanelModelDeserializer extends JsonDeserializer<LibrarypanelLibraryPanelModel> {
    
    @Override
    public LibrarypanelLibraryPanelModel deserialize(JsonParser jp, DeserializationContext cxt) throws IOException, JsonProcessingException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        // Deserialise all the fields
        LibrarypanelLibraryPanelModel librarypanelLibraryPanelModel = new LibrarypanelLibraryPanelModel();
        if (root.has("type")) {
            librarypanelLibraryPanelModel.type = mapper.convertValue(root.get("type"), new TypeReference<>() {});
        }
        if (root.has("pluginVersion")) {
            librarypanelLibraryPanelModel.pluginVersion = mapper.convertValue(root.get("pluginVersion"), new TypeReference<>() {});
        }
        if (root.has("tags")) {
            librarypanelLibraryPanelModel.tags = mapper.convertValue(root.get("tags"), new TypeReference<>() {});
        }
        if (root.has("targets")) {
            librarypanelLibraryPanelModel.targets = mapper.convertValue(root.get("targets"), new TypeReference<>() {});
        }
        if (root.has("title")) {
            librarypanelLibraryPanelModel.title = mapper.convertValue(root.get("title"), new TypeReference<>() {});
        }
        if (root.has("description")) {
            librarypanelLibraryPanelModel.description = mapper.convertValue(root.get("description"), new TypeReference<>() {});
        }
        if (root.has("transparent")) {
            librarypanelLibraryPanelModel.transparent = mapper.convertValue(root.get("transparent"), new TypeReference<>() {});
        }
        if (root.has("datasource")) {
            librarypanelLibraryPanelModel.datasource = mapper.convertValue(root.get("datasource"), new TypeReference<>() {});
        }
        if (root.has("links")) {
            librarypanelLibraryPanelModel.links = mapper.convertValue(root.get("links"), new TypeReference<>() {});
        }
        if (root.has("repeat")) {
            librarypanelLibraryPanelModel.repeat = mapper.convertValue(root.get("repeat"), new TypeReference<>() {});
        }
        if (root.has("repeatDirection")) {
            librarypanelLibraryPanelModel.repeatDirection = mapper.convertValue(root.get("repeatDirection"), new TypeReference<>() {});
        }
        if (root.has("maxPerRow")) {
            librarypanelLibraryPanelModel.maxPerRow = mapper.convertValue(root.get("maxPerRow"), new TypeReference<>() {});
        }
        if (root.has("maxDataPoints")) {
            librarypanelLibraryPanelModel.maxDataPoints = mapper.convertValue(root.get("maxDataPoints"), new TypeReference<>() {});
        }
        if (root.has("transformations")) {
            librarypanelLibraryPanelModel.transformations = mapper.convertValue(root.get("transformations"), new TypeReference<>() {});
        }
        if (root.has("interval")) {
            librarypanelLibraryPanelModel.interval = mapper.convertValue(root.get("interval"), new TypeReference<>() {});
        }
        if (root.has("timeFrom")) {
            librarypanelLibraryPanelModel.timeFrom = mapper.convertValue(root.get("timeFrom"), new TypeReference<>() {});
        }
        if (root.has("timeShift")) {
            librarypanelLibraryPanelModel.timeShift = mapper.convertValue(root.get("timeShift"), new TypeReference<>() {});
        }
        if (root.has("hideTimeOverride")) {
            librarypanelLibraryPanelModel.hideTimeOverride = mapper.convertValue(root.get("hideTimeOverride"), new TypeReference<>() {});
        }
        if (root.has("options")) {
            librarypanelLibraryPanelModel.options = mapper.convertValue(root.get("options"), new TypeReference<>() {});
        }
        if (root.has("fieldConfig")) {
            librarypanelLibraryPanelModel.fieldConfig = mapper.convertValue(root.get("fieldConfig"), new TypeReference<>() {});
        }
       
       
       // Dataquery stuff
       String datasourceType = "";
       librarypanelLibraryPanelModel.datasource = mapper.treeToValue(root.get("datasource"), DataSourceRef.class);
       if (librarypanelLibraryPanelModel.datasource != null) {
            datasourceType = librarypanelLibraryPanelModel.datasource.type;
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
      librarypanelLibraryPanelModel.targets = targets;
        
       return librarypanelLibraryPanelModel;
    }
}
