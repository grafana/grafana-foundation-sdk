"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.ManifestBuilder = void 0;
const tslib_1 = require("tslib");
const resource = tslib_1.__importStar(require("../resource"));
class ManifestBuilder {
    constructor() {
        this.internal = resource.defaultManifest();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    apiVersion(apiVersion) {
        this.internal.apiVersion = apiVersion;
        return this;
    }
    kind(kind) {
        this.internal.kind = kind;
        return this;
    }
    metadata(metadata) {
        const metadataResource = metadata.build();
        this.internal.metadata = metadataResource;
        return this;
    }
    spec(spec) {
        this.internal.spec = spec;
        return this;
    }
}
exports.ManifestBuilder = ManifestBuilder;
//# sourceMappingURL=manifestBuilder.gen.js.map