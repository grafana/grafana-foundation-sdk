// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class DataQueryKindBuilder<T extends DataQueryKindBuilder<T>> implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public DataQueryKindBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
    }
    public T group(String group) {
        this.internal.group = group;
        return (T) this;
    }
    
    public T version(String version) {
        this.internal.version = version;
        return (T) this;
    }
    
    public T datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return (T) this;
    }
    
    public T spec(Object spec) {
        this.internal.spec = spec;
        return (T) this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
