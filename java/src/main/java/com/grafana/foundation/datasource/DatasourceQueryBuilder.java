// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.datasource;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class DatasourceQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public DatasourceQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "datasource";
    }
    public DatasourceQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public DatasourceQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public DatasourceQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public DatasourceQueryBuilder panelId(Integer panelId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).panelId = panelId;
        return this;
    }
    
    public DatasourceQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public DatasourceQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public DatasourceQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.datasource.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public DatasourceQueryBuilder withTransforms(Boolean withTransforms) {
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
