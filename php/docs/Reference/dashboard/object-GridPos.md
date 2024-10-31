---
title: <span class="badge object-type-class"></span> GridPos
---
# <span class="badge object-type-class"></span> GridPos

Position and dimensions of a panel in the grid

## Definition

```php
class GridPos implements \JsonSerializable
{
    /**
     * Panel height. The height is the number of rows from the top edge of the panel.
     */
    public int $h;

    /**
     * Panel width. The width is the number of columns from the left edge of the panel.
     */
    public int $w;

    /**
     * Panel x. The x coordinate is the number of columns from the left edge of the grid
     */
    public int $x;

    /**
     * Panel y. The y coordinate is the number of rows from the top edge of the grid
     */
    public int $y;

    /**
     * Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
     */
    public ?bool $static;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

