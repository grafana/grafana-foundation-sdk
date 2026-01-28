---
title: <span class="badge builder"></span> OptionsBuilder
---
# <span class="badge builder"></span> OptionsBuilder

## Constructor

```typescript
new OptionsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> inlineEditing

Enable inline editing

```typescript
inlineEditing(inlineEditing: boolean)
```

### <span class="badge object-method"></span> panZoom

Enable pan and zoom

```typescript
panZoom(panZoom: boolean)
```

### <span class="badge object-method"></span> root

The root element of canvas (frame), where all canvas elements are nested

TODO: Figure out how to define a default value for this

```typescript
root(root: {
	// Name of the root element
	name: string;
	// Type of root element (frame)
	type: "frame";
	// The list of canvas elements attached to the root element
	elements: canvas.CanvasElementOptions[];
})
```

### <span class="badge object-method"></span> showAdvancedTypes

Show all available element types

```typescript
showAdvancedTypes(showAdvancedTypes: boolean)
```

## See also

 * <span class="badge object-type-interface"></span> [Options](./object-Options.md)
