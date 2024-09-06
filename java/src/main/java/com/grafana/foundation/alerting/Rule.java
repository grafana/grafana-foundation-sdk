// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import java.util.LinkedList;

public class Rule {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("annotations")
    public Map<String, String> annotations;
    @JsonProperty("condition")
    public String condition;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("data")
    public List<Query> data;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("execErrState")
    public RuleExecErrState execErrState;
    @JsonProperty("folderUID")
    public String folderUID;
    // The amount of time, in seconds, for which the rule must be breached for the rule to be considered to be Firing.
    // Before this time has elapsed, the rule is only considered to be Pending.
    @JsonProperty("for")
    public String forArg;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("id")
    public Long id;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("isPaused")
    public Boolean isPaused;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public Map<String, String> labels;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("noDataState")
    public RuleNoDataState noDataState;
    @JsonProperty("orgID")
    public Long orgID;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("provenance")
    public String provenance;
    @JsonProperty("ruleGroup")
    public String ruleGroup;
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("uid")
    public String uid;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("updated")
    public String updated;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Rule> {
        private final Rule internal;
        
        public Builder(String title) {
            this.internal = new Rule();
        if (!(title.length() >= 1)) {
            throw new IllegalArgumentException("title.length() must be >= 1");
        }
        if (!(title.length() <= 190)) {
            throw new IllegalArgumentException("title.length() must be <= 190");
        }
    this.internal.title = title;
        }
    public Builder annotations(Map<String, String> annotations) {
    this.internal.annotations = annotations;
        return this;
    }
    
    public Builder condition(String condition) {
    this.internal.condition = condition;
        return this;
    }
    
    public Builder queries(com.grafana.foundation.cog.Builder<List<Query>> data) {
    this.internal.data = data.build();
        return this;
    }
    
    public Builder withQuery(com.grafana.foundation.cog.Builder<Query> data) {
		if (this.internal.data == null) {
			this.internal.data = new LinkedList<>();
		}
    this.internal.data.add(data.build());
        return this;
    }
    
    public Builder execErrState(RuleExecErrState execErrState) {
    this.internal.execErrState = execErrState;
        return this;
    }
    
    public Builder folderUID(String folderUID) {
    this.internal.folderUID = folderUID;
        return this;
    }
    
    public Builder forArg(String forArg) {
    this.internal.forArg = forArg;
        return this;
    }
    
    public Builder id(Long id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder isPaused(Boolean isPaused) {
    this.internal.isPaused = isPaused;
        return this;
    }
    
    public Builder labels(Map<String, String> labels) {
    this.internal.labels = labels;
        return this;
    }
    
    public Builder noDataState(RuleNoDataState noDataState) {
    this.internal.noDataState = noDataState;
        return this;
    }
    
    public Builder orgID(Long orgID) {
    this.internal.orgID = orgID;
        return this;
    }
    
    public Builder provenance(String provenance) {
    this.internal.provenance = provenance;
        return this;
    }
    
    public Builder ruleGroup(String ruleGroup) {
        if (!(ruleGroup.length() >= 1)) {
            throw new IllegalArgumentException("ruleGroup.length() must be >= 1");
        }
        if (!(ruleGroup.length() <= 190)) {
            throw new IllegalArgumentException("ruleGroup.length() must be <= 190");
        }
    this.internal.ruleGroup = ruleGroup;
        return this;
    }
    
    public Builder title(String title) {
        if (!(title.length() >= 1)) {
            throw new IllegalArgumentException("title.length() must be >= 1");
        }
        if (!(title.length() <= 190)) {
            throw new IllegalArgumentException("title.length() must be <= 190");
        }
    this.internal.title = title;
        return this;
    }
    
    public Builder uid(String uid) {
        if (!(uid.length() >= 1)) {
            throw new IllegalArgumentException("uid.length() must be >= 1");
        }
        if (!(uid.length() <= 40)) {
            throw new IllegalArgumentException("uid.length() must be <= 40");
        }
    this.internal.uid = uid;
        return this;
    }
    
    public Builder updated(String updated) {
    this.internal.updated = updated;
        return this;
    }
    public Rule build() {
            return this.internal;
        }
    }
}
