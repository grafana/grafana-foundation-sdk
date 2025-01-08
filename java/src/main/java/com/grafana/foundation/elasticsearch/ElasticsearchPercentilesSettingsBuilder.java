// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;

import java.util.List;

public class ElasticsearchPercentilesSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchPercentilesSettings> {
    protected final ElasticsearchPercentilesSettings internal;
    
    public ElasticsearchPercentilesSettingsBuilder() {
        this.internal = new ElasticsearchPercentilesSettings();
    }
    public ElasticsearchPercentilesSettingsBuilder script(com.grafana.foundation.cog.Builder<InlineScript> script) {
    this.internal.script = script.build();
        return this;
    }
    
    public ElasticsearchPercentilesSettingsBuilder missing(String missing) {
    this.internal.missing = missing;
        return this;
    }
    
    public ElasticsearchPercentilesSettingsBuilder percents(List<String> percents) {
    this.internal.percents = percents;
        return this;
    }
    public ElasticsearchPercentilesSettings build() {
        return this.internal;
    }
}
