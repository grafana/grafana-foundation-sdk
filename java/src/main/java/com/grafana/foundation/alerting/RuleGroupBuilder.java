// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;
import java.util.LinkedList;

public class RuleGroupBuilder implements com.grafana.foundation.cog.Builder<RuleGroup> {
    protected final RuleGroup internal;
    
    public RuleGroupBuilder(String title) {
        this.internal = new RuleGroup();
    this.internal.title = title;
    }
    public RuleGroupBuilder folderUid(String folderUid) {
    this.internal.folderUid = folderUid;
        return this;
    }
    
    public RuleGroupBuilder interval(Long interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public RuleGroupBuilder rules(com.grafana.foundation.cog.Builder<List<Rule>> rules) {
    this.internal.rules = rules.build();
        return this;
    }
    
    public RuleGroupBuilder withRule(com.grafana.foundation.cog.Builder<Rule> rule) {
		if (this.internal.rules == null) {
			this.internal.rules = new LinkedList<>();
		}
    this.internal.rules.add(rule.build());
        return this;
    }
    
    public RuleGroupBuilder title(String title) {
    this.internal.title = title;
        return this;
    }
    public RuleGroup build() {
        return this.internal;
    }
}
