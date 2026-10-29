// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class NotificationSettingsDeserializer extends JsonDeserializer<NotificationSettings> {

    @Override
    public NotificationSettings deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        NotificationSettings notificationSettings = new NotificationSettings();
        if (!root.has("type")) {
            throw new IOException("Cannot find discriminator for NotificationSettings");
        }
        String discriminator = root.get("type").asText();  
        
        switch (discriminator) {
        case "NamedRoutingTree":
            notificationSettings.namedRoutingTree = mapper.convertValue(root, NamedRoutingTree.class);
            break;
        case "SimplifiedRouting":
            notificationSettings.simplifiedRouting = mapper.convertValue(root, SimplifiedRouting.class);
            break;
        }
        
        return notificationSettings;
    }
}
