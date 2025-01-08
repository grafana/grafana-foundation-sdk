// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.canvas;


public class ConnectionCoordinatesBuilder implements com.grafana.foundation.cog.Builder<ConnectionCoordinates> {
    protected final ConnectionCoordinates internal;
    
    public ConnectionCoordinatesBuilder() {
        this.internal = new ConnectionCoordinates();
    }
    public ConnectionCoordinatesBuilder x(Double x) {
    this.internal.x = x;
        return this;
    }
    
    public ConnectionCoordinatesBuilder y(Double y) {
    this.internal.y = y;
        return this;
    }
    public ConnectionCoordinates build() {
        return this.internal;
    }
}
