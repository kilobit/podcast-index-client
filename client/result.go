/* Copyright 2021 Kilobit Labs Inc. */

package client

import _ "fmt"
import _ "errors"

import "encoding/json"

import . "kilobit.ca/go/objected"

type Result struct {
	Object
}

func (r *Result) UnmarshalJSON(bs []byte) error {

	return json.Unmarshal(bs, &r.Object)
}

func (r Result) Success() bool {

	status := r.GetString("status")

	return status == "true"
}

func (r Result) Description() string {

	return r.GetString("description")
}

func (r Result) Query() Value {

	val, _ := r.Get("query")

	return val
}

func (r Result) Count() int {

	n, err := r.GetNumber("count")
	if err != nil {
		n = -1
	}

	return int(n)
}

func (r Result) Feed() *Feed {

	var feed *Feed = nil

	val, ok := r.Get("feed")
	if ok {
		obj, ok := val.(map[string]interface{})
		if ok {
			feed = &Feed{obj}
		}
	}

	return feed
}

func (r Result) Feeds() []*Feed {

	feeds := []*Feed{}

	objs, ok := r.Get("feeds")
	if ok {
		vals := ToValues(objs)
		for _, val := range vals {

			obj, ok := val.(map[string]interface{})
			feed := &Feed{obj}
			if ok {
				feeds = append(feeds, feed)
			}
		}
	}

	return feeds
}

func ToString(obj interface{}) string {
	str, _ := obj.(string)
	return str
}

func ToValues(obj interface{}) Values {

	result := Values{}

	vals, ok := obj.([]interface{})
	if ok {
		result = Values(vals)
	}

	return result
}
