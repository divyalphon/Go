package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

const token = "7pYOFkV6ws8REsviV9jiSOPN"

var names = []byte("names")

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gettingdata(w http.ResponseWriter, r *http.Request) {

	url := "http://127.0.0.1:8200/v1/secret/data/baz" //path is the key

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("X-Vault-Token", token)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error is coming.....")
		panic(err)
	} else {
		fmt.Println("Inside else statement.....")
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(resp.Body)
		fmt.Fprintln(w, string(data))
		fmt.Fprintln(w, resp.Status)
		fmt.Println(string(data) + "     " + "this is data")

		mydata := (string(data))
		r := make(map[string]interface{})
		json.Unmarshal([]byte(mydata), &r)
		fmt.Println(mydata)
		fmt.Fprintln(w, r)

		for _, record := range r {

			log.Printf(" [===>] Record: %s", record)

			if rec, ok := record.(map[string]interface{}); ok {
				for key, val := range rec {
					log.Printf(" [========>] %s = %s", key, val)
					if key == "data" {
						for _, ree := range rec {
							log.Printf(" [===>] Record: %s", ree)

							if ree1, ok := ree.(map[string]interface{}); ok {
								for inkey, inval := range ree1 {
									log.Printf(" [============>] %s=%s", inkey, inval)

									a, err := GetBytes(inval)

									fmt.Println("++++++++++++")
									fmt.Println(a, err)

									db, err := bolt.Open("bolt.db", 0644, nil)
									if err != nil {
										log.Fatal(err)
									}
									defer db.Close()
									key := []byte(inkey)
									value := a
									err = db.Update(func(tx *bolt.Tx) error {
										bucket, err := tx.CreateBucketIfNotExists(names)
										if err != nil {
											return err
										}

										err = bucket.Put(key, value)
										if err != nil {
											return err
										}
										return nil
									})
									err = db.View(func(tx *bolt.Tx) error {
										bucket := tx.Bucket(names)
										if bucket == nil {
											return fmt.Errorf("Bucket %q not found!", names)
										}

										val := bucket.Get(key)
										fmt.Println("***************************************************************")

										fmt.Println(string(val))

										return nil
									})

									if err != nil {
										log.Fatal(err)
									}
								}
							}

						}
						fmt.Println(key, " ", val)
					}
				}
			}

		}

		if err != nil {

			fmt.Println(err)
		}
		fmt.Println("         ")

	}

	defer resp.Body.Close()

}

func main() {
	http.HandleFunc("/getting", gettingdata)
	http.ListenAndServe(":8080", nil)
}
