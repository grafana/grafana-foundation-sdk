// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.grafana.foundation.cog.variants.UnknownDataquery;
import com.grafana.foundation.cog.variants.Registry;

import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;
import com.grafana.foundation.cog.variants.PanelConfig;
import java.util.Map;
import com.grafana.foundation.dashboard.DataSourceRef;


import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Iterator;
import java.util.Map;

public class QueryDeserializer extends JsonDeserializer<Query> {
    
    @Override
    public Query deserialize(JsonParser jp, DeserializationContext cxt) throws IOException, JsonProcessingException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        // Deserialise all the fields
        Query query = new Query();
        if (root.has("datasourceUid")) {
            query.datasourceUid = mapper.convertValue(root.get("datasourceUid"), new TypeReference<>() {});
        }
        if (root.has("model")) {
            query.model = mapper.convertValue(root.get("model"), new TypeReference<>() {});
        }
        if (root.has("queryType")) {
            query.queryType = mapper.convertValue(root.get("queryType"), new TypeReference<>() {});
        }
        if (root.has("refId")) {
            query.refId = mapper.convertValue(root.get("refId"), new TypeReference<>() {});
        }
        if (root.has("relativeTimeRange")) {
            query.relativeTimeRange = mapper.convertValue(root.get("relativeTimeRange"), new TypeReference<>() {});
        }
       
       
       // Dataquery stuff
       String datasourceType = "";
      Class<? extends Dataquery> clazz = Registry.getDataquery(datasourceType);
      if (clazz != null) {
          query.model = mapper.treeToValue(root.get("model"), clazz);
      } else {
          UnknownDataquery unknownDataquery = new UnknownDataquery();
          Iterator<Map.Entry<String, JsonNode>> fieldsIterator = root.get("model").fields();
          while (fieldsIterator.hasNext()) {
              Map.Entry<String, JsonNode> field = fieldsIterator.next();
              unknownDataquery.genericFields.put(field.getKey(), mapper.treeToValue(field.getValue(), Object.class));
          }
          query.model = unknownDataquery;
      }
        
       return query;
    }
}
