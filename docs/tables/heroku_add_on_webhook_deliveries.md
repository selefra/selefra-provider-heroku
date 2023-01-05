# Table: heroku_add_on_webhook_deliveries

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_attempt | json | X | √ |  | 
| next_attempt_at | timestamp | X | √ |  | 
| num_attempts | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| updated_at | timestamp | X | √ |  | 
| webhook | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| event | json | X | √ |  | 
| id | string | √ | √ |  | 
| status | string | X | √ |  | 


