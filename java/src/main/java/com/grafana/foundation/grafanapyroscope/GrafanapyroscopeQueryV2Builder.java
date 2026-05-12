// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.grafanapyroscope;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;
import java.util.List;

public class GrafanapyroscopeQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public GrafanapyroscopeQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafanapyroscope";
    }
    public GrafanapyroscopeQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder labelSelector(String labelSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labelSelector = labelSelector;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder spanSelector(List<String> spanSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spanSelector = spanSelector;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder profileTypeId(String profileTypeId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).profileTypeId = profileTypeId;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder groupBy(List<String> groupBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).groupBy = groupBy;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder limit(Long limit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).limit = limit;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder maxNodes(Long maxNodes) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxNodes = maxNodes;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public GrafanapyroscopeQueryV2Builder queryType(String queryType) {
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
