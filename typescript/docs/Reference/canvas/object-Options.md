---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Enable inline editing
	inlineEditing: boolean;
	// Show all available element types
	showAdvancedTypes: boolean;
	// Enable pan and zoom
	panZoom: boolean;
	// The root element of canvas (frame), where all canvas elements are nested
	// TODO: Figure out how to define a default value for this
	root: {
		// Name of the root element
		name: string;
		// Type of root element (frame)
		type: "frame";
		// The list of canvas elements attached to the root element
		elements: canvas.CanvasElementOptions[];
	};
}

```
## Methods

No methods.
