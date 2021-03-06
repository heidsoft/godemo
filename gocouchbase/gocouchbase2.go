package main

import (
	"encoding/json"
	"fmt"

	"github.com/couchbase/gocb"
)

var bucket *gocb.Bucket

type Person struct {
	FirstName        string            `json:"firstname,omitempty"`
	LastName         string            `json:"lastname,omitempty"`
	SocialNetworking *SocialNetworking `json:"socialNetworking,omitempty"`
}

type SocialNetworking struct {
	Twitter string `json:"twitter,omitempty"`
	Website string `json:"website,omitempty"`
}

func updateDocument(documentId string) {
	fmt.Println("Update document by id...")
	var person Person
	_, error := bucket.Get(documentId, &person)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	jsonPerson, _ := json.Marshal(&person)
	fmt.Println(string(jsonPerson))
}


func getDocument(documentId string) {
	fmt.Println("Getting the full document by id...")
	var person Person
	_, error := bucket.Get(documentId, &person)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	jsonPerson, _ := json.Marshal(&person)
	fmt.Println(string(jsonPerson))
}

func createDocument(documentId string, person *Person) {
	fmt.Println("Upserting a full document...")
	_, error := bucket.Upsert(documentId, person, 0)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	getDocument(documentId)
	getSubDocument(documentId)
}

func getSubDocument(documentId string) {
	fmt.Println("Getting part of a document by id...")
	fragment, error := bucket.LookupIn(documentId).Get("socialNetworking").Execute()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	var socialNetworking SocialNetworking
	fragment.Content("socialNetworking", &socialNetworking)
	jsonSocialNetworking, _ := json.Marshal(&socialNetworking)
	fmt.Println(string(jsonSocialNetworking))
	upsertSubDocument(documentId, "thepolyglotdeveloper.com")
}

func upsertSubDocument(documentId string, website string) {
	fmt.Println("Upserting part of a document...")
	_, error := bucket.MutateIn(documentId, 0, 0).Upsert("socialNetworking.website", website, true).Execute()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	getDocument(documentId)
}

type Datacenter struct{
	Name        string            `json:"name"`
	Description string            `json:"description"`
}
type DcMeta struct {
	Class       string            `json:"class"`
	Description string            `json:"description"`
	Data 	    []*Datacenter     `json:"data"`
}

type Class struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Inherits 	string          `json:"inherits"`
	Superclass 	bool        `json:"superclass"`
	Active 		bool            `json:"active"`
	CiData 	   	[]*Ci 	 `json:"cidata"`
}
type Ci struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Class 	    string     `json:"class"`
	Show 	    bool       `json:"show"`
	Active 	    bool       `json:"active"`
	Editmode    string     `json:"editmode"`
}



func main() {
	fmt.Println("Starting the app...")
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucket, _ = cluster.OpenBucket("default", "")
	//person := Person{FirstName: "liu", LastName: "bin", SocialNetworking: &SocialNetworking{Twitter: "heidsoft"}}
	//createDocument("liubin", &person)

	//cas,err := bucket.Upsert("jake", person, 0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//person.FirstName = "wen"
	//_,err = bucket.Replace("jake",person,cas,0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}

	//bucket.Remove("datacenter_meta",0)
	//dcs := []*Datacenter{} //空切片
	//
	//dc1 := new(Datacenter)
	//dc1.Name = "上海"
	//dc1.Description = "上海数据中心"
	//dcs = append(dcs,dc1)
	//
	//dc2 := new(Datacenter)
	//dc2.Name = "北京"
	//dc2.Description = "北京数据中心"
	//dcs = append(dcs,dc2)
	//
	//dc3 := new(Datacenter)
	//dc3.Name = "新加坡"
	//dc3.Description = "新加坡数据中心"
	//dcs = append(dcs,dc3)
	//
	//dcMeta := new(DcMeta)
	//dcMeta.Description = "数据中心元数据"
	//dcMeta.Class = "datacenter"
	//dcMeta.Data = dcs
	//
	//fmt.Println("插入数据中心元数据....")
	//cas,err := bucket.Upsert("datacenter_meta", dcMeta, 0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//
	//dc4 := new(Datacenter)
	//dc4.Name = "纽约"
	//dc4.Description = "纽约数据中心"
	//dcs = append(dcs,dc4)
	//dcMeta.Data = dcs
	//
	//fmt.Println("更新数据中心元数据....")
	//_,err = bucket.Replace("datacenter_meta",dcMeta,cas,0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}

	//var myDcMeta DcMeta
	//_, error := bucket.Get("datacenter_meta", &myDcMeta)
	//if error != nil {
	//	fmt.Println(error.Error())
	//	return
	//}
	//jsonMyDcMeta, _ := json.Marshal(&myDcMeta)
	//fmt.Println(string(jsonMyDcMeta))



	//CiList := []*Ci{}
	//c1 := new(Ci)
	//c1.Class="Host"
	//c1.Description="CPU"
	//c1.Active=true
	//c1.Editmode="只读"
	//c1.Name="cpu"
	//c1.Show=true
	//
	//CiList = append(CiList,c1)
	//
	//c2 := new(Ci)
	//c2.Class="Host"
	//c2.Description="内存"
	//c2.Active=true
	//c2.Editmode="只读"
	//c2.Name="memory"
	//c2.Show=true
	//CiList = append(CiList,c2)
	//
	//hostClass := Class{
	//	Name:"host",
	//	Description:"DELL主机",
	//	Inherits:"host",
	//	Superclass:false,
	//	Active:true,
	//	CiData:CiList,
	//}
	//
	//_,err := bucket.Upsert("class_host_dell", hostClass, 0)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("创建资产成功")


	//json 方式
	_,err := bucket.Upsert("map_json2",map[string]string{
		"a":"1",
		"b":"2",
		"c":"3",
	}, 0)

	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("创建资产成功")

	var jsonMap map[string]string
	cas, error := bucket.Get("map_json2", &jsonMap)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	json1, _ := json.Marshal(&jsonMap)
	fmt.Println(string(json1))

	jsonMap["d"]="4"
	jsonMap["e"]="5"
	jsonMap["a"]="100"

	cas,err = bucket.Replace("map_json2",&jsonMap,cas,0)
	bucket.ExecuteN1qlQuery()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	json2, _ := json.Marshal(&jsonMap)
	fmt.Println(string(json2))

}