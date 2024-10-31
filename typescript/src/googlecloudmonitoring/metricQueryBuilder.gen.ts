// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
export class MetricQueryBuilder implements cog.Builder<googlecloudmonitoring.MetricQuery> {
    protected readonly internal: googlecloudmonitoring.MetricQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultMetricQuery();
    }

    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.MetricQuery {
        return this.internal;
    }

    // GCP project to execute the query against.
    projectName(projectName: string): this {
        this.internal.projectName = projectName;
        return this;
    }

    // Alignment function to be used. Defaults to ALIGN_MEAN.
    perSeriesAligner(perSeriesAligner: string): this {
        this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }

    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignmentPeriod(alignmentPeriod: string): this {
        this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }

    // Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    aliasBy(aliasBy: string): this {
        this.internal.aliasBy = aliasBy;
        return this;
    }

    editorMode(editorMode: string): this {
        this.internal.editorMode = editorMode;
        return this;
    }

    metricType(metricType: string): this {
        this.internal.metricType = metricType;
        return this;
    }

    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    crossSeriesReducer(crossSeriesReducer: string): this {
        this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }

    // Array of labels to group data by.
    groupBys(groupBys: string[]): this {
        this.internal.groupBys = groupBys;
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

    view(view: string): this {
        this.internal.view = view;
        return this;
    }

    // MQL query to be executed.
    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor(preprocessor: googlecloudmonitoring.PreprocessorType): this {
        this.internal.preprocessor = preprocessor;
        return this;
    }

    // To disable the graphPeriod, it should explictly be set to 'disabled'.
    graphPeriod(graphPeriod: string): this {
        this.internal.graphPeriod = graphPeriod;
        return this;
    }
}
