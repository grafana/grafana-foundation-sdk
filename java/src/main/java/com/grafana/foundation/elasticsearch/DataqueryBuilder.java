// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;
import com.grafana.foundation.common.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder alias(String alias) {
        this.internal.alias = alias;
        return this;
    }
    
    public DataqueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public DataqueryBuilder timeField(String timeField) {
        this.internal.timeField = timeField;
        return this;
    }
    
    public DataqueryBuilder bucketAggs(List<BucketAggregation> bucketAggs) {
        this.internal.bucketAggs = bucketAggs;
        return this;
    }
    
    public DataqueryBuilder metrics(List<MetricAggregation> metrics) {
        this.internal.metrics = metrics;
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
