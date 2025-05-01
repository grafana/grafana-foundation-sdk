// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class AccessRule {
    // The kind this rule applies to (dashboards, alert, etc)
    @JsonProperty("kind")
    public String kind;
    // READ, WRITE, CREATE, DELETE, ...
    // should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
    @JsonProperty("verb")
    public String verb;
    // Specific sub-elements like "alert.rules" or "dashboard.permissions"????
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("target")
    public String target;
    public AccessRule() {
        this.kind = "*";
        this.verb = "";
    }
    public AccessRule(String kind,String verb,String target) {
        this.kind = kind;
        this.verb = verb;
        this.target = target;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
