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
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("none");
        }

        return self::$instances["None"];
    }

    public static function exact(): self
    {
        if (!isset(self::$instances["Exact"])) {
            self::$instances["Exact"] = new self("exact");
        }

        return self::$instances["Exact"];
    }

    public static function numbers(): self
    {
        if (!isset(self::$instances["Numbers"])) {
            self::$instances["Numbers"] = new self("numbers");
        }

        return self::$instances["Numbers"];
    }

    public static function signature(): self
    {
        if (!isset(self::$instances["Signature"])) {
            self::$instances["Signature"] = new self("signature");
        }

        return self::$instances["Signature"];
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
