<?php

namespace Grafana\Foundation\Canvas;

final class BackgroundConfigConverter
{
    public static function convert(\Grafana\Foundation\Canvas\BackgroundConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Canvas\BackgroundConfigBuilder())',
        ];
            if ($input->color !== null) {
    
        
    $buffer = 'color(';
        $arg0 = \Grafana\Foundation\Common\ColorDimensionConfigConverter::convert($input->color);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->image !== null) {
    
        
    $buffer = 'image(';
        $arg0 = \Grafana\Foundation\Common\ResourceDimensionConfigConverter::convert($input->image);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->size !== null) {
    
        
    $buffer = 'size(';
        $arg0 ='\Grafana\Foundation\Canvas\BackgroundImageSize::fromValue("'.$input->size.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

