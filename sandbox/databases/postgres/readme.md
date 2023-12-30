# PostgreSQL

```python
# Jupyter require additional python modules in order to make sql magick working...
#> pip install "pgspecial<2" -q
#> pip install ipython-sql pgspecial psycopg2 -q
```

```python
%load_ext sql
```

Northwind db can be downloaded [here](https://raw.githubusercontent.com/sjh37/EntityFramework-Reverse-POCO-Code-First-Generator/master/TestDatabases/PostgreSQL/northwind.sql) or anywhere else... its northwind after all...

```python
#> curl -s https://gist.githubusercontent.com/butuzov/ef20f47c020ecdbbde19090fbd9e8fc7/raw/74983701f1ead70f053e7590e754b7528e4aaa27/northwind.sql --output _tmp_northwind2.sql
#> docker-compose up -d

> [1A[1B[0G[?25l[+] Running 1/0
> [32m✔[0m Container postgres-postgresql-1  [32mR...[0m                                   [34m0.0s [0m
> [?25h
```

```sql
postgresql://postgres:postgres@localhost:5432/northwind
```

## `psql`

`docker exec -it postgres-postgresql-1 psql -U postgres -h localhost -d northwind`

<div class="alert alert-warning">WARNING: 
    Not all commands are supported via <code>jupyter</code> / <code>pgspecial</code></div>


```sql
\d

> 14 rows affected.
```

<table>
    <thead>
        <tr>
            <th>Schema</th>
            <th>Name</th>
            <th>Type</th>
            <th>Owner</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>public</td>
            <td>categories</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customer_customer_demo</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customer_demographics</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customers</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>employee_territories</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>employees</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>order_details</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>orders</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>products</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>region</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>shippers</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>suppliers</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>territories</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>us_states</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
    </tbody>
</table>

Some other examples of [psql](https://www.postgresql.org/docs/current/app-psql.html)

- `\t` - tuples only
- `\a` - no extra breaking elements (added by psql)
- `\q query.sql` - rediect to sql.
- `\o`  stop redirection.
- `\i query.sql` - execute generated script
- `\du` - users
- etc...

### psql Customizations

```
\pset null 'NULL'
\encoding latin1
\set PROMPT1 '%n@%M:%>%x %/# '
\pset pager always
\timing on
\set qstats92 '
    SELECT usename, datname, left(query,100) || ''...'' as  query
    FROM pg_stat_activity WHERE state != ''idle'' ;
```

 - `\timing` setting timing for queryies on
 - `\set AUTOCOMMIT off`

### Describe & List 

```sql
\dt

> 14 rows affected.
```

<table>
    <thead>
        <tr>
            <th>Schema</th>
            <th>Name</th>
            <th>Type</th>
            <th>Owner</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>public</td>
            <td>categories</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customer_customer_demo</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customer_demographics</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customers</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>employee_territories</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>employees</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>order_details</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>orders</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>products</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>region</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>shippers</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>suppliers</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>territories</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>public</td>
            <td>us_states</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
    </tbody>
</table>

```sql
\dt pg_catalog.pg_t*

> 9 rows affected.
```

<table>
    <thead>
        <tr>
            <th>Schema</th>
            <th>Name</th>
            <th>Type</th>
            <th>Owner</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>pg_catalog</td>
            <td>pg_tablespace</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_transform</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_trigger</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_ts_config</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_ts_config_map</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_ts_dict</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_ts_parser</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_ts_template</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
        <tr>
            <td>pg_catalog</td>
            <td>pg_type</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
    </tbody>
</table>

```sql
# db list
 \l

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Owner</th>
            <th>Encoding</th>
            <th>Collate</th>
            <th>Ctype</th>
            <th>Access privileges</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>northwind</td>
            <td>postgres</td>
            <td>UTF8</td>
            <td>en_US.utf8</td>
            <td>en_US.utf8</td>
            <td>None</td>
        </tr>
        <tr>
            <td>postgres</td>
            <td>postgres</td>
            <td>UTF8</td>
            <td>en_US.utf8</td>
            <td>en_US.utf8</td>
            <td>None</td>
        </tr>
        <tr>
            <td>template0</td>
            <td>postgres</td>
            <td>UTF8</td>
            <td>en_US.utf8</td>
            <td>en_US.utf8</td>
            <td>=c/postgres<br>postgres=CTc/postgres</td>
        </tr>
        <tr>
            <td>template1</td>
            <td>postgres</td>
            <td>UTF8</td>
            <td>en_US.utf8</td>
            <td>en_US.utf8</td>
            <td>=c/postgres<br>postgres=CTc/postgres</td>
        </tr>
    </tbody>
</table>

```sql
\d+

> 14 rows affected.
```

<table>
    <thead>
        <tr>
            <th>Schema</th>
            <th>Name</th>
            <th>Type</th>
            <th>Owner</th>
            <th>Size</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>public</td>
            <td>categories</td>
            <td>table</td>
            <td>postgres</td>
            <td>16 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customer_customer_demo</td>
            <td>table</td>
            <td>postgres</td>
            <td>8192 bytes</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customer_demographics</td>
            <td>table</td>
            <td>postgres</td>
            <td>8192 bytes</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>customers</td>
            <td>table</td>
            <td>postgres</td>
            <td>48 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>employee_territories</td>
            <td>table</td>
            <td>postgres</td>
            <td>8192 bytes</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>employees</td>
            <td>table</td>
            <td>postgres</td>
            <td>16 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>order_details</td>
            <td>table</td>
            <td>postgres</td>
            <td>120 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>orders</td>
            <td>table</td>
            <td>postgres</td>
            <td>144 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>products</td>
            <td>table</td>
            <td>postgres</td>
            <td>8192 bytes</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>region</td>
            <td>table</td>
            <td>postgres</td>
            <td>16 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>shippers</td>
            <td>table</td>
            <td>postgres</td>
            <td>8192 bytes</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>suppliers</td>
            <td>table</td>
            <td>postgres</td>
            <td>16 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>territories</td>
            <td>table</td>
            <td>postgres</td>
            <td>16 kB</td>
            <td>None</td>
        </tr>
        <tr>
            <td>public</td>
            <td>us_states</td>
            <td>table</td>
            <td>postgres</td>
            <td>8192 bytes</td>
            <td>None</td>
        </tr>
    </tbody>
</table>

### Connect

`\c pg_catalog` - coonect ot db

### Import & export

* `copy table_name FROM file.txt DELIMITER '|';` 
* `copy table_name FROM file.csv;`
* `copy table_name FROM file.csv NULL as '';`

```
\copy (SELECT * FROM staging.factfinder_import  WHERE s01 ~ E'^[0-9]+' ) 
TO '/test.tab'
WITH DELIMITER E'\t' CSV HEADER
```

```
-- double quote all columns
\copy staging.factfinder_import TO '/test.csv'
WITH CSV HEADER QUOTE '"' FORCE QUOTE *
```

Export as html with `-H` option.

## Tooling

* https://rextester.com/l/postgresql_online_compiler
* https://explain.depesz.com/ - (explain)
* https://dbfiddle.uk

## Data Types

### Numerics

#### Serials / Sequences

```sql
DROP SEQUENCE IF EXISTS s; 
CREATE SEQUENCE s START 1;

DROP TABLE IF EXISTS stuff ;
CREATE TABLE stuff(
    id bigint DEFAULT nextval('s') PRIMARY KEY, name text);

> Done.
> Done.
> Done.
> Done.

result >>> []
```

```sql
SELECT nextval('s')

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>nextval</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
        </tr>
    </tbody>
</table>

```sql
SELECT x FROM generate_series(1,51,20) as x;

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
        </tr>
        <tr>
            <td>21</td>
        </tr>
        <tr>
            <td>41</td>
        </tr>
    </tbody>
</table>

### Textual

#### String Functions

```sql
SELECT
    lpad('ab', 4, '0') as ab_lpad,
    rpad('ab', 4, '0') as ab_rpad,
    lpad('abcde', 4, '0') as ab_lpad_trunc;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ab_lpad</th>
            <th>ab_rpad</th>
            <th>ab_lpad_trunc</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>00ab</td>
            <td>ab00</td>
            <td>abcd</td>
        </tr>
    </tbody>
</table>

```sql
SELECT
    -- different trim types
    a             as a_before, 
    trim(a)       as a_trim, 
    rtrim(a)      as a_rt,
    i             as i_before, 
    ltrim(i, '0') as i_lt_0,
    rtrim(i, '0') as i_rt_0, 
    trim(i, '0')  as i_t_0
FROM (
    -- concat with padding 4 spacing a number...
	SELECT 
        repeat(' ', 4) || i || repeat(' ', 4) as a, 
        '0' || i as i
	FROM generate_series(0, 200, 50) as i
) as  x;

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>a_before</th>
            <th>a_trim</th>
            <th>a_rt</th>
            <th>i_before</th>
            <th>i_lt_0</th>
            <th>i_rt_0</th>
            <th>i_t_0</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;0    </td>
            <td>0</td>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;0</td>
            <td>00</td>
            <td></td>
            <td></td>
            <td></td>
        </tr>
        <tr>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;50    </td>
            <td>50</td>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;50</td>
            <td>050</td>
            <td>50</td>
            <td>05</td>
            <td>5</td>
        </tr>
        <tr>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;100    </td>
            <td>100</td>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;100</td>
            <td>0100</td>
            <td>100</td>
            <td>01</td>
            <td>1</td>
        </tr>
        <tr>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;150    </td>
            <td>150</td>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;150</td>
            <td>0150</td>
            <td>150</td>
            <td>015</td>
            <td>15</td>
        </tr>
        <tr>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;200    </td>
            <td>200</td>
            <td>&nbsp;&nbsp;&nbsp;&nbsp;200</td>
            <td>0200</td>
            <td>200</td>
            <td>02</td>
            <td>2</td>
        </tr>
    </tbody>
</table>

```sql
SELECT split_part('abc.123.z45','.',2) as x;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>123</td>
        </tr>
    </tbody>
</table>

```sql
SELECT unnest(string_to_array('abc.123.z45', '.')) as x;

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>abc</td>
        </tr>
        <tr>
            <td>123</td>
        </tr>
        <tr>
            <td>z45</td>
        </tr>
    </tbody>
</table>

```sql
SELECT regexp_replace( '6197306254',
    '([0-9]{3})([0-9]{3})([0-9]{4})',
     E'\(\\1\) \\2-\\3'
) as x;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>6197306254</td>
        </tr>
    </tbody>
</table>

```sql
SELECT unnest(
    regexp_matches(
    'Cell (619) 852-5083. Work (619)123-4567 , Casa 619-730-6254. Bésame mucho.',
    E'[(]{0,1}[0-9]{3}[)-.]{0,1}[\\s]{0,1}[0-9]{3}[-.]{0,1}[0-9]{4}', 'g')
) as x

> 0 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
    </tbody>
</table>

```sql
SELECT substring(
    'Cell (619) 852-5083. Work (619)123-4567 , Casa 619-730-6254. Bésame mucho.'
    FROM E'[(]{0,1}[0-9]{3}[)-.]{0,1}[\\s]{0,1}[0-9]{3}[-.]{0,1}[0-9]{4}'
) as x;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>None</td>
        </tr>
    </tbody>
</table>

### Temporal Data

- `date` - Stores the month, day, and year, with no time zone awareness and no concept of hours, minutes, or seconds.
- `time` (aka `time without time zone`): Stores hours, minutes, and seconds with no awareness of time zone or calendar dates.
- `timestamp` (aka `timestamp without time zone`): Stores both calendar dates and time (hours, minutes, seconds) but does not care about the time zone.
- `timestamptz` (aka `timestamp with time zone`): A time zone−aware date and time data type. Internally, timestamptz is stored in Coordinated Universal Time (UTC), but its display defaults to the time zone of the server, the service config, the database, the user, or the session. Yes, you can observe different time zones at different levels. If you input a timestamp with no time zone and cast it to one with the time zone, PostgreSQL assumes the default time zone in effect. If you don’t set your time zone in postgresql.conf, the server’s default takes effect. This means that if you change your server’s time zone, you’ll see all the displayed times change after the PostgreSQL restarts.
- `timetz` (aka `time with time zone`): The lesser-used sister of timestamptz. It is time zone−aware but does not store the date. It always assumes DST of the current date and time. Some programming languages with no concept of time without date might map `timetz` to a timestamp with some arbitrary date such as Unix Epoch 1970, resulting in year 1970 being assumed.
- `interval`: A duration of time in hours, days, months, minutes, and others. It comes in handy for datetime arithmetic. For example, if the world is supposed to end in exactly 666 days from now, all you have to do is add an interval of 666 days to the current time to get the exact moment (and plan accordingly).
- `tsrange`: Allows you to define opened and closed ranges of timestamp with no timezone. The type consists of two timestamps and opened/closed range qualifiers. For example, `'[2012-01-01 14:00, 2012-01-01 15:00)'::tsrange` defines a period starting at 14:00 but ending before 15:00. Refer to Range Types for details.
- `tstzrange`: Allows you to define opened and closed ranges of timestamp with timezone.
- `daterange`: Allows you to define opened and closed ranges of dates.

```sql
SELECT '2012-03-11 3:10 AM America/Los_Angeles'::timestamptz
 - '2012-03-11 1:50 AM America/Los_Angeles'::timestamptz;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>0:20:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '2012-03-11 3:10 AM'::timestamp - '2012-03-11 1:50 AM'::timestamp;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1:20:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '2012-02-28 10:00 PM America/Los_Angeles'::timestamptz

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>timestamptz</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2012-02-29 06:00:00+00:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '2012-02-28 10:00 PM America/Los_Angeles'::timestamptz AT TIME ZONE 'Europe/Paris';

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>timezone</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2012-02-29 07:00:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '2012-02-10 11:00 PM'::timestamp + interval '1 hour'

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2012-02-11 00:00:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '23 hours 20 minutes'::interval + '1 hour'::interval;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1 day, 0:20:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '2012-02-10 11:00 PM'::timestamptz - interval '1 hour';

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2012-02-10 22:00:00+00:00</td>
        </tr>
    </tbody>
</table>

```sql
SELECT
('2012-10-25 10:00 AM'::timestamp, '2012-10-25 2:00 PM'::timestamp)
OVERLAPS
('2012-10-25 11:00 AM'::timestamp,'2012-10-26 2:00 PM'::timestamp) as x,
--- other variation
('2012-10-25'::date,'2012-10-26'::date)
OVERLAPS
('2012-10-26'::date,'2012-10-27'::date) as y;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
            <th>y</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
            <td>False</td>
        </tr>
    </tbody>
</table>

```sql
SELECT (dt - interval '1 day')::date as eom
FROM generate_series('2/1/2012', '4/30/2012', interval '1 month') as dt;

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>eom</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2012-01-31</td>
        </tr>
        <tr>
            <td>2012-02-29</td>
        </tr>
        <tr>
            <td>2012-03-31</td>
        </tr>
    </tbody>
</table>

```sql
SELECT dt, 
date_part('hour', dt) as hr, 
to_char(dt,'HH12:MI AM') as mn
FROM
generate_series(
	'2012-03-11 12:30 AM',
	'2012-03-11 1:20 AM',
	interval '15 minutes'
) as dt;

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>dt</th>
            <th>hr</th>
            <th>mn</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2012-03-11 00:30:00+00:00</td>
            <td>0.0</td>
            <td>12:30 AM</td>
        </tr>
        <tr>
            <td>2012-03-11 00:45:00+00:00</td>
            <td>0.0</td>
            <td>12:45 AM</td>
        </tr>
        <tr>
            <td>2012-03-11 01:00:00+00:00</td>
            <td>1.0</td>
            <td>01:00 AM</td>
        </tr>
        <tr>
            <td>2012-03-11 01:15:00+00:00</td>
            <td>1.0</td>
            <td>01:15 AM</td>
        </tr>
    </tbody>
</table>

### Arrays

```sql
SELECT ARRAY[2001, 2002, 2003] as yrs;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>yrs</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[2001, 2002, 2003]</td>
        </tr>
    </tbody>
</table>

```sql
SELECT string_to_array('CA.MA.TX', '.') as estados;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>estados</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[&#x27;CA&#x27;, &#x27;MA&#x27;, &#x27;TX&#x27;]</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '{Alex,Sonia}'::text[] as name, '{46,43}'::smallint[] as age;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>name</th>
            <th>age</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[&#x27;Alex&#x27;, &#x27;Sonia&#x27;]</td>
            <td>[46, 43]</td>
        </tr>
    </tbody>
</table>

```sql
-- column to array
SELECT array_agg(data.region_description) as x
FROM 
    (SELECT * FROM region) as data;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[&#x27;Eastern&#x27;, &#x27;Western&#x27;, &#x27;Northern&#x27;, &#x27;Southern&#x27;]</td>
        </tr>
    </tbody>
</table>

```sql
-- select into arrray from values
SELECT array_agg(f.t)
FROM ( VALUES ('{Alex,Sonia}'::text[]), ('{46,43}'::text[] ) ) as f(t)

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>array_agg</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[[&#x27;Alex&#x27;, &#x27;Sonia&#x27;], [&#x27;46&#x27;, &#x27;43&#x27;]]</td>
        </tr>
    </tbody>
</table>

```sql
SELECT unnest('{XOX, OXO, XOX}'::char(3)[]) as tic_tac_toe;

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>tic_tac_toe</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>XOX</td>
        </tr>
        <tr>
            <td>OXO</td>
        </tr>
        <tr>
            <td>XOX</td>
        </tr>
    </tbody>
</table>

```sql
SELECT
    unnest( '{blind,mouse}'::varchar[]) as v,
    unnest('{1,2,3}'::smallint[]) as i;

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>v</th>
            <th>i</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>blind</td>
            <td>1</td>
        </tr>
        <tr>
            <td>mouse</td>
            <td>2</td>
        </tr>
        <tr>
            <td>None</td>
            <td>3</td>
        </tr>
    </tbody>
</table>

```sql
SELECT * FROM unnest('{blind,mouse}'::text[], '{1,2,3}'::int[]) as f(t,i);

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>t</th>
            <th>i</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>blind</td>
            <td>1</td>
        </tr>
        <tr>
            <td>mouse</td>
            <td>2</td>
        </tr>
        <tr>
            <td>None</td>
            <td>3</td>
        </tr>
    </tbody>
</table>

#### Array Slicing and Splicing

```sql
SELECT d.y[2:3] FROM (SELECT ARRAY[2001, 2002, 2003, 2004] as y ) as d

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>y</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[2002, 2003]</td>
        </tr>
    </tbody>
</table>

```sql
-- will fail in postgress>=14
-- SELECT '{1, 2, 3}'::integer[] || 4 || 5 as x;

SELECT ARRAY[1,2,3]::integer[] || 4 || 5 as x

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[1, 2, 3, 4, 5]</td>
        </tr>
    </tbody>
</table>

```sql
SELECT
    d.c[1] as primero,
    d.c[array_upper(d.c, 1)] as segundo
    FROM (SELECT ARRAY[0,1,2,3,4] as c ) as d

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>primero</th>
            <th>segundo</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>0</td>
            <td>4</td>
        </tr>
    </tbody>
</table>

#### Array Containment Checks

`=`, `<>`, `<`, `>`, `@>`, `<@`, and `&&`. 

```sql
SELECT ARRAY[1,2,3]

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>array</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[1, 2, 3]</td>
        </tr>
    </tbody>
</table>

```sql
-- contains - is all of 2,3 subset of the 1,2,3
SELECT ARRAY[1,2,3]::int[] @> ARRAY[3,2]::int[] as contains

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>contains</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

```sql
-- contained - is all of 1,2,3 subset of the 2,3
SELECT ARRAY[1,2,3]::int[] <@ ARRAY[3,2]::int[] as contains

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>contains</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>False</td>
        </tr>
    </tbody>
</table>

```sql
-- && any elements in common
SELECT 
    ARRAY[1,2,3]::int[] && ARRAY[3,2]::int[] as contains

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>contains</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

### Range Types

* `int4range`, `int8range` - A range of integers. Integer ranges are discrete and subject to canonicalization.
* `numrange` - A continuous range of decimals, floating-point numbers, or double-precision numbers.
* `daterange` - A discrete date range of calendar dates without time zone awareness.
* `tsrange`, `tstzrange` - A continuous date and time (`timestamp`) range allowing for fractional seconds. tstrange is not time zone−aware; `tstzrange` is time zone−aware.

```sql
SELECT '[2013-01-05,2013-08-13]'::daterange;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>daterange</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[2013-01-05, 2013-08-14)</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '(0,)'::int8range;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>int8range</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[1, None)</td>
        </tr>
    </tbody>
</table>

```sql
SELECT '(2013-01-05 10:00,2013-08-13 14:00]'::tsrange;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>tsrange</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>(2013-01-05 10:00:00, 2013-08-13 14:00:00]</td>
        </tr>
    </tbody>
</table>

```sql
SELECT daterange('2013-01-05','infinity','[]');

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>daterange</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>[2013-01-05, 9999-12-31]</td>
        </tr>
    </tbody>
</table>

#### Tables

```sql
DROP TABLE IF EXISTS employment;
CREATE TABLE IF NOT EXISTS employment (
    id serial PRIMARY KEY, 
    employee varchar(20), 
    period daterange);

CREATE INDEX ix_employment_period ON employment USING gist (period); 

INSERT INTO employment (employee,period)
VALUES
	('Alex','[2012-04-24, infinity)'::daterange),
	('Sonia','[2011-04-24, 2012-06-01)'::daterange),
	('Leo','[2012-06-20, 2013-04-20)'::daterange),
	('Regina','[2012-06-20, 2013-04-20)'::daterange);

> Done.
> Done.
> Done.
> 4 rows affected.

result >>> []
```

#### `overlap` operator

```sql
SELECT
    e1.employee,
    string_agg(DISTINCT e2.employee, ', ' ORDER BY e2.employee) as  colleagues
FROM employment as e1 
INNER JOIN employment as e2
-- first opverlats in perion with e2
ON e1.period && e2.period
-- excluding employee itself
WHERE e1.employee <> e2.employee
GROUP BY e1.employee;

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>employee</th>
            <th>colleagues</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Alex</td>
            <td>Leo, Regina, Sonia</td>
        </tr>
        <tr>
            <td>Leo</td>
            <td>Alex, Regina</td>
        </tr>
        <tr>
            <td>Regina</td>
            <td>Alex, Leo</td>
        </tr>
        <tr>
            <td>Sonia</td>
            <td>Alex</td>
        </tr>
    </tbody>
</table>

#### `contains` Operator

```sql
SELECT employee FROM employment 
WHERE period @> CURRENT_DATE 
GROUP BY employee;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>employee</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Alex</td>
        </tr>
    </tbody>
</table>

### `JSON`

- `TODO` https://dou.ua/forums/topic/46086/

```sql
DROP TABLE IF EXISTS persons;
CREATE TABLE persons (id serial PRIMARY KEY, person json);

INSERT INTO persons (person)
VALUES (
    '{
        "name":"Sonia",
        "spouse":
        {
            "name":"Alex",
            "parents":
            {
                "father":"Rafael",
                "mother":"Ofelia"
            },
            "phones":
            [
                {
                    "type":"work",
                    "number":"619-722-6719"
                },
                {
                    "type":"cell",
                    "number":"619-852-5083"
                }
            ]
        },
        "children":
        [
            {
                "name":"Brandon",
                "gender":"M"
            },
            {
                "name":"Azaleah",
                "girl": true,
                "phones": []
            }
        ]
    }'
);

> Done.
> Done.
> 1 rows affected.

result >>> []
```

```sql
SELECT person->'name' FROM persons;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Sonia</td>
        </tr>
    </tbody>
</table>

```sql
SELECT person->'spouse'->'parents'->'father' FROM persons;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Rafael</td>
        </tr>
    </tbody>
</table>

```sql
SELECT person#>array['spouse','parents','father'] FROM persons;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Rafael</td>
        </tr>
    </tbody>
</table>

```sql
SELECT person->'children'->0->'name' FROM persons;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Brandon</td>
        </tr>
    </tbody>
</table>

```sql
SELECT person->'spouse'->'parents'->>'father' FROM persons

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Rafael</td>
        </tr>
    </tbody>
</table>

```sql
SELECT json_array_elements(person->'children')->>'name' as  name FROM persons;

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Brandon</td>
        </tr>
        <tr>
            <td>Azaleah</td>
        </tr>
    </tbody>
</table>

#### Selecting JOSN

```sql
SELECT row_to_json(f) as  x
FROM (
    SELECT id, json_array_elements(person->'children')->>'name' as  cname FROM persons
) as  f;

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>x</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{&#x27;id&#x27;: 1, &#x27;cname&#x27;: &#x27;Brandon&#x27;}</td>
        </tr>
        <tr>
            <td>{&#x27;id&#x27;: 1, &#x27;cname&#x27;: &#x27;Azaleah&#x27;}</td>
        </tr>
    </tbody>
</table>

```sql
SELECT row_to_json(f) as  jsoned_row FROM persons as  f;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>jsoned_row</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{&#x27;id&#x27;: 1, &#x27;person&#x27;: {&#x27;name&#x27;: &#x27;Sonia&#x27;, &#x27;spouse&#x27;: {&#x27;name&#x27;: &#x27;Alex&#x27;, &#x27;parents&#x27;: {&#x27;father&#x27;: &#x27;Rafael&#x27;, &#x27;mother&#x27;: &#x27;Ofelia&#x27;}, &#x27;phones&#x27;: [{&#x27;type&#x27;: &#x27;work&#x27;, &#x27;number&#x27;: &#x27;619-722-6719&#x27;}, {&#x27;type&#x27;: &#x27;cell&#x27;, &#x27;number&#x27;: &#x27;619-852-5083&#x27;}]}, &#x27;children&#x27;: [{&#x27;name&#x27;: &#x27;Brandon&#x27;, &#x27;gender&#x27;: &#x27;M&#x27;}, {&#x27;name&#x27;: &#x27;Azaleah&#x27;, &#x27;girl&#x27;: True, &#x27;phones&#x27;: []}]}}</td>
        </tr>
    </tbody>
</table>

```sql
DROP TABLE IF EXISTS persons;

> Done.

result >>> []
```

### `JSONB`

```sql
CREATE TABLE IF NOT EXISTS persons_b (id serial PRIMARY KEY, person jsonb);

INSERT INTO persons_b (person)
VALUES (
    '{
        "name":"Sonia",
        "spouse":
        {
            "name":"Alex",
            "parents":
            {
                "father":"Rafael",
                "mother":"Ofelia"
            },
            "phones":
            [
                {
                    "type":"work",
                    "number":"619-722-6719"
                },
                {
                    "type":"cell",
                    "number":"619-852-5083"
                }
            ]
        },
        "children":
        [
            {
                "name":"Brandon",
                "gender":"M"
            },
            {
                "name":"Azaleah",
                "girl": true,
                "phones": []
            }
        ]
    }'
);

> Done.
> 1 rows affected.

result >>> []
```

```sql
SELECT person as  b FROM persons_b WHERE id = 1

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>b</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{&#x27;name&#x27;: &#x27;Sonia&#x27;, &#x27;spouse&#x27;: {&#x27;name&#x27;: &#x27;Alex&#x27;, &#x27;phones&#x27;: [{&#x27;type&#x27;: &#x27;work&#x27;, &#x27;number&#x27;: &#x27;619-722-6719&#x27;}, {&#x27;type&#x27;: &#x27;cell&#x27;, &#x27;number&#x27;: &#x27;619-852-5083&#x27;}], &#x27;parents&#x27;: {&#x27;father&#x27;: &#x27;Rafael&#x27;, &#x27;mother&#x27;: &#x27;Ofelia&#x27;}}, &#x27;children&#x27;: [{&#x27;name&#x27;: &#x27;Brandon&#x27;, &#x27;gender&#x27;: &#x27;M&#x27;}, {&#x27;girl&#x27;: True, &#x27;name&#x27;: &#x27;Azaleah&#x27;, &#x27;phones&#x27;: []}]}</td>
        </tr>
    </tbody>
</table>

```sql
SELECT person->>'name' as  name
FROM persons_b
WHERE person @> '{"children":[{"name":"Brandon"}]}';

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Sonia</td>
        </tr>
    </tbody>
</table>

```sql
DROP INDEX IF EXISTS ix_persons_jb_person_gin;
CREATE INDEX ix_persons_jb_person_gin ON persons_b USING gin (person);

> Done.
> Done.

result >>> []
```

#### Editing JSONB data

```sql
-- Concatenation

UPDATE persons_b
SET person = person || '{"address": "Somewhere in San Diego, CA"}'::jsonb
WHERE person @> '{"name":"Sonia"}'
RETURNING person

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>person</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{&#x27;name&#x27;: &#x27;Sonia&#x27;, &#x27;spouse&#x27;: {&#x27;name&#x27;: &#x27;Alex&#x27;, &#x27;phones&#x27;: [{&#x27;type&#x27;: &#x27;work&#x27;, &#x27;number&#x27;: &#x27;619-722-6719&#x27;}, {&#x27;type&#x27;: &#x27;cell&#x27;, &#x27;number&#x27;: &#x27;619-852-5083&#x27;}], &#x27;parents&#x27;: {&#x27;father&#x27;: &#x27;Rafael&#x27;, &#x27;mother&#x27;: &#x27;Ofelia&#x27;}}, &#x27;address&#x27;: &#x27;Somewhere in San Diego, CA&#x27;, &#x27;children&#x27;: [{&#x27;name&#x27;: &#x27;Brandon&#x27;, &#x27;gender&#x27;: &#x27;M&#x27;}, {&#x27;girl&#x27;: True, &#x27;name&#x27;: &#x27;Azaleah&#x27;, &#x27;phones&#x27;: []}]}</td>
        </tr>
    </tbody>
</table>

```sql
-- drop adress from rows where name is sonia
UPDATE persons_b
    SET person = person - 'address'
WHERE person @> '{"name":"Sonia"}';

> 1 rows affected.

result >>> []
```

```sql
UPDATE persons_b
-- path denotion, update will work for second child of sonia, girsl wil be removed
SET  person = person #- '{children,1,girl}'::text[]
WHERE person @> '{"name":"Sonia"}'
RETURNING person->'children'->1;

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{&#x27;name&#x27;: &#x27;Azaleah&#x27;, &#x27;phones&#x27;: []}</td>
        </tr>
    </tbody>
</table>

```sql
UPDATE persons_b 
SET 
person = jsonb_set(person,'{children,1,gender}'::text[],'"F"'::jsonb, true)
WHERE person @> '{"name":"Sonia"}'
RETURNING person

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>person</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>{&#x27;name&#x27;: &#x27;Sonia&#x27;, &#x27;spouse&#x27;: {&#x27;name&#x27;: &#x27;Alex&#x27;, &#x27;phones&#x27;: [{&#x27;type&#x27;: &#x27;work&#x27;, &#x27;number&#x27;: &#x27;619-722-6719&#x27;}, {&#x27;type&#x27;: &#x27;cell&#x27;, &#x27;number&#x27;: &#x27;619-852-5083&#x27;}], &#x27;parents&#x27;: {&#x27;father&#x27;: &#x27;Rafael&#x27;, &#x27;mother&#x27;: &#x27;Ofelia&#x27;}}, &#x27;children&#x27;: [{&#x27;name&#x27;: &#x27;Brandon&#x27;, &#x27;gender&#x27;: &#x27;M&#x27;}, {&#x27;name&#x27;: &#x27;Azaleah&#x27;, &#x27;gender&#x27;: &#x27;F&#x27;, &#x27;phones&#x27;: []}]}</td>
        </tr>
    </tbody>
</table>

### `TODO` XML

### TSVector

```sql
SELECT 'When a good man goes to war'::tsvector @@ 'good & war'::tsquery as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

```sql
SELECT 'When a good man goes to war'::tsvector @@ to_tsquery('good & war') as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

```sql
SELECT 'When a good man goes to war'::tsvector @@ 'good & war'::tsquery as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

Search adjacent string combination

```sql
SELECT to_tsvector('When a good man goes to war') @@ to_tsquery('good <-> man') as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

```sql
SELECT 'When a good man goes to war'::tsvector @@ to_tsquery('good <-> man') as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>False</td>
        </tr>
    </tbody>
</table>

```sql
SELECT to_tsvector('When a good man goes to war') @@ to_tsquery('good <2> goes')  as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

#### TSQueries

```sql
SELECT to_tsvector('good man goes') @@ to_tsquery('good <2> patel') as Result

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>result</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>False</td>
        </tr>
    </tbody>
</table>

```sql
SELECT to_tsquery('business & analytics');

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>to_tsquery</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>&#x27;busi&#x27; &amp; &#x27;analyt&#x27;</td>
        </tr>
    </tbody>
</table>

```sql
SELECT to_tsquery('english','business & analytics');

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>to_tsquery</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>&#x27;busi&#x27; &amp; &#x27;analyt&#x27;</td>
        </tr>
    </tbody>
</table>

#### Search Dictionary

- https://stackoverflow.com/questions/41715609

```sql
\dt pg_catalog.pg_language

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>Schema</th>
            <th>Name</th>
            <th>Type</th>
            <th>Owner</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>pg_catalog</td>
            <td>pg_language</td>
            <td>table</td>
            <td>postgres</td>
        </tr>
    </tbody>
</table>

```sql
SELECT plainto_tsquery('business analytics');

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>plainto_tsquery</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>&#x27;busi&#x27; &amp; &#x27;analyt&#x27;</td>
        </tr>
    </tbody>
</table>

```sql
-- combine tsquery (|| or condition)
SELECT plainto_tsquery('business analyst') || phraseto_tsquery('data scientist');

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>&#x27;busi&#x27; &amp; &#x27;analyst&#x27; | &#x27;data&#x27; &lt;-&gt; &#x27;scientist&#x27;</td>
        </tr>
    </tbody>
</table>

```sql
-- failing since 
-- SELECT "lazy dog and datascientist"::tsquery 
SELECT 
    to_tsvector('english', 'a fat  cat sat on a mat - it ate a fat rats') as to_tsvector

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>to_tsvector</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>&#x27;ate&#x27;:9 &#x27;cat&#x27;:3 &#x27;fat&#x27;:2,11 &#x27;mat&#x27;:7 &#x27;rat&#x27;:12 &#x27;sat&#x27;:4</td>
        </tr>
    </tbody>
</table>

```sql
-- add dataset
SELECT left(title, 50) as  title, left(notes,120) as notes
FROM employees
WHERE notes @@ to_tsquery('sales & (served | degree)') AND title > ''

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title</th>
            <th>notes</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Sales Representative</td>
            <td>Janet has a BS degree in chemistry from Boston College (1984).  She has also completed a certificate program in food ret</td>
        </tr>
        <tr>
            <td>Sales Manager</td>
            <td>Steven Buchanan graduated from St. Andrews University, Scotland, with a BSC degree in 1976.  Upon joining the company as</td>
        </tr>
    </tbody>
</table>

## Querying for Data

```
SELECT values_to_display
 FROM table_anme
 WHERE expression
 GROUP BY how_to_group
 HAVING expression
 ORDER BY how_to_group
 LIMIT row_limit;
```

```sql
SELECT product_id, product_name, quantity_per_unit, unit_price FROM "products" LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_id</th>
            <th>product_name</th>
            <th>quantity_per_unit</th>
            <th>unit_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>Chai</td>
            <td>10 boxes x 30 bags</td>
            <td>18.0</td>
        </tr>
        <tr>
            <td>2</td>
            <td>Chang</td>
            <td>24 - 12 oz bottles</td>
            <td>19.0</td>
        </tr>
        <tr>
            <td>3</td>
            <td>Aniseed Syrup</td>
            <td>12 - 550 ml bottles</td>
            <td>10.0</td>
        </tr>
        <tr>
            <td>4</td>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>48 - 6 oz jars</td>
            <td>22.0</td>
        </tr>
        <tr>
            <td>5</td>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>36 boxes</td>
            <td>21.35</td>
        </tr>
    </tbody>
</table>

Using math in SQL

* `+` - addition
* `-` - subtraction
* `*` - multiplier
* `/` - division
* `^` - power
* `|/` - quadratic power

```sql
SELECT product_id, product_name, units_in_stock * unit_price as price  FROM "products" LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_id</th>
            <th>product_name</th>
            <th>price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>Chai</td>
            <td>702.0</td>
        </tr>
        <tr>
            <td>2</td>
            <td>Chang</td>
            <td>323.0</td>
        </tr>
        <tr>
            <td>3</td>
            <td>Aniseed Syrup</td>
            <td>130.0</td>
        </tr>
        <tr>
            <td>4</td>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>1166.0</td>
        </tr>
        <tr>
            <td>5</td>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>0.0</td>
        </tr>
    </tbody>
</table>

### `DISTINCT`

```sql
SELECT DISTINCT city FROM employees

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>city</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Redmond</td>
        </tr>
        <tr>
            <td>London</td>
        </tr>
        <tr>
            <td>Tacoma</td>
        </tr>
        <tr>
            <td>Kirkland</td>
        </tr>
        <tr>
            <td>Seattle</td>
        </tr>
    </tbody>
</table>

```sql
SELECT DISTINCT city, country FROM employees

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>city</th>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Seattle</td>
            <td>USA</td>
        </tr>
        <tr>
            <td>Kirkland</td>
            <td>USA</td>
        </tr>
        <tr>
            <td>London</td>
            <td>UK</td>
        </tr>
        <tr>
            <td>Redmond</td>
            <td>USA</td>
        </tr>
        <tr>
            <td>Tacoma</td>
            <td>USA</td>
        </tr>
    </tbody>
</table>

### `WHERE` & `OREDER BY`

```sql
SELECT order_id, customer_id, order_date, shipped_date FROM orders WHERE ship_city = 'Oulu' LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>order_id</th>
            <th>customer_id</th>
            <th>order_date</th>
            <th>shipped_date</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>10266</td>
            <td>WARTH</td>
            <td>1996-07-26</td>
            <td>1996-07-31</td>
        </tr>
        <tr>
            <td>10270</td>
            <td>WARTH</td>
            <td>1996-08-01</td>
            <td>1996-08-02</td>
        </tr>
        <tr>
            <td>10320</td>
            <td>WARTH</td>
            <td>1996-10-03</td>
            <td>1996-10-18</td>
        </tr>
        <tr>
            <td>10333</td>
            <td>WARTH</td>
            <td>1996-10-18</td>
            <td>1996-10-25</td>
        </tr>
        <tr>
            <td>10412</td>
            <td>WARTH</td>
            <td>1997-01-13</td>
            <td>1997-01-15</td>
        </tr>
    </tbody>
</table>

```sql
SELECT ship_country, ship_city, freight FROM orders WHERE ship_city != 'Oulu' AND freight > 30 ORDER BY freight DESC LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ship_country</th>
            <th>ship_city</th>
            <th>freight</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Germany</td>
            <td>Cunewalde</td>
            <td>1007.64</td>
        </tr>
        <tr>
            <td>Brazil</td>
            <td>Sao Paulo</td>
            <td>890.78</td>
        </tr>
        <tr>
            <td>USA</td>
            <td>Boise</td>
            <td>830.75</td>
        </tr>
        <tr>
            <td>Germany</td>
            <td>Cunewalde</td>
            <td>810.05</td>
        </tr>
    </tbody>
</table>

```sql
SELECT ship_country, ship_city, freight FROM orders WHERE ship_city <> 'Oulu' AND freight > 30 LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ship_country</th>
            <th>ship_city</th>
            <th>freight</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>France</td>
            <td>Reims</td>
            <td>32.38</td>
        </tr>
        <tr>
            <td>Brazil</td>
            <td>Rio de Janeiro</td>
            <td>65.83</td>
        </tr>
        <tr>
            <td>France</td>
            <td>Lyon</td>
            <td>41.34</td>
        </tr>
        <tr>
            <td>Belgium</td>
            <td>Charleroi</td>
            <td>51.3</td>
        </tr>
    </tbody>
</table>

```sql
SELECT ship_country, ship_city, freight FROM orders WHERE order_date > '1998-04-10' LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ship_country</th>
            <th>ship_city</th>
            <th>freight</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Austria</td>
            <td>Graz</td>
            <td>754.26</td>
        </tr>
        <tr>
            <td>USA</td>
            <td>Portland</td>
            <td>11.65</td>
        </tr>
        <tr>
            <td>Argentina</td>
            <td>Buenos Aires</td>
            <td>3.17</td>
        </tr>
        <tr>
            <td>Germany</td>
            <td>Köln</td>
            <td>43.3</td>
        </tr>
    </tbody>
</table>

### `AND`,  `OR` and `NOT`

```sql
SELECT customer_id, company_name, contact_name, contact_title, address, city 
FROM customers 
WHERE city = 'London' OR city = 'Berlin' LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>company_name</th>
            <th>contact_name</th>
            <th>contact_title</th>
            <th>address</th>
            <th>city</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>ALFKI</td>
            <td>Alfreds Futterkiste</td>
            <td>Maria Anders</td>
            <td>Sales Representative</td>
            <td>Obere Str. 57</td>
            <td>Berlin</td>
        </tr>
        <tr>
            <td>AROUT</td>
            <td>Around the Horn</td>
            <td>Thomas Hardy</td>
            <td>Sales Representative</td>
            <td>120 Hanover Sq.</td>
            <td>London</td>
        </tr>
        <tr>
            <td>BSBEV</td>
            <td>B&#x27;s Beverages</td>
            <td>Victoria Ashworth</td>
            <td>Sales Representative</td>
            <td>Fauntleroy Circus</td>
            <td>London</td>
        </tr>
        <tr>
            <td>CONSH</td>
            <td>Consolidated Holdings</td>
            <td>Elizabeth Brown</td>
            <td>Sales Representative</td>
            <td>Berkeley Gardens 12  Brewery</td>
            <td>London</td>
        </tr>
    </tbody>
</table>

```sql
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock 
FROM products WHERE (units_in_stock < units_on_order ) OR ( discontinued = 1 ) LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_id</th>
            <th>product_name</th>
            <th>supplier_id</th>
            <th>category_id</th>
            <th>quantity_per_unit</th>
            <th>unit_price</th>
            <th>units_in_stock</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>Chai</td>
            <td>8</td>
            <td>1</td>
            <td>10 boxes x 30 bags</td>
            <td>18.0</td>
            <td>39</td>
        </tr>
        <tr>
            <td>2</td>
            <td>Chang</td>
            <td>1</td>
            <td>1</td>
            <td>24 - 12 oz bottles</td>
            <td>19.0</td>
            <td>17</td>
        </tr>
        <tr>
            <td>3</td>
            <td>Aniseed Syrup</td>
            <td>1</td>
            <td>2</td>
            <td>12 - 550 ml bottles</td>
            <td>10.0</td>
            <td>13</td>
        </tr>
        <tr>
            <td>5</td>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>2</td>
            <td>2</td>
            <td>36 boxes</td>
            <td>21.35</td>
            <td>0</td>
        </tr>
    </tbody>
</table>

```sql
SELECT customer_id, company_name, contact_name, contact_title, address, city 
FROM customers WHERE city = 'London' OR city = 'Berlin' LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>company_name</th>
            <th>contact_name</th>
            <th>contact_title</th>
            <th>address</th>
            <th>city</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>ALFKI</td>
            <td>Alfreds Futterkiste</td>
            <td>Maria Anders</td>
            <td>Sales Representative</td>
            <td>Obere Str. 57</td>
            <td>Berlin</td>
        </tr>
        <tr>
            <td>AROUT</td>
            <td>Around the Horn</td>
            <td>Thomas Hardy</td>
            <td>Sales Representative</td>
            <td>120 Hanover Sq.</td>
            <td>London</td>
        </tr>
        <tr>
            <td>BSBEV</td>
            <td>B&#x27;s Beverages</td>
            <td>Victoria Ashworth</td>
            <td>Sales Representative</td>
            <td>Fauntleroy Circus</td>
            <td>London</td>
        </tr>
        <tr>
            <td>CONSH</td>
            <td>Consolidated Holdings</td>
            <td>Elizabeth Brown</td>
            <td>Sales Representative</td>
            <td>Berkeley Gardens 12  Brewery</td>
            <td>London</td>
        </tr>
    </tbody>
</table>

```sql
SELECT customer_id, company_name, contact_name, contact_title, address, city 
FROM customers WHERE city NOT IN ('London', 'Berlin') LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>company_name</th>
            <th>contact_name</th>
            <th>contact_title</th>
            <th>address</th>
            <th>city</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>ANATR</td>
            <td>Ana Trujillo Emparedados y helados</td>
            <td>Ana Trujillo</td>
            <td>Owner</td>
            <td>Avda. de la Constitución 2222</td>
            <td>México D.F.</td>
        </tr>
        <tr>
            <td>ANTON</td>
            <td>Antonio Moreno Taquería</td>
            <td>Antonio Moreno</td>
            <td>Owner</td>
            <td>Mataderos  2312</td>
            <td>México D.F.</td>
        </tr>
        <tr>
            <td>BERGS</td>
            <td>Berglunds snabbköp</td>
            <td>Christina Berglund</td>
            <td>Order Administrator</td>
            <td>Berguvsvägen  8</td>
            <td>Luleå</td>
        </tr>
        <tr>
            <td>BLAUS</td>
            <td>Blauer See Delikatessen</td>
            <td>Hanna Moos</td>
            <td>Sales Representative</td>
            <td>Forsterstr. 57</td>
            <td>Mannheim</td>
        </tr>
    </tbody>
</table>

### `BETWEEN` 

```sql
SELECT * FROM (
    SELECT o.customer_id, SUM(od.unit_price * od.quantity) as tot 
    FROM orders as o, order_details as od 
    WHERE od.order_id = o.order_id 
    GROUP BY o.customer_id
) as data
WHERE data.tot BETWEEN 5000 AND 6500 LIMIT  3

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>tot</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>MORGK</td>
            <td>5042.200035095215</td>
        </tr>
        <tr>
            <td>BOLID</td>
            <td>5297.800024032593</td>
        </tr>
        <tr>
            <td>SANTG</td>
            <td>5735.1500153541565</td>
        </tr>
    </tbody>
</table>

### Agregate with `MIN`, `MAX`, `SUM` & `AVG`

```sql
SELECT MIN(order_date), MAX(order_date) FROM orders  WHERE ship_city = 'London'

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>min</th>
            <th>max</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1996-08-26</td>
            <td>1998-04-29</td>
        </tr>
    </tbody>
</table>

```sql
SELECT AVG(unit_price) FROM products

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>avg</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>28.83389609200614</td>
        </tr>
    </tbody>
</table>

```sql
SELECT SUM(units_in_stock*unit_price) as price FROM products

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>73953.34993171692</td>
        </tr>
    </tbody>
</table>

### Agregate with  `COUNT(*)`

```sql
SELECT COUNT(DISTINCT ship_city) FROM orders

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>count</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>70</td>
        </tr>
    </tbody>
</table>

### Agregate Group Contcat

Alternative to `GROUP_CONCAT`

#### `CREATE AGGREGATE`

mechanix changed in [pg14](https://stackoverflow.com/questions/69632898)

```sql
CREATE OR REPLACE FUNCTION _group_concat_finalize(anyarray)
RETURNS text AS $$
    SELECT array_to_string($1,',')
$$ IMMUTABLE LANGUAGE SQL;

CREATE OR REPLACE AGGREGATE group_concat(anycompatible) (
   SFUNC=array_append,
   STYPE=anycompatiblearray,
   FFUNC=_group_concat_finalize,
   INITCOND='{}'
);

SELECT customer_id, group_concat(DISTINCT ship_city) FROM orders   GROUP BY customer_id LIMIT 5;

> Done.
> Done.
> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>group_concat</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>ALFKI</td>
            <td>[&#x27;Berlin&#x27;]</td>
        </tr>
        <tr>
            <td>ANATR</td>
            <td>[&#x27;México D.F.&#x27;]</td>
        </tr>
        <tr>
            <td>ANTON</td>
            <td>[&#x27;México D.F.&#x27;]</td>
        </tr>
        <tr>
            <td>AROUT</td>
            <td>[&#x27;Colchester&#x27;]</td>
        </tr>
        <tr>
            <td>BERGS</td>
            <td>[&#x27;Luleå&#x27;]</td>
        </tr>
    </tbody>
</table>

### `LIKE`

* `%` (0...)
* `_` (1)

Examples:
* LIKE "%john%" -- contains john
* LIKE "john%" -- start with john
* LIKE "%john" -- ends on john

```sql
SELECT first_name, last_name FROM employees WHERE last_name LIKE '_avolio' OR last_name LIKE  '%King%'

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>first_name</th>
            <th>last_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Nancy</td>
            <td>Davolio</td>
        </tr>
        <tr>
            <td>Robert</td>
            <td>King</td>
        </tr>
    </tbody>
</table>

### SIMILAR TO `~`

```sql
SELECT customer_id, phone
FROM customers
WHERE phone ~  
E'[(]{0,1}[0-9]{3}[)-.]{0,1}[\\s]{0,1}[0-9]{3}[-.]{0,1}[0-9]{4}'
LIMIT 4

> 0 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>phone</th>
        </tr>
    </thead>
    <tbody>
    </tbody>
</table>

### Dealing with `NULL`

```sql
SELECT ship_city, ship_region, ship_country FROM orders WHERE ship_region is NULL LIMIT 2

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ship_city</th>
            <th>ship_region</th>
            <th>ship_country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Reims</td>
            <td>None</td>
            <td>France</td>
        </tr>
        <tr>
            <td>Münster</td>
            <td>None</td>
            <td>Germany</td>
        </tr>
    </tbody>
</table>

```sql
SELECT ship_city, ship_region, ship_country FROM orders WHERE ship_region is NOT NULL LIMIT 2

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ship_city</th>
            <th>ship_region</th>
            <th>ship_country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Rio de Janeiro</td>
            <td>RJ</td>
            <td>Brazil</td>
        </tr>
        <tr>
            <td>Rio de Janeiro</td>
            <td>RJ</td>
            <td>Brazil</td>
        </tr>
    </tbody>
</table>

### `GROUP BY`

```sql
SELECT ship_country, count(*) FROM orders WHERE freight > 10 GROUP by ship_country ORDER BY count DESC LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>ship_country</th>
            <th>count</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Germany</td>
            <td>104</td>
        </tr>
        <tr>
            <td>USA</td>
            <td>104</td>
        </tr>
        <tr>
            <td>Brazil</td>
            <td>62</td>
        </tr>
        <tr>
            <td>France</td>
            <td>55</td>
        </tr>
        <tr>
            <td>UK</td>
            <td>41</td>
        </tr>
    </tbody>
</table>

#### `TODO: WITH ROLLUP`

### `HAVING` 

_Post selection group_

```sql
SELECT * FROM (
    SELECT o.customer_id, SUM(od.unit_price * od.quantity) as tot 
    FROM orders as o, order_details as od 
    WHERE od.order_id = o.order_id 
    GROUP BY o.customer_id
) as data
WHERE data.tot BETWEEN 5000 AND 6500

> 7 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>tot</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>MORGK</td>
            <td>5042.200035095215</td>
        </tr>
        <tr>
            <td>BOLID</td>
            <td>5297.800024032593</td>
        </tr>
        <tr>
            <td>SANTG</td>
            <td>5735.1500153541565</td>
        </tr>
        <tr>
            <td>BSBEV</td>
            <td>6089.899990081787</td>
        </tr>
        <tr>
            <td>WELLI</td>
            <td>6480.69997215271</td>
        </tr>
        <tr>
            <td>PRINI</td>
            <td>5317.100019454956</td>
        </tr>
        <tr>
            <td>ISLAT</td>
            <td>6146.299984931946</td>
        </tr>
    </tbody>
</table>

```sql
SELECT o.customer_id, SUM(od.unit_price * od.quantity)
FROM orders as o, order_details as od 
WHERE od.order_id = o.order_id 
GROUP BY o.customer_id
HAVING SUM(od.unit_price * od.quantity) BETWEEN 5000 AND 6500
ORDER BY sum

> 7 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer_id</th>
            <th>sum</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>MORGK</td>
            <td>5042.200035095215</td>
        </tr>
        <tr>
            <td>BOLID</td>
            <td>5297.800024032593</td>
        </tr>
        <tr>
            <td>PRINI</td>
            <td>5317.100019454956</td>
        </tr>
        <tr>
            <td>SANTG</td>
            <td>5735.1500153541565</td>
        </tr>
        <tr>
            <td>BSBEV</td>
            <td>6089.899990081787</td>
        </tr>
        <tr>
            <td>ISLAT</td>
            <td>6146.299984931946</td>
        </tr>
        <tr>
            <td>WELLI</td>
            <td>6480.69997215271</td>
        </tr>
    </tbody>
</table>

### `UNION`

_The SQL UNION clause/operator is used to combine the results of two or more SELECT statements without returning any duplicate rows._

```sql
SELECT * FROM (
    SELECT country FROM customers
    UNION
    SELECT country FROM employees
) as source LIMIT 2

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Italy</td>
        </tr>
        <tr>
            <td>Venezuela</td>
        </tr>
    </tbody>
</table>

Th `UNION ALL` command combines the result set of two or more SELECT statements (allows duplicate values).

```sql
-- do not remove duplicates

SELECT * FROM (
    SELECT country FROM customers
    UNION ALL
    SELECT country FROM employees
) as source LIMIT 3

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Germany</td>
        </tr>
        <tr>
            <td>Mexico</td>
        </tr>
        <tr>
            <td>Mexico</td>
        </tr>
    </tbody>
</table>

### `INTERSECT`

_The SQL INTERSECT clause/operator is used to combine two SELECT statements, but returns rows only from the first SELECT statement that are identical to a row in the second SELECT statement. This means INTERSECT returns only common rows returned by the two SELECT statements._

```sql
SELECT * FROM (
    SELECT country FROM customers
    INTERSECT
    SELECT country FROM employees
) as source LIMIT 2

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>UK</td>
        </tr>
        <tr>
            <td>USA</td>
        </tr>
    </tbody>
</table>

### `EXCEPT` 

_The SQL EXCEPT clause/operator is used to combine two SELECT statements and returns rows from the first SELECT statement that are not returned by the second SELECT statement. This means EXCEPT returns only rows, which are not available in the second SELECT statement._

```sql
-- do not remove duplicates

SELECT * FROM (
    SELECT country FROM customers
    EXCEPT
    SELECT country FROM employees
) as source LIMIT 2

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Italy</td>
        </tr>
        <tr>
            <td>Venezuela</td>
        </tr>
    </tbody>
</table>

```sql
-- do not remove duplicates

SELECT * FROM (
    SELECT country FROM customers
    EXCEPT ALL
    SELECT country FROM employees
) as source LIMIT 2

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Italy</td>
        </tr>
        <tr>
            <td>Italy</td>
        </tr>
    </tbody>
</table>

### `RETURNING`

<div class="alert alert-warning">WARNING: Jupyter doesnt work correctly with <code>returning *</code>, so we are wrpapping it</div>

* https://www.postgresql.org/docs/13/dml-returning.html

```sql
# can't perform magic with returning *
try:
    data =  UPDATE employees SET title_of_courtesy = 'Mrs' WHERE employee_id = 1 RETURNING *
    print(data)
except:
    pass

> 1 rows affected.
```

```sql
UPDATE employees SET title_of_courtesy = 'Mrs' WHERE employee_id = 1 RETURNING title_of_courtesy, last_name

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title_of_courtesy</th>
            <th>last_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Mrs</td>
            <td>Davolio</td>
        </tr>
    </tbody>
</table>

```sql
UPDATE employees SET title_of_courtesy = 'Ms' WHERE employee_id = 1 RETURNING title_of_courtesy, last_name

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title_of_courtesy</th>
            <th>last_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Ms</td>
            <td>Davolio</td>
        </tr>
    </tbody>
</table>

```sql
INSERT INTO employees (employee_id, first_name, last_name)
VALUES (10, 'John', 'Dow')
RETURNING title_of_courtesy, first_name, last_name

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title_of_courtesy</th>
            <th>first_name</th>
            <th>last_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>None</td>
            <td>John</td>
            <td>Dow</td>
        </tr>
    </tbody>
</table>

```sql
UPDATE employees SET title_of_courtesy = 'Dr' WHERE employee_id = 10 RETURNING title_of_courtesy, last_name

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title_of_courtesy</th>
            <th>last_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Dr</td>
            <td>Dow</td>
        </tr>
    </tbody>
</table>

```sql
DELETE FROM employees WHERE employee_id = 10 RETURNING title_of_courtesy, last_name, first_name

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title_of_courtesy</th>
            <th>last_name</th>
            <th>first_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Dr</td>
            <td>Dow</td>
            <td>John</td>
        </tr>
    </tbody>
</table>

## `TODO` Indexes

## SQL Expressions

```sql
SELECT VERSION()

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>version</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>PostgreSQL 15.4 (Debian 15.4-2.pgdg120+1) on aarch64-unknown-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit</td>
        </tr>
    </tbody>
</table>

```sql
SELECT 1+2+4.0

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>7.0</td>
        </tr>
    </tbody>
</table>

```sql
SELECT 1+2+4.0 = 7

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>True</td>
        </tr>
    </tbody>
</table>

## `TODO` Updating Data

- Insert, Updates, Upserts

## Joins

### `INNER JOIN`

<center>
    <i>Only <b>Common Values</b> of two tables.</i>
</center>

![](images/join-inner.gif)

```sql
SELECT product_name, company_name, units_in_stock 
FROM products p
INNER JOIN suppliers s ON (p.supplier_id = s.supplier_id )
LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_name</th>
            <th>company_name</th>
            <th>units_in_stock</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Chai</td>
            <td>Specialty Biscuits, Ltd.</td>
            <td>39</td>
        </tr>
        <tr>
            <td>Chang</td>
            <td>Exotic Liquids</td>
            <td>17</td>
        </tr>
        <tr>
            <td>Aniseed Syrup</td>
            <td>Exotic Liquids</td>
            <td>13</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>New Orleans Cajun Delights</td>
            <td>53</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>New Orleans Cajun Delights</td>
            <td>0</td>
        </tr>
    </tbody>
</table>

```sql
SELECT c.category_name, SUM(p.units_in_stock)
FROM products p
INNER JOIN categories c ON (p.category_id = c.category_id)
GROUP BY c.category_id
ORDER BY sum

> 8 rows affected.
```

<table>
    <thead>
        <tr>
            <th>category_name</th>
            <th>sum</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Produce</td>
            <td>100</td>
        </tr>
        <tr>
            <td>Meat/Poultry</td>
            <td>165</td>
        </tr>
        <tr>
            <td>Grains/Cereals</td>
            <td>308</td>
        </tr>
        <tr>
            <td>Confections</td>
            <td>386</td>
        </tr>
        <tr>
            <td>Dairy Products</td>
            <td>393</td>
        </tr>
        <tr>
            <td>Condiments</td>
            <td>507</td>
        </tr>
        <tr>
            <td>Beverages</td>
            <td>559</td>
        </tr>
        <tr>
            <td>Seafood</td>
            <td>701</td>
        </tr>
    </tbody>
</table>

```sql
SELECT c.category_name, SUM(p.units_in_stock * p.unit_price)
FROM products p
INNER JOIN categories c ON (p.category_id = c.category_id)
WHERE p.discontinued != 1
GROUP BY c.category_id
ORDER BY sum 
LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>category_name</th>
            <th>sum</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Produce</td>
            <td>2363.75</td>
        </tr>
        <tr>
            <td>Meat/Poultry</td>
            <td>2916.449995994568</td>
        </tr>
        <tr>
            <td>Grains/Cereals</td>
            <td>5230.5</td>
        </tr>
        <tr>
            <td>Confections</td>
            <td>10392.200072288513</td>
        </tr>
    </tbody>
</table>

```sql
SELECT order_date, product_name, ship_country, p.unit_price
FROM orders
INNER JOIN order_details as od  ON (orders.order_id = od.order_id) 
INNER JOIN products as p ON (od.product_id = p.product_id)
LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>order_date</th>
            <th>product_name</th>
            <th>ship_country</th>
            <th>unit_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1996-07-04</td>
            <td>Queso Cabrales</td>
            <td>France</td>
            <td>21.0</td>
        </tr>
        <tr>
            <td>1996-07-04</td>
            <td>Singaporean Hokkien Fried Mee</td>
            <td>France</td>
            <td>14.0</td>
        </tr>
        <tr>
            <td>1996-07-04</td>
            <td>Mozzarella di Giovanni</td>
            <td>France</td>
            <td>34.8</td>
        </tr>
        <tr>
            <td>1996-07-05</td>
            <td>Tofu</td>
            <td>Germany</td>
            <td>23.25</td>
        </tr>
        <tr>
            <td>1996-07-05</td>
            <td>Manjimup Dried Apples</td>
            <td>Germany</td>
            <td>53.0</td>
        </tr>
    </tbody>
</table>

```sql
SELECT 
    order_date, product_name, ship_country, p.unit_price,
    contact_name, company_name
FROM orders as o
JOIN order_details as od  ON (o.order_id = od.order_id) 
INNER JOIN products as p ON (od.product_id = p.product_id)
JOIN customers as c ON (o.customer_id = c.customer_id)
JOIN employees as e ON (o.employee_id = e.employee_id)
WHERE ship_country = 'USA'
LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>order_date</th>
            <th>product_name</th>
            <th>ship_country</th>
            <th>unit_price</th>
            <th>contact_name</th>
            <th>company_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1997-05-06</td>
            <td>Queso Cabrales</td>
            <td>USA</td>
            <td>21.0</td>
            <td>Howard Snyder</td>
            <td>Great Lakes Food Market</td>
        </tr>
        <tr>
            <td>1997-05-06</td>
            <td>Geitost</td>
            <td>USA</td>
            <td>2.5</td>
            <td>Howard Snyder</td>
            <td>Great Lakes Food Market</td>
        </tr>
        <tr>
            <td>1997-05-06</td>
            <td>Mozzarella di Giovanni</td>
            <td>USA</td>
            <td>34.8</td>
            <td>Howard Snyder</td>
            <td>Great Lakes Food Market</td>
        </tr>
        <tr>
            <td>1997-07-04</td>
            <td>Steeleye Stout</td>
            <td>USA</td>
            <td>18.0</td>
            <td>Howard Snyder</td>
            <td>Great Lakes Food Market</td>
        </tr>
        <tr>
            <td>1997-07-31</td>
            <td>Côte de Blaye</td>
            <td>USA</td>
            <td>263.5</td>
            <td>Howard Snyder</td>
            <td>Great Lakes Food Market</td>
        </tr>
    </tbody>
</table>

### `LEFT (OUTTER) JOIN` 

<center>
    <i>Only values of the <b>the Left</b> table and <b>matches of the Right</b> tables.</i>
</center>

![](images/join-left.gif)

```sql
SELECT company_name as supplier, product_name as product
FROM suppliers as s
LEFT JOIN products as p ON (p.supplier_id = s.supplier_id )
LIMIT 3

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>supplier</th>
            <th>product</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Exotic Liquids</td>
            <td>Chang</td>
        </tr>
        <tr>
            <td>Exotic Liquids</td>
            <td>Aniseed Syrup</td>
        </tr>
        <tr>
            <td>New Orleans Cajun Delights</td>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
        </tr>
    </tbody>
</table>

### `LEFT (OUTTER) JOIN` <sup>*</sup>

<center>
    <i>Only values of the <b>the Left</b> table and <b>not matches of the Right</b> tables.</i>
</center>

![](images/join-left-null.gif)

```sql
SELECT company_name as customer, c.customer_id
FROM customers as c
LEFT JOIN orders as o ON (o.customer_id = c.customer_id )
WHERE order_id is NULL

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer</th>
            <th>customer_id</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Paris spécialités</td>
            <td>PARIS</td>
        </tr>
        <tr>
            <td>FISSA Fabrica Inter. Salchichas S.A.</td>
            <td>FISSA</td>
        </tr>
    </tbody>
</table>

```sql
SELECT last_name, order_id
FROM employees as e
LEFT JOIN orders as o ON (o.employee_id = e.employee_id )
WHERE order_id is NULL

> 0 rows affected.
```

<table>
    <thead>
        <tr>
            <th>last_name</th>
            <th>order_id</th>
        </tr>
    </thead>
    <tbody>
    </tbody>
</table>

### `RIGHT (OUTTER) JOIN`

<center>
    <i>Only values of the <b>the Right</b> table and <b>matches of the Left</b> table.</i>
</center>

![](images/join-right.gif)

```sql
SELECT company_name as supplier, product_name as product
FROM products as p
LEFT JOIN suppliers as s ON (p.supplier_id = s.supplier_id )
LIMIT 3

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>supplier</th>
            <th>product</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Specialty Biscuits, Ltd.</td>
            <td>Chai</td>
        </tr>
        <tr>
            <td>Exotic Liquids</td>
            <td>Chang</td>
        </tr>
        <tr>
            <td>Exotic Liquids</td>
            <td>Aniseed Syrup</td>
        </tr>
    </tbody>
</table>

### `RIGHT (OUTTER) JOIN` <sup>*</sup>

<center>
    <i>Only values of the <b>the Right</b> table and <b>not matches of the Left</b> table.</i>
</center>

![](images/join-right-null.gif)

```sql
-- same request that we did for left join.

SELECT company_name as customer, c.customer_id
FROM orders as o
RIGHT JOIN customers as c ON (o.customer_id = c.customer_id )
WHERE order_id is NULL

> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer</th>
            <th>customer_id</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Paris spécialités</td>
            <td>PARIS</td>
        </tr>
        <tr>
            <td>FISSA Fabrica Inter. Salchichas S.A.</td>
            <td>FISSA</td>
        </tr>
    </tbody>
</table>

### `FULL JOIN`

<center>
    <i><b>All</b> values <b>from left and right</b> (<b>with nulls</b> on <b>non matching values</b>) </i>
</center>

![](images/join-full.gif)

```sql
SELECT company_name as customer, c.customer_id
FROM orders as o
FULL JOIN customers as c ON (o.customer_id = c.customer_id )
LIMIT 10

> 10 rows affected.
```

<table>
    <thead>
        <tr>
            <th>customer</th>
            <th>customer_id</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Vins et alcools Chevalier</td>
            <td>VINET</td>
        </tr>
        <tr>
            <td>Toms Spezialitäten</td>
            <td>TOMSP</td>
        </tr>
        <tr>
            <td>Hanari Carnes</td>
            <td>HANAR</td>
        </tr>
        <tr>
            <td>Victuailles en stock</td>
            <td>VICTE</td>
        </tr>
        <tr>
            <td>Suprêmes délices</td>
            <td>SUPRD</td>
        </tr>
        <tr>
            <td>Hanari Carnes</td>
            <td>HANAR</td>
        </tr>
        <tr>
            <td>Chop-suey Chinese</td>
            <td>CHOPS</td>
        </tr>
        <tr>
            <td>Richter Supermarkt</td>
            <td>RICSU</td>
        </tr>
        <tr>
            <td>Wellington Importadora</td>
            <td>WELLI</td>
        </tr>
        <tr>
            <td>HILARION-Abastos</td>
            <td>HILAA</td>
        </tr>
    </tbody>
</table>

### `CROSS JOIN`

<center>
    <i><b>All combinations</b> of values from <b>left</b> and <b>right</b></i>
</center>

![](images/join-cross.gif)

```sql
SELECT count(*)
FROM orders as o
CROSS JOIN customers as c

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>count</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>75530</td>
        </tr>
    </tbody>
</table>

### `SELF JOIN`

```sql
CREATE TABLE IF NOT EXISTS extras_self_join_example (
    id   INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    pid  INT,
    
    FOREIGN KEY (pid) REFERENCES extras_self_join_example (id)
);

> Done.

result >>> []
```

```sql
INSERT INTO extras_self_join_example (
    id, name, pid
) VALUES
(1, 'Foo', NULL),
(2, 'Bar', 1),
(3, 'Baz', 2),
(4, 'Quz', 1);

> 4 rows affected.

result >>> []
```

```sql
SELECT e.name || '/' || e2.name, e2.pid
FROM extras_self_join_example as e
LEFT JOIN extras_self_join_example as e2 ON e2.pid = e.id
WHERE e2.pid IS NOT NULL
ORDER BY e2.pid DESC

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>?column?</th>
            <th>pid</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Bar/Baz</td>
            <td>2</td>
        </tr>
        <tr>
            <td>Foo/Bar</td>
            <td>1</td>
        </tr>
        <tr>
            <td>Foo/Quz</td>
            <td>1</td>
        </tr>
    </tbody>
</table>

```sql
DROP TABLE extras_self_join_example

> Done.

result >>> []
```

### `USING` keyword

_Allows to use shared column to make join_

```sql
SELECT product_name, company_name, units_in_stock 
FROM products p
INNER JOIN suppliers USING(supplier_id)
LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_name</th>
            <th>company_name</th>
            <th>units_in_stock</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Chai</td>
            <td>Specialty Biscuits, Ltd.</td>
            <td>39</td>
        </tr>
        <tr>
            <td>Chang</td>
            <td>Exotic Liquids</td>
            <td>17</td>
        </tr>
        <tr>
            <td>Aniseed Syrup</td>
            <td>Exotic Liquids</td>
            <td>13</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>New Orleans Cajun Delights</td>
            <td>53</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>New Orleans Cajun Delights</td>
            <td>0</td>
        </tr>
    </tbody>
</table>

### `NATURAL JOIN`

_Joining tables using columns with same name_

```sql
SELECT product_name, company_name, units_in_stock 
FROM products p
NATURAL JOIN suppliers 
LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_name</th>
            <th>company_name</th>
            <th>units_in_stock</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Chai</td>
            <td>Specialty Biscuits, Ltd.</td>
            <td>39</td>
        </tr>
        <tr>
            <td>Chang</td>
            <td>Exotic Liquids</td>
            <td>17</td>
        </tr>
        <tr>
            <td>Aniseed Syrup</td>
            <td>Exotic Liquids</td>
            <td>13</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>New Orleans Cajun Delights</td>
            <td>53</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>New Orleans Cajun Delights</td>
            <td>0</td>
        </tr>
    </tbody>
</table>

## SubQueries

```sql
SELECT company_name, country 
FROM suppliers as s
WHERE s.country IN (
    SELECT DISTINCT(country) FROM customers
) 
ORDER BY s.company_name 
LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>company_name</th>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Aux joyeux ecclésiastiques</td>
            <td>France</td>
        </tr>
        <tr>
            <td>Bigfoot Breweries</td>
            <td>USA</td>
        </tr>
        <tr>
            <td>Cooperativa de Quesos &#x27;Las Cabras&#x27;</td>
            <td>Spain</td>
        </tr>
        <tr>
            <td>Escargots Nouveaux</td>
            <td>France</td>
        </tr>
    </tbody>
</table>

```sql
SELECT DISTINCT(s.company_name), country 
FROM suppliers as s 
JOIN customers as c USING(country) 
ORDER BY s.company_name  
LIMIT 4

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>company_name</th>
            <th>country</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Aux joyeux ecclésiastiques</td>
            <td>France</td>
        </tr>
        <tr>
            <td>Bigfoot Breweries</td>
            <td>USA</td>
        </tr>
        <tr>
            <td>Cooperativa de Quesos &#x27;Las Cabras&#x27;</td>
            <td>Spain</td>
        </tr>
        <tr>
            <td>Escargots Nouveaux</td>
            <td>France</td>
        </tr>
    </tbody>
</table>

```sql
SELECT category_name, SUM(units_in_stock)
FROM products
INNER JOIN categories USING(category_id)
GROUP BY category_name
ORDER BY sum desc
LIMIT (SELECT MIN(product_id)+4 FROM PRODUCTS)

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>category_name</th>
            <th>sum</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Seafood</td>
            <td>701</td>
        </tr>
        <tr>
            <td>Beverages</td>
            <td>559</td>
        </tr>
        <tr>
            <td>Condiments</td>
            <td>507</td>
        </tr>
        <tr>
            <td>Dairy Products</td>
            <td>393</td>
        </tr>
        <tr>
            <td>Confections</td>
            <td>386</td>
        </tr>
    </tbody>
</table>

```sql
SELECT product_name, units_in_stock
FROM products
WHERE units_in_stock > (SELECT AVG(units_in_stock) FROM products)
ORDER BY units_in_stock DESC
LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_name</th>
            <th>units_in_stock</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Rhönbräu Klosterbier</td>
            <td>125</td>
        </tr>
        <tr>
            <td>Boston Crab Meat</td>
            <td>123</td>
        </tr>
        <tr>
            <td>Grandma&#x27;s Boysenberry Spread</td>
            <td>120</td>
        </tr>
        <tr>
            <td>Pâté chinois</td>
            <td>115</td>
        </tr>
        <tr>
            <td>Sirop d&#x27;érable</td>
            <td>113</td>
        </tr>
    </tbody>
</table>

### `EXISTS`

```sql
-- join like experience

SELECT company_name, contact_name
FROM customers
WHERE EXISTS (
    SELECT customer_id FROM orders WHERE customer_id = customers.customer_id and freight > 800
)

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>company_name</th>
            <th>contact_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Queen Cozinha</td>
            <td>Lúcia Carvalho</td>
        </tr>
        <tr>
            <td>QUICK-Stop</td>
            <td>Horst Kloss</td>
        </tr>
        <tr>
            <td>Save-a-lot Markets</td>
            <td>Jose Pavarotti</td>
        </tr>
    </tbody>
</table>

```sql
SELECT c.company_name, c.contact_name, o.freight
FROM customers as c
INNER JOIN orders as o ON (o.customer_id = c.customer_id and o.freight > 800)

> 4 rows affected.
```

<table>
    <thead>
        <tr>
            <th>company_name</th>
            <th>contact_name</th>
            <th>freight</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Queen Cozinha</td>
            <td>Lúcia Carvalho</td>
            <td>890.78</td>
        </tr>
        <tr>
            <td>QUICK-Stop</td>
            <td>Horst Kloss</td>
            <td>1007.64</td>
        </tr>
        <tr>
            <td>QUICK-Stop</td>
            <td>Horst Kloss</td>
            <td>810.05</td>
        </tr>
        <tr>
            <td>Save-a-lot Markets</td>
            <td>Jose Pavarotti</td>
            <td>830.75</td>
        </tr>
    </tbody>
</table>

### `ANY`

```sql
SELECT DISTINCT(company_name)
FROM customers
JOIN orders USING (customer_id)
JOIN order_details USING (order_id)
WHERE quantity > 100

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>company_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Ernst Handel</td>
        </tr>
        <tr>
            <td>Save-a-lot Markets</td>
        </tr>
        <tr>
            <td>QUICK-Stop</td>
        </tr>
    </tbody>
</table>

```sql
SELECT DISTINCT(company_name)
FROM customers
WHERE customer_id = ANY(SELECT customer_id FROM orders JOIN order_details USING(order_id) WHERE quantity > 100)

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>company_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Ernst Handel</td>
        </tr>
        <tr>
            <td>Save-a-lot Markets</td>
        </tr>
        <tr>
            <td>QUICK-Stop</td>
        </tr>
    </tbody>
</table>

```sql
SELECT DISTINCT p.product_name, od.quantity
FROM products as p
JOIN order_details as od USING(product_id)
WHERE od.quantity > (SELECT AVG(quantity) from order_details)
ORDER BY quantity
LIMIT 3

> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_name</th>
            <th>quantity</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Jack&#x27;s New England Clam Chowder</td>
            <td>24</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>24</td>
        </tr>
        <tr>
            <td>Pavlova</td>
            <td>24</td>
        </tr>
    </tbody>
</table>

## Views

```sql
CREATE VIEW view_products_of_suppliers as 
SELECT product_name, company_name, units_in_stock 
FROM products p
INNER JOIN suppliers s ON (p.supplier_id = s.supplier_id )

> Done.

result >>> []
```

```sql
SELECT * FROM view_products_of_suppliers LIMIT 5

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_name</th>
            <th>company_name</th>
            <th>units_in_stock</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Chai</td>
            <td>Specialty Biscuits, Ltd.</td>
            <td>39</td>
        </tr>
        <tr>
            <td>Chang</td>
            <td>Exotic Liquids</td>
            <td>17</td>
        </tr>
        <tr>
            <td>Aniseed Syrup</td>
            <td>Exotic Liquids</td>
            <td>13</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Cajun Seasoning</td>
            <td>New Orleans Cajun Delights</td>
            <td>53</td>
        </tr>
        <tr>
            <td>Chef Anton&#x27;s Gumbo Mix</td>
            <td>New Orleans Cajun Delights</td>
            <td>0</td>
        </tr>
    </tbody>
</table>

### Droping Views

```sql
DROP VIEW view_products_of_suppliers;

> Done.

result >>> []
```

### Create Or Replace Views

```sql
CREATE OR REPLACE VIEW view_products_of_suppliers as 
SELECT  company_name, product_name, units_in_stock 
FROM products p
INNER JOIN suppliers s ON (p.supplier_id = s.supplier_id )

> Done.

result >>> []
```

https://www.postgresql.org/docs/13/sql-alterview.html

```sql
ALTER VIEW IF EXISTS view_products_of_suppliers 
RENAME TO products_of_suppliers_relation;

> Done.

result >>> []
```

### Inserting Into the Views

```sql
DROP VIEW IF EXISTS "empl_mrs";
CREATE VIEW "empl_mrs" as SELECT * FROM employees WHERE title_of_courtesy = 'Mrs'

> Done.
> Done.

result >>> []
```

```sql
INSERT INTO employees (employee_id, first_name, last_name, title_of_courtesy)
VALUES (12, 'Johnna', 'Dow', 'Ms')
RETURNING title_of_courtesy, first_name, last_name

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>title_of_courtesy</th>
            <th>first_name</th>
            <th>last_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Ms</td>
            <td>Johnna</td>
            <td>Dow</td>
        </tr>
    </tbody>
</table>

#### `TODO`: Adding checks to prevent bypassing filters

* WITH LOCAL CHECK OPTION;
* WITH CASCADE CHECK OPTION;

## `TODO` Importing and Exporting data

## `TODO` User Variables

## `TODO` Prepared Statements

## Stored Procedures and Functions

### Procedures

```sql
DROP PROCEDURE IF EXISTS test_proc();

CREATE PROCEDURE test_proc()
as $$ 
    BEGIN
        CREATE TABLE IF NOT EXISTS a (aid int);
        CREATE TABLE IF NOT EXISTS b (bid int); 
    COMMIT;
        CREATE TABLE IF NOT EXISTS c (cid int);
    ROLLBACK;
    END; 
$$ LANGUAGE plpgsql;

> Done.
> Done.

result >>> []
```

```sql
CALL test_proc();

> (psycopg2.errors.InvalidTransactionTermination) invalid transaction termination
> CONTEXT:  PL/pgSQL function test_proc() line 5 at COMMIT
> 
> [SQL: CALL test_proc();]
> (Background on this error at: https://sqlalche.me/e/20/2j85)
```

```sql
DROP table IF EXISTS a;
DROP table IF EXISTS b;
DROP table IF EXISTS c;

> Done.
> Done.
> Done.

result >>> []
```

### Functions: SQL

Selecting data from table

```sql
CREATE OR REPLACE FUNCTION get_product_price_by_name(name varchar) RETURNS real as $$
    SELECT unit_price FROM products WHERE product_name = name
$$  LANGUAGE SQL;

SELECT * FROM get_product_price_by_name('Tofu')

> Done.
> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>get_product_price_by_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>23.25</td>
        </tr>
    </tbody>
</table>

```sql
CREATE OR REPLACE FUNCTION get_total_number_of_goods() RETURNS bigint as $$
    SELECT sum(units_in_stock) FROM products
$$ LANGUAGE sql;

SELECT get_total_number_of_goods();

> Done.
> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>get_total_number_of_goods</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>3119</td>
        </tr>
    </tbody>
</table>

#### `OUT` parameters

```sql
CREATE OR REPLACE FUNCTION get_product_priceboundaries(OUT max_price real, OUT min_price real) as $$
    SELECT MAX(unit_price), MIN(unit_price) FROM products
$$  LANGUAGE SQL;

> Done.

result >>> []
```

```sql
SELECT get_product_priceboundaries()

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>get_product_priceboundaries</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>(263.5,2.5)</td>
        </tr>
    </tbody>
</table>

```sql
SELECT * FROM get_product_priceboundaries()

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>max_price</th>
            <th>min_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>263.5</td>
            <td>2.5</td>
        </tr>
    </tbody>
</table>

```sql
SELECT min_price, max_price FROM get_product_priceboundaries()

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>min_price</th>
            <th>max_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>2.5</td>
            <td>263.5</td>
        </tr>
    </tbody>
</table>

#### DEFAULT parameters

```sql
CREATE OR REPLACE FUNCTION get_product_price_boundaries_by_discontinueity(
    IN is_discontinued INT DEFAULT 0, 
    OUT max_price real, 
    OUT min_price real
) as $$
    SELECT MAX(unit_price), MIN(unit_price) FROM products
    WHERE discontinued = is_discontinued
$$  LANGUAGE SQL;

> Done.

result >>> []
```

```sql
SELECT * FROM get_product_price_boundaries_by_discontinueity()

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>max_price</th>
            <th>min_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>263.5</td>
            <td>2.5</td>
        </tr>
    </tbody>
</table>

```sql
SELECT * FROM get_product_price_boundaries_by_discontinueity(0)

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>max_price</th>
            <th>min_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>263.5</td>
            <td>2.5</td>
        </tr>
    </tbody>
</table>

```sql
SELECT * FROM get_product_price_boundaries_by_discontinueity(1)

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>max_price</th>
            <th>min_price</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>123.79</td>
            <td>4.5</td>
        </tr>
    </tbody>
</table>

#### `DECLARE` Variables

```sql
CREATE OR REPLACE FUNCTION get_triangle_square(a real, b real, c real) RETURNS real as $$
    
    -- p for perimeter
    DECLARE p real;
    BEGIN
        p = (a+b+c) / 2;
        return sqrt(p) * (p-a) * (p-b) * (p-c);
    END;
$$  LANGUAGE plpgsql;

> Done.

result >>> []
```

```sql
SELECT get_triangle_square(6,6,6)

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>get_triangle_square</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>81.0</td>
        </tr>
    </tbody>
</table>

#### `TODO` IF/ELSEIF

#### SETOF `[table]`

```sql
CREATE OR REPLACE FUNCTION get_price_in_range() RETURNS SETOF products as $$
    
    DECLARE 
        avg_price real;
        max_price real;
    
    BEGIN
        SELECT avg(unit_price) 
        INTO avg_price 
        FROM products;
        
        -- in case if need somewhere else
        max_price = avg_price * 1.05 ;  
        
        RETURN QUERY
        SELECT * FROM products
        WHERE unit_price BETWEEN (avg_price * 0.95) and max_price;
    END;
$$  LANGUAGE plpgsql;

SELECT product_id, product_name FROM get_price_in_range();

> Done.
> 2 rows affected.
```

<table>
    <thead>
        <tr>
            <th>product_id</th>
            <th>product_name</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>7</td>
            <td>Uncle Bob&#x27;s Organic Dried Pears</td>
        </tr>
        <tr>
            <td>61</td>
            <td>Sirop d&#x27;érable</td>
        </tr>
    </tbody>
</table>

### `RETURN NEXT`

```sql
CREATE OR REPLACE FUNCTION return_ints() RETURNS SETOF int as $$
    BEGIN
        RETURN NEXT 1;
        RETURN NEXT 3;
        RETURN NEXT 2;
    END
$$ LANGUAGE plpgsql;

SELECT return_ints();

> Done.
> 3 rows affected.
```

<table>
    <thead>
        <tr>
            <th>return_ints</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
        </tr>
        <tr>
            <td>3</td>
        </tr>
        <tr>
            <td>2</td>
        </tr>
    </tbody>
</table>

### Quick Proofing 

```sql
DROP TABLE IF EXISTS tmp_customers;
SELECT * INTO tmp_customers FROM customers;

CREATE OR REPLACE FUNCTION update_null_region_in_customers() RETURNS void as $$
    UPDATE tmp_customers SET region = 'unknown' WHERE region is NULL
$$ LANGUAGE SQL;

SELECT update_null_region_in_customers();

> Done.
> 91 rows affected.
> Done.
> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>update_null_region_in_customers</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>None</td>
        </tr>
    </tbody>
</table>

## Counter

```sql
DO $$
BEGIN
    FOR counter IN 1...10 BY 2
    LOOP
        RAISE NOTICE 'counter:... %', counter;
    END LOOP;
END$$

> Done.

result >>> []
```

## Built-in Functions

### `regexp_matches`

```sql
SELECT regexp_matches ('<user2@mail.de,user3@mail.de>', '\w[\w.+-]*@[\w.-]+\.[\w-]{2,63}', 'g');

> 0 rows affected.
```

<table>
    <thead>
        <tr>
            <th>regexp_matches</th>
        </tr>
    </thead>
    <tbody>
    </tbody>
</table>

### `NULLIF(a,b)` - returns `NULL` if `a == b`

```sql
SELECT NULLIF('foo', 'bar'), NULLIF('foo', 'foo'), NULLIF(1, 0)

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>nullif</th>
            <th>nullif_1</th>
            <th>nullif_2</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>foo</td>
            <td>None</td>
            <td>1</td>
        </tr>
    </tbody>
</table>

### `COALESCE(a,b)` - prefer `b` if `a == null`

```sql
SELECT COALESCE (Null, 'var'), COALESCE ('foo', 'bar')

> 1 rows affected.
```

<table>
    <thead>
        <tr>
            <th>coalesce</th>
            <th>coalesce_1</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>var</td>
            <td>foo</td>
        </tr>
    </tbody>
</table>

### `CASE` ... `END` - if else blocks

```sql
SELECT 
    order_id, 
    customer_id,
    quantity,
    CASE
        WHEN quantity >= 100 THEN 'a lot' 
        WHEN quantity >= 40 THEN 'average' 
        WHEN quantity >= 20 THEN 'normal amount' 
        ELSE 'too few'
    END as how_much
FROM orders 
INNER JOIN order_details USING(order_id)
WHERE quantity > 10
     and order_id in (11072, 10571)
ORDER BY how_much
LIMIT 10

> 5 rows affected.
```

<table>
    <thead>
        <tr>
            <th>order_id</th>
            <th>customer_id</th>
            <th>quantity</th>
            <th>how_much</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>11072</td>
            <td>ERNSH</td>
            <td>130</td>
            <td>a lot</td>
        </tr>
        <tr>
            <td>11072</td>
            <td>ERNSH</td>
            <td>40</td>
            <td>average</td>
        </tr>
        <tr>
            <td>10571</td>
            <td>ERNSH</td>
            <td>28</td>
            <td>normal amount</td>
        </tr>
        <tr>
            <td>11072</td>
            <td>ERNSH</td>
            <td>22</td>
            <td>normal amount</td>
        </tr>
        <tr>
            <td>10571</td>
            <td>ERNSH</td>
            <td>11</td>
            <td>too few</td>
        </tr>
    </tbody>
</table>