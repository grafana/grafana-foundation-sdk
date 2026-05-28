// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class PrometheusQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public PrometheusQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "prometheus";
    }
    public PrometheusQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public PrometheusQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public PrometheusQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public PrometheusQueryV2Builder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public PrometheusQueryV2Builder instant(Boolean instant) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = instant;
        return this;
    }
    
    public PrometheusQueryV2Builder range(Boolean range) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = range;
        return this;
    }
    
    public PrometheusQueryV2Builder exemplar(Boolean exemplar) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).exemplar = exemplar;
        return this;
    }
    
    public PrometheusQueryV2Builder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public PrometheusQueryV2Builder format(PromQueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public PrometheusQueryV2Builder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public PrometheusQueryV2Builder intervalFactor(Double intervalFactor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalFactor = intervalFactor;
        return this;
    }
    
    public PrometheusQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public PrometheusQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public PrometheusQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public PrometheusQueryV2Builder interval(String interval) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).interval = interval;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
