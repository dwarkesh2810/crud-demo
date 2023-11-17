package main

import (
	"crud/pkg/config"
	"crud/pkg/routes"
)

func init() {
	config.LoadEnvVariable()
	// config.Init()
	// routes.Init()
}

func main() {
	routes.RegisterRoutes()
	routes.Serve()

	// // // // var posts []test1

	// // // // result := config.DB.Raw("SELECT body, created_at, id, title, user_id, category_type FROM posts JOIN categories ON posts.category_id = categories.category_id;").Scan(&posts)
	// // // // if result.Error != nil {
	// // // // 	log.Print(result.Error)
	// // // // }
	// // // // log.Print(posts)

	// var posts postmodel.Posts

	// // // // config.DB.Table("posts").Select("id, user_id as UserID, title, body, categories.category_type as CategoryType, created_at").
	// // // // 	Joins("LEFT JOIN categories ON posts.category_id = categories.category_id").
	// // // // 	Find(&posts)
	// // // // config.DB.Table("posts").
	// // // // 	Select("id, user_id as UserID, title, body, categories.category_type as CategoryType, created_at").
	// // // // 	Joins("LEFT JOIN categories ON posts.category_id = categories.id").
	// // // // 	Find(&posts)
	// // category := make(map[int]string)
	// i := 1
	// for i < 4 {
	// 	config.DB.Preload("Category").Find(&posts)
	// 	data := posts.Category
	// 	newData := helper.StructToMap(data)

	// 	d := newData["CategoryID"]
	// 	v := newData["CategoryType"]
	// 	log.Print(d)
	// 	log.Print(v)

	// 	log.Printf("%T, %T", d, v)
	// 	// keys := make([]string, 0, len(newData))
	// 	// for k := range newData {
	// 	// 	keys = append(keys, k)
	// 	// }
	// 	// log.Print(keys)
	// // 	i++
	// myMap := make(map[uint]string)
	// myMap[1] = "vbchjvhjvbvbjvbkjbvnj"

	// log.Print(myMap[1])
	// to := []string{"dwarkesh0007@gmail.com", "dwarkeshbpatel@gmail.com"}
	// helper.SendMail(to, "test mail", "hello")
	// from := os.Getenv("FROM")
	// pas := os.Getenv("PASSWORD")

	// sender := mail.NewGmailSender(from, pas)

	// subject := "A test mail"
	// content := "Hello!"

	// to := []string{"dwarkesh0007@gmail.com", "dwarkeshbpatel@gmail.com"}
	// attachFile := []string{"README.md"}

	// err := sender.SendEmail(subject, content, to, nil, nil, attachFile)
	// if err != nil {
	// 	fmt.Println("failed to send mail", err)
	// }
}
