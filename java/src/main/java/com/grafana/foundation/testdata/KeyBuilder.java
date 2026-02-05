// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;


public class KeyBuilder implements com.grafana.foundation.cog.Builder<Key> {
    protected final Key internal;
    
    public KeyBuilder() {
        this.internal = new Key();
    }
    public KeyBuilder tick(Double tick) {
        this.internal.tick = tick;
        return this;
    }
    
    public KeyBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public KeyBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    public Key build() {
        return this.internal;
    }
}
