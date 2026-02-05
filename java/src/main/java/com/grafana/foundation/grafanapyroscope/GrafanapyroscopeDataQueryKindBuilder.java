// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.grafanapyroscope;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import java.util.List;

public class GrafanapyroscopeDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public GrafanapyroscopeDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafanapyroscope";
    }
    public GrafanapyroscopeDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder labelSelector(String labelSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).labelSelector = labelSelector;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder spanSelector(List<String> spanSelector) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).spanSelector = spanSelector;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder profileTypeId(String profileTypeId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).profileTypeId = profileTypeId;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder groupBy(List<String> groupBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).groupBy = groupBy;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder limit(Long limit) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).limit = limit;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder maxNodes(Long maxNodes) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxNodes = maxNodes;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.grafanapyroscope.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public GrafanapyroscopeDataQueryKindBuilder queryType(String queryType) {
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
