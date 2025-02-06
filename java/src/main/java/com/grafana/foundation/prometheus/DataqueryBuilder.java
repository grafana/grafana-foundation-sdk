// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.grafana.foundation.dashboard.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder expr(String expr) {
        this.internal.expr = expr;
        return this;
    }
    
    public DataqueryBuilder instant() {
        this.internal.instant = true;
        this.internal.range = false;
        return this;
    }
    
    public DataqueryBuilder range() {
        this.internal.range = true;
        this.internal.instant = false;
        return this;
    }
    
    public DataqueryBuilder exemplar(Boolean exemplar) {
        this.internal.exemplar = exemplar;
        return this;
    }
    
    public DataqueryBuilder editorMode(QueryEditorMode editorMode) {
        this.internal.editorMode = editorMode;
        return this;
    }
    
    public DataqueryBuilder format(PromQueryFormat format) {
        this.internal.format = format;
        return this;
    }
    
    public DataqueryBuilder legendFormat(String legendFormat) {
        this.internal.legendFormat = legendFormat;
        return this;
    }
    
    public DataqueryBuilder intervalFactor(Double intervalFactor) {
        this.internal.intervalFactor = intervalFactor;
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
    
    public DataqueryBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public DataqueryBuilder rangeAndInstant() {
        this.internal.range = true;
        this.internal.instant = true;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
