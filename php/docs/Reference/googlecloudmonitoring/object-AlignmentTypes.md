---
title: <span class="badge object-type-enum"></span> AlignmentTypes
---
# <span class="badge object-type-enum"></span> AlignmentTypes

## Definition

```php
final class AlignmentTypes implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AlignmentTypes>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function aLIGNDELTA(): self
    {
        if (!isset(self::$instances["ALIGN_DELTA"])) {
            self::$instances["ALIGN_DELTA"] = new self("ALIGN_DELTA");
        }

        return self::$instances["ALIGN_DELTA"];
    }

    public static function aLIGNRATE(): self
    {
        if (!isset(self::$instances["ALIGN_RATE"])) {
            self::$instances["ALIGN_RATE"] = new self("ALIGN_RATE");
        }

        return self::$instances["ALIGN_RATE"];
    }

    public static function aLIGNINTERPOLATE(): self
    {
        if (!isset(self::$instances["ALIGN_INTERPOLATE"])) {
            self::$instances["ALIGN_INTERPOLATE"] = new self("ALIGN_INTERPOLATE");
        }

        return self::$instances["ALIGN_INTERPOLATE"];
    }

    public static function aLIGNNEXTOLDER(): self
    {
        if (!isset(self::$instances["ALIGN_NEXT_OLDER"])) {
            self::$instances["ALIGN_NEXT_OLDER"] = new self("ALIGN_NEXT_OLDER");
        }

        return self::$instances["ALIGN_NEXT_OLDER"];
    }

    public static function aLIGNMIN(): self
    {
        if (!isset(self::$instances["ALIGN_MIN"])) {
            self::$instances["ALIGN_MIN"] = new self("ALIGN_MIN");
        }

        return self::$instances["ALIGN_MIN"];
    }

    public static function aLIGNMAX(): self
    {
        if (!isset(self::$instances["ALIGN_MAX"])) {
            self::$instances["ALIGN_MAX"] = new self("ALIGN_MAX");
        }

        return self::$instances["ALIGN_MAX"];
    }

    public static function aLIGNMEAN(): self
    {
        if (!isset(self::$instances["ALIGN_MEAN"])) {
            self::$instances["ALIGN_MEAN"] = new self("ALIGN_MEAN");
        }

        return self::$instances["ALIGN_MEAN"];
    }

    public static function aLIGNCOUNT(): self
    {
        if (!isset(self::$instances["ALIGN_COUNT"])) {
            self::$instances["ALIGN_COUNT"] = new self("ALIGN_COUNT");
        }

        return self::$instances["ALIGN_COUNT"];
    }

    public static function aLIGNSUM(): self
    {
        if (!isset(self::$instances["ALIGN_SUM"])) {
            self::$instances["ALIGN_SUM"] = new self("ALIGN_SUM");
        }

        return self::$instances["ALIGN_SUM"];
    }

    public static function aLIGNSTDDEV(): self
    {
        if (!isset(self::$instances["ALIGN_STDDEV"])) {
            self::$instances["ALIGN_STDDEV"] = new self("ALIGN_STDDEV");
        }

        return self::$instances["ALIGN_STDDEV"];
    }

    public static function aLIGNCOUNTTRUE(): self
    {
        if (!isset(self::$instances["ALIGN_COUNT_TRUE"])) {
            self::$instances["ALIGN_COUNT_TRUE"] = new self("ALIGN_COUNT_TRUE");
        }

        return self::$instances["ALIGN_COUNT_TRUE"];
    }

    public static function aLIGNCOUNTFALSE(): self
    {
        if (!isset(self::$instances["ALIGN_COUNT_FALSE"])) {
            self::$instances["ALIGN_COUNT_FALSE"] = new self("ALIGN_COUNT_FALSE");
        }

        return self::$instances["ALIGN_COUNT_FALSE"];
    }

    public static function aLIGNFRACTIONTRUE(): self
    {
        if (!isset(self::$instances["ALIGN_FRACTION_TRUE"])) {
            self::$instances["ALIGN_FRACTION_TRUE"] = new self("ALIGN_FRACTION_TRUE");
        }

        return self::$instances["ALIGN_FRACTION_TRUE"];
    }

    public static function aLIGNPERCENTILE99(): self
    {
        if (!isset(self::$instances["ALIGN_PERCENTILE_99"])) {
            self::$instances["ALIGN_PERCENTILE_99"] = new self("ALIGN_PERCENTILE_99");
        }

        return self::$instances["ALIGN_PERCENTILE_99"];
    }

    public static function aLIGNPERCENTILE95(): self
    {
        if (!isset(self::$instances["ALIGN_PERCENTILE_95"])) {
            self::$instances["ALIGN_PERCENTILE_95"] = new self("ALIGN_PERCENTILE_95");
        }

        return self::$instances["ALIGN_PERCENTILE_95"];
    }

    public static function aLIGNPERCENTILE50(): self
    {
        if (!isset(self::$instances["ALIGN_PERCENTILE_50"])) {
            self::$instances["ALIGN_PERCENTILE_50"] = new self("ALIGN_PERCENTILE_50");
        }

        return self::$instances["ALIGN_PERCENTILE_50"];
    }

    public static function aLIGNPERCENTILE05(): self
    {
        if (!isset(self::$instances["ALIGN_PERCENTILE_05"])) {
            self::$instances["ALIGN_PERCENTILE_05"] = new self("ALIGN_PERCENTILE_05");
        }

        return self::$instances["ALIGN_PERCENTILE_05"];
    }

    public static function aLIGNPERCENTCHANGE(): self
    {
        if (!isset(self::$instances["ALIGN_PERCENT_CHANGE"])) {
            self::$instances["ALIGN_PERCENT_CHANGE"] = new self("ALIGN_PERCENT_CHANGE");
        }

        return self::$instances["ALIGN_PERCENT_CHANGE"];
    }

    public static function aLIGNNONE(): self
    {
        if (!isset(self::$instances["ALIGN_NONE"])) {
            self::$instances["ALIGN_NONE"] = new self("ALIGN_NONE");
        }

        return self::$instances["ALIGN_NONE"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "ALIGN_DELTA") {
            return self::aLIGNDELTA();
        }

        if ($value === "ALIGN_RATE") {
            return self::aLIGNRATE();
        }

        if ($value === "ALIGN_INTERPOLATE") {
            return self::aLIGNINTERPOLATE();
        }

        if ($value === "ALIGN_NEXT_OLDER") {
            return self::aLIGNNEXTOLDER();
        }

        if ($value === "ALIGN_MIN") {
            return self::aLIGNMIN();
        }

        if ($value === "ALIGN_MAX") {
            return self::aLIGNMAX();
        }

        if ($value === "ALIGN_MEAN") {
            return self::aLIGNMEAN();
        }

        if ($value === "ALIGN_COUNT") {
            return self::aLIGNCOUNT();
        }

        if ($value === "ALIGN_SUM") {
            return self::aLIGNSUM();
        }

        if ($value === "ALIGN_STDDEV") {
            return self::aLIGNSTDDEV();
        }

        if ($value === "ALIGN_COUNT_TRUE") {
            return self::aLIGNCOUNTTRUE();
        }

        if ($value === "ALIGN_COUNT_FALSE") {
            return self::aLIGNCOUNTFALSE();
        }

        if ($value === "ALIGN_FRACTION_TRUE") {
            return self::aLIGNFRACTIONTRUE();
        }

        if ($value === "ALIGN_PERCENTILE_99") {
            return self::aLIGNPERCENTILE99();
        }

        if ($value === "ALIGN_PERCENTILE_95") {
            return self::aLIGNPERCENTILE95();
        }

        if ($value === "ALIGN_PERCENTILE_50") {
            return self::aLIGNPERCENTILE50();
        }

        if ($value === "ALIGN_PERCENTILE_05") {
            return self::aLIGNPERCENTILE05();
        }

        if ($value === "ALIGN_PERCENT_CHANGE") {
            return self::aLIGNPERCENTCHANGE();
        }

        if ($value === "ALIGN_NONE") {
            return self::aLIGNNONE();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AlignmentTypes");
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
