---
title: <span class="badge object-type-enum"></span> TraceqlSearchScope
---
# <span class="badge object-type-enum"></span> TraceqlSearchScope

static fields are pre-set in the UI, dynamic fields are added by the user

## Definition

```php
final class TraceqlSearchScope implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, TraceqlSearchScope>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function intrinsic(): self
    {
        if (!isset(self::$instances["intrinsic"])) {
            self::$instances["intrinsic"] = new self("intrinsic");
        }

        return self::$instances["intrinsic"];
    }

    public static function unscoped(): self
    {
        if (!isset(self::$instances["unscoped"])) {
            self::$instances["unscoped"] = new self("unscoped");
        }

        return self::$instances["unscoped"];
    }

    public static function resource(): self
    {
        if (!isset(self::$instances["resource"])) {
            self::$instances["resource"] = new self("resource");
        }

        return self::$instances["resource"];
    }

    public static function span(): self
    {
        if (!isset(self::$instances["span"])) {
            self::$instances["span"] = new self("span");
        }

        return self::$instances["span"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "intrinsic") {
            return self::intrinsic();
        }

        if ($value === "unscoped") {
            return self::unscoped();
        }

        if ($value === "resource") {
            return self::resource();
        }

        if ($value === "span") {
            return self::span();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum TraceqlSearchScope");
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
