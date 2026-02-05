// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.Map;
import java.util.List;
import java.util.LinkedList;

public class RuleBuilder implements com.grafana.foundation.cog.Builder<Rule> {
    protected final Rule internal;
    
    public RuleBuilder(String title) {
        this.internal = new Rule();
        if (!(title.length() >= 1)) {
            throw new IllegalArgumentException("title.length() must be >= 1");
        }
        if (!(title.length() <= 190)) {
            throw new IllegalArgumentException("title.length() must be <= 190");
        }
        this.internal.title = title;
    }
    public RuleBuilder annotations(Map<String, String> annotations) {
        this.internal.annotations = annotations;
        return this;
    }
    
    public RuleBuilder condition(String condition) {
        this.internal.condition = condition;
        return this;
    }
    
    public RuleBuilder queries(List<com.grafana.foundation.cog.Builder<Query>> data) {
        List<Query> dataResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Query> r1 : data) {
                Query dataDepth1 = r1.build();
                dataResources.add(dataDepth1); 
        }
        this.internal.data = dataResources;
        return this;
    }
    
    public RuleBuilder withQuery(com.grafana.foundation.cog.Builder<Query> data) {
		if (this.internal.data == null) {
			this.internal.data = new LinkedList<>();
		}
    Query dataResource = data.build();
        this.internal.data.add(dataResource);
        return this;
    }
    
    public RuleBuilder execErrState(RuleExecErrState execErrState) {
        this.internal.execErrState = execErrState;
        return this;
    }
    
    public RuleBuilder folderUID(String folderUID) {
        this.internal.folderUID = folderUID;
        return this;
    }
    
    public RuleBuilder forArg(String forArg) {
        this.internal.forArg = forArg;
        return this;
    }
    
    public RuleBuilder id(Long id) {
        this.internal.id = id;
        return this;
    }
    
    public RuleBuilder isPaused(Boolean isPaused) {
        this.internal.isPaused = isPaused;
        return this;
    }
    
    public RuleBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public RuleBuilder noDataState(RuleNoDataState noDataState) {
        this.internal.noDataState = noDataState;
        return this;
    }
    
    public RuleBuilder notificationSettings(com.grafana.foundation.cog.Builder<NotificationSettings> notificationSettings) {
    NotificationSettings notificationSettingsResource = notificationSettings.build();
        this.internal.notificationSettings = notificationSettingsResource;
        return this;
    }
    
    public RuleBuilder orgID(Long orgID) {
        this.internal.orgID = orgID;
        return this;
    }
    
    public RuleBuilder provenance(String provenance) {
        this.internal.provenance = provenance;
        return this;
    }
    
    public RuleBuilder record(com.grafana.foundation.cog.Builder<RecordRule> record) {
    RecordRule recordResource = record.build();
        this.internal.record = recordResource;
        return this;
    }
    
    public RuleBuilder ruleGroup(String ruleGroup) {
        if (!(ruleGroup.length() >= 1)) {
            throw new IllegalArgumentException("ruleGroup.length() must be >= 1");
        }
        if (!(ruleGroup.length() <= 190)) {
            throw new IllegalArgumentException("ruleGroup.length() must be <= 190");
        }
        this.internal.ruleGroup = ruleGroup;
        return this;
    }
    
    public RuleBuilder title(String title) {
        if (!(title.length() >= 1)) {
            throw new IllegalArgumentException("title.length() must be >= 1");
        }
        if (!(title.length() <= 190)) {
            throw new IllegalArgumentException("title.length() must be <= 190");
        }
        this.internal.title = title;
        return this;
    }
    
    public RuleBuilder uid(String uid) {
        if (!(uid.length() >= 1)) {
            throw new IllegalArgumentException("uid.length() must be >= 1");
        }
        if (!(uid.length() <= 40)) {
            throw new IllegalArgumentException("uid.length() must be <= 40");
        }
        this.internal.uid = uid;
        return this;
    }
    
    public RuleBuilder updated(String updated) {
        this.internal.updated = updated;
        return this;
    }
    public Rule build() {
        return this.internal;
    }
}
