<?php

namespace Grafana\Foundation\Common;

final class LineStyleConverter
{
    public static function convert(\Grafana\Foundation\Common\LineStyle $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\LineStyleBuilder())',
        ];
            if ($input->fill !== null) {
    
        
    $buffer = 'fill(';
        $arg0 ='\Grafana\Foundation\Common\LineStyleFill::fromValue("'.$input->fill.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->dash !== null && count($input->dash) >= 1) {
    
        
    $buffer = 'dash(';
        $tmparg0 = [];
        foreach ($input->dash as $arg1) {
        $tmpdasharg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpdasharg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

