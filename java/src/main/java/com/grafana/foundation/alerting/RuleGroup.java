// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;
import java.util.LinkedList;

public class RuleGroup {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("folderUid")
    public String folderUid;
    // The interval, in seconds, at which all rules in the group are evaluated.
    // If a group contains many rules, the rules are evaluated sequentially.
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("interval")
    public Long interval;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("rules")
    public List<Rule> rules;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("title")
    public String title;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RuleGroup> {
        private final RuleGroup internal;
        
        public Builder(String title) {
            this.internal = new RuleGroup();
    this.internal.title = title;
        }
    public Builder folderUid(String folderUid) {
    this.internal.folderUid = folderUid;
        return this;
    }
    
    public Builder interval(Long interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public Builder rules(com.grafana.foundation.cog.Builder<List<Rule>> rules) {
    this.internal.rules = rules.build();
        return this;
    }
    
    public Builder withRule(com.grafana.foundation.cog.Builder<Rule> rules) {
		if (this.internal.rules == null) {
			this.internal.rules = new LinkedList<>();
		}
    this.internal.rules.add(rules.build());
        return this;
    }
    
    public Builder title(String title) {
    this.internal.title = title;
        return this;
    }
    public RuleGroup build() {
            return this.internal;
        }
    }
}
