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
import com.grafana.foundation.common.DataSourceRef;


import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;
import java.util.Map;

public class AnnotationQueryDeserializer extends JsonDeserializer<AnnotationQuery> {
    
    @Override
    public AnnotationQuery deserialize(JsonParser jp, DeserializationContext cxt) throws IOException, JsonProcessingException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        // Deserialise all the fields
        AnnotationQuery annotationQuery = new AnnotationQuery();
        if (root.has("name")) {
            annotationQuery.name = mapper.convertValue(root.get("name"), new TypeReference<>() {});
        }
        if (root.has("datasource")) {
            annotationQuery.datasource = mapper.convertValue(root.get("datasource"), new TypeReference<>() {});
        }
        if (root.has("enable")) {
            annotationQuery.enable = mapper.convertValue(root.get("enable"), new TypeReference<>() {});
        }
        if (root.has("hide")) {
            annotationQuery.hide = mapper.convertValue(root.get("hide"), new TypeReference<>() {});
        }
        if (root.has("iconColor")) {
            annotationQuery.iconColor = mapper.convertValue(root.get("iconColor"), new TypeReference<>() {});
        }
        if (root.has("filter")) {
            annotationQuery.filter = mapper.convertValue(root.get("filter"), new TypeReference<>() {});
        }
        if (root.has("target")) {
            annotationQuery.target = mapper.convertValue(root.get("target"), new TypeReference<>() {});
        }
        if (root.has("type")) {
            annotationQuery.type = mapper.convertValue(root.get("type"), new TypeReference<>() {});
        }
        if (root.has("builtIn")) {
            annotationQuery.builtIn = mapper.convertValue(root.get("builtIn"), new TypeReference<>() {});
        }
        if (root.has("expr")) {
            annotationQuery.expr = mapper.convertValue(root.get("expr"), new TypeReference<>() {});
        }
       
       
       // Dataquery stuff
       String datasourceType = "";
       annotationQuery.datasource = mapper.treeToValue(root.get("datasource"), DataSourceRef.class);
       if (annotationQuery.datasource != null) {
            datasourceType = annotationQuery.datasource.type;
       }
      Class<? extends Dataquery> clazz = Registry.getDataquery(datasourceType);
      if (clazz != null) {
          annotationQuery.target = mapper.treeToValue(root.get("target"), clazz);
      } else {
          UnknownDataquery unknownDataquery = new UnknownDataquery();
          Iterator<Map.Entry<String, JsonNode>> fieldsIterator = root.get("target").fields();
          while (fieldsIterator.hasNext()) {
              Map.Entry<String, JsonNode> field = fieldsIterator.next();
              unknownDataquery.genericFields.put(field.getKey(), mapper.treeToValue(field.getValue(), Object.class));
          }
          annotationQuery.target = unknownDataquery;
      }
        
       return annotationQuery;
    }
}
