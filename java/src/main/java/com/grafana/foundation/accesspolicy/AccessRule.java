// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class AccessRule {
    // The kind this rule applies to (dashboards, alert, etc)
    @JsonProperty("kind")
    public String kind;
    // READ, WRITE, CREATE, DELETE, ...
    // should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
    @JsonProperty("verb")
    public String verb;
    // Specific sub-elements like "alert.rules" or "dashboard.permissions"????
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("target")
    public String target;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AccessRule> {
        private final AccessRule internal;
        
        public Builder() {
            this.internal = new AccessRule();
        this.kind("*");
        }
    public Builder kind(String kind) {
    this.internal.kind = kind;
        return this;
    }
    
    public Builder verb(String verb) {
    this.internal.verb = verb;
        return this;
    }
    
    public Builder target(String target) {
    this.internal.target = target;
        return this;
    }
    public AccessRule build() {
            return this.internal;
        }
    }
}
