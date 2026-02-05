"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CanvasConnectionBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class CanvasConnectionBuilder {
    constructor() {
        this.internal = canvas.defaultCanvasConnection();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    source(source) {
        const sourceResource = source.build();
        this.internal.source = sourceResource;
        return this;
    }
    target(target) {
        const targetResource = target.build();
        this.internal.target = targetResource;
        return this;
    }
    targetName(targetName) {
        this.internal.targetName = targetName;
        return this;
    }
    color(color) {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    size(size) {
        const sizeResource = size.build();
        this.internal.size = sizeResource;
        return this;
    }
    vertices(vertices) {
        const verticesResources = vertices.map(builder1 => builder1.build());
        this.internal.vertices = verticesResources;
        return this;
    }
    sourceOriginal(sourceOriginal) {
        const sourceOriginalResource = sourceOriginal.build();
        this.internal.sourceOriginal = sourceOriginalResource;
        return this;
    }
    targetOriginal(targetOriginal) {
        const targetOriginalResource = targetOriginal.build();
        this.internal.targetOriginal = targetOriginalResource;
        return this;
    }
}
exports.CanvasConnectionBuilder = CanvasConnectionBuilder;
//# sourceMappingURL=canvasConnectionBuilder.gen.js.map