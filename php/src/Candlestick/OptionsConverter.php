<?php

namespace Grafana\Foundation\Candlestick;

final class OptionsConverter
{
    public static function convert(\Grafana\Foundation\Candlestick\Options $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Candlestick\OptionsBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Candlestick\VizDisplayMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'candleStyle(';
        $arg0 ='\Grafana\Foundation\Candlestick\CandleStyle::fromValue("'.$input->candleStyle.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'colorStrategy(';
        $arg0 ='\Grafana\Foundation\Candlestick\ColorStrategy::fromValue("'.$input->colorStrategy.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'fields(';
        $arg0 = \Grafana\Foundation\Candlestick\CandlestickFieldMapConverter::convert($input->fields);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'colors(';
        $arg0 = \Grafana\Foundation\Candlestick\CandlestickColorsConverter::convert($input->colors);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'legend(';
        $arg0 = \Grafana\Foundation\Common\VizLegendOptionsConverter::convert($input->legend);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->includeAllFields !== null && $input->includeAllFields !== false) {
    
        
    $buffer = 'includeAllFields(';
        $arg0 =\var_export($input->includeAllFields, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

