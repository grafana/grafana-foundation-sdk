// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// @deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.
export class LegacyCloudMonitoringAnnotationQueryBuilder implements cog.Builder<googlecloudmonitoring.LegacyCloudMonitoringAnnotationQuery> {
    protected readonly internal: googlecloudmonitoring.LegacyCloudMonitoringAnnotationQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultLegacyCloudMonitoringAnnotationQuery();
    }

    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.LegacyCloudMonitoringAnnotationQuery {
        return this.internal;
    }

    // GCP project to execute the query against.
    projectName(projectName: string): this {
        this.internal.projectName = projectName;
        return this;
    }

    metricType(metricType: string): this {
        this.internal.metricType = metricType;
        return this;
    }

    // Query refId.
    refId(refId: string): this {
        this.internal.refId = refId;
        return this;
    }

    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters(filters: string[]): this {
        this.internal.filters = filters;
        return this;
    }

    metricKind(metricKind: googlecloudmonitoring.MetricKind): this {
        this.internal.metricKind = metricKind;
        return this;
    }

    valueType(valueType: string): this {
        this.internal.valueType = valueType;
        return this;
    }

    // Annotation title.
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Annotation text.
    text(text: string): this {
        this.internal.text = text;
        return this;
    }
}
