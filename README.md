# DBWork
A personal work with SQL databases write in GO.

# Some conclusions
• go-sql-driver/mysql dosn't work with MySQL 5.0 - there are problems with auth plugin because the <code>mysql.user</code> don't have a <code>pugin</code> column;

• the driver from <b>ziutek</b> can be used with MySQl 5.0 but only in <i>native</i> form <code>mysql.New("tcp", "", "HOST", "USER", "PASSWORD", "DATABASE")</code> available within the package github.com/ziutek/mymysql/mysql.

# Final conclusion
For this purpose #Java# seams to be a better choise.
