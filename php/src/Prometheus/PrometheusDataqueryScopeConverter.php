<?php

namespace Grafana\Foundation\Prometheus;

final class PrometheusDataqueryScopeConverter
{
    public static function convert(\Grafana\Foundation\Prometheus\PrometheusDataqueryScope $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Prometheus\PrometheusDataqueryScopeBuilder())',
        ];
            if ($input->matchers !== "") {
    
        
    $buffer = 'matchers(';
        $arg0 =\var_export($input->matchers, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

