package entity

type Photo struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Title     string `gorm:"type:varchar(100)"`
	Caption   string `gorm:"type:text" json:"-"`
	Photo_url string `gorm:"type:text" json:"-"`
	UserID    int64  `gorm:"not null" json:"-"`
	User      User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
