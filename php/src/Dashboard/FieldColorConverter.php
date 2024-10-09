<?php

namespace Grafana\Foundation\Dashboard;

final class FieldColorConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\FieldColor $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\FieldColorBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Dashboard\FieldColorModeId::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->fixedColor !== null && $input->fixedColor !== "") {
    
        
    $buffer = 'fixedColor(';
        $arg0 =\var_export($input->fixedColor, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->seriesBy !== null) {
    
        
    $buffer = 'seriesBy(';
        $arg0 ='\Grafana\Foundation\Dashboard\FieldColorSeriesByMode::fromValue("'.$input->seriesBy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

