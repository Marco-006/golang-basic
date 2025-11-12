# 1. go.mod文件中的require依赖，理论上是自动生成的； 
# 2. 执行的时候打印SQL，单条SQL语句： db.Debug(); 另外还有全局的配置
# 3. 重点参考官方指导文档
# 4. 这个一对多的关系中，User是“一”，CreditCard是“多”；  并没有表现在User结构体中的字段上，而是通过CreditCard结构体中的UserID字段来体现的
# 5. 更新用户并完全更新其所有关联 db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
# 6. 预加载全部  clause.Associations
# 7. 拓展： golang，GORM， 如何保证初始化表结构的语句只执行一次？
# 8. 引用的时候，真正有效果的是，go文件首行的：package task6gormadvance
# 9. 使用range的时候，不知道应该使用一个参数  还是两个参数? 用一个参数index
# 10. SQL高阶查询Traditional API中， Raw()适合SELECT查询和需要处理结果集的场景，而Exec()适合INSERT、UPDATE、DELETE等数据变更操作
# 11. GORM的官方文档，不要再花时间去逐行读，遇到问题再深入地查；减少时间浪费
# 12. Hook ：在指定节点前后  分别执行逻辑
# 13. 当子类钩子没有启用的时候：父类中的钩子  覆盖子类中的
# 14. sqlx
## sqlx = sql + 便捷映射 + 命名参数，零魔法、零学习成本，适合不想用 ORM又不想手写 rows.Scan 地狱的场景。
## sqlx 需要引入mysql
# 15. GORM, 打印的时候不友好
