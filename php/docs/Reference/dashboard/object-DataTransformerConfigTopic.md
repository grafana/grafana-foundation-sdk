---
title: <span class="badge object-type-enum"></span> DataTransformerConfigTopic
---
# <span class="badge object-type-enum"></span> DataTransformerConfigTopic

## Definition

```php
final class DataTransformerConfigTopic implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, DataTransformerConfigTopic>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function series(): self
    {
        if (!isset(self::$instances["Series"])) {
            self::$instances["Series"] = new self("series");
        }

        return self::$instances["Series"];
    }

    public static function annotations(): self
    {
        if (!isset(self::$instances["Annotations"])) {
            self::$instances["Annotations"] = new self("annotations");
        }

        return self::$instances["Annotations"];
    }

    public static function alertStates(): self
    {
        if (!isset(self::$instances["AlertStates"])) {
            self::$instances["AlertStates"] = new self("alertStates");
        }

        return self::$instances["AlertStates"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "series") {
            return self::series();
        }

        if ($value === "annotations") {
            return self::annotations();
        }

        if ($value === "alertStates") {
            return self::alertStates();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum DataTransformerConfigTopic");
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
