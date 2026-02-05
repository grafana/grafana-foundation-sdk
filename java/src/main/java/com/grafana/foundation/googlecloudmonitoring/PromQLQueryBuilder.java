// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;


public class PromQLQueryBuilder implements com.grafana.foundation.cog.Builder<PromQLQuery> {
    protected final PromQLQuery internal;
    
    public PromQLQueryBuilder() {
        this.internal = new PromQLQuery();
    }
    public PromQLQueryBuilder projectName(String projectName) {
        this.internal.projectName = projectName;
        return this;
    }
    
    public PromQLQueryBuilder expr(String expr) {
        this.internal.expr = expr;
        return this;
    }
    
    public PromQLQueryBuilder step(String step) {
        this.internal.step = step;
        return this;
    }
    public PromQLQuery build() {
        return this.internal;
    }
}
