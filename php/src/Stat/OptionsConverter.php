<?php

namespace Grafana\Foundation\Stat;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Stat\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Stat\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'graphMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueGraphMode::fromValue("'.$input->graphMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'colorMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueColorMode::fromValue("'.$input->colorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'justifyMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueJustifyMode::fromValue("'.$input->justifyMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'reduceOptions(';
        $arg0 = \Grafana\Foundation\Common\ReduceDataOptionsConverter::convert($input->reduceOptions);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->text !== null) {
    
        
    $buffer = 'text(';
        $arg0 = \Grafana\Foundation\Common\VizTextDisplayOptionsConverter::convert($input->text);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'textMode(';
        $arg0 ='\Grafana\Foundation\Common\BigValueTextMode::fromValue("'.$input->textMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'orientation(';
        $arg0 ='\Grafana\Foundation\Common\VizOrientation::fromValue("'.$input->orientation.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

