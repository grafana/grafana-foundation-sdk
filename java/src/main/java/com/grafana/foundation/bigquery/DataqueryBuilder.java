// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bigquery;

import com.grafana.foundation.dashboard.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder dataset(String dataset) {
        this.internal.dataset = dataset;
        return this;
    }
    
    public DataqueryBuilder table(String table) {
        this.internal.table = table;
        return this;
    }
    
    public DataqueryBuilder project(String project) {
        this.internal.project = project;
        return this;
    }
    
    public DataqueryBuilder format(QueryFormat format) {
        this.internal.format = format;
        return this;
    }
    
    public DataqueryBuilder rawQuery(Boolean rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    
    public DataqueryBuilder rawSql(String rawSql) {
        this.internal.rawSql = rawSql;
        return this;
    }
    
    public DataqueryBuilder location(String location) {
        this.internal.location = location;
        return this;
    }
    
    public DataqueryBuilder partitioned(Boolean partitioned) {
        this.internal.partitioned = partitioned;
        return this;
    }
    
    public DataqueryBuilder partitionedField(String partitionedField) {
        this.internal.partitionedField = partitionedField;
        return this;
    }
    
    public DataqueryBuilder convertToUTC(Boolean convertToUTC) {
        this.internal.convertToUTC = convertToUTC;
        return this;
    }
    
    public DataqueryBuilder sharded(Boolean sharded) {
        this.internal.sharded = sharded;
        return this;
    }
    
    public DataqueryBuilder queryPriority(QueryPriority queryPriority) {
        this.internal.queryPriority = queryPriority;
        return this;
    }
    
    public DataqueryBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public DataqueryBuilder editorMode(EditorMode editorMode) {
        this.internal.editorMode = editorMode;
        return this;
    }
    
    public DataqueryBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
    SQLExpression sqlResource = sql.build();
        this.internal.sql = sqlResource;
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
    
    public DataqueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
