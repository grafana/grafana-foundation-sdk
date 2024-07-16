// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.rolebinding;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class BuiltinRoleRefOrCustomRoleRefDeserializer extends JsonDeserializer<BuiltinRoleRefOrCustomRoleRef> {

    @Override
    public BuiltinRoleRefOrCustomRoleRef deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        BuiltinRoleRefOrCustomRoleRef builtinRoleRefOrCustomRoleRef = new BuiltinRoleRefOrCustomRoleRef();
        if (!root.has("kind")) {
            throw new IOException("Cannot find discriminator for BuiltinRoleRefOrCustomRoleRef");
        }
        String discriminator = root.get("kind").asText();  
        
        switch (discriminator) {
        case "BuiltinRole":
            builtinRoleRefOrCustomRoleRef.builtinRoleRef = mapper.convertValue(root, BuiltinRoleRef.class);
            break;
        case "Role":
            builtinRoleRefOrCustomRoleRef.customRoleRef = mapper.convertValue(root, CustomRoleRef.class);
            break;
        }
        
        return builtinRoleRefOrCustomRoleRef;
    }
}
