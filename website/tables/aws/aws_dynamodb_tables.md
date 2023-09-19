# Table: aws_dynamodb_tables

This table shows data for Amazon DynamoDB Tables.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_dynamodb_tables:
  - [aws_dynamodb_table_continuous_backups](aws_dynamodb_table_continuous_backups)
  - [aws_dynamodb_table_replica_auto_scalings](aws_dynamodb_table_replica_auto_scalings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|archival_summary|`json`|
|attribute_definitions|`json`|
|billing_mode_summary|`json`|
|creation_date_time|`timestamp[us, tz=UTC]`|
|deletion_protection_enabled|`bool`|
|global_secondary_indexes|`json`|
|global_table_version|`utf8`|
|item_count|`int64`|
|key_schema|`json`|
|latest_stream_arn|`utf8`|
|latest_stream_label|`utf8`|
|local_secondary_indexes|`json`|
|provisioned_throughput|`json`|
|replicas|`json`|
|restore_summary|`json`|
|sse_description|`json`|
|stream_specification|`json`|
|table_arn|`utf8`|
|table_class_summary|`json`|
|table_id|`utf8`|
|table_name|`utf8`|
|table_size_bytes|`int64`|
|table_status|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### DynamoDB tables should automatically scale capacity with demand

```sql
WITH
  dynamodb_tables
    AS (
      SELECT
        account_id,
        arn,
        table_name,
        billing_mode_summary->>'BillingMode' IS DISTINCT FROM 'PAY_PER_REQUEST'
          AS is_not_on_demand
      FROM
        aws_dynamodb_tables
    ),
  replica_auto_scalings
    AS (
      SELECT
        s.table_arn,
        s.region,
        s.region_name,
        (
          s.replica_provisioned_read_capacity_auto_scaling_settings->>'AutoScalingDisabled'
        )::BOOL
          AS read_auto_scaling_disabled,
        (
          s.replica_provisioned_write_capacity_auto_scaling_settings->>'AutoScalingDisabled'
        )::BOOL
          AS write_auto_scaling_disabled
      FROM
        aws_dynamodb_table_replica_auto_scalings AS s
    ),
  auto_scaling_disabled_count_in_replica_a_s
    AS (
      SELECT
        table_arn,
        region,
        region_name,
        CASE
        WHEN (
          read_auto_scaling_disabled IS NOT NULL
          AND read_auto_scaling_disabled IS true
        )
        THEN 1
        ELSE 0
        END
          AS read_auto_scaling_disabled_count,
        CASE
        WHEN (
          write_auto_scaling_disabled IS NOT NULL
          AND write_auto_scaling_disabled IS true
        )
        THEN 1
        ELSE 0
        END
          AS write_auto_scaling_disabled_count
      FROM
        replica_auto_scalings
    ),
  sum_of_auto_scaling_disabled_count_in_replica_a_s
    AS (
      SELECT
        table_arn,
        sum(
          read_auto_scaling_disabled_count + write_auto_scaling_disabled_count
        )
          AS auto_scaling_disabled_count
      FROM
        auto_scaling_disabled_count_in_replica_a_s
      GROUP BY
        table_arn
    ),
  dynamodb_tables_with_replica_auto_scaling_disabled_count
    AS (
      SELECT
        account_id,
        arn,
        table_name,
        is_not_on_demand,
        CASE
        WHEN (auto_scaling_disabled_count IS NULL) THEN 0
        ELSE auto_scaling_disabled_count
        END
          AS replica_auto_scaling_disabled_count
      FROM
        dynamodb_tables AS t1
        LEFT JOIN sum_of_auto_scaling_disabled_count_in_replica_a_s AS t2 ON
            t1.arn = t2.table_arn
    ),
  policies_r
    AS (
      SELECT
        resource_id
      FROM
        aws_applicationautoscaling_policies
      WHERE
        service_namespace = 'dynamodb'
        AND policy_type = 'TargetTrackingScaling'
        AND scalable_dimension = 'dynamodb:table:ReadCapacityUnits'
      GROUP BY
        resource_id
    ),
  policies_w
    AS (
      SELECT
        resource_id
      FROM
        aws_applicationautoscaling_policies
      WHERE
        service_namespace = 'dynamodb'
        AND policy_type = 'TargetTrackingScaling'
        AND scalable_dimension = 'dynamodb:table:WriteCapacityUnits'
      GROUP BY
        resource_id
    )
SELECT
  'DynamoDB tables should automatically scale capacity with demand' AS title,
  t.account_id,
  t.arn AS resource_id,
  CASE
  WHEN (
    t.is_not_on_demand IS true
    AND t.replica_auto_scaling_disabled_count > 0
  )
  THEN 'fail'
  WHEN (
    t.is_not_on_demand IS true
    AND t.replica_auto_scaling_disabled_count = 0
    AND (pr.resource_id IS NULL OR pw.resource_id IS NULL)
  )
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  dynamodb_tables_with_replica_auto_scaling_disabled_count AS t
  LEFT JOIN policies_r AS pr ON pr.resource_id = concat('table/', t.table_name)
  LEFT JOIN policies_w AS pw ON pw.resource_id = concat('table/', t.table_name);
```

### DynamoDB tables should have point-in-time recovery enabled

```sql
SELECT
  'DynamoDB tables should have point-in-time recovery enabled' AS title,
  t.account_id,
  t.arn AS resource_id,
  CASE
  WHEN b.point_in_time_recovery_description->>'PointInTimeRecoveryStatus'
  IS DISTINCT FROM 'ENABLED'
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_dynamodb_tables AS t
  LEFT JOIN aws_dynamodb_table_continuous_backups AS b ON b.table_arn = t.arn;
```

### DynamoDB table with no items

```sql
SELECT
  'DynamoDB table with no items' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_dynamodb_tables
WHERE
  item_count = 0;
```


