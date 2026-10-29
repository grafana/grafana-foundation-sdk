---
title: <span class="badge object-type-enum"></span> TotalHitsRelation
---
# <span class="badge object-type-enum"></span> TotalHitsRelation

#TotalHitsRelation says how totalHits relates to the real number of matching

rules the caller may see: "eq" when it is exact, "lte" when it is an upper

bound because authorisation was applied after the search ranked its results.

## Definition

```php
final class TotalHitsRelation implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TotalHitsRelation>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function eq(): self
    {
        if (!isset(self::$instances["eq"])) {
            self::$instances["eq"] = new self("eq");
        }

        return self::$instances["eq"];
    }

    public static function lte(): self
    {
        if (!isset(self::$instances["lte"])) {
            self::$instances["lte"] = new self("lte");
        }

        return self::$instances["lte"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "eq") {
            return self::eq();
        }

        if ($value === "lte") {
            return self::lte();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TotalHitsRelation");
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
