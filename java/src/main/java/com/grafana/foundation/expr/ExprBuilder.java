// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Expr internal;
    
    public ExprBuilder() {
        this.internal = new Expr();
    }
    public ExprBuilder typeMath(com.grafana.foundation.cog.Builder<TypeMath> typeMath) {
    TypeMath typeMathResource = typeMath.build();
        this.internal.typeMath = typeMathResource;
        return this;
    }
    
    public ExprBuilder typeReduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce) {
    TypeReduce typeReduceResource = typeReduce.build();
        this.internal.typeReduce = typeReduceResource;
        return this;
    }
    
    public ExprBuilder typeResample(com.grafana.foundation.cog.Builder<TypeResample> typeResample) {
    TypeResample typeResampleResource = typeResample.build();
        this.internal.typeResample = typeResampleResource;
        return this;
    }
    
    public ExprBuilder typeClassicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions) {
    TypeClassicConditions typeClassicConditionsResource = typeClassicConditions.build();
        this.internal.typeClassicConditions = typeClassicConditionsResource;
        return this;
    }
    
    public ExprBuilder typeThreshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold) {
    TypeThreshold typeThresholdResource = typeThreshold.build();
        this.internal.typeThreshold = typeThresholdResource;
        return this;
    }
    
    public ExprBuilder typeSql(com.grafana.foundation.cog.Builder<TypeSql> typeSql) {
    TypeSql typeSqlResource = typeSql.build();
        this.internal.typeSql = typeSqlResource;
        return this;
    }
    public Expr build() {
        return this.internal;
    }
}
