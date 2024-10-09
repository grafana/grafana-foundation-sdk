<?php

namespace Grafana\Foundation\Common;

final class OptionsWithLegendConverter
{
    public static function convert(\Grafana\Foundation\Common\OptionsWithLegend $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\OptionsWithLegendBuilder())',
        ];
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

