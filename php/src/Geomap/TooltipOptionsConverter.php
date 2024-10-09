<?php

namespace Grafana\Foundation\Geomap;

final class TooltipOptionsConverter
{
    public static function convert(\Grafana\Foundation\Geomap\TooltipOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Geomap\TooltipOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Geomap\TooltipMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

