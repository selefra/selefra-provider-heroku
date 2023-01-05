# Table: heroku_add_on_region_capabilities

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | json | X | √ |  | 
| supports_private_networking | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| addon_service | json | X | √ |  | 
| id | string | √ | √ |  | 


