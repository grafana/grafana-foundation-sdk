---
title: <span class="badge object-type-enum"></span> BackgroundImageSize
---
# <span class="badge object-type-enum"></span> BackgroundImageSize

## Definition

```php
final class BackgroundImageSize implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, BackgroundImageSize>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function original(): self
    {
        if (!isset(self::$instances["Original"])) {
            self::$instances["Original"] = new self("original");
        }

        return self::$instances["Original"];
    }

    public static function contain(): self
    {
        if (!isset(self::$instances["Contain"])) {
            self::$instances["Contain"] = new self("contain");
        }

        return self::$instances["Contain"];
    }

    public static function cover(): self
    {
        if (!isset(self::$instances["Cover"])) {
            self::$instances["Cover"] = new self("cover");
        }

        return self::$instances["Cover"];
    }

    public static function fill(): self
    {
        if (!isset(self::$instances["Fill"])) {
            self::$instances["Fill"] = new self("fill");
        }

        return self::$instances["Fill"];
    }

    public static function tile(): self
    {
        if (!isset(self::$instances["Tile"])) {
            self::$instances["Tile"] = new self("tile");
        }

        return self::$instances["Tile"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "original") {
            return self::original();
        }

        if ($value === "contain") {
            return self::contain();
        }

        if ($value === "cover") {
            return self::cover();
        }

        if ($value === "fill") {
            return self::fill();
        }

        if ($value === "tile") {
            return self::tile();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum BackgroundImageSize");
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
