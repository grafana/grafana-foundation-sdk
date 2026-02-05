"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.InfinityOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const dashboardv2beta1 = tslib_1.__importStar(require("../dashboardv2beta1"));
class InfinityOptionsBuilder {
    constructor() {
        this.internal = dashboardv2beta1.defaultInfinityOptions();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    method(method) {
        this.internal.method = method;
        return this;
    }
    url(url) {
        this.internal.url = url;
        return this;
    }
    body(body) {
        this.internal.body = body;
        return this;
    }
    // These are 2D arrays of strings, each representing a key-value pair
    // We are defining them this way because we can't generate a go struct that
    // that would have exactly two strings in each sub-array
    queryParams(queryParams) {
        this.internal.queryParams = queryParams;
        return this;
    }
    datasourceUid(datasourceUid) {
        this.internal.datasourceUid = datasourceUid;
        return this;
    }
    headers(headers) {
        this.internal.headers = headers;
        return this;
    }
}
exports.InfinityOptionsBuilder = InfinityOptionsBuilder;
//# sourceMappingURL=infinityOptionsBuilder.gen.js.map