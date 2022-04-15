package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/sayo/Genperson/Gen_series"
	"github.com/sayo/Genperson/Person"
)

func main() {
	rand.Seed(time.Now().Unix())
	age := Gen_series.Choose(16, 60)
	gender := Gen_series.Choose(0, 1)
	per := person.Person{
		Id:         Gen_series.Gen_id(),
		Name:       Gen_series.Gen_name(),
		Pinyin:     "",
		Age:        age,
		Sex:        gender,
		Sex_str: 	"",
		IdCard:     "",
		Mobile:     Gen_series.Gen_mobile(),
		OrgCode:    Gen_series.Gen_orgcode(),
		CreditCode: Gen_series.Gen_creditcode(),
	}
	per.Pinyin=Gen_series.Gen_pinyin(per.Name)
	per.IdCard= Gen_series.Gen_id_card(per.Age,per.Sex)
	sex:=""
	if per.Sex ==1 {
		sex="男"
	}else if per.Sex==0{
		sex="女"
	}
	per.Sex_str=sex
	json_data,err:=json.MarshalIndent(per,"","\t")
	if err != nil{
		fmt.Println("josn err",err)
	}
	fmt.Printf("%s\n", json_data)
}
