// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class SubscriptionsQueryBuilder implements com.grafana.foundation.cog.Builder<SubscriptionsQuery> {
    protected final SubscriptionsQuery internal;
    
    public SubscriptionsQueryBuilder() {
        this.internal = new SubscriptionsQuery();
        this.internal.kind = "SubscriptionsQuery";
    }
    public SubscriptionsQueryBuilder rawQuery(String rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    public SubscriptionsQuery build() {
        return this.internal;
    }
}
