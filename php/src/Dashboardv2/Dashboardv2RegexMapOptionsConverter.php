<?php

namespace Grafana\Foundation\Dashboardv2;

final class Dashboardv2RegexMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\Dashboardv2RegexMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\Dashboardv2RegexMapOptionsBuilder())',
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
        $arg0 = \Grafana\Foundation\Dashboardv2\ValueMappingResultConverter::convert($input->result);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

