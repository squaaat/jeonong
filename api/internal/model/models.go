package model

type Article struct {
	DefaultModel
}

func (m *Article) TableName() string {
	return "j_article"
}

type ArticlesDiscount struct {
	DefaultModel

	ArticleID string  `gorm:"primaryKey;type:CHAR(36);not null"`
	Article   Article `gorm:"foreignKey:ArticleID"`

	SaleID string `gorm:"primaryKey;type:CHAR(36);not null"`
	Sale   Sale   `gorm:"foreignKey:SaleID"`
}

func (m *ArticlesDiscount) TableName() string {
	return "j_articles_discount"
}

type Sale struct {
	DefaultModel

	ProductID string  `gorm:"type:CHAR(36);not null"`
	Product   Product `gorm:"foreignKey:ProductID"`

	MarketID string `gorm:"type:CHAR(36);not null"`
	Market   Market `gorm:"foreignKey:MarketID"`
}

func (m *Sale) TableName() string {
	return "j_sale"
}

type Product struct {
	DefaultModel

	CategoryID string   `gorm:"type:CHAR(36);not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`

	ManufactureID string      `gorm:"type:CHAR(36);not null"`
	Manufacture   Manufacture `gorm:"foreignKey:ManufactureID"`
}

func (m *Product) TableName() string {
	return "j_product"
}

type Keyword struct {
	DefaultModel

	Name string `gorm:"type:VARCHAR(100);not null;uniqueIndex"`
	Code string `gorm:"type:VARCHAR(100);not null"`
}

func (m *Keyword) TableName() string {
	return "j_keyword"
}

type Category struct {
	DefaultModel

	ParentCategoryID string `gorm:"type:CHAR(36);not null"`
	ParentCategory interface{} `gorm:"foreignKey:KeywordID"`

	ParentKeywordID string  `gorm:"type:CHAR(36)"`
	ParentKeyword   Keyword `gorm:"foreignKey:ParentKeywordID"`

	KeywordID string  `gorm:"type:CHAR(36);not null"`
	Keyword   Keyword `gorm:"foreignKey:KeywordID"`
}

func (m *Category) TableName() string {
	return "j_category"
}

type Manufacture struct {
	DefaultModel

	KeywordID string  `gorm:"type:CHAR(36);not null"`
	Keyword   Keyword `gorm:"foreignKey:KeywordID"`

	CompanyRegistrationNumber string `gorm:"type:VARCHAR(100)"`
}

func (m *Manufacture) TableName() string {
	return "j_manufacture"
}

type Market struct {
	DefaultModel
}

func (m *Market) TableName() string {
	return "j_market"
}
