package tool

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func Base64(name,img  string)  string {

	//fmt.Println(img[22:])
	//data, err := enc.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAKwAAACrCAYAAAD7NHdgAAABYWlDQ1BrQ0dDb2xvclNwYWNlRGlzcGxheVAzAAAokWNgYFJJLCjIYWFgYMjNKykKcndSiIiMUmB/yMAOhLwMYgwKicnFBY4BAT5AJQwwGhV8u8bACKIv64LMOiU1tUm1XsDXYqbw1YuvRJsw1aMArpTU4mQg/QeIU5MLikoYGBhTgGzl8pICELsDyBYpAjoKyJ4DYqdD2BtA7CQI+whYTUiQM5B9A8hWSM5IBJrB+API1klCEk9HYkPtBQFul8zigpzESoUAYwKuJQOUpFaUgGjn/ILKosz0jBIFR2AopSp45iXr6SgYGRiaMzCAwhyi+nMgOCwZxc4gxJrvMzDY7v////9uhJjXfgaGjUCdXDsRYhoWDAyC3AwMJ3YWJBYlgoWYgZgpLY2B4dNyBgbeSAYG4QtAPdHFacZGYHlGHicGBtZ7//9/VmNgYJ/MwPB3wv//vxf9//93MVDzHQaGA3kAFSFl7jXH0fsAABmmSURBVHgB7V0HzBVF1x4Eg4CIAlYUC4LYsCUWFAsvxRcx9hIBC8Eo9hKMDSsYFGMDwWDAEkURYwlKU7Eh0oyNF2ygKKi02BBF1P3Os/+/17m7M7Ntdu/svXOSzd079cyZZ2fOnGmNHCJmyUrAYAnwEN3MYD4ta1YCAQlYwAZEYh1MloAFrMm1Y3kLSMACNiAS62CyBCxgTa4dy1tAAhawAZFYB5MlYAFrcu1Y3gISsIANiMQ6mCwBC1iTa8fyFpCABWxAJNbBZAlYwJpcO5a3gAQsYAMisQ4mS8AC1uTasbwFJGABGxCJdTBZAhawJteO5S0gAQvYgEisg8kSsIA1uXYsbwEJWMAGRGIdTJaABazJtWN5C0jAAjYgEutgsgQsYDOsnT/++IPddNNN7Lnnnsswl9pKukltFTddaVetWsXmz5/PTjzxxNCE/v77b7b//vuzpUuXumEXLFjARo4cGRpvxYoV7NNPP2X19fWhYWsyALZ5WwqXwJIlS5xOnTphS7yzzTbbOARAZaQhQ4a4YRHee95//31lnOHDh5fCDho0SBm2ljz//fdfx3tYLRU8TVmvuOKKEpgAwF122cUBiEU0d+7csrAeYI877jhRcNftoYceCsQJA7g0sSrz8MCKXwvYiJXbtm3bAKD23ntvh7rwQArdu3cPhPVAO2rUqED4CRMmCMPjI7HklFrXmgfsmDFjnP79+zu33XabEheTJk0SAgogPPTQQ52ff/65FH/06NHSsAjfsmVLZ+XKlaXwNCCThsdHoqJFixY5APUpp5ziLF++XBW00H62haXqo8FTGVC6devmfPLJJ8KKPfnkk8vCeq2l94sWde3atc5nn33mtGjRQhkWcU477TTnp59+cqZOnRoaFoAW0bhx45wmTZqU4tfV1YmCVYWbBSxV44MPPliqbA94m222mfPII4+UVTJaQ89f9bv55ptHCqdKQ+SHj4Wn33//3TnvvPOEeeEjqEaygKVavfjii4WVDtAMGDDA+fXXX926v/POO6XhRADLws1TIWbOnOl07NhRys+cOXOqEa9Wh0WtHn300dKKB+jatGnjdOjQwWncuLEyXBYAFaXZuXPnUD7Gjx9f9YCt2YkDMkkRLuS0bt06hscUIv04lJWwMoUmUIAAVTc1e8MNNzAaQDEa2LA33nhDWAWYsVqzZo3QT7cjtdC6k5SmJwMsNbvs4YcfZgcddBDr2bMne/3116VpGO9RTX0IzDwk8LIHXf+zzz5bVsw333yzLIw/jo7/559/vmsJ8DK+//77M89zjz328LJzf1evXu2a7LbbbruyvHv16lUWzvQ/VTvoeuWVV8oqhgcejPww2n/wwQeOyrDPx0n6ftVVVwkxkMeHcumllzoNDQ0OfmH1EJVh2223FfJnqmPVApZUAGEF8ZXWqFGj0DB8+CTvKvPSgQcemHn+YTy3b9/eVGwK+eIBW1U6bLNmzaiu1EQSUQdI6UuAZFtvvbU0lWOPPVbql5dHFDnlxUvcfGoOsHEFFDf8N998o4xC07hK/zw8LWDzkHKEPEyoCADyo48+EnILv7feekvol6ejCXJKWt6qssPmVRHo8tG1o/t/6aWXAgClxSiMBlhst912K9ULwHr11VczUQtMi2/ccABzHoDOS06lwut8EWq5BjpOnjzZOfjgg93FJWRLdG699VZnxowZzvr160vcfvnll5kMaAigDgHUNRF9+OGHZeYqDLBkAylYCx577DE3nigM0oW/n1588UUHZjBRHKr71GXEQnSeli1b5jz11FPOJZdc4hxwwAFO06ZNnb59+zobN27kg1XsnR90FWY97PHHHy+tKCzx69evn9O8eXNpmCQVDZACUKpRP2pR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")
	if img[22]==','{
		var enc = base64.StdEncoding
		data, err := enc.DecodeString(img[23:])
		if err != nil{
			fmt.Println(err.Error())
		}
		CreateDateDir("upload")
		//f, _ := os.OpenFile("E:\\newsimg"+"\\"+time.Now().Format("2006-1-02")+"\\"+name+".png",os.O_RDWR|os.O_CREATE,os.ModePerm)
		f, _ := os.OpenFile("upload\\"+time.Now().Format("20060102")+"/"+name+"."+img[11:15],os.O_RDWR|os.O_CREATE,os.ModePerm)

		defer f.Close()
		f.Write(data)
		var imgurl string= "/upload/"+time.Now().Format("20060102")+"/"+name+"."+img[11:15]
		return imgurl
	}else{
		var enc = base64.StdEncoding
		data, err := enc.DecodeString(img[22:])
		if err != nil{
			fmt.Println(err.Error())
		}
		CreateDateDir("upload")
		//f, _ := os.OpenFile("E:\\newsimg"+"\\"+time.Now().Format("2006-1-02")+"\\"+name+".png",os.O_RDWR|os.O_CREATE,os.ModePerm)
		f, _ := os.OpenFile("upload\\"+time.Now().Format("20060102")+"/"+name+"."+img[11:14],os.O_RDWR|os.O_CREATE,os.ModePerm)

		defer f.Close()
		f.Write(data)
		var imgurl string= "/upload/"+time.Now().Format("20060102")+"/"+name+"."+img[11:14]
		return imgurl
	}
	//data, err := enc.DecodeString(img[22:])
}
