# Table: heroku_formations

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| command | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| quantity | int | X | √ |  | 
| size | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| app | json | X | √ |  | 


