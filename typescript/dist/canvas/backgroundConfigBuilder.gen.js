"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.BackgroundConfigBuilder = void 0;
const tslib_1 = require("tslib");
const canvas = tslib_1.__importStar(require("../canvas"));
class BackgroundConfigBuilder {
    constructor() {
        this.internal = canvas.defaultBackgroundConfig();
    }
    /**
     * Builds the object.
     */
    build() {
        return this.internal;
    }
    color(color) {
        const colorResource = color.build();
        this.internal.color = colorResource;
        return this;
    }
    image(image) {
        const imageResource = image.build();
        this.internal.image = imageResource;
        return this;
    }
    size(size) {
        this.internal.size = size;
        return this;
    }
}
exports.BackgroundConfigBuilder = BackgroundConfigBuilder;
//# sourceMappingURL=backgroundConfigBuilder.gen.js.map