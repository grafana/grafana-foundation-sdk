// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.athena;


public class ConnectionArgsBuilder implements com.grafana.foundation.cog.Builder<ConnectionArgs> {
    protected final ConnectionArgs internal;
    
    public ConnectionArgsBuilder() {
        this.internal = new ConnectionArgs();
    }
    public ConnectionArgsBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public ConnectionArgsBuilder catalog(String catalog) {
        this.internal.catalog = catalog;
        return this;
    }
    
    public ConnectionArgsBuilder database(String database) {
        this.internal.database = database;
        return this;
    }
    
    public ConnectionArgsBuilder resultReuseEnabled(Boolean resultReuseEnabled) {
        this.internal.resultReuseEnabled = resultReuseEnabled;
        return this;
    }
    
    public ConnectionArgsBuilder resultReuseMaxAgeInMinutes(Double resultReuseMaxAgeInMinutes) {
        this.internal.resultReuseMaxAgeInMinutes = resultReuseMaxAgeInMinutes;
        return this;
    }
    public ConnectionArgs build() {
        return this.internal;
    }
}
