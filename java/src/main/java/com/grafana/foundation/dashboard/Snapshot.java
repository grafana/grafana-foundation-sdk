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
    // original url, url of the dashboard that was snapshotted
    @JsonProperty("originalUrl")
    public String originalUrl;
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
    public Snapshot() {
    }
    
    public Snapshot(String created,String expires,Boolean external,String externalUrl,String originalUrl,Integer id,String key,String name,Integer orgId,String updated,String url,Integer userId,Dashboard dashboard) {
        this.created = created;
        this.expires = expires;
        this.external = external;
        this.externalUrl = externalUrl;
        this.originalUrl = originalUrl;
        this.id = id;
        this.key = key;
        this.name = name;
        this.orgId = orgId;
        this.updated = updated;
        this.url = url;
        this.userId = userId;
        this.dashboard = dashboard;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
