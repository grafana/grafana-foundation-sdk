---
title: <span class="badge object-type-enum"></span> QueryEditorMode
---
# <span class="badge object-type-enum"></span> QueryEditorMode

## Definition

```php
final class QueryEditorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, QueryEditorMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function code(): self
    {
        if (!isset(self::$instances["code"])) {
            self::$instances["code"] = new self("code");
        }

        return self::$instances["code"];
    }

    public static function builder(): self
    {
        if (!isset(self::$instances["builder"])) {
            self::$instances["builder"] = new self("builder");
        }

        return self::$instances["builder"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "code") {
            return self::code();
        }

        if ($value === "builder") {
            return self::builder();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum QueryEditorMode");
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
