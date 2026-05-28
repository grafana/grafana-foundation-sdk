// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.grafana.foundation.dashboardv2.DataQueryKind;
import java.util.Map;
import com.grafana.foundation.dashboardv2.Dashboardv2DataQueryKindDatasource;

public class ExprQueryV2Builder implements com.grafana.foundation.cog.Builder<DataQueryKind> {
    protected final DataQueryKind internal;
    
    public ExprQueryV2Builder() {
        this.internal = new DataQueryKind();
        this.internal.kind = "DataQuery";
        this.internal.group = "__expr__";
    }
    public ExprQueryV2Builder version(String version) {
        this.internal.version = version;
        return this;
    }
    
    public ExprQueryV2Builder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public ExprQueryV2Builder datasource(com.grafana.foundation.cog.Builder<Dashboardv2DataQueryKindDatasource> datasource) {
    Dashboardv2DataQueryKindDatasource datasourceResource = datasource.build();
        this.internal.datasource = datasourceResource;
        return this;
    }
    
    public ExprQueryV2Builder typeMath(com.grafana.foundation.cog.Builder<TypeMath> typeMath) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeMath typeMathResource = typeMath.build();
        ((Expr) this.internal.spec).typeMath = typeMathResource;
        return this;
    }
    
    public ExprQueryV2Builder typeReduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeReduce typeReduceResource = typeReduce.build();
        ((Expr) this.internal.spec).typeReduce = typeReduceResource;
        return this;
    }
    
    public ExprQueryV2Builder typeResample(com.grafana.foundation.cog.Builder<TypeResample> typeResample) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeResample typeResampleResource = typeResample.build();
        ((Expr) this.internal.spec).typeResample = typeResampleResource;
        return this;
    }
    
    public ExprQueryV2Builder typeClassicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeClassicConditions typeClassicConditionsResource = typeClassicConditions.build();
        ((Expr) this.internal.spec).typeClassicConditions = typeClassicConditionsResource;
        return this;
    }
    
    public ExprQueryV2Builder typeThreshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.expr.Expr();
		}
    TypeThreshold typeThresholdResource = typeThreshold.build();
        ((Expr) this.internal.spec).typeThreshold = typeThresholdResource;
        return this;
    }
    
    public ExprQueryV2Builder typeSql(com.grafana.foundation.cog.Builder<TypeSql> typeSql) {
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
