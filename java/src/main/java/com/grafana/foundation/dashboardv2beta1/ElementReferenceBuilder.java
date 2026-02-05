// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ElementReferenceBuilder implements com.grafana.foundation.cog.Builder<ElementReference> {
    protected final ElementReference internal;
    
    public ElementReferenceBuilder() {
        this.internal = new ElementReference();
        this.internal.kind = "ElementReference";
    }
    public ElementReferenceBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public ElementReference build() {
        return this.internal;
    }
}
