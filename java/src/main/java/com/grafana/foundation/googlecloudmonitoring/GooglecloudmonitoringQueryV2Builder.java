// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class GooglecloudmonitoringQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public GooglecloudmonitoringQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloud-monitoring";
    }
    public GooglecloudmonitoringQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).refId = refId;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).hide = hide;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder aliasBy(String aliasBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).aliasBy = aliasBy;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    TimeSeriesList timeSeriesListResource = timeSeriesList.build();
        ((CloudMonitoringQuery) this.internal.spec).timeSeriesList = timeSeriesListResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    TimeSeriesQuery timeSeriesQueryResource = timeSeriesQuery.build();
        ((CloudMonitoringQuery) this.internal.spec).timeSeriesQuery = timeSeriesQueryResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    SLOQuery sloQueryResource = sloQuery.build();
        ((CloudMonitoringQuery) this.internal.spec).sloQuery = sloQueryResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder promQLQuery(com.grafana.foundation.cog.Builder<PromQLQuery> promQLQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    PromQLQuery promQLQueryResource = promQLQuery.build();
        ((CloudMonitoringQuery) this.internal.spec).promQLQuery = promQLQueryResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryV2Builder intervalMs(Double intervalMs) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).intervalMs = intervalMs;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
