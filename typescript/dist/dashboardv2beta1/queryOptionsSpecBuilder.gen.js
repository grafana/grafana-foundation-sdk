"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.QueryOptionsSpecBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class QueryOptionsSpecBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultQueryOptionsSpec();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    timeFrom(timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    maxDataPoints(maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    timeShift(timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    queryCachingTTL(queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    interval(interval) {
        this.internal.interval = interval;
        return this;
    }
    cacheTimeout(cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    hideTimeOverride(hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    timeCompare(timeCompare) {
        this.internal.timeCompare = timeCompare;
        return this;
    }
}
exports.QueryOptionsSpecBuilder = QueryOptionsSpecBuilder;
//# sourceMappingURL=queryOptionsSpecBuilder.gen.js.map