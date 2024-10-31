---
title: <span class="badge object-type-enum"></span> BigValueColorMode
---
# <span class="badge object-type-enum"></span> BigValueColorMode

TODO docs

## Definition

```php
final class BigValueColorMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BigValueColorMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function value(): self
    {
        if (!isset(self::$instances["Value"])) {
            self::$instances["Value"] = new self("value");
        }

        return self::$instances["Value"];
    }

    public static function background(): self
    {
        if (!isset(self::$instances["Background"])) {
            self::$instances["Background"] = new self("background");
        }

        return self::$instances["Background"];
    }

    public static function backgroundSolid(): self
    {
        if (!isset(self::$instances["BackgroundSolid"])) {
            self::$instances["BackgroundSolid"] = new self("background_solid");
        }

        return self::$instances["BackgroundSolid"];
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
        if ($value === "value") {
            return self::value();
        }

        if ($value === "background") {
            return self::background();
        }

        if ($value === "background_solid") {
            return self::backgroundSolid();
        }

        if ($value === "none") {
            return self::none();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BigValueColorMode");
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
