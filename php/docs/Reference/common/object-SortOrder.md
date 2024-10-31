---
title: <span class="badge object-type-enum"></span> SortOrder
---
# <span class="badge object-type-enum"></span> SortOrder

TODO docs

## Definition

```php
final class SortOrder implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SortOrder>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function ascending(): self
    {
        if (!isset(self::$instances["Ascending"])) {
            self::$instances["Ascending"] = new self("asc");
        }

        return self::$instances["Ascending"];
    }

    public static function descending(): self
    {
        if (!isset(self::$instances["Descending"])) {
            self::$instances["Descending"] = new self("desc");
        }

        return self::$instances["Descending"];
    }

    public static function none(): self
    {
        if (!isset(self::$instances["None"])) {
            self::$instances["None"] = new self("none");
        }

        return self::$instances["None"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "asc") {
            return self::ascending();
        }

        if ($value === "desc") {
            return self::descending();
        }

        if ($value === "none") {
            return self::none();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SortOrder");
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
