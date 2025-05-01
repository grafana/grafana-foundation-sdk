// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.grafana.foundation.cog.variants.Dataquery;

public class QueryBuilder implements com.grafana.foundation.cog.Builder<Query> {
    protected final Query internal;
    
    public QueryBuilder(String refId) {
        this.internal = new Query();
        this.internal.refId = refId;
    }
    public QueryBuilder datasourceUid(String datasourceUid) {
        this.internal.datasourceUid = datasourceUid;
        return this;
    }
    
    public QueryBuilder model(com.grafana.foundation.cog.Builder<Dataquery> model) {
    Dataquery modelResource = model.build();
        this.internal.model = modelResource;
        return this;
    }
    
    public QueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public QueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public QueryBuilder relativeTimeRange(RelativeTimeRange relativeTimeRange) {
        this.internal.relativeTimeRange = relativeTimeRange;
        return this;
    }
    public Query build() {
        return this.internal;
    }
}
