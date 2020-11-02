package factory

import "fmt"
//CommonInterface CommonInterface
type CommonInterface interface {
	Producer(name string) string
}

//NewCommonInterface NewCommonInterface
func NewCommonInterface(t int) CommonInterface {
	switch t {
	case 2:
		return &riceProducer{}
	case 1:
		fallthrough
	default:
		return &fruitsProducer{}
	}
}

type fruitsProducer struct{}

func (p *fruitsProducer) Producer(name string) string {
	return fmt.Sprintf("%s生产水果!", name)
}

type riceProducer struct{}

func (r *riceProducer) Producer(name string) string {
	return fmt.Sprintf("%s生产大米!", name)
}
