<?php

namespace Grafana\Foundation\Testdata;

final class StreamingQueryConverter
{
    public static function convert(\Grafana\Foundation\Testdata\StreamingQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\StreamingQueryBuilder())',
        ];
            if ($input->bands !== null) {
    
        
    $buffer = 'bands(';
        $arg0 =\var_export($input->bands, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'noise(';
        $arg0 =\var_export($input->noise, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'speed(';
        $arg0 =\var_export($input->speed, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'spread(';
        $arg0 =\var_export($input->spread, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Testdata\StreamingQueryType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->url !== null && $input->url !== "") {
    
        
    $buffer = 'url(';
        $arg0 =\var_export($input->url, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

