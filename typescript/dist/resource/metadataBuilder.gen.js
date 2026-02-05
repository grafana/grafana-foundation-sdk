"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetadataBuilder = void 0;
const tslib_1 = require("tslib");
const resource = tslib_1.__importStar(require("../resource"));
class MetadataBuilder {
    constructor() {
        this.internal = resource.defaultMetadata();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    name(name) {
        this.internal.name = name;
        return this;
    }
    namespace(namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    labels(labels) {
        this.internal.labels = labels;
        return this;
    }
    annotations(annotations) {
        this.internal.annotations = annotations;
        return this;
    }
    uid(uid) {
        this.internal.uid = uid;
        return this;
    }
    resourceVersion(resourceVersion) {
        this.internal.resourceVersion = resourceVersion;
        return this;
    }
    generation(generation) {
        this.internal.generation = generation;
        return this;
    }
    creationTimestamp(creationTimestamp) {
        this.internal.creationTimestamp = creationTimestamp;
        return this;
    }
    updateTimestamp(updateTimestamp) {
        this.internal.updateTimestamp = updateTimestamp;
        return this;
    }
    deletionTimestamp(deletionTimestamp) {
        this.internal.deletionTimestamp = deletionTimestamp;
        return this;
    }
}
exports.MetadataBuilder = MetadataBuilder;
//# sourceMappingURL=metadataBuilder.gen.js.map