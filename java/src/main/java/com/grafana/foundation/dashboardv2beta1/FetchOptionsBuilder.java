// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class FetchOptionsBuilder implements com.grafana.foundation.cog.Builder<FetchOptions> {
    protected final FetchOptions internal;
    
    public FetchOptionsBuilder() {
        this.internal = new FetchOptions();
    }
    public FetchOptionsBuilder method(HttpRequestMethod method) {
        this.internal.method = method;
        return this;
    }
    
    public FetchOptionsBuilder url(String url) {
        this.internal.url = url;
        return this;
    }
    
    public FetchOptionsBuilder body(String body) {
        this.internal.body = body;
        return this;
    }
    
    public FetchOptionsBuilder queryParams(List<List<String>> queryParams) {
        this.internal.queryParams = queryParams;
        return this;
    }
    
    public FetchOptionsBuilder headers(List<List<String>> headers) {
        this.internal.headers = headers;
        return this;
    }
    public FetchOptions build() {
        return this.internal;
    }
}
