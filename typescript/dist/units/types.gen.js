"use strict";
// Code generated - EDITING IS FUTILE. DO NOT EDIT.
Object.defineProperty(exports, "__esModule", { value: true });
exports.Pounds = exports.Dollars = exports.MillimolesPerLiter = exports.MilligramsPerDecilitre = exports.GramPerNormalCubicMeter = exports.GramPerCubicMeter = exports.MilligramPerNormalCubicMeter = exports.MilligramPerCubicMeter = exports.MicrogramPerNormalCubicMeter = exports.MicrogramPerCubicMeter = exports.NanogramPerNormalCubicMeter = exports.NanogramPerCubicMeter = exports.PartsPerBillion = exports.PartsPerMillion = exports.YottaFlopsPerSecond = exports.ZettaFlopsPerSecond = exports.ExaFlopsPerSecond = exports.PetaFlopsPerSecond = exports.TeraFlopsPerSecond = exports.GigaFlopsPerSecond = exports.MegaFlopsPerSecond = exports.FlopsPerSecond = exports.Hectares = exports.Acres = exports.SquareMiles = exports.SquareFeet = exports.SquareMeters = exports.ArcSeconds = exports.ArcMinutes = exports.Gradian = exports.Radians = exports.Degrees = exports.GUnit = exports.FeetSecSquared = exports.MetersSecSquared = exports.Pixels = exports.LocaleFormat = exports.ScientificNotation = exports.Hexadecimal = exports.HexadecimalOx = exports.Candela = exports.Decibel = exports.Humidity = exports.PercentUnit = exports.Percent = exports.SiShort = exports.Short = exports.String = exports.Number = exports.NoUnit = void 0;
exports.KilobytesPerSecond = exports.KibibitsPerSecond = exports.KibibytesPerSecond = exports.BitsPerSecondSI = exports.BitsPerSecondIEC = exports.BytesPerSecondSI = exports.BytesPerSecondIEC = exports.PacketsPerSecond = exports.Petabytes = exports.Pebibytes = exports.Terabytes = exports.Tebibytes = exports.Gigabytes = exports.Gibibytes = exports.Megabytes = exports.Mebibytes = exports.Kilobytes = exports.Kibibytes = exports.BitsSI = exports.BitsIEC = exports.BytesSI = exports.BytesIEC = exports.IsraeliNewShekels = exports.UruguayPeso = exports.Guarani = exports.BulgarianLev = exports.CFPFranc = exports.MalaysianRinggit = exports.TurkishLira = exports.VietnameseDong = exports.PhilippinePeso = exports.IndonesianRupiah = exports.SouthKoreanWon = exports.IndianRupee = exports.SouthAfricanRand = exports.MicroBitcoin = exports.MilliBitcoin = exports.Bictoin = exports.PolishZloty = exports.SwissFranc = exports.CzechKoruna = exports.SwedishKrona = exports.NorwegianKrone = exports.IcelandicKrona = exports.DanishKrone = exports.Real = exports.Hryvnias = exports.Rubles = exports.Yen = exports.Euro = void 0;
exports.KiloVolt = exports.Volt = exports.MilliAmpere = exports.KiloAmpere = exports.Ampere = exports.ElectronVolt = exports.Joule = exports.MilliAmpereHour = exports.KiloAmpereHour = exports.AmpereHour = exports.MegaWattHour = exports.KiloWattMinute = exports.KiloWattHour = exports.WattHourPerKilogram = exports.WattHour = exports.KiloVoltAmpereReactive = exports.VoltAmpereReactive = exports.KiloVoltAmpere = exports.VoltAmpere = exports.WattPerSquareMeter = exports.MilliWatt = exports.GigaWatt = exports.MegaWatt = exports.KiloWatt = exports.Watt = exports.DateTimeFromNow = exports.DatetimeDefault = exports.DatetimeLocalNotToday = exports.DatetimeLocal = exports.DatetimeUSNotToday = exports.DatetimeUS = exports.DatetimeISONotToday = exports.DatetimeISO = exports.PetabitsPerSecond = exports.PetabytesPerSecond = exports.PebibitsPerSecond = exports.PebibytesPerSecond = exports.TerabitsPerSecond = exports.TerabytesPerSecond = exports.TebibitsPerSecond = exports.TebibytesPerSecond = exports.GigabitsPerSecond = exports.GigabytesPerSecond = exports.GibibitsPerSecond = exports.GibibytesPerSecond = exports.MegabitsPerSecond = exports.MegabytesPerSecond = exports.MebibitsPerSecond = exports.MebibytesPerSecond = exports.KilobitsPerSecond = void 0;
exports.Hectopascals = exports.Pascals = exports.Kilobars = exports.Bars = exports.Millibars = exports.Mile = exports.Kilometer = exports.Meter = exports.Feet = exports.Inch = exports.Millimeter = exports.MetricTon = exports.Kilogram = exports.Pound = exports.Gram = exports.Milligram = exports.ExaHashesPerSecond = exports.PetaHashesPerSecond = exports.TeraHashesPerSecond = exports.GigaHashesPerSecond = exports.MegaHashesPerSecond = exports.KiloHashesPerSecond = exports.HashesPerSecond = exports.KiloNewtons = exports.Newtons = exports.KiloNewtonMeters = exports.NewtonMeters = exports.Lux = exports.MilliLitrePerMinute = exports.LitrePerMinute = exports.LitrePerHour = exports.CubicFeetPerMinute = exports.CubicFeetPerSecond = exports.CubicMetersPerSecond = exports.GallonsPerMinute = exports.Lumens = exports.MicroHenry = exports.MilliHenry = exports.Henry = exports.FemtoFarad = exports.PicoFarad = exports.NanoFarad = exports.MicroFarad = exports.Farad = exports.MegaOhm = exports.KiloOhm = exports.Ohm = exports.MilliOhm = exports.DecibelMilliWatt = exports.MilliVolt = void 0;
exports.RecordsPerSecond = exports.MessagesPerSecond = exports.EventsPerSecond = exports.IOOpsPerSecond = exports.WritesPerSecond = exports.ReadsPerSecond = exports.RequestsPerSecond = exports.OpsPerSecond = exports.CountsPerSecond = exports.ClockSeconds = exports.ClockMilliseconds = exports.Timeticks = exports.DurationInDaysHoursMinutesSeconds = exports.DurationInHoursMinutesSeconds = exports.DurationSeconds = exports.DurationMilliseconds = exports.Days = exports.Hours = exports.Minutes = exports.Seconds = exports.Milliseconds = exports.Microseconds = exports.Nanoseconds = exports.Hertz = exports.Kelvin = exports.Fahrenheit = exports.Celsius = exports.DegreesPerSecond = exports.RadiansPerSecond = exports.GigahertzRotation = exports.MegahertzRotation = exports.KilohertzRotation = exports.HertzRotation = exports.RevolutionsPerMinute = exports.MicrosievertPerHour = exports.MillisievertPerHour = exports.SievertPerHour = exports.Roentgen = exports.Exposure = exports.Rem = exports.Microsievert = exports.Millisievert = exports.Sievert = exports.Rad = exports.Gray = exports.Curie = exports.Becquerel = exports.PSI = exports.InchesOfMercury = exports.Kilopascals = void 0;
exports.BoolOnOff = exports.BoolYesNo = exports.Bool = exports.Gallons = exports.CubicDecimeter = exports.NormalCubicMeter = exports.CubicMeter = exports.Litre = exports.Millilitre = exports.Knot = exports.MilesPerHour = exports.KilometersPerHour = exports.MetersPerSecond = exports.RowsPerMinute = exports.RecordsPerMinute = exports.MessagesPerMinute = exports.EventsPerMinute = exports.WritesPerMinute = exports.ReadsPerMinute = exports.RequestsPerMinute = exports.OpsPerMinute = exports.CountsPerMinute = exports.RowsPerSecond = void 0;
exports.NoUnit = "none";
exports.Number = "none";
exports.String = "string";
exports.Short = "short";
// SI short
exports.SiShort = "sishort";
// Percent (0-100)
exports.Percent = "percent";
// Percent (0.0-1.0)
exports.PercentUnit = "percentunit";
// Humidity (%H)
exports.Humidity = "humidity";
// Decibel (dB)
exports.Decibel = "dB";
// Candela (cd)
exports.Candela = "candela";
// Hexadecimal (0x)
exports.HexadecimalOx = "hex0x";
exports.Hexadecimal = "hex";
exports.ScientificNotation = "sci";
exports.LocaleFormat = "locale";
// Pixels (px)
exports.Pixels = "pixel";
// Acceleration: Meters/sec² (m/sec²)
exports.MetersSecSquared = "accMS2";
// Acceleration: Feet/sec² (f/sec²)
exports.FeetSecSquared = "accFS2";
// G unit (g)
exports.GUnit = "accG";
// Angle: Degrees (°)
exports.Degrees = "degree";
// Angle: Radians
exports.Radians = "radian";
// Angle: Gradian
exports.Gradian = "grad";
// Angle: Arc Minutes
exports.ArcMinutes = "arcmin";
// Angle: Arc Seconds
exports.ArcSeconds = "arcsec";
// Area: Square Meters (m²)
exports.SquareMeters = "areaM2";
// Area: Square Feet (ft²)
exports.SquareFeet = "areaF2";
// Area: Square Miles (mi²)
exports.SquareMiles = "areaMI2";
// Area: Acres (ac)
exports.Acres = "acres";
// Area: Hectares (ha)
exports.Hectares = "hectares";
// Computation: FLOP/s
exports.FlopsPerSecond = "flops";
// Computation: MFLOP/s
exports.MegaFlopsPerSecond = "mflops";
// Computation: GFLOP/s
exports.GigaFlopsPerSecond = "gflops";
// Computation: TFLOP/s
exports.TeraFlopsPerSecond = "tflops";
// Computation: PFLOP/s
exports.PetaFlopsPerSecond = "pflops";
// Computation: EFLOP/s
exports.ExaFlopsPerSecond = "eflops";
// Computation: ZFLOP/s
exports.ZettaFlopsPerSecond = "zflops";
// Computation: YFLOP/s
exports.YottaFlopsPerSecond = "yflops";
// Concentration: parts-per-million (ppm)
exports.PartsPerMillion = "ppm";
// Concentration: parts-per-billion (ppb)
exports.PartsPerBillion = "conppb";
// Concentration: nanogram per cubic meter (ng/m³)
exports.NanogramPerCubicMeter = "conngm3";
// Concentration: nanogram per normal cubic meter (ng/Nm³)
exports.NanogramPerNormalCubicMeter = "conngNm3";
// Concentration: microgram per cubic meter (μg/m³)
exports.MicrogramPerCubicMeter = "conμgm3";
// Concentration: microgram per normal cubic meter (μg/Nm³)
exports.MicrogramPerNormalCubicMeter = "conμgNm3";
// Concentration: milligram per cubic meter (mg/m³)
exports.MilligramPerCubicMeter = "conmgm3";
// Concentration: milligram per normal cubic meter (mg/Nm³)
exports.MilligramPerNormalCubicMeter = "conmgNm3";
// Concentration: gram per cubic meter (g/m³)
exports.GramPerCubicMeter = "congm3";
// Concentration: gram per normal cubic meter (g/Nm³)
exports.GramPerNormalCubicMeter = "congNm3";
// Concentration: milligrams per decilitre (mg/dL)
exports.MilligramsPerDecilitre = "conmgdL";
// Concentration: millimoles per litre (mmol/L)
exports.MillimolesPerLiter = "conmmolL";
// Currency: Dollars ($)
exports.Dollars = "currencyUSD";
// Currency: Pounds (£)
exports.Pounds = "currencyGBP";
// Currency: Euro (€)
exports.Euro = "currencyEUR";
// Currency: Yen (¥)
exports.Yen = "currencyJPY";
// Currency: Rubles (₽)
exports.Rubles = "currencyRUB";
// Currency: Hryvnias (₴)
exports.Hryvnias = "currencyUAH";
// Currency: Real (R$)
exports.Real = "currencyBRL";
// Currency: Danish Krone (kr)
exports.DanishKrone = "currencyDKK";
// Currency: Icelandic Króna (kr)
exports.IcelandicKrona = "currencyISK";
// Currency: Norwegian Krone (kr)
exports.NorwegianKrone = "currencyNOK";
// Currency: Swedish Krona (kr)
exports.SwedishKrona = "currencySEK";
// Currency: Czech koruna (czk)
exports.CzechKoruna = "currencyCZK";
// Currency: Swiss franc (CHF)
exports.SwissFranc = "currencyCHF";
// Currency: Polish Złoty (PLN)
exports.PolishZloty = "currencyPLN";
// Currency: Bitcoin (฿)
exports.Bictoin = "currencyBTC";
// Currency: Milli Bitcoin (฿)
exports.MilliBitcoin = "currencymBTC";
// Currency: Micro Bitcoin (฿)
exports.MicroBitcoin = "currencyμBTC";
// Currency: South African Rand (R)
exports.SouthAfricanRand = "currencyZAR";
// Currency: Indian Rupee (₹)
exports.IndianRupee = "currencyINR";
// Currency: South Korean Won (₩)
exports.SouthKoreanWon = "currencyKRW";
// Currency: Indonesian Rupiah (Rp)
exports.IndonesianRupiah = "currencyIDR";
// Currency: Philippine Peso (PHP)
exports.PhilippinePeso = "currencyPHP";
// Currency: Vietnamese Dong (VND)
exports.VietnameseDong = "currencyVND";
// Currency: Turkish Lira (₺)
exports.TurkishLira = "currencyTRY";
// Currency: Malaysian Ringgit (RM)
exports.MalaysianRinggit = "currencyMYR";
// Currency: CFP franc (XPF)
exports.CFPFranc = "currencyXPF";
// Currency: Bulgarian Lev (BGN)
exports.BulgarianLev = "currencyBGN";
// Currency: Guaraní (₲)
exports.Guarani = "currencyPYG";
// Currency: Uruguay Peso (UYU)
exports.UruguayPeso = "currencyUYU";
// Currency: Israeli New Shekels (₪)
exports.IsraeliNewShekels = "currencyILS";
// Data: bytes(IEC)
exports.BytesIEC = "bytes";
// Data: bytes(SI)
exports.BytesSI = "decbytes";
// Data: bits(IEC)
exports.BitsIEC = "bits";
// Data: bits(SI)
exports.BitsSI = "bydecbitstes";
// Data: kibibytes
exports.Kibibytes = "kbytes";
// Data: kilobytes
exports.Kilobytes = "deckbytes";
// Data: mebibytes
exports.Mebibytes = "mbytes";
// Data: megabytes
exports.Megabytes = "decmbytes";
// Data: gibibytes
exports.Gibibytes = "gbytes";
// Data: gigabytes
exports.Gigabytes = "decgbytes";
// Data: tebibytes
exports.Tebibytes = "tbytes";
// Data: terabytes
exports.Terabytes = "dectbytes";
// Data: pebibytes
exports.Pebibytes = "pbytes";
// Data: petabytes
exports.Petabytes = "decpbytes";
// Data rate: packets/sec
exports.PacketsPerSecond = "pps";
// Data rate: bytes/sec(IEC)
exports.BytesPerSecondIEC = "binBps";
// Data rate: bytes/sec(SI)
exports.BytesPerSecondSI = "Bps";
// Data rate: bits/sec(IEC)
exports.BitsPerSecondIEC = "binbps";
// Data rate: bits/sec(SI)
exports.BitsPerSecondSI = "bps";
// Data rate: kibibytes/sec
exports.KibibytesPerSecond = "KiBs";
// Data rate: kibibits/sec
exports.KibibitsPerSecond = "Kibits";
// Data rate: kilobytes/sec
exports.KilobytesPerSecond = "KBs";
// Data rate: kilobits/sec
exports.KilobitsPerSecond = "Kbits";
// Data rate: mebibytes/sec
exports.MebibytesPerSecond = "MiBs";
// Data rate: mebibits/sec
exports.MebibitsPerSecond = "Mibits";
// Data rate: megabytes/sec
exports.MegabytesPerSecond = "MBs";
// Data rate: megabits/sec
exports.MegabitsPerSecond = "Mbits";
// Data rate: gibibytes/sec
exports.GibibytesPerSecond = "GiBs";
// Data rate: gibibits/sec
exports.GibibitsPerSecond = "Gibits";
// Data rate: gigabytes/sec
exports.GigabytesPerSecond = "GBs";
// Data rate: gigabits/sec
exports.GigabitsPerSecond = "Gbits";
// Data rate: tebibytes/sec
exports.TebibytesPerSecond = "TiBs";
// Data rate: tebibits/sec
exports.TebibitsPerSecond = "Tibits";
// Data rate: terabytes/sec
exports.TerabytesPerSecond = "TBs";
// Data rate: terabits/sec
exports.TerabitsPerSecond = "Tbits";
// Data rate: pebibytes/sec
exports.PebibytesPerSecond = "PiBs";
// Data rate: pebibits/sec
exports.PebibitsPerSecond = "Pibits";
// Data rate: petabytes/sec
exports.PetabytesPerSecond = "PBs";
// Data rate: petabits/sec
exports.PetabitsPerSecond = "Pbits";
// Date & time: Datetime ISO
exports.DatetimeISO = "dateTimeAsIso";
// Date & time: Datetime ISO (No date if today)
exports.DatetimeISONotToday = "dateTimeAsIsoNoDateIfToday";
// Date & time: Datetime US
exports.DatetimeUS = "dateTimeAsUS";
// Date & time: Datetime US (No date if today)
exports.DatetimeUSNotToday = "dateTimeAsUSNoDateIfToday";
// Date & time: Datetime local
exports.DatetimeLocal = "dateTimeAsLocal";
// Date & time: Datetime local (No date if today)
exports.DatetimeLocalNotToday = "dateTimeAsLocalNoDateIfToday";
// Date & time: Datetime default
exports.DatetimeDefault = "dateTimeAsSystem";
// Date & time: From Now
exports.DateTimeFromNow = "dateTimeFromNow";
// Energy: Watt (W)
exports.Watt = "watt";
// Energy: Kilowatt (kW)
exports.KiloWatt = "kwatt";
// Energy: Megawatt (MW)
exports.MegaWatt = "megwatt";
// Energy: Gigawatt (GW)
exports.GigaWatt = "gwatt";
// Energy: Milliwatt (mW)
exports.MilliWatt = "mwatt";
// Energy: Watt per square meter (W/m²)
exports.WattPerSquareMeter = "Wm2";
// Energy: Volt-Ampere (VA)
exports.VoltAmpere = "voltamp";
// Energy: Kilovolt-Ampere (kVA)
exports.KiloVoltAmpere = "kvoltamp";
// Energy: Volt-Ampere reactive (VAr)
exports.VoltAmpereReactive = "voltampreact";
// Energy: Kilovolt-Ampere reactive (kVAr)
exports.KiloVoltAmpereReactive = "kvoltampreact";
// Energy: Watt-hour (Wh)
exports.WattHour = "watth";
// Energy: Watt-hour per Kilogram (Wh/kg)
exports.WattHourPerKilogram = "watthperkg";
// Energy: Kilowatt-hour (kWh)
exports.KiloWattHour = "kwatth";
// Energy: Kilowatt-min (kWm)
exports.KiloWattMinute = "kwattm";
// Energy: Megawatt-hour (MWh)
exports.MegaWattHour = "mwatth";
// Energy: Ampere-hour (Ah)
exports.AmpereHour = "amph";
// Energy: Kiloampere-hour (kAh)
exports.KiloAmpereHour = "kamph";
// Energy: Milliampere-hour (mAh)
exports.MilliAmpereHour = "mamph";
// Energy: Joule (J)
exports.Joule = "joule";
// Energy: Electron volt (eV)
exports.ElectronVolt = "ev";
// Energy: Ampere (A)
exports.Ampere = "amp";
// Energy: Kiloampere (kA)
exports.KiloAmpere = "kamp";
// Energy: Milliampere (mA)
exports.MilliAmpere = "mamp";
// Energy: Volt (V)
exports.Volt = "volt";
// Energy: Kilovolt (kV)
exports.KiloVolt = "kvolt";
// Energy: Millivolt (mV)
exports.MilliVolt = "mvolt";
// Energy: Decibel-milliwatt (dBm)
exports.DecibelMilliWatt = "dBm";
// Energy: Milliohm (mΩ)
exports.MilliOhm = "mohm";
// Energy: Ohm (Ω)
exports.Ohm = "ohm";
// Energy: Kiloohm (kΩ)
exports.KiloOhm = "kohm";
// Energy: Megaohm (MΩ)
exports.MegaOhm = "Mohm";
// Energy: Farad (F)
exports.Farad = "farad";
// Energy: Microfarad (µF)
exports.MicroFarad = "watt";
// Energy: Nanofarad (nF)
exports.NanoFarad = "nfarad";
// Energy: Picofarad (pF)
exports.PicoFarad = "pfarad";
// Energy: Femtofarad (fF)
exports.FemtoFarad = "ffarad";
// Energy: Henry (H)
exports.Henry = "henry";
// Energy: Millihenry (mH)
exports.MilliHenry = "mhenry";
// Energy: Microhenry (µH)
exports.MicroHenry = "µhenry";
// Energy: Lumens (Lm)
exports.Lumens = "lumens";
// Flow: Gallons/min (gpm)
exports.GallonsPerMinute = "flowgpm";
// Flow: Cubic meters/sec (cms)
exports.CubicMetersPerSecond = "flowcms";
// Flow: Cubic feet/sec (cfs)
exports.CubicFeetPerSecond = "flowcms";
// Flow: Cubic feet/min (cfm)
exports.CubicFeetPerMinute = "flowcms";
// Flow: Litre/hour
exports.LitrePerHour = "litreh";
// Flow: Litre/min (L/min)
exports.LitrePerMinute = "flowlpm";
// Flow: milliLitre/min (mL/min)
exports.MilliLitrePerMinute = "flowmlpm";
// Flow: Lux (lx)
exports.Lux = "lux";
// Force: Newton-meters (Nm)
exports.NewtonMeters = "forceNm";
// Force: Kilonewton-meters (kNm)
exports.KiloNewtonMeters = "forcekNm";
// Force: Newtons (N)
exports.Newtons = "forceN";
// Force: Kilonewtons (kN)
exports.KiloNewtons = "forcekN";
// Hash rate: hashes/sec
exports.HashesPerSecond = "Hs";
// Hash rate: kilohashes/sec
exports.KiloHashesPerSecond = "KHs";
// Hash rate: megahashes/sec
exports.MegaHashesPerSecond = "MHs";
// Hash rate: gigahashes/sec
exports.GigaHashesPerSecond = "GHs";
// Hash rate: terahashes/sec
exports.TeraHashesPerSecond = "THs";
// Hash rate: petahashes/sec
exports.PetaHashesPerSecond = "PHs";
// Hash rate: exahashes/sec
exports.ExaHashesPerSecond = "EHs";
// Mass: milligram (mg)
exports.Milligram = "massmg";
// Mass: gram (g)
exports.Gram = "massg";
// Mass: pound (lb)
exports.Pound = "masslb";
// Mass: kilogram (kg)
exports.Kilogram = "masskg";
// Mass: metric ton (t)
exports.MetricTon = "masst";
// Length: millimeter (mm)
exports.Millimeter = "lengthmm";
// Length: inch (in)
exports.Inch = "lengthin";
// Length: feet (ft)
exports.Feet = "lengthft";
// Length: meter (m)
exports.Meter = "lengthm";
// Length: kilometer (km)
exports.Kilometer = "lengthkm";
// Length: mile (mi)
exports.Mile = "lengthmi";
// Pressure: Millibars
exports.Millibars = "pressurembar";
// Pressure: Bars
exports.Bars = "pressurebar";
// Pressure: Kilobars
exports.Kilobars = "pressurekbar";
// Pressure: Pascals
exports.Pascals = "pressurepa";
// Pressure: Hectopascals
exports.Hectopascals = "pressurehpa";
// Pressure: Kilopascals
exports.Kilopascals = "pressurekpa";
// Pressure: Inches of mercury
exports.InchesOfMercury = "pressurehg";
// Pressure: PSI
exports.PSI = "pressurepsi";
// Radiation: Becquerel (Bq)
exports.Becquerel = "radbq";
// Radiation: curie (Ci)
exports.Curie = "radci";
// Radiation: Gray (Gy)
exports.Gray = "radgy";
// Radiation: rad
exports.Rad = "radrad";
// Radiation: Sievert (Sv)
exports.Sievert = "radsv";
// Radiation: milliSievert (mSv)
exports.Millisievert = "radmsv";
// Radiation: microSievert (µSv)
exports.Microsievert = "radusv";
// Radiation: rem
exports.Rem = "radrem";
// Radiation: Exposure (C/kg)
exports.Exposure = "radexpckg";
// Radiation: roentgen (R)
exports.Roentgen = "radr";
// Radiation: Sievert/hour (Sv/h)
exports.SievertPerHour = "radsvh";
// Radiation: milliSievert/hour (mSv/h)
exports.MillisievertPerHour = "radmsvh";
// Radiation: microSievert/hour (µSv/h)
exports.MicrosievertPerHour = "radusvh";
// Rotational Speed: Revolutions per minute (rpm)
exports.RevolutionsPerMinute = "rotrpm";
// Rotational Speed: Hertz (Hz)
exports.HertzRotation = "rothz";
// Rotational Speed: Kilohertz (kHz)
exports.KilohertzRotation = "rotkhz";
// Rotational Speed: Megahertz (MHz)
exports.MegahertzRotation = "rotmhz";
// Rotational Speed: Gigahertz (GHz)
exports.GigahertzRotation = "rotghz";
// Rotational Speed: Radians per second (rad/s)
exports.RadiansPerSecond = "rotrads";
// Rotational Speed: Degrees per second (°/s)
exports.DegreesPerSecond = "rotdegs";
// Temperature: Celsius (°C)
exports.Celsius = "celsius";
// Temperature: Fahrenheit (°F)
exports.Fahrenheit = "fahrenheit";
// Temperature: Kelvin (K)
exports.Kelvin = "kelvin";
// Time: Hertz (1/s)
exports.Hertz = "hertz";
// Time: nanoseconds (ns)
exports.Nanoseconds = "ns";
// Time: microseconds (µs)
exports.Microseconds = "µs";
// Time: milliseconds (ms)
exports.Milliseconds = "ms";
// Time: seconds (s)
exports.Seconds = "s";
// Time: minutes (m)
exports.Minutes = "m";
// Time: hours (h)
exports.Hours = "h";
// Time: days (d)
exports.Days = "d";
// Time: duration (ms)
exports.DurationMilliseconds = "dtdurationms";
// Time: duration (s)
exports.DurationSeconds = "dtdurations";
// Time: duration (hh:mm:ss)
exports.DurationInHoursMinutesSeconds = "dthms";
// Time: duration (d hh:mm:ss)
exports.DurationInDaysHoursMinutesSeconds = "dtdhms";
// Time: Timeticks (s/100)
exports.Timeticks = "timeticks";
// Time: clock (ms)
exports.ClockMilliseconds = "clockms";
// Time: clock (s)
exports.ClockSeconds = "clocks";
// Throughput: counts/sec (cps)
exports.CountsPerSecond = "cps";
// Throughput: ops/sec (ops)
exports.OpsPerSecond = "ops";
// Throughput: requests/sec (rps)
exports.RequestsPerSecond = "reqps";
// Throughput: reads/sec (rps)
exports.ReadsPerSecond = "rps";
// Throughput: writes/sec (wps)
exports.WritesPerSecond = "wps";
// Throughput: I/O ops/sec (iops)
exports.IOOpsPerSecond = "iops";
// Throughput: events/sec (eps)
exports.EventsPerSecond = "eps";
// Throughput: messages/sec (mps)
exports.MessagesPerSecond = "mps";
// Throughput: records/sec (rps)
exports.RecordsPerSecond = "recps";
// Throughput: rows/sec (rps)
exports.RowsPerSecond = "rowsps";
// Throughput: counts/min (cpm)
exports.CountsPerMinute = "cpm";
// Throughput: ops/min (opm)
exports.OpsPerMinute = "opm";
// Throughput: requests/min (rpm)
exports.RequestsPerMinute = "reqpm";
// Throughput: reads/min (rpm)
exports.ReadsPerMinute = "rpm";
// Throughput: writes/min (wpm)
exports.WritesPerMinute = "wpm";
// Throughput: events/min (epm)
exports.EventsPerMinute = "epm";
// Throughput: messages/min (mpm)
exports.MessagesPerMinute = "mpm";
// Throughput: records/min (rpm)
exports.RecordsPerMinute = "recpm";
// Throughput: rows/min (rpm)
exports.RowsPerMinute = "rowspm";
// Velocity: meters/second (m/s)
exports.MetersPerSecond = "velocityms";
// Velocity: kilometers/hour (km/h)
exports.KilometersPerHour = "velocitykmh";
// Velocity: miles/hour (mph)
exports.MilesPerHour = "velocitymph";
// Velocity: knot (kn)
exports.Knot = "velocityknot";
// Volume: millilitre (mL)
exports.Millilitre = "mlitre";
// Volume: litre (L)
exports.Litre = "litre";
// Volume: cubic meter
exports.CubicMeter = "m3";
// Volume: Normal cubic meter
exports.NormalCubicMeter = "Nm3";
// Volume: cubic decimeter
exports.CubicDecimeter = "dm3";
// Volume: gallons
exports.Gallons = "gallons";
// Boolean: True / False
exports.Bool = "bool";
// Boolean: Yes / No
exports.BoolYesNo = "bool_yes_no";
// Boolean: On / Off
exports.BoolOnOff = "bool_on_off";
//# sourceMappingURL=types.gen.js.map