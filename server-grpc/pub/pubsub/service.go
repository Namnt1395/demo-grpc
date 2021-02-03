package pub

import (
	"github.com/docker/docker/pkg/pubsub"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}
func (p *PubsubService) Publish(ctx context.Context, arg *String, ) (*String, error) {
	p.pub.Publish(arg)
	return &String{}, nil
}

func (p *PubsubService) Subscribe(arg *String, stream PubsubService_SubscribeServer, ) error {
	arrKey := strings.Split(arg.Key, ",")
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		data, ok := v.(*String)
		if ok {
			for _, keySub := range arrKey {
				if strings.HasPrefix(data.Key, keySub) {
					return true
				}
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&String{Key: v.(*String).Key, Value: v.(*String).Value}); err != nil {
			return err
		}
	}

	return nil
}
