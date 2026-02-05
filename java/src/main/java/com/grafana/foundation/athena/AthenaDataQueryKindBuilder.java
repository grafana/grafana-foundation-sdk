// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class AthenaDataQueryKindBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public AthenaDataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-athena-datasource";
    }
    public AthenaDataQueryKindBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public AthenaDataQueryKindBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public AthenaDataQueryKindBuilder format(FormatOptions format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public AthenaDataQueryKindBuilder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
    ConnectionArgs connectionArgsResource = connectionArgs.build();
        ((Dataquery) this.internal.spec).connectionArgs = connectionArgsResource;
        return this;
    }
    
    public AthenaDataQueryKindBuilder table(String table) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).table = table;
        return this;
    }
    
    public AthenaDataQueryKindBuilder column(String column) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).column = column;
        return this;
    }
    
    public AthenaDataQueryKindBuilder queryID(String queryID) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryID = queryID;
        return this;
    }
    
    public AthenaDataQueryKindBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public AthenaDataQueryKindBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public AthenaDataQueryKindBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public AthenaDataQueryKindBuilder rawSQL(String rawSQL) {
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
