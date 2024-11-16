<?php

namespace Grafana\Foundation\Xychart;

final class XYSeriesConfigConverter
{
    public static function convert(\Grafana\Foundation\Xychart\XYSeriesConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Xychart\XYSeriesConfigBuilder())',
        ];
            if ($input->name !== null) {
    
        
    $buffer = 'name(';
        $arg0 = \Grafana\Foundation\Xychart\XychartXYSeriesConfigNameConverter::convert($input->name);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->frame !== null) {
    
        
    $buffer = 'frame(';
        $arg0 = \Grafana\Foundation\Xychart\XychartXYSeriesConfigFrameConverter::convert($input->frame);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->x !== null) {
    
        
    $buffer = 'x(';
        $arg0 = \Grafana\Foundation\Xychart\XychartXYSeriesConfigXConverter::convert($input->x);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->y !== null) {
    
        
    $buffer = 'y(';
        $arg0 = \Grafana\Foundation\Xychart\XychartXYSeriesConfigYConverter::convert($input->y);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->color !== null) {
    
        
    $buffer = 'color(';
        $arg0 = \Grafana\Foundation\Xychart\XychartXYSeriesConfigColorConverter::convert($input->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->size !== null) {
    
        
    $buffer = 'size(';
        $arg0 = \Grafana\Foundation\Xychart\XychartXYSeriesConfigSizeConverter::convert($input->size);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

