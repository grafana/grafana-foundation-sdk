// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// A dashboard snapshot shares an interactive dashboard publicly.
// It is a read-only version of a dashboard, and is not editable.
// It is possible to create a snapshot of a snapshot.
// Grafana strips away all sensitive information from the dashboard.
// Sensitive information stripped: queries (metric, template,annotation) and panel links.
public class Snapshot {
    // Time when the snapshot was created
    @JsonProperty("created")
    public String created;
    // Time when the snapshot expires, default is never to expire
    @JsonProperty("expires")
    public String expires;
    // Is the snapshot saved in an external grafana instance
    @JsonProperty("external")
    public Boolean external;
    // external url, if snapshot was shared in external grafana instance
    @JsonProperty("externalUrl")
    public String externalUrl;
    // Unique identifier of the snapshot
    @JsonProperty("id")
    public Integer id;
    // Optional, defined the unique key of the snapshot, required if external is true
    @JsonProperty("key")
    public String key;
    // Optional, name of the snapshot
    @JsonProperty("name")
    public String name;
    // org id of the snapshot
    @JsonProperty("orgId")
    public Integer orgId;
    // last time when the snapshot was updated
    @JsonProperty("updated")
    public String updated;
    // url of the snapshot, if snapshot was shared internally
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("url")
    public String url;
    // user id of the snapshot creator
    @JsonProperty("userId")
    public Integer userId;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("dashboard")
    public Dashboard dashboard;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<Snapshot> {
        private final Snapshot internal;
        
        public Builder() {
            this.internal = new Snapshot();
        }
    public Builder expires(String expires) {
    this.internal.expires = expires;
        return this;
    }
    
    public Builder external(Boolean external) {
    this.internal.external = external;
        return this;
    }
    
    public Builder externalUrl(String externalUrl) {
    this.internal.externalUrl = externalUrl;
        return this;
    }
    
    public Builder id(Integer id) {
    this.internal.id = id;
        return this;
    }
    
    public Builder key(String key) {
    this.internal.key = key;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder orgId(Integer orgId) {
    this.internal.orgId = orgId;
        return this;
    }
    
    public Builder url(String url) {
    this.internal.url = url;
        return this;
    }
    
    public Builder dashboard(com.grafana.foundation.cog.Builder<Dashboard> dashboard) {
    this.internal.dashboard = dashboard.build();
        return this;
    }
    public Snapshot build() {
            return this.internal;
        }
    }
}
