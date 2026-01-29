// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class DataLinkBuilder implements com.grafana.foundation.cog.Builder<DataLink> {
    protected final DataLink internal;
    
    public DataLinkBuilder() {
        this.internal = new DataLink();
    }
    public DataLinkBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public DataLinkBuilder url(String url) {
        this.internal.url = url;
        return this;
    }
    
    public DataLinkBuilder targetBlank(Boolean targetBlank) {
        this.internal.targetBlank = targetBlank;
        return this;
    }
    public DataLink build() {
        return this.internal;
    }
}
