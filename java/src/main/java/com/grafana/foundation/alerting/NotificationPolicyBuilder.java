// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;
import java.util.Map;
import java.util.LinkedList;

public class NotificationPolicyBuilder implements com.grafana.foundation.cog.Builder<NotificationPolicy> {
    protected final NotificationPolicy internal;
    
    public NotificationPolicyBuilder() {
        this.internal = new NotificationPolicy();
    }
    public NotificationPolicyBuilder continueArg(Boolean continueArg) {
        this.internal.continueArg = continueArg;
        return this;
    }
    
    public NotificationPolicyBuilder groupBy(List<String> groupBy) {
        this.internal.groupBy = groupBy;
        return this;
    }
    
    public NotificationPolicyBuilder groupInterval(String groupInterval) {
        this.internal.groupInterval = groupInterval;
        return this;
    }
    
    public NotificationPolicyBuilder groupWait(String groupWait) {
        this.internal.groupWait = groupWait;
        return this;
    }
    
    public NotificationPolicyBuilder match(Map<String, String> match) {
        this.internal.match = match;
        return this;
    }
    
    public NotificationPolicyBuilder matchRe(Map<String, String> matchRe) {
        this.internal.matchRe = matchRe;
        return this;
    }
    
    public NotificationPolicyBuilder matchers(List<com.grafana.foundation.cog.Builder<Matcher>> matchers) {
        List<Matcher> matchersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Matcher> r1 : matchers) {
                Matcher matchersDepth1 = r1.build();
                matchersResources.add(matchersDepth1); 
        }
        this.internal.matchers = matchersResources;
        return this;
    }
    
    public NotificationPolicyBuilder muteTimeIntervals(List<String> muteTimeIntervals) {
        this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }
    
    public NotificationPolicyBuilder objectMatchers(List<List<String>> objectMatchers) {
        this.internal.objectMatchers = objectMatchers;
        return this;
    }
    
    public NotificationPolicyBuilder provenance(String provenance) {
        this.internal.provenance = provenance;
        return this;
    }
    
    public NotificationPolicyBuilder receiver(String receiver) {
        this.internal.receiver = receiver;
        return this;
    }
    
    public NotificationPolicyBuilder repeatInterval(String repeatInterval) {
        this.internal.repeatInterval = repeatInterval;
        return this;
    }
    
    public NotificationPolicyBuilder routes(List<com.grafana.foundation.cog.Builder<NotificationPolicy>> routes) {
        List<NotificationPolicy> routesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<NotificationPolicy> r1 : routes) {
                NotificationPolicy routesDepth1 = r1.build();
                routesResources.add(routesDepth1); 
        }
        this.internal.routes = routesResources;
        return this;
    }
    public NotificationPolicy build() {
        return this.internal;
    }
}
