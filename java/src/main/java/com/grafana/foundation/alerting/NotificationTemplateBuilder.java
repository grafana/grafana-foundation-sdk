// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class NotificationTemplateBuilder implements com.grafana.foundation.cog.Builder<NotificationTemplate> {
    protected final NotificationTemplate internal;
    
    public NotificationTemplateBuilder() {
        this.internal = new NotificationTemplate();
    }
    public NotificationTemplateBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public NotificationTemplateBuilder provenance(String provenance) {
    this.internal.provenance = provenance;
        return this;
    }
    
    public NotificationTemplateBuilder template(String template) {
    this.internal.template = template;
        return this;
    }
    public NotificationTemplate build() {
        return this.internal;
    }
}
