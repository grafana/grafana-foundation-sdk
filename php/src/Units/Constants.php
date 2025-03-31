<?php

namespace Grafana\Foundation\Units;

final class Constants
{
    const NO_UNIT = "none";
    const NUMBER = "none";
    const STRING = "string";
    const SHORT = "short";
    /**
     * SI short
     */
    const SI_SHORT = "sishort";
    /**
     * Percent (0-100)
     */
    const PERCENT = "percent";
    /**
     * Percent (0.0-1.0)
     */
    const PERCENT_UNIT = "percentunit";
    /**
     * Humidity (%H)
     */
    const HUMIDITY = "humidity";
    /**
     * Decibel (dB)
     */
    const DECIBEL = "dB";
    /**
     * Candela (cd)
     */
    const CANDELA = "candela";
    /**
     * Hexadecimal (0x)
     */
    const HEXADECIMAL_OX = "hex0x";
    const HEXADECIMAL = "hex";
    const SCIENTIFIC_NOTATION = "sci";
    const LOCALE_FORMAT = "locale";
    /**
     * Pixels (px)
     */
    const PIXELS = "pixel";
    /**
     * Acceleration: Meters/sec² (m/sec²)
     */
    const METERS_SEC_SQUARED = "accMS2";
    /**
     * Acceleration: Feet/sec² (f/sec²)
     */
    const FEET_SEC_SQUARED = "accFS2";
    /**
     * G unit (g)
     */
    const G_UNIT = "accG";
    /**
     * Angle: Degrees (°)
     */
    const DEGREES = "degree";
    /**
     * Angle: Radians
     */
    const RADIANS = "radian";
    /**
     * Angle: Gradian
     */
    const GRADIAN = "grad";
    /**
     * Angle: Arc Minutes
     */
    const ARC_MINUTES = "arcmin";
    /**
     * Angle: Arc Seconds
     */
    const ARC_SECONDS = "arcsec";
    /**
     * Area: Square Meters (m²)
     */
    const SQUARE_METERS = "areaM2";
    /**
     * Area: Square Feet (ft²)
     */
    const SQUARE_FEET = "areaF2";
    /**
     * Area: Square Miles (mi²)
     */
    const SQUARE_MILES = "areaMI2";
    /**
     * Area: Acres (ac)
     */
    const ACRES = "acres";
    /**
     * Area: Hectares (ha)
     */
    const HECTARES = "hectares";
    /**
     * Computation: FLOP/s
     */
    const FLOPS_PER_SECOND = "flops";
    /**
     * Computation: MFLOP/s
     */
    const MEGA_FLOPS_PER_SECOND = "mflops";
    /**
     * Computation: GFLOP/s
     */
    const GIGA_FLOPS_PER_SECOND = "gflops";
    /**
     * Computation: TFLOP/s
     */
    const TERA_FLOPS_PER_SECOND = "tflops";
    /**
     * Computation: PFLOP/s
     */
    const PETA_FLOPS_PER_SECOND = "pflops";
    /**
     * Computation: EFLOP/s
     */
    const EXA_FLOPS_PER_SECOND = "eflops";
    /**
     * Computation: ZFLOP/s
     */
    const ZETTA_FLOPS_PER_SECOND = "zflops";
    /**
     * Computation: YFLOP/s
     */
    const YOTTA_FLOPS_PER_SECOND = "yflops";
    /**
     * Concentration: parts-per-million (ppm)
     */
    const PARTS_PER_MILLION = "ppm";
    /**
     * Concentration: parts-per-billion (ppb)
     */
    const PARTS_PER_BILLION = "conppb";
    /**
     * Concentration: nanogram per cubic meter (ng/m³)
     */
    const NANOGRAM_PER_CUBIC_METER = "conngm3";
    /**
     * Concentration: nanogram per normal cubic meter (ng/Nm³)
     */
    const NANOGRAM_PER_NORMAL_CUBIC_METER = "conngNm3";
    /**
     * Concentration: microgram per cubic meter (μg/m³)
     */
    const MICROGRAM_PER_CUBIC_METER = "conμgm3";
    /**
     * Concentration: microgram per normal cubic meter (μg/Nm³)
     */
    const MICROGRAM_PER_NORMAL_CUBIC_METER = "conμgNm3";
    /**
     * Concentration: milligram per cubic meter (mg/m³)
     */
    const MILLIGRAM_PER_CUBIC_METER = "conmgm3";
    /**
     * Concentration: milligram per normal cubic meter (mg/Nm³)
     */
    const MILLIGRAM_PER_NORMAL_CUBIC_METER = "conmgNm3";
    /**
     * Concentration: gram per cubic meter (g/m³)
     */
    const GRAM_PER_CUBIC_METER = "congm3";
    /**
     * Concentration: gram per normal cubic meter (g/Nm³)
     */
    const GRAM_PER_NORMAL_CUBIC_METER = "congNm3";
    /**
     * Concentration: milligrams per decilitre (mg/dL)
     */
    const MILLIGRAMS_PER_DECILITRE = "conmgdL";
    /**
     * Concentration: millimoles per litre (mmol/L)
     */
    const MILLIMOLES_PER_LITER = "conmmolL";
    /**
     * Currency: Dollars ($)
     */
    const DOLLARS = "currencyUSD";
    /**
     * Currency: Pounds (£)
     */
    const POUNDS = "currencyGBP";
    /**
     * Currency: Euro (€)
     */
    const EURO = "currencyEUR";
    /**
     * Currency: Yen (¥)
     */
    const YEN = "currencyJPY";
    /**
     * Currency: Rubles (₽)
     */
    const RUBLES = "currencyRUB";
    /**
     * Currency: Hryvnias (₴)
     */
    const HRYVNIAS = "currencyUAH";
    /**
     * Currency: Real (R$)
     */
    const REAL = "currencyBRL";
    /**
     * Currency: Danish Krone (kr)
     */
    const DANISH_KRONE = "currencyDKK";
    /**
     * Currency: Icelandic Króna (kr)
     */
    const ICELANDIC_KRONA = "currencyISK";
    /**
     * Currency: Norwegian Krone (kr)
     */
    const NORWEGIAN_KRONE = "currencyNOK";
    /**
     * Currency: Swedish Krona (kr)
     */
    const SWEDISH_KRONA = "currencySEK";
    /**
     * Currency: Czech koruna (czk)
     */
    const CZECH_KORUNA = "currencyCZK";
    /**
     * Currency: Swiss franc (CHF)
     */
    const SWISS_FRANC = "currencyCHF";
    /**
     * Currency: Polish Złoty (PLN)
     */
    const POLISH_ZLOTY = "currencyPLN";
    /**
     * Currency: Bitcoin (฿)
     */
    const BICTOIN = "currencyBTC";
    /**
     * Currency: Milli Bitcoin (฿)
     */
    const MILLI_BITCOIN = "currencymBTC";
    /**
     * Currency: Micro Bitcoin (฿)
     */
    const MICRO_BITCOIN = "currencyμBTC";
    /**
     * Currency: South African Rand (R)
     */
    const SOUTH_AFRICAN_RAND = "currencyZAR";
    /**
     * Currency: Indian Rupee (₹)
     */
    const INDIAN_RUPEE = "currencyINR";
    /**
     * Currency: South Korean Won (₩)
     */
    const SOUTH_KOREAN_WON = "currencyKRW";
    /**
     * Currency: Indonesian Rupiah (Rp)
     */
    const INDONESIAN_RUPIAH = "currencyIDR";
    /**
     * Currency: Philippine Peso (PHP)
     */
    const PHILIPPINE_PESO = "currencyPHP";
    /**
     * Currency: Vietnamese Dong (VND)
     */
    const VIETNAMESE_DONG = "currencyVND";
    /**
     * Currency: Turkish Lira (₺)
     */
    const TURKISH_LIRA = "currencyTRY";
    /**
     * Currency: Malaysian Ringgit (RM)
     */
    const MALAYSIAN_RINGGIT = "currencyMYR";
    /**
     * Currency: CFP franc (XPF)
     */
    const CFP_FRANC = "currencyXPF";
    /**
     * Currency: Bulgarian Lev (BGN)
     */
    const BULGARIAN_LEV = "currencyBGN";
    /**
     * Currency: Guaraní (₲)
     */
    const GUARANI = "currencyPYG";
    /**
     * Currency: Uruguay Peso (UYU)
     */
    const URUGUAY_PESO = "currencyUYU";
    /**
     * Currency: Israeli New Shekels (₪)
     */
    const ISRAELI_NEW_SHEKELS = "currencyILS";
    /**
     * Data: bytes(IEC)
     */
    const BYTES_IEC = "bytes";
    /**
     * Data: bytes(SI)
     */
    const BYTES_SI = "decbytes";
    /**
     * Data: bits(IEC)
     */
    const BITS_IEC = "bits";
    /**
     * Data: bits(SI)
     */
    const BITS_SI = "bydecbitstes";
    /**
     * Data: kibibytes
     */
    const KIBIBYTES = "kbytes";
    /**
     * Data: kilobytes
     */
    const KILOBYTES = "deckbytes";
    /**
     * Data: mebibytes
     */
    const MEBIBYTES = "mbytes";
    /**
     * Data: megabytes
     */
    const MEGABYTES = "decmbytes";
    /**
     * Data: gibibytes
     */
    const GIBIBYTES = "gbytes";
    /**
     * Data: gigabytes
     */
    const GIGABYTES = "decgbytes";
    /**
     * Data: tebibytes
     */
    const TEBIBYTES = "tbytes";
    /**
     * Data: terabytes
     */
    const TERABYTES = "dectbytes";
    /**
     * Data: pebibytes
     */
    const PEBIBYTES = "pbytes";
    /**
     * Data: petabytes
     */
    const PETABYTES = "decpbytes";
    /**
     * Data rate: packets/sec
     */
    const PACKETS_PER_SECOND = "pps";
    /**
     * Data rate: bytes/sec(IEC)
     */
    const BYTES_PER_SECOND_IEC = "binBps";
    /**
     * Data rate: bytes/sec(SI)
     */
    const BYTES_PER_SECOND_SI = "Bps";
    /**
     * Data rate: bits/sec(IEC)
     */
    const BITS_PER_SECOND_IEC = "binbps";
    /**
     * Data rate: bits/sec(SI)
     */
    const BITS_PER_SECOND_SI = "bps";
    /**
     * Data rate: kibibytes/sec
     */
    const KIBIBYTES_PER_SECOND = "KiBs";
    /**
     * Data rate: kibibits/sec
     */
    const KIBIBITS_PER_SECOND = "Kibits";
    /**
     * Data rate: kilobytes/sec
     */
    const KILOBYTES_PER_SECOND = "KBs";
    /**
     * Data rate: kilobits/sec
     */
    const KILOBITS_PER_SECOND = "Kbits";
    /**
     * Data rate: mebibytes/sec
     */
    const MEBIBYTES_PER_SECOND = "MiBs";
    /**
     * Data rate: mebibits/sec
     */
    const MEBIBITS_PER_SECOND = "Mibits";
    /**
     * Data rate: megabytes/sec
     */
    const MEGABYTES_PER_SECOND = "MBs";
    /**
     * Data rate: megabits/sec
     */
    const MEGABITS_PER_SECOND = "Mbits";
    /**
     * Data rate: gibibytes/sec
     */
    const GIBIBYTES_PER_SECOND = "GiBs";
    /**
     * Data rate: gibibits/sec
     */
    const GIBIBITS_PER_SECOND = "Gibits";
    /**
     * Data rate: gigabytes/sec
     */
    const GIGABYTES_PER_SECOND = "GBs";
    /**
     * Data rate: gigabits/sec
     */
    const GIGABITS_PER_SECOND = "Gbits";
    /**
     * Data rate: tebibytes/sec
     */
    const TEBIBYTES_PER_SECOND = "TiBs";
    /**
     * Data rate: tebibits/sec
     */
    const TEBIBITS_PER_SECOND = "Tibits";
    /**
     * Data rate: terabytes/sec
     */
    const TERABYTES_PER_SECOND = "TBs";
    /**
     * Data rate: terabits/sec
     */
    const TERABITS_PER_SECOND = "Tbits";
    /**
     * Data rate: pebibytes/sec
     */
    const PEBIBYTES_PER_SECOND = "PiBs";
    /**
     * Data rate: pebibits/sec
     */
    const PEBIBITS_PER_SECOND = "Pibits";
    /**
     * Data rate: petabytes/sec
     */
    const PETABYTES_PER_SECOND = "PBs";
    /**
     * Data rate: petabits/sec
     */
    const PETABITS_PER_SECOND = "Pbits";
    /**
     * Date & time: Datetime ISO
     */
    const DATETIME_ISO = "dateTimeAsIso";
    /**
     * Date & time: Datetime ISO (No date if today)
     */
    const DATETIME_ISO_NOT_TODAY = "dateTimeAsIsoNoDateIfToday";
    /**
     * Date & time: Datetime US
     */
    const DATETIME_US = "dateTimeAsUS";
    /**
     * Date & time: Datetime US (No date if today)
     */
    const DATETIME_US_NOT_TODAY = "dateTimeAsUSNoDateIfToday";
    /**
     * Date & time: Datetime local
     */
    const DATETIME_LOCAL = "dateTimeAsLocal";
    /**
     * Date & time: Datetime local (No date if today)
     */
    const DATETIME_LOCAL_NOT_TODAY = "dateTimeAsLocalNoDateIfToday";
    /**
     * Date & time: Datetime default
     */
    const DATETIME_DEFAULT = "dateTimeAsSystem";
    /**
     * Date & time: From Now
     */
    const DATE_TIME_FROM_NOW = "dateTimeFromNow";
    /**
     * Energy: Watt (W)
     */
    const WATT = "watt";
    /**
     * Energy: Kilowatt (kW)
     */
    const KILO_WATT = "kwatt";
    /**
     * Energy: Megawatt (MW)
     */
    const MEGA_WATT = "megwatt";
    /**
     * Energy: Gigawatt (GW)
     */
    const GIGA_WATT = "gwatt";
    /**
     * Energy: Milliwatt (mW)
     */
    const MILLI_WATT = "mwatt";
    /**
     * Energy: Watt per square meter (W/m²)
     */
    const WATT_PER_SQUARE_METER = "Wm2";
    /**
     * Energy: Volt-Ampere (VA)
     */
    const VOLT_AMPERE = "voltamp";
    /**
     * Energy: Kilovolt-Ampere (kVA)
     */
    const KILO_VOLT_AMPERE = "kvoltamp";
    /**
     * Energy: Volt-Ampere reactive (VAr)
     */
    const VOLT_AMPERE_REACTIVE = "voltampreact";
    /**
     * Energy: Kilovolt-Ampere reactive (kVAr)
     */
    const KILO_VOLT_AMPERE_REACTIVE = "kvoltampreact";
    /**
     * Energy: Watt-hour (Wh)
     */
    const WATT_HOUR = "watth";
    /**
     * Energy: Watt-hour per Kilogram (Wh/kg)
     */
    const WATT_HOUR_PER_KILOGRAM = "watthperkg";
    /**
     * Energy: Kilowatt-hour (kWh)
     */
    const KILO_WATT_HOUR = "kwatth";
    /**
     * Energy: Kilowatt-min (kWm)
     */
    const KILO_WATT_MINUTE = "kwattm";
    /**
     * Energy: Megawatt-hour (MWh)
     */
    const MEGA_WATT_HOUR = "mwatth";
    /**
     * Energy: Ampere-hour (Ah)
     */
    const AMPERE_HOUR = "amph";
    /**
     * Energy: Kiloampere-hour (kAh)
     */
    const KILO_AMPERE_HOUR = "kamph";
    /**
     * Energy: Milliampere-hour (mAh)
     */
    const MILLI_AMPERE_HOUR = "mamph";
    /**
     * Energy: Joule (J)
     */
    const JOULE = "joule";
    /**
     * Energy: Electron volt (eV)
     */
    const ELECTRON_VOLT = "ev";
    /**
     * Energy: Ampere (A)
     */
    const AMPERE = "amp";
    /**
     * Energy: Kiloampere (kA)
     */
    const KILO_AMPERE = "kamp";
    /**
     * Energy: Milliampere (mA)
     */
    const MILLI_AMPERE = "mamp";
    /**
     * Energy: Volt (V)
     */
    const VOLT = "volt";
    /**
     * Energy: Kilovolt (kV)
     */
    const KILO_VOLT = "kvolt";
    /**
     * Energy: Millivolt (mV)
     */
    const MILLI_VOLT = "mvolt";
    /**
     * Energy: Decibel-milliwatt (dBm)
     */
    const DECIBEL_MILLI_WATT = "dBm";
    /**
     * Energy: Milliohm (mΩ)
     */
    const MILLI_OHM = "mohm";
    /**
     * Energy: Ohm (Ω)
     */
    const OHM = "ohm";
    /**
     * Energy: Kiloohm (kΩ)
     */
    const KILO_OHM = "kohm";
    /**
     * Energy: Megaohm (MΩ)
     */
    const MEGA_OHM = "Mohm";
    /**
     * Energy: Farad (F)
     */
    const FARAD = "farad";
    /**
     * Energy: Microfarad (µF)
     */
    const MICRO_FARAD = "watt";
    /**
     * Energy: Nanofarad (nF)
     */
    const NANO_FARAD = "nfarad";
    /**
     * Energy: Picofarad (pF)
     */
    const PICO_FARAD = "pfarad";
    /**
     * Energy: Femtofarad (fF)
     */
    const FEMTO_FARAD = "ffarad";
    /**
     * Energy: Henry (H)
     */
    const HENRY = "henry";
    /**
     * Energy: Millihenry (mH)
     */
    const MILLI_HENRY = "mhenry";
    /**
     * Energy: Microhenry (µH)
     */
    const MICRO_HENRY = "µhenry";
    /**
     * Energy: Lumens (Lm)
     */
    const LUMENS = "lumens";
    /**
     * Flow: Gallons/min (gpm)
     */
    const GALLONS_PER_MINUTE = "flowgpm";
    /**
     * Flow: Cubic meters/sec (cms)
     */
    const CUBIC_METERS_PER_SECOND = "flowcms";
    /**
     * Flow: Cubic feet/sec (cfs)
     */
    const CUBIC_FEET_PER_SECOND = "flowcms";
    /**
     * Flow: Cubic feet/min (cfm)
     */
    const CUBIC_FEET_PER_MINUTE = "flowcms";
    /**
     * Flow: Litre/hour
     */
    const LITRE_PER_HOUR = "litreh";
    /**
     * Flow: Litre/min (L/min)
     */
    const LITRE_PER_MINUTE = "flowlpm";
    /**
     * Flow: milliLitre/min (mL/min)
     */
    const MILLI_LITRE_PER_MINUTE = "flowmlpm";
    /**
     * Flow: Lux (lx)
     */
    const LUX = "lux";
    /**
     * Force: Newton-meters (Nm)
     */
    const NEWTON_METERS = "forceNm";
    /**
     * Force: Kilonewton-meters (kNm)
     */
    const KILO_NEWTON_METERS = "forcekNm";
    /**
     * Force: Newtons (N)
     */
    const NEWTONS = "forceN";
    /**
     * Force: Kilonewtons (kN)
     */
    const KILO_NEWTONS = "forcekN";
    /**
     * Hash rate: hashes/sec
     */
    const HASHES_PER_SECOND = "Hs";
    /**
     * Hash rate: kilohashes/sec
     */
    const KILO_HASHES_PER_SECOND = "KHs";
    /**
     * Hash rate: megahashes/sec
     */
    const MEGA_HASHES_PER_SECOND = "MHs";
    /**
     * Hash rate: gigahashes/sec
     */
    const GIGA_HASHES_PER_SECOND = "GHs";
    /**
     * Hash rate: terahashes/sec
     */
    const TERA_HASHES_PER_SECOND = "THs";
    /**
     * Hash rate: petahashes/sec
     */
    const PETA_HASHES_PER_SECOND = "PHs";
    /**
     * Hash rate: exahashes/sec
     */
    const EXA_HASHES_PER_SECOND = "EHs";
    /**
     * Mass: milligram (mg)
     */
    const MILLIGRAM = "massmg";
    /**
     * Mass: gram (g)
     */
    const GRAM = "massg";
    /**
     * Mass: pound (lb)
     */
    const POUND = "masslb";
    /**
     * Mass: kilogram (kg)
     */
    const KILOGRAM = "masskg";
    /**
     * Mass: metric ton (t)
     */
    const METRIC_TON = "masst";
    /**
     * Length: millimeter (mm)
     */
    const MILLIMETER = "lengthmm";
    /**
     * Length: inch (in)
     */
    const INCH = "lengthin";
    /**
     * Length: feet (ft)
     */
    const FEET = "lengthft";
    /**
     * Length: meter (m)
     */
    const METER = "lengthm";
    /**
     * Length: kilometer (km)
     */
    const KILOMETER = "lengthkm";
    /**
     * Length: mile (mi)
     */
    const MILE = "lengthmi";
    /**
     * Pressure: Millibars
     */
    const MILLIBARS = "pressurembar";
    /**
     * Pressure: Bars
     */
    const BARS = "pressurebar";
    /**
     * Pressure: Kilobars
     */
    const KILOBARS = "pressurekbar";
    /**
     * Pressure: Pascals
     */
    const PASCALS = "pressurepa";
    /**
     * Pressure: Hectopascals
     */
    const HECTOPASCALS = "pressurehpa";
    /**
     * Pressure: Kilopascals
     */
    const KILOPASCALS = "pressurekpa";
    /**
     * Pressure: Inches of mercury
     */
    const INCHES_OF_MERCURY = "pressurehg";
    /**
     * Pressure: PSI
     */
    const PSI = "pressurepsi";
    /**
     * Radiation: Becquerel (Bq)
     */
    const BECQUEREL = "radbq";
    /**
     * Radiation: curie (Ci)
     */
    const CURIE = "radci";
    /**
     * Radiation: Gray (Gy)
     */
    const GRAY = "radgy";
    /**
     * Radiation: rad
     */
    const RAD = "radrad";
    /**
     * Radiation: Sievert (Sv)
     */
    const SIEVERT = "radsv";
    /**
     * Radiation: milliSievert (mSv)
     */
    const MILLISIEVERT = "radmsv";
    /**
     * Radiation: microSievert (µSv)
     */
    const MICROSIEVERT = "radusv";
    /**
     * Radiation: rem
     */
    const REM = "radrem";
    /**
     * Radiation: Exposure (C/kg)
     */
    const EXPOSURE = "radexpckg";
    /**
     * Radiation: roentgen (R)
     */
    const ROENTGEN = "radr";
    /**
     * Radiation: Sievert/hour (Sv/h)
     */
    const SIEVERT_PER_HOUR = "radsvh";
    /**
     * Radiation: milliSievert/hour (mSv/h)
     */
    const MILLISIEVERT_PER_HOUR = "radmsvh";
    /**
     * Radiation: microSievert/hour (µSv/h)
     */
    const MICROSIEVERT_PER_HOUR = "radusvh";
    /**
     * Rotational Speed: Revolutions per minute (rpm)
     */
    const REVOLUTIONS_PER_MINUTE = "rotrpm";
    /**
     * Rotational Speed: Hertz (Hz)
     */
    const HERTZ_ROTATION = "rothz";
    /**
     * Rotational Speed: Kilohertz (kHz)
     */
    const KILOHERTZ_ROTATION = "rotkhz";
    /**
     * Rotational Speed: Megahertz (MHz)
     */
    const MEGAHERTZ_ROTATION = "rotmhz";
    /**
     * Rotational Speed: Gigahertz (GHz)
     */
    const GIGAHERTZ_ROTATION = "rotghz";
    /**
     * Rotational Speed: Radians per second (rad/s)
     */
    const RADIANS_PER_SECOND = "rotrads";
    /**
     * Rotational Speed: Degrees per second (°/s)
     */
    const DEGREES_PER_SECOND = "rotdegs";
    /**
     * Temperature: Celsius (°C)
     */
    const CELSIUS = "celsius";
    /**
     * Temperature: Fahrenheit (°F)
     */
    const FAHRENHEIT = "fahrenheit";
    /**
     * Temperature: Kelvin (K)
     */
    const KELVIN = "kelvin";
    /**
     * Time: Hertz (1/s)
     */
    const HERTZ = "hertz";
    /**
     * Time: nanoseconds (ns)
     */
    const NANOSECONDS = "ns";
    /**
     * Time: microseconds (µs)
     */
    const MICROSECONDS = "µs";
    /**
     * Time: milliseconds (ms)
     */
    const MILLISECONDS = "ms";
    /**
     * Time: seconds (s)
     */
    const SECONDS = "s";
    /**
     * Time: minutes (m)
     */
    const MINUTES = "m";
    /**
     * Time: hours (h)
     */
    const HOURS = "h";
    /**
     * Time: days (d)
     */
    const DAYS = "d";
    /**
     * Time: duration (ms)
     */
    const DURATION_MILLISECONDS = "dtdurationms";
    /**
     * Time: duration (s)
     */
    const DURATION_SECONDS = "dtdurations";
    /**
     * Time: duration (hh:mm:ss)
     */
    const DURATION_IN_HOURS_MINUTES_SECONDS = "dthms";
    /**
     * Time: duration (d hh:mm:ss)
     */
    const DURATION_IN_DAYS_HOURS_MINUTES_SECONDS = "dtdhms";
    /**
     * Time: Timeticks (s/100)
     */
    const TIMETICKS = "timeticks";
    /**
     * Time: clock (ms)
     */
    const CLOCK_MILLISECONDS = "clockms";
    /**
     * Time: clock (s)
     */
    const CLOCK_SECONDS = "clocks";
    /**
     * Throughput: counts/sec (cps)
     */
    const COUNTS_PER_SECOND = "cps";
    /**
     * Throughput: ops/sec (ops)
     */
    const OPS_PER_SECOND = "ops";
    /**
     * Throughput: requests/sec (rps)
     */
    const REQUESTS_PER_SECOND = "reqps";
    /**
     * Throughput: reads/sec (rps)
     */
    const READS_PER_SECOND = "rps";
    /**
     * Throughput: writes/sec (wps)
     */
    const WRITES_PER_SECOND = "wps";
    /**
     * Throughput: I/O ops/sec (iops)
     */
    const IO_OPS_PER_SECOND = "iops";
    /**
     * Throughput: events/sec (eps)
     */
    const EVENTS_PER_SECOND = "eps";
    /**
     * Throughput: messages/sec (mps)
     */
    const MESSAGES_PER_SECOND = "mps";
    /**
     * Throughput: records/sec (rps)
     */
    const RECORDS_PER_SECOND = "recps";
    /**
     * Throughput: rows/sec (rps)
     */
    const ROWS_PER_SECOND = "rowsps";
    /**
     * Throughput: counts/min (cpm)
     */
    const COUNTS_PER_MINUTE = "cpm";
    /**
     * Throughput: ops/min (opm)
     */
    const OPS_PER_MINUTE = "opm";
    /**
     * Throughput: requests/min (rpm)
     */
    const REQUESTS_PER_MINUTE = "reqpm";
    /**
     * Throughput: reads/min (rpm)
     */
    const READS_PER_MINUTE = "rpm";
    /**
     * Throughput: writes/min (wpm)
     */
    const WRITES_PER_MINUTE = "wpm";
    /**
     * Throughput: events/min (epm)
     */
    const EVENTS_PER_MINUTE = "epm";
    /**
     * Throughput: messages/min (mpm)
     */
    const MESSAGES_PER_MINUTE = "mpm";
    /**
     * Throughput: records/min (rpm)
     */
    const RECORDS_PER_MINUTE = "recpm";
    /**
     * Throughput: rows/min (rpm)
     */
    const ROWS_PER_MINUTE = "rowspm";
    /**
     * Velocity: meters/second (m/s)
     */
    const METERS_PER_SECOND = "velocityms";
    /**
     * Velocity: kilometers/hour (km/h)
     */
    const KILOMETERS_PER_HOUR = "velocitykmh";
    /**
     * Velocity: miles/hour (mph)
     */
    const MILES_PER_HOUR = "velocitymph";
    /**
     * Velocity: knot (kn)
     */
    const KNOT = "velocityknot";
    /**
     * Volume: millilitre (mL)
     */
    const MILLILITRE = "mlitre";
    /**
     * Volume: litre (L)
     */
    const LITRE = "litre";
    /**
     * Volume: cubic meter
     */
    const CUBIC_METER = "m3";
    /**
     * Volume: Normal cubic meter
     */
    const NORMAL_CUBIC_METER = "Nm3";
    /**
     * Volume: cubic decimeter
     */
    const CUBIC_DECIMETER = "dm3";
    /**
     * Volume: gallons
     */
    const GALLONS = "gallons";
    /**
     * Boolean: True / False
     */
    const BOOL = "bool";
    /**
     * Boolean: Yes / No
     */
    const BOOL_YES_NO = "bool_yes_no";
    /**
     * Boolean: On / Off
     */
    const BOOL_ON_OFF = "bool_on_off";
}