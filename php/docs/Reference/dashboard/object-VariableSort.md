---
title: <span class="badge object-type-enum"></span> VariableSort
---
# <span class="badge object-type-enum"></span> VariableSort

Sort variable options

Accepted values are:

`0`: No sorting

`1`: Alphabetical ASC

`2`: Alphabetical DESC

`3`: Numerical ASC

`4`: Numerical DESC

`5`: Alphabetical Case Insensitive ASC

`6`: Alphabetical Case Insensitive DESC

`7`: Natural ASC

`8`: Natural DESC

## Definition

```php
final class VariableSort implements \JsonSerializable, \Stringable {
    /**
     * @var int
     */
    private $value;

    /**
     * @var array<string, VariableSort>
     */
    private static $instances = [];

    private function __construct(int $value)
    {
        $this->value = $value;
    }

    public static function disabled(): self
    {
        if (!isset(self::$instances["disabled"])) {
            self::$instances["disabled"] = new self(0);
        }

        return self::$instances["disabled"];
    }

    public static function alphabeticalAsc(): self
    {
        if (!isset(self::$instances["alphabeticalAsc"])) {
            self::$instances["alphabeticalAsc"] = new self(1);
        }

        return self::$instances["alphabeticalAsc"];
    }

    public static function alphabeticalDesc(): self
    {
        if (!isset(self::$instances["alphabeticalDesc"])) {
            self::$instances["alphabeticalDesc"] = new self(2);
        }

        return self::$instances["alphabeticalDesc"];
    }

    public static function numericalAsc(): self
    {
        if (!isset(self::$instances["numericalAsc"])) {
            self::$instances["numericalAsc"] = new self(3);
        }

        return self::$instances["numericalAsc"];
    }

    public static function numericalDesc(): self
    {
        if (!isset(self::$instances["numericalDesc"])) {
            self::$instances["numericalDesc"] = new self(4);
        }

        return self::$instances["numericalDesc"];
    }

    public static function alphabeticalCaseInsensitiveAsc(): self
    {
        if (!isset(self::$instances["alphabeticalCaseInsensitiveAsc"])) {
            self::$instances["alphabeticalCaseInsensitiveAsc"] = new self(5);
        }

        return self::$instances["alphabeticalCaseInsensitiveAsc"];
    }

    public static function alphabeticalCaseInsensitiveDesc(): self
    {
        if (!isset(self::$instances["alphabeticalCaseInsensitiveDesc"])) {
            self::$instances["alphabeticalCaseInsensitiveDesc"] = new self(6);
        }

        return self::$instances["alphabeticalCaseInsensitiveDesc"];
    }

    public static function naturalAsc(): self
    {
        if (!isset(self::$instances["naturalAsc"])) {
            self::$instances["naturalAsc"] = new self(7);
        }

        return self::$instances["naturalAsc"];
    }

    public static function naturalDesc(): self
    {
        if (!isset(self::$instances["naturalDesc"])) {
            self::$instances["naturalDesc"] = new self(8);
        }

        return self::$instances["naturalDesc"];
    }

    public static function fromValue(int $value): self
    {
        if ($value === 0) {
            return self::disabled();
        }

        if ($value === 1) {
            return self::alphabeticalAsc();
        }

        if ($value === 2) {
            return self::alphabeticalDesc();
        }

        if ($value === 3) {
            return self::numericalAsc();
        }

        if ($value === 4) {
            return self::numericalDesc();
        }

        if ($value === 5) {
            return self::alphabeticalCaseInsensitiveAsc();
        }

        if ($value === 6) {
            return self::alphabeticalCaseInsensitiveDesc();
        }

        if ($value === 7) {
            return self::naturalAsc();
        }

        if ($value === 8) {
            return self::naturalDesc();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum VariableSort");
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
