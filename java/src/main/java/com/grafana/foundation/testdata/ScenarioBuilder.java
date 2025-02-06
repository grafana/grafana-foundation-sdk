// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;


public class ScenarioBuilder implements com.grafana.foundation.cog.Builder<Scenario> {
    protected final Scenario internal;
    
    public ScenarioBuilder() {
        this.internal = new Scenario();
    }
    public ScenarioBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public ScenarioBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public ScenarioBuilder stringInput(String stringInput) {
        this.internal.stringInput = stringInput;
        return this;
    }
    
    public ScenarioBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public ScenarioBuilder hideAliasField(Boolean hideAliasField) {
        this.internal.hideAliasField = hideAliasField;
        return this;
    }
    public Scenario build() {
        return this.internal;
    }
}
