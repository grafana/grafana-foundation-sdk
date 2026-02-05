// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.parca;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class ParcaDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ParcaDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "parca";
    }
    public ParcaDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ParcaDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ParcaDataQueryKindBuilder labelSelector(String labelSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labelSelector = labelSelector;
        return this;
    }
    
    public ParcaDataQueryKindBuilder profileTypeId(String profileTypeId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).profileTypeId = profileTypeId;
        return this;
    }
    
    public ParcaDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public ParcaDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public ParcaDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
