// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class NamedRoutingTree {
    @JsonProperty("type")
    public NotificationSettingsType type;
    @JsonProperty("routingTree")
    public String routingTree;
    public NamedRoutingTree() {
        this.type = NotificationSettingsType.NAMED_ROUTING_TREE;
        this.routingTree = "";
    }
    public NamedRoutingTree(String routingTree) {
        this.type = NotificationSettingsType.NAMED_ROUTING_TREE;
        this.routingTree = routingTree;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
