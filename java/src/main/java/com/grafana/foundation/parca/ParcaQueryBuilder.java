// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.parca;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class ParcaQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ParcaQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "parca";
    }
    public ParcaQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ParcaQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public ParcaQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ParcaQueryBuilder labelSelector(String labelSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labelSelector = labelSelector;
        return this;
    }
    
    public ParcaQueryBuilder profileTypeId(String profileTypeId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).profileTypeId = profileTypeId;
        return this;
    }
    
    public ParcaQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public ParcaQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public ParcaQueryBuilder queryType(String queryType) {
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
