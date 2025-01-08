// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.preferences;


public class QueryHistoryPreferenceBuilder implements com.grafana.foundation.cog.Builder<QueryHistoryPreference> {
    protected final QueryHistoryPreference internal;
    
    public QueryHistoryPreferenceBuilder() {
        this.internal = new QueryHistoryPreference();
    }
    public QueryHistoryPreferenceBuilder homeTab(String homeTab) {
    this.internal.homeTab = homeTab;
        return this;
    }
    public QueryHistoryPreference build() {
        return this.internal;
    }
}
