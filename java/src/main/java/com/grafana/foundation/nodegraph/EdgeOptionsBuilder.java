// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.nodegraph;


public class EdgeOptionsBuilder implements com.grafana.foundation.cog.Builder<EdgeOptions> {
    protected final EdgeOptions internal;
    
    public EdgeOptionsBuilder() {
        this.internal = new EdgeOptions();
    }
    public EdgeOptionsBuilder mainStatUnit(String mainStatUnit) {
    this.internal.mainStatUnit = mainStatUnit;
        return this;
    }
    
    public EdgeOptionsBuilder secondaryStatUnit(String secondaryStatUnit) {
    this.internal.secondaryStatUnit = secondaryStatUnit;
        return this;
    }
    public EdgeOptions build() {
        return this.internal;
    }
}
