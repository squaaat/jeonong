package model

type Articles struct {
	DefaultModel
}

type ArticlesDiscounts struct {
	DefaultModel

	ArticleID string   `gorm:"primaryKey;type:CHAR(36);not null"`
	Article   Articles `gorm:"foreignKey:ArticleID"`

	SaleID string `gorm:"primaryKey;type:CHAR(36);not null"`
	Sale   Sales  `gorm:"foreignKey:SaleID"`
}

type Sales struct {
	DefaultModel

	ProductID string   `gorm:"type:CHAR(36);not null"`
	Product   Products `gorm:"foreignKey:ProductID"`

	MarketID string  `gorm:"type:CHAR(36);not null"`
	Market   Markets `gorm:"foreignKey:MarketID"`
}

type Products struct {
	DefaultModel

	CategoryID string     `gorm:"type:CHAR(36);not null"`
	Category   Categories `gorm:"foreignKey:CategoryID"`

	ManufactureID string       `gorm:"type:CHAR(36);not null"`
	Manufacture   Manufactures `gorm:"foreignKey:ManufactureID"`
}

type Categories struct {
	DefaultModel
}

type CategoryHierarchy struct {
	DefaultModel

	CategoryID string     `gorm:"type:CHAR(36);not null"`
	Category   Categories `gorm:"foreignKey:CategoryID"`

	ParentCategoryID string     `gorm:"type:CHAR(36);not null"`
	ParentCategory   Categories `gorm:"foreignKey:ParentCategoryID"`
}

type Manufactures struct {
	DefaultModel
}

type Markets struct {
	DefaultModel
}
