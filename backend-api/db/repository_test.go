package db

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/yusukemisa/go-vue-pay-grpc/backend-api/domain"
)

func TestSelectAllItems(t *testing.T) {
	want := &domain.Items{
		{
			ID:          1,
			Name:        "toy",
			Description: "test-toy",
			Amount:      2000,
		},
		{
			ID:          2,
			Name:        "game",
			Description: "test-game",
			Amount:      6000,
		},
	}
	items, err := SelectAllItems()
	if err != nil {
		t.Fatalf("DB Connect Fail:%v", err.Error())
	}
	log.Printf(toString(items))
	if !reflect.DeepEqual(want, &items) {
		t.Fatalf("want %v, got %v\n", toString(want), toString(items))
	}
}

func toString(v interface{}) string {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		var strs []string
		for i := 0; i < s.Len(); i++ {
			strs = append(strs, fmt.Sprintf("%+v", s.Index(i)))
		}
		return "[" + strings.Join(strs, " ") + "]"
	default:
		return fmt.Sprintf("%+v", v)
	}
}
