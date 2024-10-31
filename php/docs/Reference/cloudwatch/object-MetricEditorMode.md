---
title: <span class="badge object-type-enum"></span> MetricEditorMode
---
# <span class="badge object-type-enum"></span> MetricEditorMode

## Definition

```php
final class MetricEditorMode implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, MetricEditorMode>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function builder(): self
    {
        if (!isset(self::$instances["Builder"])) {
            self::$instances["Builder"] = new self(0);
        }

        return self::$instances["Builder"];
    }

    public static function code(): self
    {
        if (!isset(self::$instances["Code"])) {
            self::$instances["Code"] = new self(1);
        }

        return self::$instances["Code"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::builder();
        }

        if ($value === 1) {
            return self::code();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum MetricEditorMode");
    }

    public function jsonSerialize(): int
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return (string) $this->value;
    }
}

```
