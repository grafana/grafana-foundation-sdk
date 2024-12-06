// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.accesspolicy;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import java.util.LinkedList;

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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AccessPolicy> {
        protected final AccessPolicy internal;
        
        public Builder() {
            this.internal = new AccessPolicy();
        }
    public Builder scope(com.grafana.foundation.cog.Builder<ResourceRef> scope) {
    this.internal.scope = scope.build();
        return this;
    }
    
    public Builder role(com.grafana.foundation.cog.Builder<RoleRef> role) {
    this.internal.role = role.build();
        return this;
    }
    
    public Builder rules(com.grafana.foundation.cog.Builder<AccessRule> rule) {
		if (this.internal.rules == null) {
			this.internal.rules = new LinkedList<>();
		}
    this.internal.rules.add(rule.build());
        return this;
    }
    public AccessPolicy build() {
            return this.internal;
        }
    }
}
