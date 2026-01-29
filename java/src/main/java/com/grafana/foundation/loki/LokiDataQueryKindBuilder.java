// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.loki;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;
import com.grafana.foundation.common.DataSourceRef;

public class LokiDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public LokiDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "loki";
    }
    public LokiDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public LokiDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public LokiDataQueryKindBuilder expr(String expr) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).expr = expr;
        return this;
    }
    
    public LokiDataQueryKindBuilder legendFormat(String legendFormat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).legendFormat = legendFormat;
        return this;
    }
    
    public LokiDataQueryKindBuilder maxLines(Long maxLines) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).maxLines = maxLines;
        return this;
    }
    
    public LokiDataQueryKindBuilder resolution(Long resolution) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).resolution = resolution;
        return this;
    }
    
    public LokiDataQueryKindBuilder editorMode(QueryEditorMode editorMode) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).editorMode = editorMode;
        return this;
    }
    
    public LokiDataQueryKindBuilder range(Boolean range) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).range = range;
        return this;
    }
    
    public LokiDataQueryKindBuilder instant(Boolean instant) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).instant = instant;
        return this;
    }
    
    public LokiDataQueryKindBuilder step(String step) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).step = step;
        return this;
    }
    
    public LokiDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public LokiDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public LokiDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public LokiDataQueryKindBuilder oldDatasource(DataSourceRef datasource) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).datasource = datasource;
        return this;
    }
    
    public LokiDataQueryKindBuilder direction(LokiQueryDirection direction) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.loki.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).direction = direction;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
