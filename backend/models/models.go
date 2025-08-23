package models

import (
	"time"

	"gorm.io/gorm"
)

// ==================== Users ====================
type User struct {
	ID        string         `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relasi
	SocialAccounts []SocialAccount `gorm:"foreignKey:UserID;references:ID"`
	Reports        []Report        `gorm:"foreignKey:UserID;references:ID"`
}

// ==================== SocialAccounts ====================
type SocialAccount struct {
	ID              string `gorm:"primaryKey"`
	UserID          string `gorm:"not null;index"`
	User            *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Platform        string
	AccountUsername string
	AccessToken     string
	CreatedAt       time.Time

	// Relasi
	Posts []Post `gorm:"foreignKey:AccountID;references:ID"`
}

// ==================== Posts ====================
type Post struct {
	ID             string `gorm:"primaryKey"`
	AccountID      string `gorm:"not null;index"`
	SocialAccount  *SocialAccount `gorm:"foreignKey:AccountID;references:ID;constraint:OnDelete:CASCADE;"`
	PlatformPostID string
	Content        string
	MediaURL       string
	CreatedAt      time.Time
	FetchedAt      time.Time

	// Relasi
	Comments   []Comment           `gorm:"foreignKey:PostID;references:ID"`
	Engagement *Engagement         `gorm:"foreignKey:PostID;references:ID"`
	Sentiment  *SentimentAnalysis  `gorm:"foreignKey:PostID;references:ID"`
	PostTopics []PostTopic         `gorm:"foreignKey:PostID;references:ID"`
}

// ==================== Comments ====================
type Comment struct {
	ID        string `gorm:"primaryKey"`
	PostID    string `gorm:"not null;index"`
	Post      *Post  `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE;"`
	UserName  string
	Content   string
	CreatedAt time.Time
}

// ==================== Engagements ====================
type Engagement struct {
	ID            string `gorm:"primaryKey"`
	PostID        string `gorm:"uniqueIndex"`
	Post          *Post  `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE;"`
	Likes         int
	Shares        int
	CommentsCount int
	Views         int
	FetchedAt     time.Time
}

// ==================== SentimentAnalysis ====================
type SentimentAnalysis struct {
	ID         string `gorm:"primaryKey"`
	PostID     string `gorm:"uniqueIndex"`
	Post       *Post  `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE;"`
	Sentiment  string
	Confidence float64
	AnalyzedAt time.Time
}

// ==================== Topics ====================
type Topic struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Keywords  string

	// Relasi
	PostTopics []PostTopic `gorm:"foreignKey:TopicID;references:ID"`
}

// ==================== PostTopics (many-to-many) ====================
type PostTopic struct {
	ID      string `gorm:"primaryKey"`
	PostID  string `gorm:"not null;index"`
	TopicID string `gorm:"not null;index"`
	Post    *Post  `gorm:"foreignKey:PostID;references:ID;constraint:OnDelete:CASCADE;"`
	Topic   *Topic `gorm:"foreignKey:TopicID;references:ID;constraint:OnDelete:CASCADE;"`
}

// ==================== Trends ====================
type Trend struct {
	ID                    string `gorm:"primaryKey"`
	Hashtag               string
	SentimentDistribution string `gorm:"type:jsonb"`
	MentionsCount         int
	PeriodStart           time.Time
	PeriodEnd             time.Time
}

// ==================== Reports ====================
type Report struct {
	ID          string `gorm:"primaryKey"`
	UserID      string `gorm:"index"`
	User        *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:SET NULL;"`
	Title       string
	Description string
	ReportURL   string
	GeneratedAt time.Time
}
