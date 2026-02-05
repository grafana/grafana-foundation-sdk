"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.UniqueCountBuilder = void 0;
const tslib_1 = require("tslib");
const elasticsearch = tslib_1.__importStar(require("../elasticsearch"));
class UniqueCountBuilder {
    constructor() {
        this.internal = elasticsearch.defaultUniqueCount();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    field(field) {
        this.internal.field = field;
        return this;
    }
    id(id) {
        this.internal.id = id;
        return this;
    }
    settings(settings) {
        this.internal.settings = settings;
        return this;
    }
    hide(hide) {
        this.internal.hide = hide;
        return this;
    }
}
exports.UniqueCountBuilder = UniqueCountBuilder;
//# sourceMappingURL=uniqueCountBuilder.gen.js.map