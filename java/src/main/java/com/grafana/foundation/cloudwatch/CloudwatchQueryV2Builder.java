// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class CloudwatchQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public CloudwatchQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "cloudwatch";
    }
    public CloudwatchQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public CloudwatchQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public CloudwatchQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public CloudwatchQueryV2Builder metricsQuery(com.grafana.foundation.cog.Builder<MetricsQuery> metricsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.RequestBuilder().build();
		}
    MetricsQuery metricsQueryResource = metricsQuery.build();
        ((Request) this.internal.spec).metricsQuery = metricsQueryResource;
        return this;
    }
    
    public CloudwatchQueryV2Builder logsQuery(com.grafana.foundation.cog.Builder<LogsQuery> logsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.RequestBuilder().build();
		}
    LogsQuery logsQueryResource = logsQuery.build();
        ((Request) this.internal.spec).logsQuery = logsQueryResource;
        return this;
    }
    
    public CloudwatchQueryV2Builder annotationQuery(com.grafana.foundation.cog.Builder<AnnotationQuery> annotationQuery) {
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
