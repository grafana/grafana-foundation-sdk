---
title: <span class="badge object-type-enum"></span> LogsSortOrder
---
# <span class="badge object-type-enum"></span> LogsSortOrder

## Definition

```php
final class LogsSortOrder implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, LogsSortOrder>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function descending(): self
    {
        if (!isset(self::$instances["Descending"])) {
            self::$instances["Descending"] = new self("Descending");
        }

        return self::$instances["Descending"];
    }

    public static function ascending(): self
    {
        if (!isset(self::$instances["Ascending"])) {
            self::$instances["Ascending"] = new self("Ascending");
        }

        return self::$instances["Ascending"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "Descending") {
            return self::descending();
        }

        if ($value === "Ascending") {
            return self::ascending();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum LogsSortOrder");
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
