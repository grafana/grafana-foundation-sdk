// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class PrometheusDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public PrometheusDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "prometheus";
    }
    public PrometheusDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder instant() {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = true;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder range() {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = true;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder exemplar(Boolean exemplar) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).exemplar = exemplar;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder format(PromQueryFormat format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder intervalFactor(Double intervalFactor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).intervalFactor = intervalFactor;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder interval(String interval) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).interval = interval;
        return this;
    }
    
    public PrometheusDataQueryKindBuilder rangeAndInstant() {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.prometheus.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = true;
        ((Dataquery) this.internal.spec).instant = true;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
