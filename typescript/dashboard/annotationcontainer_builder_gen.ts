// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation events directly from the graph panel or via the HTTP API.
// See https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
export class AnnotationContainerBuilder implements cog.OptionsBuilder<dashboard.AnnotationContainer> {
    private readonly internal: dashboard.AnnotationContainer;

    constructor() {
        this.internal = dashboard.defaultAnnotationContainer();
    }

    build(): dashboard.AnnotationContainer {
        return this.internal;
    }

    // List of annotations
    list(list: cog.OptionsBuilder<dashboard.AnnotationQuery>[]): this {
        const listResources = list.map(builder => builder.build());
        this.internal.list = listResources;
        return this;
    }
}
