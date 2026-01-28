<?php

namespace Grafana\Foundation\Elasticsearch;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Elasticsearch\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->alias !== null && $input->spec->alias !== "") {
    
        
    $buffer = 'alias(';
        $arg0 =\var_export($input->spec->alias, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->query !== null && $input->spec->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->spec->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->timeField !== null && $input->spec->timeField !== "") {
    
        
    $buffer = 'timeField(';
        $arg0 =\var_export($input->spec->timeField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->bucketAggs !== null && count($input->spec->bucketAggs) >= 1) {
    
        
    $buffer = 'bucketAggs(';
        $tmparg0 = [];
        foreach ($input->spec->bucketAggs as $arg1) {
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->metrics !== null && count($input->spec->metrics) >= 1) {
    
        
    $buffer = 'metrics(';
        $tmparg0 = [];
        foreach ($input->spec->metrics as $arg1) {
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
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Elasticsearch\Dataquery && $input->spec->datasource !== null) {
    
        
    $buffer = 'oldDatasource(';
        $arg0 ='(new \Grafana\Foundation\Common\DataSourceRef('.(($input->spec->datasource->type !== null) ? 'type: '.\var_export($input->spec->datasource->type, true).', ' : '').''.(($input->spec->datasource->uid !== null) ? 'uid: '.\var_export($input->spec->datasource->uid, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

