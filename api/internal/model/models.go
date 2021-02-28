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

type Category struct {
	DefaultModel

	Category1ID string `gorm:"type:CHAR(36);index:category1ID_and_status"`
	Category2ID string `gorm:"type:CHAR(36);index:category2ID_and_status"`
	Category3ID string `gorm:"type:CHAR(36);index:category3ID_and_status"`
	Category4ID string `gorm:"type:CHAR(36);index:category4ID_and_status"`

	Status   EnumStatus `gorm:"type:ENUM('WAIT','IDLE','INVALID','DELETED');default:'WAIT';index:category1ID_and_status,category2ID_and_status,category3ID_and_status,category4ID_and_status"`
	Sort     int64
	Depth    int64
	Code     string `gorm:"type:VARCHAR(100);index_name_and_code"`
	Name     string `gorm:"type:VARCHAR(100);index_name_and_code"`
	FullName string
}

func (m *Category) TableName() string {
	return "j_category"
}

type Manufacture struct {
	DefaultModel

	Code                      string `gorm:"type:VARCHAR(100);index_name_and_code"`
	Name                      string `gorm:"type:VARCHAR(100);index_name_and_code"`
	CompanyRegistrationNumber string `gorm:"type:VARCHAR(100);index"`
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
