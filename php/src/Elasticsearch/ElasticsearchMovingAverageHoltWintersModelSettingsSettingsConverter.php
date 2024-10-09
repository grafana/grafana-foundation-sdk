<?php

namespace Grafana\Foundation\Elasticsearch;

final class ElasticsearchMovingAverageHoltWintersModelSettingsSettingsConverter
{
    public static function convert(\Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettings $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder())',
        ];
            if ($input->alpha !== null && $input->alpha !== "") {
    
        
    $buffer = 'alpha(';
        $arg0 =\var_export($input->alpha, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->beta !== null && $input->beta !== "") {
    
        
    $buffer = 'beta(';
        $arg0 =\var_export($input->beta, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gamma !== null && $input->gamma !== "") {
    
        
    $buffer = 'gamma(';
        $arg0 =\var_export($input->gamma, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->period !== null && $input->period !== "") {
    
        
    $buffer = 'period(';
        $arg0 =\var_export($input->period, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->pad !== null) {
    
        
    $buffer = 'pad(';
        $arg0 =\var_export($input->pad, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

