## Transactions

* `all` or `nothing`

<div class="alert alert-danger"> 
    
 `%sql` does not support transaction, so lets execute query elsewhere

</div>

* `A` - Atomacity
* C
* `I` - Isolation
* `D` - Durability 

```python
%load_ext sql
```

```sql
postgresql://user:password@postgresql:5432/northwind

result >>> 'Connected: user@northwind'
```

> no matter of are you using transactions or no, `you are using transactions`.

* RAISE EXCEPTION
* BEGIN
* EXECTION WHEN

```sql
-- General View For Transaction

BEGIN; 

    -- insert 
    -- update 
    -- delete
    -- replace

COMMIT;

stdout >>> Done.
stdout >>> Done.
result >>> []
```

### Isolation Levels

![](images/isolation-levels.png)

* `MVCC` multi version control system
* timestamp based

```sql
BEGIN; 

WITH prod_update as (
    UPDATE products 
    SET discontinued = 1
    WHERE units_in_stock < 10
    RETURNING product_id
)

SELECT * INTO last_orders_on_dicontinues FROM order_details WHERE product_id IN (SELECT product_id FROM prod_update);

DROP TABLE last_orders_on_dicontinues1;

COMMIT;

SELECT * FROM last_orders_on_dicontinues;

stdout >>> Done.
stdout >>> 296 rows affected.
stdout >>> (psycopg2.errors.UndefinedTable) table "last_orders_on_dicontinues1" does not exist
stdout >>> 
stdout >>> [SQL: DROP TABLE last_orders_on_dicontinues1;]
stdout >>> (Background on this error at: http://sqlalche.me/e/14/f405)
```

```sql
DROP TABLE last_orders_on_dicontinues;

stdout >>> Done.
result >>> []
```

```sql
BEGIN; 

WITH prod_update as (
    UPDATE products 
    SET discontinued = 1
    WHERE units_in_stock < 10
    RETURNING product_id
)

SELECT * INTO last_orders_on_dicontinues FROM order_details WHERE product_id IN (SELECT product_id FROM prod_update);

DELETE FROM order_details
WHERE product_id IN (SELECT product_id FROM last_orders_on_dicontinues);

SELECT count(*) FROM last_orders_on_dicontinues;

ROLLBACK;



stdout >>> Done.
stdout >>> (psycopg2.errors.DuplicateTable) relation "last_orders_on_dicontinues" already exists
stdout >>> 
stdout >>> [SQL: WITH prod_update as (
stdout >>>     UPDATE products 
stdout >>>     SET discontinued = 1
stdout >>>     WHERE units_in_stock < 10
stdout >>>     RETURNING product_id
stdout >>> )
stdout >>> 
stdout >>> SELECT * INTO last_orders_on_dicontinues FROM order_details WHERE product_id IN (SELECT product_id FROM prod_update);]
stdout >>> (Background on this error at: http://sqlalche.me/e/14/f405)
```

```python
BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;

SAVEPOINT backup;
ROLLBACK backup;

```