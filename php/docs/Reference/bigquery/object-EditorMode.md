---
title: <span class="badge object-type-enum"></span> EditorMode
---
# <span class="badge object-type-enum"></span> EditorMode

## Definition

```php
final class EditorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, EditorMode>
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

        throw new \UnexpectedValueException("Value '$value' is not part of the enum EditorMode");
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
