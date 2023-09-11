package dto

type Transaction struct {
	ID             int64   `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20);not null" json:"id"`
	OrderNumber    string  `gorm:"column:order_number;type:varchar(255);not null;comment:'第三方订单号'" json:"order_number"`                                                // 第三方订单号
	UserName       string  `gorm:"column:user_name;type:varchar(255);not null;comment:'会员账号'" json:"user_name"`                                                        // 会员账号
	OrderID        string  `gorm:"uniqueIndex:uniq_order_number_game_id;column:order_id;type:varchar(100);not null;default:0;comment:'id'" json:"order_id"`            // id
	GameID         int64   `gorm:"uniqueIndex:uniq_order_number_game_id;column:game_id;type:bigint(20);not null;comment:'场馆ID'" json:"game_id"`                        // 场馆ID
	GameSubID      int64   `gorm:"column:game_sub_id;type:bigint(20);not null;comment:'场馆下游戏ID'" json:"game_sub_id"`                                                   // 场馆下游戏ID
	BetAmount      float64 `gorm:"column:bet_amount;type:decimal(18,2);not null;default:0.00;comment:'下注金额'" json:"bet_amount"`                                        // 下注金额
	WinLose        float64 `gorm:"column:win_lose;type:decimal(18,2);not null;default:0.00;comment:'游戏输赢'" json:"win_lose"`                                            // 游戏输赢
	BetTime        int64   `gorm:"column:bet_time;type:bigint(20);not null;default:0;comment:'下注时间'" json:"bet_time"`                                                  // 下注时间
	SettleTime     int64   `gorm:"column:settle_time;type:bigint(20);not null;default:0;comment:'结算时间'" json:"settle_time"`                                            // 结算时间
	SettleStatus   int8    `gorm:"column:settle_status;type:tinyint(2);default:null;default:1;comment:'注单状态 Unsettled 1 /Settled 2/ Canceled 3'" json:"settle_status"` // 注单状态 Unsettled 1 /Settled 2/ Canceled 3
	UserID         int64   `gorm:"column:user_id;type:bigint(20);default:null;comment:'user ID'" json:"user_id"`                                                       // user ID
	ValidBetAmount float64 `gorm:"column:valid_bet_amount;type:decimal(18,2);not null;default:0.00;comment:'有效投注'" json:"valid_bet_amount"`                            // 有效投注
}
