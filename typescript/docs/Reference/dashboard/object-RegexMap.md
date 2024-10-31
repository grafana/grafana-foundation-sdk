---
title: <span class="badge object-type-interface"></span> RegexMap
---
# <span class="badge object-type-interface"></span> RegexMap

Maps regular expressions to replacement text and a color.

For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.

## Definition

```typescript
export interface RegexMap {
	type: "regex";
	// Regular expression to match against and the result to apply when the value matches the regex
	options: {
		// Regular expression to match against
		pattern: string;
		// Config to apply when the value matches the regex
		result: dashboard.ValueMappingResult;
	};
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RegexMapBuilder](./builder-RegexMapBuilder.md)
