---
title: <span class="badge object-type-enum"></span> ValueTypes
---
# <span class="badge object-type-enum"></span> ValueTypes

## Definition

```php
final class ValueTypes implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ValueTypes>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function vALUETYPEUNSPECIFIED(): self
    {
        if (!isset(self::$instances["VALUE_TYPE_UNSPECIFIED"])) {
            self::$instances["VALUE_TYPE_UNSPECIFIED"] = new self("VALUE_TYPE_UNSPECIFIED");
        }

        return self::$instances["VALUE_TYPE_UNSPECIFIED"];
    }

    public static function bOOL(): self
    {
        if (!isset(self::$instances["BOOL"])) {
            self::$instances["BOOL"] = new self("BOOL");
        }

        return self::$instances["BOOL"];
    }

    public static function iNT64(): self
    {
        if (!isset(self::$instances["INT64"])) {
            self::$instances["INT64"] = new self("INT64");
        }

        return self::$instances["INT64"];
    }

    public static function dOUBLE(): self
    {
        if (!isset(self::$instances["DOUBLE"])) {
            self::$instances["DOUBLE"] = new self("DOUBLE");
        }

        return self::$instances["DOUBLE"];
    }

    public static function sTRING(): self
    {
        if (!isset(self::$instances["STRING"])) {
            self::$instances["STRING"] = new self("STRING");
        }

        return self::$instances["STRING"];
    }

    public static function dISTRIBUTION(): self
    {
        if (!isset(self::$instances["DISTRIBUTION"])) {
            self::$instances["DISTRIBUTION"] = new self("DISTRIBUTION");
        }

        return self::$instances["DISTRIBUTION"];
    }

    public static function mONEY(): self
    {
        if (!isset(self::$instances["MONEY"])) {
            self::$instances["MONEY"] = new self("MONEY");
        }

        return self::$instances["MONEY"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "VALUE_TYPE_UNSPECIFIED") {
            return self::vALUETYPEUNSPECIFIED();
        }

        if ($value === "BOOL") {
            return self::bOOL();
        }

        if ($value === "INT64") {
            return self::iNT64();
        }

        if ($value === "DOUBLE") {
            return self::dOUBLE();
        }

        if ($value === "STRING") {
            return self::sTRING();
        }

        if ($value === "DISTRIBUTION") {
            return self::dISTRIBUTION();
        }

        if ($value === "MONEY") {
            return self::mONEY();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ValueTypes");
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
