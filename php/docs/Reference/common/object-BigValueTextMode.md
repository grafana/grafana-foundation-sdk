---
title: <span class="badge object-type-enum"></span> BigValueTextMode
---
# <span class="badge object-type-enum"></span> BigValueTextMode

TODO docs

## Definition

```php
final class BigValueTextMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BigValueTextMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function value(): self
    {
        if (!isset(self::$instances["Value"])) {
            self::$instances["Value"] = new self("value");
        }

        return self::$instances["Value"];
    }

    public static function valueAndName(): self
    {
        if (!isset(self::$instances["ValueAndName"])) {
            self::$instances["ValueAndName"] = new self("value_and_name");
        }

        return self::$instances["ValueAndName"];
    }

    public static function name(): self
    {
        if (!isset(self::$instances["Name"])) {
            self::$instances["Name"] = new self("name");
        }

        return self::$instances["Name"];
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
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "value") {
            return self::value();
        }

        if ($value === "value_and_name") {
            return self::valueAndName();
        }

        if ($value === "name") {
            return self::name();
        }

        if ($value === "none") {
            return self::none();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BigValueTextMode");
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
