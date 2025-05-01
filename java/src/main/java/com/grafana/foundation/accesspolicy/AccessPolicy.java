// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class AccessPolicy {
    // The scope where these policies should apply
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("scope")
    public ResourceRef scope;
    // The role that must apply this policy
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("role")
    public RoleRef role;
    // The set of rules to apply.  Note that * is required to modify
    // access policy rules, and that "none" will reject all actions
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("rules")
    public List<AccessRule> rules;
    public AccessPolicy() {
        this.scope = new com.grafana.foundation.accesspolicy.ResourceRefBuilder().build();
        this.role = new com.grafana.foundation.accesspolicy.RoleRefBuilder().build();
        this.rules = new LinkedList<>();
    }
    public AccessPolicy(ResourceRef scope,RoleRef role,List<AccessRule> rules) {
        this.scope = scope;
        this.role = role;
        this.rules = rules;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
