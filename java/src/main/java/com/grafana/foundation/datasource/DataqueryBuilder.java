// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.datasource;

import com.grafana.foundation.common.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder panelId(Integer panelId) {
        this.internal.panelId = panelId;
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
    
    public DataqueryBuilder withTransforms(Boolean withTransforms) {
        this.internal.withTransforms = withTransforms;
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
