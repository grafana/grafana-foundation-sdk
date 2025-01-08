// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class ContactPointBuilder implements com.grafana.foundation.cog.Builder<ContactPoint> {
    protected final ContactPoint internal;
    
    public ContactPointBuilder() {
        this.internal = new ContactPoint();
    }
    public ContactPointBuilder disableResolveMessage(Boolean disableResolveMessage) {
    this.internal.disableResolveMessage = disableResolveMessage;
        return this;
    }
    
    public ContactPointBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public ContactPointBuilder provenance(String provenance) {
    this.internal.provenance = provenance;
        return this;
    }
    
    public ContactPointBuilder settings(Object settings) {
    this.internal.settings = settings;
        return this;
    }
    
    public ContactPointBuilder type(ContactPointType type) {
    this.internal.type = type;
        return this;
    }
    
    public ContactPointBuilder uid(String uid) {
        if (!(uid.length() >= 1)) {
            throw new IllegalArgumentException("uid.length() must be >= 1");
        }
        if (!(uid.length() <= 40)) {
            throw new IllegalArgumentException("uid.length() must be <= 40");
        }
    this.internal.uid = uid;
        return this;
    }
    public ContactPoint build() {
        return this.internal;
    }
}
