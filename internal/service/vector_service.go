package service

type VectorService interface {
}

type vectorService struct {
}

func NewVectorService() VectorService {
	return &vectorService{}
}

//TODO add function to post data to python to process the doc to embedding
