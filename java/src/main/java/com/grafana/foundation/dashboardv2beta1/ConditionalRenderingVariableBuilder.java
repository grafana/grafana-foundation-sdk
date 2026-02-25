// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ConditionalRenderingVariableBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingVariableKind> {
    protected final ConditionalRenderingVariableKind internal;
    
    public ConditionalRenderingVariableBuilder() {
        this.internal = new ConditionalRenderingVariableKind();
        this.internal.kind = "ConditionalRenderingVariable";
    }
    public ConditionalRenderingVariableBuilder variable(String variable) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingVariableSpec();
		}
        this.internal.spec.variable = variable;
        return this;
    }
    
    public ConditionalRenderingVariableBuilder operator(ConditionalRenderingVariableSpecOperator operator) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingVariableSpec();
		}
        this.internal.spec.operator = operator;
        return this;
    }
    
    public ConditionalRenderingVariableBuilder value(String value) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingVariableSpec();
		}
        this.internal.spec.value = value;
        return this;
    }
    public ConditionalRenderingVariableKind build() {
        return this.internal;
    }
}
