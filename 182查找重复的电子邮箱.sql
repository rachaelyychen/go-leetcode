# 编写一个 SQL 查询，查找 Person 表中所有重复的电子邮箱。
select a.Email from (select count(Email) as cnt, Email from Person group by Email) a where a.cnt > 1
