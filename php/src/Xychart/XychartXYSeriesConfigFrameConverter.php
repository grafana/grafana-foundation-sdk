<?php

namespace Grafana\Foundation\Xychart;

final class XychartXYSeriesConfigFrameConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XychartXYSeriesConfigFrameBuilder())',
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

