// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class AthenaQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public AthenaQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-athena-datasource";
    }
    public AthenaQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public AthenaQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public AthenaQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public AthenaQueryBuilder format(FormatOptions format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public AthenaQueryBuilder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
    ConnectionArgs connectionArgsResource = connectionArgs.build();
        ((Dataquery) this.internal.spec).connectionArgs = connectionArgsResource;
        return this;
    }
    
    public AthenaQueryBuilder table(String table) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).table = table;
        return this;
    }
    
    public AthenaQueryBuilder column(String column) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).column = column;
        return this;
    }
    
    public AthenaQueryBuilder queryID(String queryID) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryID = queryID;
        return this;
    }
    
    public AthenaQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public AthenaQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public AthenaQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public AthenaQueryBuilder rawSQL(String rawSQL) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).rawSQL = rawSQL;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
