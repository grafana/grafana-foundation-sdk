// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class ActionBuilder implements com.grafana.foundation.cog.Builder<Action> {
    protected final Action internal;
    
    public ActionBuilder() {
        this.internal = new Action();
    }
    public ActionBuilder type(ActionType type) {
        this.internal.type = type;
        return this;
    }
    
    public ActionBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public ActionBuilder fetch(com.grafana.foundation.cog.Builder<FetchOptions> fetch) {
    FetchOptions fetchResource = fetch.build();
        this.internal.fetch = fetchResource;
        return this;
    }
    
    public ActionBuilder infinity(com.grafana.foundation.cog.Builder<InfinityOptions> infinity) {
    InfinityOptions infinityResource = infinity.build();
        this.internal.infinity = infinityResource;
        return this;
    }
    
    public ActionBuilder confirmation(String confirmation) {
        this.internal.confirmation = confirmation;
        return this;
    }
    
    public ActionBuilder oneClick(Boolean oneClick) {
        this.internal.oneClick = oneClick;
        return this;
    }
    
    public ActionBuilder variables(List<com.grafana.foundation.cog.Builder<ActionVariable>> variables) {
        List<ActionVariable> variablesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ActionVariable> r1 : variables) {
                ActionVariable variablesDepth1 = r1.build();
                variablesResources.add(variablesDepth1); 
        }
        this.internal.variables = variablesResources;
        return this;
    }
    
    public ActionBuilder style(com.grafana.foundation.cog.Builder<Dashboardv2beta1ActionStyle> style) {
    Dashboardv2beta1ActionStyle styleResource = style.build();
        this.internal.style = styleResource;
        return this;
    }
    public Action build() {
        return this.internal;
    }
}
