// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class CloudwatchQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public CloudwatchQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloudwatch";
    }
    public CloudwatchQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public CloudwatchQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public CloudwatchQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public CloudwatchQueryBuilder metricsQuery(com.grafana.foundation.cog.Builder<MetricsQuery> metricsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.RequestBuilder().build();
		}
    MetricsQuery metricsQueryResource = metricsQuery.build();
        ((Request) this.internal.spec).metricsQuery = metricsQueryResource;
        return this;
    }
    
    public CloudwatchQueryBuilder logsQuery(com.grafana.foundation.cog.Builder<LogsQuery> logsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.RequestBuilder().build();
		}
    LogsQuery logsQueryResource = logsQuery.build();
        ((Request) this.internal.spec).logsQuery = logsQueryResource;
        return this;
    }
    
    public CloudwatchQueryBuilder annotationQuery(com.grafana.foundation.cog.Builder<AnnotationQuery> annotationQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.RequestBuilder().build();
		}
    AnnotationQuery annotationQueryResource = annotationQuery.build();
        ((Request) this.internal.spec).annotationQuery = annotationQueryResource;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
