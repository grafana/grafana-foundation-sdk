// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class AzureResourceGraphQueryBuilder implements com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> {
    protected final AzureResourceGraphQuery internal;
    
    public AzureResourceGraphQueryBuilder() {
        this.internal = new AzureResourceGraphQuery();
    }
    public AzureResourceGraphQueryBuilder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public AzureResourceGraphQueryBuilder resultFormat(String resultFormat) {
    this.internal.resultFormat = resultFormat;
        return this;
    }
    public AzureResourceGraphQuery build() {
        return this.internal;
    }
}
