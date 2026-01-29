// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class VariableKindDeserializer extends JsonDeserializer<VariableKind> {

    @Override
    public VariableKind deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        VariableKind variableKind = new VariableKind();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for VariableKind");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "AdhocVariable":
            variableKind.adhocVariableKind = mapper.convertValue(root, AdhocVariableKind.class);
            break;
        case "ConstantVariable":
            variableKind.constantVariableKind = mapper.convertValue(root, ConstantVariableKind.class);
            break;
        case "CustomVariable":
            variableKind.customVariableKind = mapper.convertValue(root, CustomVariableKind.class);
            break;
        case "DatasourceVariable":
            variableKind.datasourceVariableKind = mapper.convertValue(root, DatasourceVariableKind.class);
            break;
        case "GroupByVariable":
            variableKind.groupByVariableKind = mapper.convertValue(root, GroupByVariableKind.class);
            break;
        case "IntervalVariable":
            variableKind.intervalVariableKind = mapper.convertValue(root, IntervalVariableKind.class);
            break;
        case "QueryVariable":
            variableKind.queryVariableKind = mapper.convertValue(root, QueryVariableKind.class);
            break;
        case "SwitchVariable":
            variableKind.switchVariableKind = mapper.convertValue(root, SwitchVariableKind.class);
            break;
        case "TextVariable":
            variableKind.textVariableKind = mapper.convertValue(root, TextVariableKind.class);
            break;
        }
        
        return variableKind;
    }
}
