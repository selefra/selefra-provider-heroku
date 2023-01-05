# Table: heroku_dyno_sizes

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| dyno_units | int | X | √ |  | 
| memory | float | X | √ |  | 
| name | string | X | √ |  | 
| private_space_only | bool | X | √ |  | 
| compute | int | X | √ |  | 
| cost | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| dedicated | bool | X | √ |  | 
| id | string | √ | √ |  | 


