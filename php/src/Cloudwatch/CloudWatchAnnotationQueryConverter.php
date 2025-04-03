<?php

namespace Grafana\Foundation\Cloudwatch;

final class CloudWatchAnnotationQueryConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\CloudWatchAnnotationQuery $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\CloudWatchAnnotationQueryBuilder())',
        ];
            
    
        {
    $buffer = 'queryMode(';
        $arg0 ='\Grafana\Foundation\Cloudwatch\CloudWatchQueryMode::fromValue("'.$input->queryMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->prefixMatching !== null) {
    
        
    $buffer = 'prefixMatching(';
        $arg0 =\var_export($input->prefixMatching, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->actionPrefix !== null && $input->actionPrefix !== "") {
    
        
    $buffer = 'actionPrefix(';
        $arg0 =\var_export($input->actionPrefix, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->hide, true);
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
            if ($input->region !== "") {
    
        
    $buffer = 'region(';
        $arg0 =\var_export($input->region, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->namespace !== "") {
    
        
    $buffer = 'namespace(';
        $arg0 =\var_export($input->namespace, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metricName !== null && $input->metricName !== "") {
    
        
    $buffer = 'metricName(';
        $arg0 =\var_export($input->metricName, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'dimensions(';
        $arg0 = "[";
        foreach ($input->dimensions as $key => $arg1) {
            switch (true) {
            case is_string($arg1):
                $disjunctionarg1 =\var_export($arg1, true);
                $tmpdimensionsarg1 = $disjunctionarg1;
                break;
            case is_array($arg1):
                $tmpdisjunctionarg1 = [];
        foreach ($arg1 as $arg1Value) {
        $tmparg1arg1Value =\var_export($arg1Value, true);
        $tmpdisjunctionarg1[] = $tmparg1arg1Value;
        }
        $disjunctionarg1 = "[" . implode(", \n", $tmpdisjunctionarg1) . "]";
                $tmpdimensionsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
            $arg0 .= "\t".var_export($key, true)." => $tmpdimensionsarg1,";
        }
        $arg0 .= "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->matchExact !== null) {
    
        
    $buffer = 'matchExact(';
        $arg0 =\var_export($input->matchExact, true);
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
            if ($input->accountId !== null && $input->accountId !== "") {
    
        
    $buffer = 'accountId(';
        $arg0 =\var_export($input->accountId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->statistic !== null && $input->statistic !== "") {
    
        
    $buffer = 'statistic(';
        $arg0 =\var_export($input->statistic, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->alarmNamePrefix !== null && $input->alarmNamePrefix !== "") {
    
        
    $buffer = 'alarmNamePrefix(';
        $arg0 =\var_export($input->alarmNamePrefix, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->statistics !== null && count($input->statistics) >= 1) {
    
        
    $buffer = 'statistics(';
        $tmparg0 = [];
        foreach ($input->statistics as $arg1) {
        $tmpstatisticsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpstatisticsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

