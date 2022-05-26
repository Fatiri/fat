package exchange

import (
	"reflect"
	"testing"

	"github.com/fat/models"
	"github.com/fat/repository"
	"github.com/fat/usecase/webhook"
	"github.com/gin-gonic/gin"
)

type mock struct {
	config *models.Config
}

func TestIndodaxCtx_Order(t *testing.T) {
	type fields struct {
		config  *models.Config
		webhook webhook.IndodaxWebhook
	}
	type args struct {
		ctx *gin.Context
		arg repository.CreateOrderParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    repository.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IndodaxCtx{
				config:  tt.fields.config,
				webhook: tt.fields.webhook,
			}
			got, err := i.Order(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("IndodaxCtx.Order() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndodaxCtx.Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
