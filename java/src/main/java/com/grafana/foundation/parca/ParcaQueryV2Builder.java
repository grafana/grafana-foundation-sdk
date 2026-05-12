// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.parca;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class ParcaQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ParcaQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "parca";
    }
    public ParcaQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ParcaQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public ParcaQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ParcaQueryV2Builder labelSelector(String labelSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labelSelector = labelSelector;
        return this;
    }
    
    public ParcaQueryV2Builder profileTypeId(String profileTypeId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).profileTypeId = profileTypeId;
        return this;
    }
    
    public ParcaQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public ParcaQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.parca.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public ParcaQueryV2Builder queryType(String queryType) {
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
