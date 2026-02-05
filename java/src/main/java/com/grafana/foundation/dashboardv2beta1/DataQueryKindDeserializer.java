// Code generated - EDITING IS FUTILE. DO NOT EDIT.


package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import com.grafana.foundation.cog.variants.UnknownDataquery;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.cog.variants.Registry;

import java.io.IOException;
import java.util.Iterator;
import java.util.Map;

public class DataQueryKindDeserializer extends JsonDeserializer<DataQueryKind> {
    
    @Override
    public DataQueryKind deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        DataQueryKind dataQueryKind = new DataQueryKind();
        if (root.has("kind")) {
            dataQueryKind.kind = mapper.convertValue(root.get("kind"), new TypeReference<>() {});
        }
        else if (root.has("group")) {
            dataQueryKind.group = mapper.convertValue(root.get("group"), new TypeReference<>() {});
        }
        else if (root.has("version")) {
            dataQueryKind.version = mapper.convertValue(root.get("version"), new TypeReference<>() {});
        }
        else if (root.has("datasource")) {
            dataQueryKind.datasource = mapper.convertValue(root.get("datasource"), new TypeReference<>() {});
        }
        
        if (dataQueryKind.group != null && !dataQueryKind.group.trim().isEmpty()) {
            Class<? extends Dataquery> clazz = Registry.getDataquery(dataQueryKind.group);
            if (clazz != null) {
                dataQueryKind.spec = mapper.treeToValue(root.get("spec"), clazz);
            } else {
                UnknownDataquery unknownDataquery = new UnknownDataquery();
                Iterator<Map.Entry<String, JsonNode>> fieldsIterator = root.get("spec").fields();
                while (fieldsIterator.hasNext()) {
                  Map.Entry<String, JsonNode> field = fieldsIterator.next();
                  unknownDataquery.genericFields.put(field.getKey(), mapper.treeToValue(field.getValue(), Object.class));
                }
                dataQueryKind.spec = unknownDataquery;
            }
        }
        
        return dataQueryKind;
    }
}