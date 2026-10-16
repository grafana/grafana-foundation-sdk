// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


/**
 * Dashboard v2 shape: no `id`, it is carried by the parent's `group`.
 */
public class V2TransformationSpecBuilder implements com.grafana.foundation.cog.Builder<V2TransformationSpec> {
    protected final V2TransformationSpec internal;
    
    public V2TransformationSpecBuilder() {
        this.internal = new V2TransformationSpec();
    }
    public V2TransformationSpecBuilder disabled(Boolean disabled) {
        this.internal.disabled = disabled;
        return this;
    }
    
    public V2TransformationSpecBuilder filter(MatcherConfig filter) {
        this.internal.filter = filter;
        return this;
    }
    
    public V2TransformationSpecBuilder topic(DataTopic topic) {
        this.internal.topic = topic;
        return this;
    }
    
    public V2TransformationSpecBuilder options(Object options) {
        this.internal.options = options;
        return this;
    }
    public V2TransformationSpec build() {
        return this.internal;
    }
}
