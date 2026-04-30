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
    
    public CloudwatchQueryV2Builder cloudWatchMetricsQuery(com.grafana.foundation.cog.Builder<CloudWatchMetricsQuery> cloudWatchMetricsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.CloudWatchQuery();
		}
    CloudWatchMetricsQuery cloudWatchMetricsQueryResource = cloudWatchMetricsQuery.build();
        ((CloudWatchQuery) this.internal.spec).cloudWatchMetricsQuery = cloudWatchMetricsQueryResource;
        return this;
    }
    
    public CloudwatchQueryV2Builder cloudWatchLogsQuery(com.grafana.foundation.cog.Builder<CloudWatchLogsQuery> cloudWatchLogsQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.CloudWatchQuery();
		}
    CloudWatchLogsQuery cloudWatchLogsQueryResource = cloudWatchLogsQuery.build();
        ((CloudWatchQuery) this.internal.spec).cloudWatchLogsQuery = cloudWatchLogsQueryResource;
        return this;
    }
    
    public CloudwatchQueryV2Builder cloudWatchAnnotationQuery(com.grafana.foundation.cog.Builder<CloudWatchAnnotationQuery> cloudWatchAnnotationQuery) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.cloudwatch.CloudWatchQuery();
		}
    CloudWatchAnnotationQuery cloudWatchAnnotationQueryResource = cloudWatchAnnotationQuery.build();
        ((CloudWatchQuery) this.internal.spec).cloudWatchAnnotationQuery = cloudWatchAnnotationQueryResource;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
