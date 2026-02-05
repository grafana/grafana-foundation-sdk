"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConnectionArgsBuilder = void 0;
const tslib_1 = require("tslib");
const athena = tslib_1.__importStar(require("../athena"));
class ConnectionArgsBuilder {
    constructor() {
        this.internal = athena.defaultConnectionArgs();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    region(region) {
        this.internal.region = region;
        return this;
    }
    catalog(catalog) {
        this.internal.catalog = catalog;
        return this;
    }
    database(database) {
        this.internal.database = database;
        return this;
    }
    resultReuseEnabled(resultReuseEnabled) {
        this.internal.resultReuseEnabled = resultReuseEnabled;
        return this;
    }
    resultReuseMaxAgeInMinutes(resultReuseMaxAgeInMinutes) {
        this.internal.resultReuseMaxAgeInMinutes = resultReuseMaxAgeInMinutes;
        return this;
    }
}
exports.ConnectionArgsBuilder = ConnectionArgsBuilder;
//# sourceMappingURL=connectionArgsBuilder.gen.js.map