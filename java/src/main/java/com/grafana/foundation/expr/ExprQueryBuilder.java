// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.grafana.foundation.dashboardv2beta1.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource;

public class ExprQueryBuilder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ExprQueryBuilder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "__expr__";
    }
    public ExprQueryBuilder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ExprQueryBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public ExprQueryBuilder datasource(com.grafana.foundation.cog.Builder<Dashboardv2beta1DataQueryKindDatasource> datasource) {
    Dashboardv2beta1DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ExprQueryBuilder typeMath(com.grafana.foundation.cog.Builder<TypeMath> typeMath) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeMath typeMathResource = typeMath.build();
        ((Expr) this.internal.spec).typeMath = typeMathResource;
        return this;
    }
    
    public ExprQueryBuilder typeReduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeReduce typeReduceResource = typeReduce.build();
        ((Expr) this.internal.spec).typeReduce = typeReduceResource;
        return this;
    }
    
    public ExprQueryBuilder typeResample(com.grafana.foundation.cog.Builder<TypeResample> typeResample) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeResample typeResampleResource = typeResample.build();
        ((Expr) this.internal.spec).typeResample = typeResampleResource;
        return this;
    }
    
    public ExprQueryBuilder typeClassicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeClassicConditions typeClassicConditionsResource = typeClassicConditions.build();
        ((Expr) this.internal.spec).typeClassicConditions = typeClassicConditionsResource;
        return this;
    }
    
    public ExprQueryBuilder typeThreshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeThreshold typeThresholdResource = typeThreshold.build();
        ((Expr) this.internal.spec).typeThreshold = typeThresholdResource;
        return this;
    }
    
    public ExprQueryBuilder typeSql(com.grafana.foundation.cog.Builder<TypeSql> typeSql) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeSql typeSqlResource = typeSql.build();
        ((Expr) this.internal.spec).typeSql = typeSqlResource;
        return this;
    }
    public DataQueryKind build() {
        return this.internal;
    }
}
