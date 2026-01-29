// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ConditionalRenderingVariableSpecBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingVariableSpec> {
    protected final ConditionalRenderingVariableSpec internal;
    
    public ConditionalRenderingVariableSpecBuilder() {
        this.internal = new ConditionalRenderingVariableSpec();
    }
    public ConditionalRenderingVariableSpecBuilder variable(String variable) {
        this.internal.variable = variable;
        return this;
    }
    
    public ConditionalRenderingVariableSpecBuilder operator(ConditionalRenderingVariableSpecOperator operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public ConditionalRenderingVariableSpecBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    public ConditionalRenderingVariableSpec build() {
        return this.internal;
    }
}
