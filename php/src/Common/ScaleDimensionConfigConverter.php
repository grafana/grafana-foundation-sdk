<?php

namespace Grafana\Foundation\Common;

final class ScaleDimensionConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\ScaleDimensionConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\ScaleDimensionConfigBuilder())',
        ];
            
    
        {
    $buffer = 'min(';
        $arg0 =\var_export($input->min, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'max(';
        $arg0 =\var_export($input->max, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->fixed !== null) {
    
        
    $buffer = 'fixed(';
        $arg0 =\var_export($input->fixed, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->mode !== null) {
    
        
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Common\ScaleDimensionMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

