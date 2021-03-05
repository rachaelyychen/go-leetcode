# 编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 id 。
select a.Id from  Weather a join Weather b on a.Temperature > b.Temperature and dateDiff(a.RecordDate,b.RecordDate) = 1
