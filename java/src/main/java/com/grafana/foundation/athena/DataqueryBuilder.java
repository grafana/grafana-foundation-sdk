// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;

import com.grafana.foundation.dashboard.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder format(FormatOptions format) {
        this.internal.format = format;
        return this;
    }
    
    public DataqueryBuilder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs) {
        this.internal.connectionArgs = connectionArgs.build();
        return this;
    }
    
    public DataqueryBuilder table(String table) {
        this.internal.table = table;
        return this;
    }
    
    public DataqueryBuilder column(String column) {
        this.internal.column = column;
        return this;
    }
    
    public DataqueryBuilder queryID(String queryID) {
        this.internal.queryID = queryID;
        return this;
    }
    
    public DataqueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public DataqueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public DataqueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public DataqueryBuilder rawSQL(String rawSQL) {
        this.internal.rawSQL = rawSQL;
        return this;
    }
    
    public DataqueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
