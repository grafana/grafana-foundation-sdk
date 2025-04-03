<?php

namespace Grafana\Foundation\Common;

final class OptionsWithTimezonesConverter
{
    public static function convert(\Grafana\Foundation\Common\OptionsWithTimezones $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\OptionsWithTimezonesBuilder())',
        ];
            if ($input->timezone !== null && count($input->timezone) >= 1) {
    
        
    $buffer = 'timezone(';
        $tmparg0 = [];
        foreach ($input->timezone as $arg1) {
        $tmptimezonearg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptimezonearg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

