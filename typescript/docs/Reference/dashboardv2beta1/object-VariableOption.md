---
title: <span class="badge object-type-interface"></span> VariableOption
---
# <span class="badge object-type-interface"></span> VariableOption

Variable option specification

## Definition

```typescript
export interface VariableOption {
	// Whether the option is selected or not
	selected?: boolean;
	// Text to be displayed for the option
	text: string | string[];
	// Value of the option
	value: string | string[];
}

```
