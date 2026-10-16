// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.Map;

/**
 * A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
 * settings, and a map of elements (panels and cells) referenced by the layout.
 */
public class NotebookSpecBuilder implements com.grafana.foundation.cog.Builder<NotebookSpec> {
    protected final NotebookSpec internal;
    
    public NotebookSpecBuilder() {
        this.internal = new NotebookSpec();
    }
    public NotebookSpecBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public NotebookSpecBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public NotebookSpecBuilder tags(List<String> tags) {
        this.internal.tags = tags;
        return this;
    }
    
    public NotebookSpecBuilder timeSettings(com.grafana.foundation.cog.Builder<TimeSettingsSpec> timeSettings) {
    TimeSettingsSpec timeSettingsResource = timeSettings.build();
        this.internal.timeSettings = timeSettingsResource;
        return this;
    }
    
    public NotebookSpecBuilder elements(Map<String, NotebookElement> elements) {
        this.internal.elements = elements;
        return this;
    }
    
    public NotebookSpecBuilder layout(com.grafana.foundation.cog.Builder<NotebookLayoutKind> layout) {
    NotebookLayoutKind layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }
    public NotebookSpec build() {
        return this.internal;
    }
}
