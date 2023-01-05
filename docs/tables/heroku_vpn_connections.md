# Table: heroku_vpn_connections

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| public_ip | string | X | √ |  | 
| routable_cidrs | string_array | X | √ |  | 
| space_cidr_block | string | X | √ |  | 
| status | string | X | √ |  | 
| tunnels | json | X | √ |  | 
| ike_version | int | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| status_message | string | X | √ |  | 


