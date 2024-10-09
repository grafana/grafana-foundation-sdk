<?php

namespace Grafana\Foundation\Elasticsearch;

final class DataqueryConverter
{
    public static function convert(\Grafana\Foundation\Cog\Dataquery $input): string
    {
        assert($input instanceof \Grafana\Foundation\Elasticsearch\Dataquery);
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\DataqueryBuilder())',
        ];
            if ($input->alias !== null && $input->alias !== "") {
    
        
    $buffer = 'alias(';
        $arg0 =\var_export($input->alias, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->query !== null && $input->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timeField !== null && $input->timeField !== "") {
    
        
    $buffer = 'timeField(';
        $arg0 =\var_export($input->timeField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->bucketAggs !== null && count($input->bucketAggs) >= 1) {
    
        
    $buffer = 'bucketAggs(';
        $tmparg0 = [];
        foreach ($input->bucketAggs as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\DateHistogram:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\DateHistogramConverter::convert($arg1);
                $tmpbucketAggsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Histogram:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\HistogramConverter::convert($arg1);
                $tmpbucketAggsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Terms:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\TermsConverter::convert($arg1);
                $tmpbucketAggsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Filters:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\FiltersConverter::convert($arg1);
                $tmpbucketAggsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\GeoHashGrid:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\GeoHashGridConverter::convert($arg1);
                $tmpbucketAggsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Nested:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\NestedConverter::convert($arg1);
                $tmpbucketAggsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpbucketAggsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->metrics !== null && count($input->metrics) >= 1) {
    
        
    $buffer = 'metrics(';
        $tmparg0 = [];
        foreach ($input->metrics as $arg1) {
        switch (true) {
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Count:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\CountConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\MovingAverage:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\MovingAverageConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Derivative:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\DerivativeConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\CumulativeSum:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\CumulativeSumConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\BucketScript:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\BucketScriptConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\SerialDiff:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\SerialDiffConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\RawData:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\RawDataConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\RawDocument:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\RawDocumentConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\UniqueCount:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\UniqueCountConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Percentiles:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\PercentilesConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\ExtendedStats:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\ExtendedStatsConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Min:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\MinConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Max:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\MaxConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Sum:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\SumConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Average:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\AverageConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\MovingFunction:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\MovingFunctionConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Logs:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\LogsConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\Rate:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\RateConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            case $arg1 instanceof \Grafana\Foundation\Elasticsearch\TopMetrics:
                $disjunctionarg1 = \Grafana\Foundation\Elasticsearch\TopMetricsConverter::convert($arg1);
                $tmpmetricsarg1 = $disjunctionarg1;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $tmparg0[] = $tmpmetricsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
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
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\DataSourceRef('.(($input->datasource->type !== null) ? 'type: '.\var_export($input->datasource->type, true).', ' : '').''.(($input->datasource->uid !== null) ? 'uid: '.\var_export($input->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

