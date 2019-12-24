package main

import (
	"flag"
	"fmt"
	"github.com/cch123/elasticsql"
	"github.com/olivere/elastic"
	"golang.org/x/net/context"
	"reflect"
)

var esClient *elastic.Client
var db string
var table string

func init() {
	/*host := "127.0.0.1"
	port := "9200"
	elastic.SetURL(fmt.Sprintf("http://%s:%s", host, port))*/
	var err error
	esClient, err = elastic.NewClient()
	if err != nil {
		panic(err)
	}
	flag.Parse()
	flag.StringVar(&db, "db", "mt_es_01", "数据库名称")
	flag.StringVar(&table, "table", "toby", "数据库名称")
}

func createIndexer() {
	_, err := esClient.CreateIndex(db).Do(context.Background())
	if err != nil {
		panic(fmt.Sprintf("create indexer error :%v", err))
	}
}

type UserHobby string

type People struct {
	Name  string      `json:"name"`
	Age   int         `json:age`
	Addr  string      `json:addr`
	Hobby []UserHobby `json:"hobby"`
	Mem   string      `json:"mem"`
}

func addDoc(p People, id string) {

	_, err := esClient.Index().
		Index(db).
		Type(table).
		Id(id).
		BodyJson(p).
		Refresh("wait_for").
		Do(context.Background())

	if err != nil {
		panic(fmt.Sprintf("add doc error :%v", err))
	}
}

func query(q elastic.Query) *elastic.SearchResult {
	res, err := esClient.Search().
		Index(db).                  // search in index "tweets"
		Query(q).                   // specify the query
		Sort("user.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(context.Background())    // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	return res
}

func del(id string) {
	_, err := esClient.Delete().Index(db).Type(table).Id(id).Do(context.Background())
	if err != nil {
		panic(fmt.Sprintf("del doc error :%v", err))
	}
}

func addMilti() {
	p := People{
		Name:  "拖比",
		Age:   30,
		Addr:  "宝鸡",
		Hobby: []UserHobby{"读书", "羽毛球"},
		Mem:   "这是测试数据",
	}
	addDoc(p, "1")

	p = People{
		Name:  "莉莉",
		Age:   10,
		Addr:  "宝鸡",
		Hobby: []UserHobby{"羽毛球", "篮球"},
		Mem:   "这是测试数据",
	}
	addDoc(p, "2")

	p = People{
		Name:  "萌推拖比",
		Age:   20,
		Addr:  "宝鸡",
		Hobby: []UserHobby{"足球", "篮球"},
		Mem:   "这是测试数据",
	}
	addDoc(p, "3")

	p = People{
		Name:  "和拖比heh",
		Age:   31,
		Addr:  "上海",
		Hobby: []UserHobby{"足球", "篮球"},
		Mem:   "这是测试数据",
	}
	addDoc(p, "4")

	p = People{
		Name:  "和拖比heh",
		Age:   91,
		Addr:  "宝鸡",
		Hobby: []UserHobby{"羽毛球", "篮球"},
		Mem:   "这是测试数据",
	}
	addDoc(p, "4")
}

func searchTest() {
	serach := elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("name", "拖比"))
	users := query(serach)
	var tp People
	for _, user := range users.Each(reflect.TypeOf(tp)) {
		if t, ok := user.(People); ok {
			fmt.Printf("%s,%s\n", t.Name, t.Age)
		} else {
			fmt.Println("1")
		}
	}
}

func searchTest2() {
	var sql = `
select * from t3
where name="拖比" and age > 30
and hobby in ("羽毛球") `

	dsl, _, _ := elasticsql.Convert(sql)

	res, err := esClient.Search().
		Index(db). // search in index "tweets"
		Type(table).Routing(dsl).
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	var tp People
	for _, user := range res.Each(reflect.TypeOf(tp)) {
		if t, ok := user.(People); ok {
			fmt.Printf("name:%s,age:%d,addr:%s,hobby:%v\n", t.Name, t.Age, t.Addr, t.Hobby)
		} else {
			fmt.Println("1")
		}
	}
}

func main() {

}
