package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
)

func main() {
	http.HandleFunc("/receive", func(writer http.ResponseWriter, request *http.Request) {
		compressed, err := ioutil.ReadAll(request.Body)
		if nil != err {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		reqBuf, err := snappy.Decode(nil, compressed)
		if nil != err {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		var req prompb.WriteRequest
		if err := proto.Unmarshal(reqBuf, &req); err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		for _, ts := range req.Timeseries {
			m := make(model.Metric, len(ts.Labels))
			for _, l := range ts.Labels {
				m[model.LabelName(l.Name)] = model.LabelValue(l.Value)
			}

			fmt.Println(m)

			//for _, s := range ts.Samples {
			//	fmt.Printf(" %f %d", s.Value, s.Timestamp)
			//}
		}
	})

	http.ListenAndServe(":1234", nil)
}
