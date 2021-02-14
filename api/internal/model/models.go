package model

type Article struct {
	DefaultModel
}

type ArticlesDiscount struct {
	DefaultModel

	ArticleID string  `gorm:"primaryKey;type:CHAR(36);not null"`
	Article   Article `gorm:"foreignKey:ArticleID"`

	SaleID string `gorm:"primaryKey;type:CHAR(36);not null"`
	Sale   Sale   `gorm:"foreignKey:SaleID"`
}

type Sale struct {
	DefaultModel

	ProductID string  `gorm:"type:CHAR(36);not null"`
	Product   Product `gorm:"foreignKey:ProductID"`

	MarketID string  `gorm:"type:CHAR(36);not null"`
	Market   Markets `gorm:"foreignKey:MarketID"`
}

type Product struct {
	DefaultModel

	CategoryID string   `gorm:"type:CHAR(36);not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`

	ManufactureID string      `gorm:"type:CHAR(36);not null"`
	Manufacture   Manufacture `gorm:"foreignKey:ManufactureID"`
}

type Keyword struct {
	DefaultModel

	Name string `gorm:"type:VARCHAR(100);not null;uniqueIndex"`
	Code string `gorm:"type:VARCHAR(100);not null"`
}

type Category struct {
	DefaultModel

	KeywordID string  `gorm:"type:CHAR(36);not null"`
	Keyword   Keyword `gorm:"foreignKey:KeywordID"`

	ParentKeywordID string  `gorm:"type:CHAR(36)"`
	ParentKeyword   Keyword `gorm:"foreignKey:ParentKeywordID"`
}

type Manufacture struct {
	DefaultModel

	KeywordID string  `gorm:"type:CHAR(36);not null"`
	Keyword   Keyword `gorm:"foreignKey:KeywordID"`

	CompanyRegistrationNumber string `gorm:"type:VARCHAR(100)"`
}

type Markets struct {
	DefaultModel
}
