// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class ResourceGroupsQueryBuilder implements com.grafana.foundation.cog.Builder<ResourceGroupsQuery> {
    protected final ResourceGroupsQuery internal;
    
    public ResourceGroupsQueryBuilder() {
        this.internal = new ResourceGroupsQuery();
    this.internal.kind = "ResourceGroupsQuery";
    }
    public ResourceGroupsQueryBuilder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public ResourceGroupsQueryBuilder subscription(String subscription) {
    this.internal.subscription = subscription;
        return this;
    }
    public ResourceGroupsQuery build() {
        return this.internal;
    }
}
