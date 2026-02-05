<?php

namespace Grafana\Foundation\Xychart;

final class XychartXYSeriesConfigSizeConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XychartXYSeriesConfigSize $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XychartXYSeriesConfigSizeBuilder())',
        ];
            
    
        {
    $buffer = 'matcher(';
        $arg0 = \Grafana\Foundation\Xychart\MatcherConfigConverter::convert($input->matcher);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

