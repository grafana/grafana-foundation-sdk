"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.CanvasElementOptionsBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class CanvasElementOptionsBuilder {
    constructor() {
        this.internal = canvas.defaultCanvasElementOptions();
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
    type(type) {
        this.internal.type = type;
        return this;
    }
    // TODO: figure out how to define this (element config(s))
    config(config) {
        this.internal.config = config;
        return this;
    }
    constraint(constraint) {
        const constraintResource = constraint.build();
        this.internal.constraint = constraintResource;
        return this;
    }
    placement(placement) {
        const placementResource = placement.build();
        this.internal.placement = placementResource;
        return this;
    }
    background(background) {
        const backgroundResource = background.build();
        this.internal.background = backgroundResource;
        return this;
    }
    border(border) {
        const borderResource = border.build();
        this.internal.border = borderResource;
        return this;
    }
    connections(connections) {
        const connectionsResources = connections.map(builder1 => builder1.build());
        this.internal.connections = connectionsResources;
        return this;
    }
}
exports.CanvasElementOptionsBuilder = CanvasElementOptionsBuilder;
//# sourceMappingURL=canvasElementOptionsBuilder.gen.js.map