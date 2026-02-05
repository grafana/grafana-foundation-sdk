"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.WorkspacesQueryBuilder = void 0;
const tslib_1 = require("tslib");
const azuremonitor = tslib_1.__importStar(require("../azuremonitor"));
class WorkspacesQueryBuilder {
    constructor() {
        this.internal = azuremonitor.defaultWorkspacesQuery();
        this.internal.kind = "WorkspacesQuery";
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    rawQuery(rawQuery) {
        this.internal.rawQuery = rawQuery;
        return this;
    }
    subscription(subscription) {
        this.internal.subscription = subscription;
        return this;
    }
}
exports.WorkspacesQueryBuilder = WorkspacesQueryBuilder;
//# sourceMappingURL=workspacesQueryBuilder.gen.js.map