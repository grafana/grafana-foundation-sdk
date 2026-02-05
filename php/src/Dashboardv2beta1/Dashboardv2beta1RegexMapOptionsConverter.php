<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1RegexMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1RegexMapOptionsBuilder())',
        ];
            if ($input->pattern !== "") {
    
        
    $buffer = 'pattern(';
        $arg0 =\var_export($input->pattern, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'result(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ValueMappingResultConverter::convert($input->result);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

