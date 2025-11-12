package task2

import (
	"gorm.io/gorm"
)

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
// （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。

// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
// 向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	gorm.Model
	Balance float64
}

type Transaction struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        float64
}

func Run(db *gorm.DB) {
	// db.AutoMigrate(&Account{})
	// db.AutoMigrate(&Transaction{})
}

func TransactionFunc(db *gorm.DB, fromAccountID uint, toAccountID uint, amount float64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var accountA Account
		if err := tx.Where("id = ?", fromAccountID).First(&accountA).Error; err != nil {
			return err
		}
		if accountA.Balance < amount {
			return gorm.ErrInvalidTransaction
		}
		if err := tx.Model(&accountA).Update("balance", accountA.Balance-amount).Error; err != nil {
			return err
		}
		var accountB Account
		if err := tx.Where("id = ?", toAccountID).First(&accountB).Error; err != nil {
			return err
		}
		if err := tx.Model(&accountB).Update("balance", accountB.Balance+amount).Error; err != nil {
			return err
		}
		if err := tx.Create(&Transaction{FromAccountID: fromAccountID, ToAccountID: toAccountID, Amount: amount}).Error; err != nil {
			return err
		}
		return nil
	})
}
