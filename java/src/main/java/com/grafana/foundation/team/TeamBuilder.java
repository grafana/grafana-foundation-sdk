// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.team;


public class TeamBuilder implements com.grafana.foundation.cog.Builder<Team> {
    protected final Team internal;
    
    public TeamBuilder(String name) {
        this.internal = new Team();
        this.internal.name = name;
    }
    public TeamBuilder email(String email) {
        this.internal.email = email;
        return this;
    }
    
    public TeamBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    public Team build() {
        return this.internal;
    }
}
