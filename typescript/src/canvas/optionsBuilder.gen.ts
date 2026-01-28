// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as canvas from '../canvas';

export class OptionsBuilder implements cog.Builder<canvas.Options> {
    protected readonly internal: canvas.Options;

    constructor() {
        this.internal = canvas.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): canvas.Options {
        return this.internal;
    }

    // Enable inline editing
    inlineEditing(inlineEditing: boolean): this {
        this.internal.inlineEditing = inlineEditing;
        return this;
    }

    // Show all available element types
    showAdvancedTypes(showAdvancedTypes: boolean): this {
        this.internal.showAdvancedTypes = showAdvancedTypes;
        return this;
    }

    // The root element of canvas (frame), where all canvas elements are nested
    // TODO: Figure out how to define a default value for this
    root(root: {
	// Name of the root element
	name: string;
	// Type of root element (frame)
	type: "frame";
	// The list of canvas elements attached to the root element
	elements: canvas.CanvasElementOptions[];
}): this {
        this.internal.root = root;
        return this;
    }
}

