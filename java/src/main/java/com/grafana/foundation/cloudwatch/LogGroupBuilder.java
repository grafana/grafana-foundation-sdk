// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;


public class LogGroupBuilder implements com.grafana.foundation.cog.Builder<LogGroup> {
    protected final LogGroup internal;
    
    public LogGroupBuilder() {
        this.internal = new LogGroup();
    }
    public LogGroupBuilder arn(String arn) {
    this.internal.arn = arn;
        return this;
    }
    
    public LogGroupBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public LogGroupBuilder accountId(String accountId) {
    this.internal.accountId = accountId;
        return this;
    }
    
    public LogGroupBuilder accountLabel(String accountLabel) {
    this.internal.accountLabel = accountLabel;
        return this;
    }
    public LogGroup build() {
        return this.internal;
    }
}
