base项目

表a, id自增主键, date索引
语句1: select * from a where id > x and id < y and date < xx
语句2: select * from a where date < xx and id > x and id < y
语句3: select * from a where id between x and y and date < xx
语句4: select * from a where date < xx and id between x and y

1, 2, 3, 4分别会使用哪些索引