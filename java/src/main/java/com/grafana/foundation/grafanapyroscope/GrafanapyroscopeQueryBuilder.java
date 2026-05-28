// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.grafanapyroscope;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;

public class GrafanapyroscopeQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public GrafanapyroscopeQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafanapyroscope";
    }
    public GrafanapyroscopeQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder labelSelector(String labelSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labelSelector = labelSelector;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder spanSelector(List<String> spanSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spanSelector = spanSelector;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder profileTypeId(String profileTypeId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).profileTypeId = profileTypeId;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder groupBy(List<String> groupBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).groupBy = groupBy;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder limit(Long limit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).limit = limit;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder maxNodes(Long maxNodes) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxNodes = maxNodes;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public GrafanapyroscopeQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
