<?php

namespace Grafana\Foundation\Alerting;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Alerting\Query $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Alerting\QueryBuilder('.\var_export($input->refId, true).'))',
        ];
            if ($input->datasourceUid !== null && $input->datasourceUid !== "") {
    
        
    $buffer = 'datasourceUid(';
        $arg0 =\var_export($input->datasourceUid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->model !== null) {
    
        
    $buffer = 'model(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertDataqueryToCode($input->model, );
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->queryType !== null && $input->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== null && $input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->relativeTimeRange !== null) {
    
        
    $buffer = 'relativeTimeRange(';
        $arg0 ='(new \Grafana\Foundation\Alerting\RelativeTimeRange(from: '.\var_export($input->relativeTimeRange->from, true).',to: '.\var_export($input->relativeTimeRange->to, true).',))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

