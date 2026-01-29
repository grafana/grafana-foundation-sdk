// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class DataSourceRefBuilder implements com.grafana.foundation.cog.Builder<DataSourceRef> {
    protected final DataSourceRef internal;
    
    public DataSourceRefBuilder() {
        this.internal = new DataSourceRef();
    }
    public DataSourceRefBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public DataSourceRefBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    public DataSourceRef build() {
        return this.internal;
    }
}
