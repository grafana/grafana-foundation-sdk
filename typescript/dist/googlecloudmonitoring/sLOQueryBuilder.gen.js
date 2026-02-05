"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.SLOQueryBuilder = void 0;
const tslib_1 = require("tslib");
const googlecloudmonitoring = tslib_1.__importStar(require("../googlecloudmonitoring"));
// SLO sub-query properties.
class SLOQueryBuilder {
    constructor() {
        this.internal = googlecloudmonitoring.defaultSLOQuery();
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
    // Alignment function to be used. Defaults to ALIGN_MEAN.
    perSeriesAligner(perSeriesAligner) {
        this.internal.perSeriesAligner = perSeriesAligner;
        return this;
    }
    // Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignmentPeriod(alignmentPeriod) {
        this.internal.alignmentPeriod = alignmentPeriod;
        return this;
    }
    // SLO selector.
    selectorName(selectorName) {
        this.internal.selectorName = selectorName;
        return this;
    }
    // ID for the service the SLO is in.
    serviceId(serviceId) {
        this.internal.serviceId = serviceId;
        return this;
    }
    // Name for the service the SLO is in.
    serviceName(serviceName) {
        this.internal.serviceName = serviceName;
        return this;
    }
    // ID for the SLO.
    sloId(sloId) {
        this.internal.sloId = sloId;
        return this;
    }
    // Name of the SLO.
    sloName(sloName) {
        this.internal.sloName = sloName;
        return this;
    }
    // SLO goal value.
    goal(goal) {
        this.internal.goal = goal;
        return this;
    }
    // Specific lookback period for the SLO.
    lookbackPeriod(lookbackPeriod) {
        this.internal.lookbackPeriod = lookbackPeriod;
        return this;
    }
}
exports.SLOQueryBuilder = SLOQueryBuilder;
//# sourceMappingURL=sLOQueryBuilder.gen.js.map