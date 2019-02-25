package db

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Gender 構造体
type Gender struct {
	gorm.Model
	GenderName string
}

// Hometown 構造体
type Hometown struct {
	gorm.Model
	HometownName string
}

// BloodType 構造体
type BloodType struct {
	gorm.Model
	BloodTypeName string
}

// User 構造体
type User struct {
	gorm.Model
	Name        string
	Katakana    string
	Gender      Gender
	GenderID    uint
	Tel         string
	Mail        string
	Birthday    string
	Age         uint
	Hometown    Hometown
	HometownID  uint
	BloodType   BloodType
	BloodTypeID uint
}

// DBopen DBオープン
func DBopen() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

// InitializeTable テーブル作成
func InitializeTable() {
	db := DBopen()
	defer db.Close()

	//あれば削除
	db.DropTableIfExists(&Gender{}, &Hometown{}, &BloodType{}, &User{}, "sqlite_sequence")

	db.AutoMigrate(&Gender{}, &Hometown{}, &BloodType{}, &User{})

	//性別作成
	gender := []string{
		"男",
		"女",
	}
	for _, v := range gender {
		db.Create(&Gender{GenderName: v})
	}

	//都道府県作成
	hometown := []string{
		"北海道",
		"青森県",
		"岩手県",
		"宮城県",
		"秋田県",
		"山形県",
		"福島県",
		"茨城県",
		"栃木県",
		"群馬県",
		"埼玉県",
		"千葉県",
		"東京都",
		"神奈川県",
		"新潟県",
		"富山県",
		"石川県",
		"福井県",
		"山梨県",
		"長野県",
		"岐阜県",
		"静岡県",
		"愛知県",
		"三重県",
		"滋賀県",
		"京都府",
		"大阪府",
		"兵庫県",
		"奈良県",
		"和歌山県",
		"鳥取県",
		"島根県",
		"岡山県",
		"広島県",
		"山口県",
		"徳島県",
		"香川県",
		"愛媛県",
		"高知県",
		"福岡県",
		"佐賀県",
		"長崎県",
		"熊本県",
		"大分県",
		"宮崎県",
		"鹿児島県",
		"沖縄県",
	}
	for _, v := range hometown {
		db.Create(&Hometown{HometownName: v})
	}

	//血液型作成
	bloodType := []string{
		"A",
		"B",
		"AB",
		"O",
	}
	for _, v := range bloodType {
		db.Create(&BloodType{BloodTypeName: v})
	}

	// csv読み込み
	fp, err := os.Open("personal_infomation.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(fp)

	//var err error
	i := 0
	for {
		i++
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if i == 1 {
			continue
		}

		gender := Gender{}
		hometown := Hometown{}
		bloodType := BloodType{}
		user := User{}

		db.Where("gender_name = ?", record[3]).First(&gender)
		db.Where("hometown_name = ?", record[8]).First(&hometown)
		db.Where("blood_type_name = ?", record[9]).First(&bloodType)

		user.Name = record[1]
		user.Katakana = record[2]
		user.GenderID = gender.ID
		user.Tel = record[4]
		user.Mail = record[5]
		user.Birthday = record[6]
		//https://stackoverflow.com/questions/35154875/convert-string-to-uint-in-go-lang
		u64, _ := strconv.ParseUint(record[7], 10, 32)
		user.Age = uint(u64)
		user.HometownID = hometown.ID
		user.BloodTypeID = bloodType.ID

		db.Create(&user)

		fmt.Println(record[1])
	}

	fmt.Println("完了")
}

// SelectAll Userモデルの全セレクト
func SelectAll() {
	db := DBopen()
	defer db.Close()

	users := []User{}

	db.Preload("Gender").Preload("Hometown").Preload("BloodType").Order("id desc").Find(&users)

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	fmt.Println(string(jsonBytes))
}
