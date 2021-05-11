package main

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"lessbutter.co/mealkit/config"
	product "lessbutter.co/mealkit/domains"
	"lessbutter.co/mealkit/external"
)

func addCategories(conn *mongo.Client) {
	categories := []map[string]string{
		{
			"name":             "Hamultang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Hamultang.png",
		},
		{
			"name":             "Yukgyejang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Yukgyejang.png",
		},
		{
			"name":             "Maratang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Maratang.png",
		},
		{
			"name":             "Duonjangzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Duonjangzzigye.png",
		},
		{
			"name":             "Kimchizzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kimchizzigye.png",
		},
		{
			"name":             "Gambas",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gambas.png",
		},
		{

			"name":             "Etcjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etcjeongol.png",
		},
		{

			"name":             "Steak",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Steak.png",
		},
		{
			"name":             "Gogi",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gogi.png",
		},
		{
			"name":             "Umooktang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Umooktang.png",
		},
		{
			"name":             "Churtang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Churtang.png",
		},
		{

			"name":             "Bibbimbap",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bibbimbap.png",
		},
		{
			"name":             "Gobchangjeongol",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Gobchangjeongol.png",
		},
		{
			"name":             "Chunggukjang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Chunggukjang.png",
		},
		{
			"name":             "Budaezzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Budaezzigye.png",
		},
		{

			"name":             "Etc",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Etc.png",
		},
		{
			"name":             "Altang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Altang.png",
		},
		{
			"name":             "Myun",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Myun.png",
		},
		{
			"name":             "Millefeuille",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Millefeuille.png",
		},
		{
			"name":             "Uguzytang",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Uguzytang.png",
		},
		{
			"name":             "Bunsik",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Bunsik.png",
		},
		{

			"name":             "Pasta",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Pasta.png",
		},
		{
			"name":             "Sundubuzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Sundubuzzigye.png",
		},
		{
			"name":             "Kongbeasyzzigye",
			"categoryimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/categories/Kongbeasyzzigye.png",
		},
	}

	cats := make([]interface{}, 0)
	for _, result := range categories {
		cats = append(cats, result)
	}
	log.Println(cats)
	ret, _ := product.AddCategories(conn, cats)
	log.Println(ret)
}

// func addBrands(conn *mongo.Client) {
// 	brands := []map[string]string{
// 		{
// 			"name": "맛수러움",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/matsuruum.png",
// 		},
// 		{
// 			"name":          "프레시지",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/fresheasy.png",
// 		},
// 		{
// 			"name": "프레시몰",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/freshmeal.png",
// 		},
// 		{
// 			"name":          "푸드어셈블",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/foodassemble.png",
// 		},
// 		{
// 			"name": "네이쳐푸드",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/naturefood.png",
// 		},
// 		{
// 			"name":          "얌테이블",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/yamtable.png",
// 		},
// 		{
// 			"name":          "이츠웰",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/itswell.png",
// 		},
// 		{
// 			"name": "쿡솜씨",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/cooksomssi.png",
// 		},
// 		{
// 			"name": "피콕",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/peacock.png",
// 		},
// 		{
// 			"name": "마이셰프",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/mychef.png",
// 		},
// 		{
// 			"name": "에슐리",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/ashley.png",
// 		},
// 		{
// 			"name":          "파우즈",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/pause.png",
// 		},
// 		{
// 			"name":          "앙트레",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/entree.png",
// 		},
// 		{
// 			"name": "올쿡",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/allcook.png",
// 		},
// 		{
// 			"name": "모노키친",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/monokitchen.png",
// 		},
// 		{
// 			"name": "닥터키친",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/drkitchen.png",
// 		},
// 		{
// 			"name":          "파파쿡",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/papacook.png",
// 		},
// 		{
// 			"name": "기타",
// 			"brandimageurl": "https://mealkit.s3.ap-northeast-2.amazonaws.com/brands/momil.png",
// 		},
// 	}
// 	brandsInterface := make([]interface{}, 0)
// 	for _, result := range brands {
// 		brandsInterface = append(brandsInterface, result)
// 	}

// 	ret, _ := product.AddCategories(conn, brandsInterface)
// 	log.Println(ret)

// }

func main() {
	conf := config.GetConfiguration()
	conn := external.MongoConn(conf)
	addCategories(conn)
	// addBrands(conn)

}
