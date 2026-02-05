<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class RepeatOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\RepeatOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\RepeatOptionsBuilder())',
        ];
            if ($input->value !== "") {
    
        
    $buffer = 'value(';
        $arg0 =\var_export($input->value, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->direction !== null) {
    
        
    $buffer = 'direction(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\RepeatOptionsDirection::fromValue("'.$input->direction.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->maxPerRow !== null) {
    
        
    $buffer = 'maxPerRow(';
        $arg0 =\var_export($input->maxPerRow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

