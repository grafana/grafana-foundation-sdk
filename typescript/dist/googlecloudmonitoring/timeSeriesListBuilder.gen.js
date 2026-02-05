"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeSeriesListBuilder = void 0;
const tslib_1 = require("tslib");
const googlecloudmonitoring = tslib_1.__importStar(require("../googlecloudmonitoring"));
// Time Series List sub-query properties.
class TimeSeriesListBuilder {
    constructor() {
        this.internal = googlecloudmonitoring.defaultTimeSeriesList();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    // GCP project to execute the query against.
    projectName(projectName) {
        this.internal.projectName = projectName;
        return this;
    }
    // Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    crossSeriesReducer(crossSeriesReducer) {
        this.internal.crossSeriesReducer = crossSeriesReducer;
        return this;
    }
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignmentPeriod(alignmentPeriod) {
        this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    perSeriesAligner(perSeriesAligner) {
        this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    // Array of labels to group data by.
    groupBys(groupBys) {
        this.internal.groupBys = groupBys;
        return this;
    }
    // Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters(filters) {
        this.internal.filters = filters;
        return this;
    }
    // Data view, defaults to FULL.
    view(view) {
        this.internal.view = view;
        return this;
    }
    // Annotation title.
    title(title) {
        this.internal.title = title;
        return this;
    }
    // Annotation text.
    text(text) {
        this.internal.text = text;
        return this;
    }
    // Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    secondaryCrossSeriesReducer(secondaryCrossSeriesReducer) {
        this.internal.secondaryCrossSeriesReducer = secondaryCrossSeriesReducer;
        return this;
    }
    // Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    secondaryAlignmentPeriod(secondaryAlignmentPeriod) {
        this.internal.secondaryAlignmentPeriod = secondaryAlignmentPeriod;
        return this;
    }
    // Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
    secondaryPerSeriesAligner(secondaryPerSeriesAligner) {
        this.internal.secondaryPerSeriesAligner = secondaryPerSeriesAligner;
        return this;
    }
    // Only present if a preprocessor is selected. Array of labels to group data by.
    secondaryGroupBys(secondaryGroupBys) {
        this.internal.secondaryGroupBys = secondaryGroupBys;
        return this;
    }
    // Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor(preprocessor) {
        this.internal.preprocessor = preprocessor;
        return this;
    }
}
exports.TimeSeriesListBuilder = TimeSeriesListBuilder;
//# sourceMappingURL=timeSeriesListBuilder.gen.js.map