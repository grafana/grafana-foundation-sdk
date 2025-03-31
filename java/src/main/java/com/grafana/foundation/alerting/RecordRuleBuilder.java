// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class RecordRuleBuilder implements com.grafana.foundation.cog.Builder<RecordRule> {
    protected final RecordRule internal;
    
    public RecordRuleBuilder() {
        this.internal = new RecordRule();
    }
    public RecordRuleBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public RecordRuleBuilder metric(String metric) {
        this.internal.metric = metric;
        return this;
    }
    
    public RecordRuleBuilder targetDatasourceUid(String targetDatasourceUid) {
        this.internal.targetDatasourceUid = targetDatasourceUid;
        return this;
    }
    public RecordRule build() {
        return this.internal;
    }
}
