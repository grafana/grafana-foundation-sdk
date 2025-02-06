// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;


public class BaseGrafanaTemplateVariableQueryBuilder implements com.grafana.foundation.cog.Builder<BaseGrafanaTemplateVariableQuery> {
    protected final BaseGrafanaTemplateVariableQuery internal;
    
    public BaseGrafanaTemplateVariableQueryBuilder() {
        this.internal = new BaseGrafanaTemplateVariableQuery();
    }
    public BaseGrafanaTemplateVariableQueryBuilder rawQuery(String rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    public BaseGrafanaTemplateVariableQuery build() {
        return this.internal;
    }
}
