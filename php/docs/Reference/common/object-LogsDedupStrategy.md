---
title: <span class="badge object-type-enum"></span> LogsDedupStrategy
---
# <span class="badge object-type-enum"></span> LogsDedupStrategy

## Definition

```php
final class LogsDedupStrategy implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LogsDedupStrategy>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function none(): self
    {
        if (!isset(self::$instances["none"])) {
            self::$instances["none"] = new self("none");
        }

        return self::$instances["none"];
    }

    public static function exact(): self
    {
        if (!isset(self::$instances["exact"])) {
            self::$instances["exact"] = new self("exact");
        }

        return self::$instances["exact"];
    }

    public static function numbers(): self
    {
        if (!isset(self::$instances["numbers"])) {
            self::$instances["numbers"] = new self("numbers");
        }

        return self::$instances["numbers"];
    }

    public static function signature(): self
    {
        if (!isset(self::$instances["signature"])) {
            self::$instances["signature"] = new self("signature");
        }

        return self::$instances["signature"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "none") {
            return self::none();
        }

        if ($value === "exact") {
            return self::exact();
        }

        if ($value === "numbers") {
            return self::numbers();
        }

        if ($value === "signature") {
            return self::signature();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LogsDedupStrategy");
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
