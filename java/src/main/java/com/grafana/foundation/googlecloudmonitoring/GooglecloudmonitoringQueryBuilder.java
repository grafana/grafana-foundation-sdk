// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class GooglecloudmonitoringQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public GooglecloudmonitoringQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloud-monitoring";
    }
    public GooglecloudmonitoringQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).refId = refId;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).hide = hide;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder queryType(String queryType) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).queryType = queryType;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder aliasBy(String aliasBy) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
        ((CloudMonitoringQuery) this.internal.spec).aliasBy = aliasBy;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    TimeSeriesList timeSeriesListResource = timeSeriesList.build();
        ((CloudMonitoringQuery) this.internal.spec).timeSeriesList = timeSeriesListResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    TimeSeriesQuery timeSeriesQueryResource = timeSeriesQuery.build();
        ((CloudMonitoringQuery) this.internal.spec).timeSeriesQuery = timeSeriesQueryResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    SLOQuery sloQueryResource = sloQuery.build();
        ((CloudMonitoringQuery) this.internal.spec).sloQuery = sloQueryResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder promQLQuery(com.grafana.foundation.cog.Builder<PromQLQuery> promQLQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.googlecloudmonitoring.CloudMonitoringQueryBuilder().build();
		}
    PromQLQuery promQLQueryResource = promQLQuery.build();
        ((CloudMonitoringQuery) this.internal.spec).promQLQuery = promQLQueryResource;
        return this;
    }
    
    public GooglecloudmonitoringQueryBuilder intervalMs(Double intervalMs) {
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
