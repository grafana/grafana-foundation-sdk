// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class WorkspacesQueryBuilder implements com.grafana.foundation.cog.Builder<WorkspacesQuery> {
    protected final WorkspacesQuery internal;
    
    public WorkspacesQueryBuilder() {
        this.internal = new WorkspacesQuery();
        this.internal.kind = "WorkspacesQuery";
    }
    public WorkspacesQueryBuilder rawQuery(String rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public WorkspacesQueryBuilder subscription(String subscription) {
        this.internal.subscription = subscription;
        return this;
    }
    public WorkspacesQuery build() {
        return this.internal;
    }
}
