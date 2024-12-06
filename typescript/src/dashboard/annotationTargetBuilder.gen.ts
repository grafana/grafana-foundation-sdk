// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most dashboards
export class AnnotationTargetBuilder implements cog.Builder<dashboard.AnnotationTarget> {
    protected readonly internal: dashboard.AnnotationTarget;

    constructor() {
        this.internal = dashboard.defaultAnnotationTarget();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationTarget {
        return this.internal;
    }

    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }

    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    matchAny(matchAny: boolean): this {
        this.internal.matchAny = matchAny;
        return this;
    }

    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    // Only required/valid for the grafana datasource...
    // but code+tests is already depending on it so hard to change
    type(type: string): this {
        this.internal.type = type;
        return this;
    }
}
