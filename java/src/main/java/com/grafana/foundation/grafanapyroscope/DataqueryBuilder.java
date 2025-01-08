// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.grafanapyroscope;

import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder labelSelector(String labelSelector) {
    this.internal.labelSelector = labelSelector;
        return this;
    }
    
    public DataqueryBuilder spanSelector(List<String> spanSelector) {
    this.internal.spanSelector = spanSelector;
        return this;
    }
    
    public DataqueryBuilder profileTypeId(String profileTypeId) {
    this.internal.profileTypeId = profileTypeId;
        return this;
    }
    
    public DataqueryBuilder groupBy(List<String> groupBy) {
    this.internal.groupBy = groupBy;
        return this;
    }
    
    public DataqueryBuilder maxNodes(Long maxNodes) {
    this.internal.maxNodes = maxNodes;
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
