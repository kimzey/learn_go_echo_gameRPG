package exception

type ItemListing struct {

}

func (e *ItemListing) Error() string {
	return "item Lising failed"
}