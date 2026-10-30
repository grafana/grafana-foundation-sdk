// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = NotificationSettingsDeserializer.class)
public class NotificationSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected SimplifiedRouting simplifiedRouting;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected NamedRoutingTree namedRoutingTree;
    protected NotificationSettings() {}
    public static NotificationSettings createSimplifiedRouting(SimplifiedRouting simplifiedRouting) {
        NotificationSettings notificationSettings = new NotificationSettings();
        notificationSettings.simplifiedRouting = simplifiedRouting;
        return notificationSettings;
    }
    public static NotificationSettings createNamedRoutingTree(NamedRoutingTree namedRoutingTree) {
        NotificationSettings notificationSettings = new NotificationSettings();
        notificationSettings.namedRoutingTree = namedRoutingTree;
        return notificationSettings;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
