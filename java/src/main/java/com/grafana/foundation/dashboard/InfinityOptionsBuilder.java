// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class InfinityOptionsBuilder implements com.grafana.foundation.cog.Builder<InfinityOptions> {
    protected final InfinityOptions internal;
    
    public InfinityOptionsBuilder() {
        this.internal = new InfinityOptions();
    }
    public InfinityOptionsBuilder method(HttpRequestMethod method) {
        this.internal.method = method;
        return this;
    }
    
    public InfinityOptionsBuilder url(String url) {
        this.internal.url = url;
        return this;
    }
    
    public InfinityOptionsBuilder body(String body) {
        this.internal.body = body;
        return this;
    }
    
    public InfinityOptionsBuilder queryParams(List<List<String>> queryParams) {
        this.internal.queryParams = queryParams;
        return this;
    }
    
    public InfinityOptionsBuilder headers(List<List<String>> headers) {
        this.internal.headers = headers;
        return this;
    }
    
    public InfinityOptionsBuilder datasourceUid(String datasourceUid) {
        this.internal.datasourceUid = datasourceUid;
        return this;
    }
    public InfinityOptions build() {
        return this.internal;
    }
}
