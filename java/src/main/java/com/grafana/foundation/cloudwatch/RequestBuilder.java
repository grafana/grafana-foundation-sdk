// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class RequestBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Request internal;
    
    public RequestBuilder() {
        this.internal = new Request();
    }
    public RequestBuilder metricsQuery(com.grafana.foundation.cog.Builder<MetricsQuery> metricsQuery) {
    MetricsQuery metricsQueryResource = metricsQuery.build();
        this.internal.metricsQuery = metricsQueryResource;
        return this;
    }
    
    public RequestBuilder logsQuery(com.grafana.foundation.cog.Builder<LogsQuery> logsQuery) {
    LogsQuery logsQueryResource = logsQuery.build();
        this.internal.logsQuery = logsQueryResource;
        return this;
    }
    
    public RequestBuilder annotationQuery(com.grafana.foundation.cog.Builder<AnnotationQuery> annotationQuery) {
    AnnotationQuery annotationQueryResource = annotationQuery.build();
        this.internal.annotationQuery = annotationQueryResource;
        return this;
    }
    public Request build() {
        return this.internal;
    }
}
