<?php

namespace Grafana\Foundation\Dashboardv2;

final class Dashboardv2SpecialValueMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\Dashboardv2SpecialValueMapOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'match(';
        $arg0 ='\Grafana\Foundation\Dashboardv2\SpecialValueMatch::fromValue("'.$input->match.'")';
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

