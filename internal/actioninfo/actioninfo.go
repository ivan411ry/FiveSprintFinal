package actioninfo
import "log"

type DataParser interface {
Parse(datastring string) error
ActionInfo() (string, error)
}

func Info(dataset []string, factory func() DataParser) {
for _, data := range dataset {
	dp := factory()
	err := dp.Parse(data)
	if err != nil {
		log.Printf("parsing error: %v", err)
		continue
	}
	info, err := dp.ActionInfo()
	if err != nil {
	log.Printf("action info error: %v", err)
	continue
}
log.Println(info)
}
}
