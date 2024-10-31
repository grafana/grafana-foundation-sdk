// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as googlecloudmonitoring from '../googlecloudmonitoring';

// Annotation sub-query properties.
export class AnnotationQueryBuilder implements cog.Builder<googlecloudmonitoring.AnnotationQuery> {
    protected readonly internal: googlecloudmonitoring.AnnotationQuery;

    constructor() {
        this.internal = googlecloudmonitoring.defaultAnnotationQuery();
    }

    /**
     * Builds the object.
     */
    build(): googlecloudmonitoring.AnnotationQuery {
        return this.internal;
    }

    // GCP project to execute the query against.
    projectName(projectName: string): this {
        this.internal.projectName = projectName;
        return this;
    }

    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    crossSeriesReducer(crossSeriesReducer: string): this {
        this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }

    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignmentPeriod(alignmentPeriod: string): this {
        this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }

    // Alignment function to be used. Defaults to ALIGN_MEAN.
    perSeriesAligner(perSeriesAligner: string): this {
        this.internal.perSeriesAligner = perSeriesAligner;
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

    // Data view, defaults to FULL.
    view(view: string): this {
        this.internal.view = view;
        return this;
    }

    // Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    secondaryCrossSeriesReducer(secondaryCrossSeriesReducer: string): this {
        this.internal.secondaryCrossSeriesReducer = secondaryCrossSeriesReducer;
        return this;
    }

    // Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    secondaryAlignmentPeriod(secondaryAlignmentPeriod: string): this {
        this.internal.secondaryAlignmentPeriod = secondaryAlignmentPeriod;
        return this;
    }

    // Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
    secondaryPerSeriesAligner(secondaryPerSeriesAligner: string): this {
        this.internal.secondaryPerSeriesAligner = secondaryPerSeriesAligner;
        return this;
    }

    // Only present if a preprocessor is selected. Array of labels to group data by.
    secondaryGroupBys(secondaryGroupBys: string[]): this {
        this.internal.secondaryGroupBys = secondaryGroupBys;
        return this;
    }

    // Annotation title.
    title(title: string): this {
        this.internal.title = title;
        return this;
    }

    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor(preprocessor: googlecloudmonitoring.PreprocessorType): this {
        this.internal.preprocessor = preprocessor;
        return this;
    }

    // Annotation text.
    text(text: string): this {
        this.internal.text = text;
        return this;
    }
}
