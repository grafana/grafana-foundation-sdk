// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class AthenaQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public AthenaQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "grafana-athena-datasource";
    }
    public AthenaQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public AthenaQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public AthenaQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public AthenaQueryV2Builder format(FormatOptions format) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).format = format;
        return this;
    }
    
    public AthenaQueryV2Builder connectionArgs(com.grafana.foundation.cog.Builder<ConnectionArgs> connectionArgs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
    ConnectionArgs connectionArgsResource = connectionArgs.build();
        ((Dataquery) this.internal.spec).connectionArgs = connectionArgsResource;
        return this;
    }
    
    public AthenaQueryV2Builder table(String table) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).table = table;
        return this;
    }
    
    public AthenaQueryV2Builder column(String column) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).column = column;
        return this;
    }
    
    public AthenaQueryV2Builder queryID(String queryID) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryID = queryID;
        return this;
    }
    
    public AthenaQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).refId = refId;
        return this;
    }
    
    public AthenaQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).hide = hide;
        return this;
    }
    
    public AthenaQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.athena.DataqueryBuilder().build();
		}
        ((Dataquery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public AthenaQueryV2Builder rawSQL(String rawSQL) {
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
