package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	Username         string     `json:"username" gorm:"type:varchar(60);UNIQUE;NOT NULL;index"`
	FirstName        string     `json:"firstName" gorm:"type:varchar(70);NOT NULL"`
	LastName         string     `json:"lastName" gorm:"type:varchar(70);NOT NULL"`
	Email            string     `json:"email" gorm:"type:varchar(70);UNIQUE;NOT NULL"`
	Password         string     `json:"password" gorm:"type:varchar(60);NOT NULL"`
	Photo            string     `json:"photo" gorm:"type:text"`
	Birthday         *time.Time `json:"birthday" gorm:"type:timestamp;NOT NULL"`
	Phones           []Phone
	Friends          []User            `gorm:"foreignKey:UserFriend"`
	UserAnimeActions []UserAnimeAction `gorm:"many2many:user_anime_actions"`
	Comments         []Comment         `gorm:"foreignKey:OwnerID"`
	CreatedAt        time.Time         `json:"createdAt"`
	DeletedAt        *time.Time        `json:"deletedAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
}

type Phone struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Phone     string     `json:"phone" gorm:"type:varchar(15);UNIQUE;index"`
	UserID    uint       `json:"ownerID"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type UserFriend struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	UserID     int  `json:"user_id" gorm:"column:user_id;uniqueIndex:userFriendUnique"`
	UserFriend int  `json:"friend_id" gorm:"column:friend_id;uniqueIndex:userFriendUnique"`
	IsBlocked  bool `json:"is_blocked" gorm:"type:bool;default:false"`
}

type Action struct {
	ID               uint              `json:"id" gorm:"primaryKey"`
	Name             string            `json:"name" gorm:"type:varchar(70);unique;index;not null"`
	UserAnimeActions []UserAnimeAction `gorm:"foreignKey:ActionID"`
	CreatedAt        time.Time         `json:"createdAt"`
	DeletedAt        *time.Time        `json:"deletedAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
}

type Genre struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(70);unique;index;not null"`
	Animes    []*Anime   `json:"animes" gorm:"many2many:anime_genres;"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type Tag struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(70);unique;index;not null"`
	Animes    []*Anime   `json:"animes" gorm:"many2many:anime_tags;"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type Studio struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(70);unique;index;not null"`
	Animes    []Anime    `gorm:"foreignKey:StudioID"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type AnimeGroup struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"type:varchar(70);unique;index;not null"`
	Animes    []Anime    `gorm:"foreignKey:AnimeGroupID"`
	CreatedAt time.Time  `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type Type struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name" gorm:"type varchar(70);unique;index;not null"`
	Animes []Anime `json:"animes" gorm:"foreignKey:TypeID"`
}

type Anime struct {
	ID           uint       `json:"id" gorm:"PRIMARY_KEY"`
	Name         string     `json:"name" gorm:"type:varchar(200);unique;not null"`
	Description  string     `json:"description" gorm:"type:text;not null"`
	ReleaseDate  time.Time  `json:"releaseDate" gorm:"not null"`
	AnimeGroupID uint       `json:"animeGroupID" gorm:"not null"`
	Rating       float32    `json:"ratingRelease" gorm:"type:decimal(4,2);not null;check:rating >= 0.0"`
	ViewsNumber  uint64     `json:"viewsNumber" gorm:"not null;check:views_number >= 0"`
	StudioID     uint       `json:"stdioID" gorm:"not null"`
	Genres       []*Genre   `json:"genres" gorm:"many2many:anime_genres;"`
	Tags         []*Tag     `json:"tags" gorm:"many2many:anime_tags;"`
	TypeID       uint       `json:"typeID" gorm:"not null"`
	Comments     []Comment  `json:"comments" gorm:"foreignKey:AnimeID"`
	CreatedAt    time.Time  `json:"createdAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type UserAnimeAction struct {
	UserID      int     `json:"userID" gorm:"primaryKey"`
	AnimeID     int     `json:"animeID" gorm:"primaryKey"`
	ActionID    int     `json:"actionID" gorm:"not null"`
	IsFavourite bool    `json:"isFavourite" gorm:"not null;default:false"`
	Rating      float32 `json:"userRating" gorm:"default:null;check:rating >= 0.0"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type Comment struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	Content           string     `json:"content" gorm:"type:text; not null"`
	ReppliedCommentId *uint      `json:"repliedCommentID" gorm:"default:null"`
	OwnerID           int        `json:"ownerID" gorm:"not null"`
	AnimeID           int        `json:"animeID" gorm:"not null"`
	ReppliedComments  []Comment  `json:"repliedComments" gorm:"foreignkey:ReppliedCommentId"`
	CreatedAt         time.Time  `json:"createdAt"`
	DeletedAt         *time.Time `json:"deletedAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}
