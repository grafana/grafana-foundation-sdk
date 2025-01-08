// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class DataSourceJsonDataBuilder implements com.grafana.foundation.cog.Builder<DataSourceJsonData> {
    protected final DataSourceJsonData internal;
    
    public DataSourceJsonDataBuilder() {
        this.internal = new DataSourceJsonData();
    }
    public DataSourceJsonDataBuilder authType(String authType) {
    this.internal.authType = authType;
        return this;
    }
    
    public DataSourceJsonDataBuilder defaultRegion(String defaultRegion) {
    this.internal.defaultRegion = defaultRegion;
        return this;
    }
    
    public DataSourceJsonDataBuilder profile(String profile) {
    this.internal.profile = profile;
        return this;
    }
    
    public DataSourceJsonDataBuilder manageAlerts(Boolean manageAlerts) {
    this.internal.manageAlerts = manageAlerts;
        return this;
    }
    
    public DataSourceJsonDataBuilder alertmanagerUid(String alertmanagerUid) {
    this.internal.alertmanagerUid = alertmanagerUid;
        return this;
    }
    public DataSourceJsonData build() {
        return this.internal;
    }
}
