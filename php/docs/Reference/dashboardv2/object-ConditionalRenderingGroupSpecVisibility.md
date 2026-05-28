---
title: <span class="badge object-type-enum"></span> ConditionalRenderingGroupSpecVisibility
---
# <span class="badge object-type-enum"></span> ConditionalRenderingGroupSpecVisibility

## Definition

```php
final class ConditionalRenderingGroupSpecVisibility implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ConditionalRenderingGroupSpecVisibility>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function show(): self
    {
        if (!isset(self::$instances["Show"])) {
            self::$instances["Show"] = new self("show");
        }

        return self::$instances["Show"];
    }

    public static function hide(): self
    {
        if (!isset(self::$instances["Hide"])) {
            self::$instances["Hide"] = new self("hide");
        }

        return self::$instances["Hide"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "show") {
            return self::show();
        }

        if ($value === "hide") {
            return self::hide();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ConditionalRenderingGroupSpecVisibility");
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
