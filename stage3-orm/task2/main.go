package task2

// import (
// 	"fmt"

// 	"gorm.io/gorm"
// )

// // 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
// // （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

// // 要求 ：
// // 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
// // 向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

// type Account struct {
// 	gorm.Model
// 	Balance float64
// }

// type Transaction struct {
// 	gorm.Model
// 	FromAccountID uint
// 	ToAccountID   uint
// 	Amount        float64
// }

// func Run(db *gorm.DB) {
// 	// db.AutoMigrate(&Account{})
// 	// db.AutoMigrate(&Transaction{})

// 	//造数据
// 	db.Create(&Account{Balance: 50})
// 	db.Create(&Account{Balance: 50})

// 	TransactionFunc(db, 1, 2, 10)
// }

// func TransactionFunc(db *gorm.DB, fromAccountID uint, toAccountID uint, amount float64) {

// 	var accountA Account
// 	db.Debug().Where("id = ?", fromAccountID).Find(&accountA)
// 	fmt.Println(accountA)

// 	if accountA.Balance >= amount {
// 		db.Debug().Create(&Transaction{FromAccountID: fromAccountID, ToAccountID: toAccountID, Amount: amount})
// 		db.Debug().Model(&accountA).Update("balance", accountA.Balance-amount)
// 	} else {
// 		fmt.Println("余额不足")
// 	}
// }
