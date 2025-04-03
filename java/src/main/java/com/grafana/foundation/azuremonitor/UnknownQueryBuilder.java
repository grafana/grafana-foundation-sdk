// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class UnknownQueryBuilder implements com.grafana.foundation.cog.Builder<UnknownQuery> {
    protected final UnknownQuery internal;
    
    public UnknownQueryBuilder() {
        this.internal = new UnknownQuery();
        this.internal.kind = "UnknownQuery";
    }
    public UnknownQueryBuilder rawQuery(String rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    public UnknownQuery build() {
        return this.internal;
    }
}
