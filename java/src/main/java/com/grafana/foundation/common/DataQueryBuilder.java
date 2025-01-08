// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class DataQueryBuilder implements com.grafana.foundation.cog.Builder<DataQuery> {
    protected final DataQuery internal;
    
    public DataQueryBuilder() {
        this.internal = new DataQuery();
    }
    public DataQueryBuilder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public DataQueryBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public DataQueryBuilder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public DataQueryBuilder datasource(Object datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    public DataQuery build() {
        return this.internal;
    }
}
