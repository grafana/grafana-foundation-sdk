// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class ResourceGraphQueryBuilder implements com.grafana.foundation.cog.Builder<ResourceGraphQuery> {
    protected final ResourceGraphQuery internal;
    
    public ResourceGraphQueryBuilder() {
        this.internal = new ResourceGraphQuery();
    }
    public ResourceGraphQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public ResourceGraphQueryBuilder resultFormat(String resultFormat) {
        this.internal.resultFormat = resultFormat;
        return this;
    }
    public ResourceGraphQuery build() {
        return this.internal;
    }
}
