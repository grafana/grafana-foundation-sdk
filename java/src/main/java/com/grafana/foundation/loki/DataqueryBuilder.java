// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.grafana.foundation.common.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder expr(String expr) {
        this.internal.expr = expr;
        return this;
    }
    
    public DataqueryBuilder legendFormat(String legendFormat) {
        this.internal.legendFormat = legendFormat;
        return this;
    }
    
    public DataqueryBuilder maxLines(Long maxLines) {
        this.internal.maxLines = maxLines;
        return this;
    }
    
    public DataqueryBuilder resolution(Long resolution) {
        this.internal.resolution = resolution;
        return this;
    }
    
    public DataqueryBuilder editorMode(QueryEditorMode editorMode) {
        this.internal.editorMode = editorMode;
        return this;
    }
    
    public DataqueryBuilder range(Boolean range) {
        this.internal.range = range;
        return this;
    }
    
    public DataqueryBuilder instant(Boolean instant) {
        this.internal.instant = instant;
        return this;
    }
    
    public DataqueryBuilder step(String step) {
        this.internal.step = step;
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
    
    public DataqueryBuilder direction(LokiQueryDirection direction) {
        this.internal.direction = direction;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
