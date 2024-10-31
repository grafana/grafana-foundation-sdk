---
title: <span class="badge object-type-enum"></span> TermsOrder
---
# <span class="badge object-type-enum"></span> TermsOrder

## Definition

```php
final class TermsOrder implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TermsOrder>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function desc(): self
    {
        if (!isset(self::$instances["desc"])) {
            self::$instances["desc"] = new self("desc");
        }

        return self::$instances["desc"];
    }

    public static function asc(): self
    {
        if (!isset(self::$instances["asc"])) {
            self::$instances["asc"] = new self("asc");
        }

        return self::$instances["asc"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "desc") {
            return self::desc();
        }

        if ($value === "asc") {
            return self::asc();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TermsOrder");
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
