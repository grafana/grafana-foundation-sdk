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
    
    public RuleGroupBuilder rules(List<com.grafana.foundation.cog.Builder<Rule>> rules) {
        List<Rule> rulesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Rule> r1 : rules) {
                Rule rulesDepth1 = r1.build();
                rulesResources.add(rulesDepth1); 
        }
        this.internal.rules = rulesResources;
        return this;
    }
    
    public RuleGroupBuilder withRule(com.grafana.foundation.cog.Builder<Rule> rule) {
		if (this.internal.rules == null) {
			this.internal.rules = new LinkedList<>();
		}
    Rule ruleResource = rule.build();
        this.internal.rules.add(ruleResource);
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
