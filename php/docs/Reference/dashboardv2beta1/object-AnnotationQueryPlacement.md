---
title: <span class="badge object-type-enum"></span> AnnotationQueryPlacement
---
# <span class="badge object-type-enum"></span> AnnotationQueryPlacement

Annotation Query placement. Defines where the annotation query should be displayed.

- "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu

## Definition

```php
final class AnnotationQueryPlacement implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AnnotationQueryPlacement>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function inControlsMenu(): self
    {
        if (!isset(self::$instances["inControlsMenu"])) {
            self::$instances["inControlsMenu"] = new self("inControlsMenu");
        }

        return self::$instances["inControlsMenu"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "inControlsMenu") {
            return self::inControlsMenu();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AnnotationQueryPlacement");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

```
