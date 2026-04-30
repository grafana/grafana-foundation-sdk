// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.datasource;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class DatasourceQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public DatasourceQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "datasource";
    }
    public DatasourceQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public DatasourceQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public DatasourceQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public DatasourceQueryV2Builder panelId(Integer panelId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).panelId = panelId;
        return this;
    }
    
    public DatasourceQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public DatasourceQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public DatasourceQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public DatasourceQueryV2Builder withTransforms(Boolean withTransforms) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).withTransforms = withTransforms;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
