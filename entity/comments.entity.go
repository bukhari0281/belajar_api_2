package entity

type Comments struct {
	ID       int64  `gorm:"primary_key:auto_increment" jason:"-"`
	Message  string `gorm:"type:text"`
	Photo_id int64  `gorm:"not null" json:"-"`
	UserID   int64  `gorm:"not null" json:"-"`
	Photo    Photo  `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	User     User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
