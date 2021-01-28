package zoo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/maapteh/graphql-golang/model"
)

var (
	zooInstances = initZooInstances()
)

// Zoo Abstraction of zoo's
type Zoo interface {
	ListCrocodiles() []*model.Crocodile
	AddCrocodile(input *model.CrocodileInput) *model.Crocodile
	GetCrocodile(id int) *model.Crocodile
}

// ArtisZoo Amsterdams Zoo
type ArtisZoo struct {
	mux           sync.Mutex
	crocodileNest crocodileNest
}

func initZooInstances() map[Instance]Zoo {
	proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		log.Printf("Configuring HTTP client with Proxy %s", proxy)
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			panic(fmt.Errorf("Unable to parse proxy URL [%v]", err))
		}
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	zooInstances := make(map[Instance]Zoo)
	zooInstances[Artis] = &ArtisZoo{
		crocodileNest: &artisCrocodileNest{
			base: "https://test-api.loadimpact.com/",
		},
	}
	return zooInstances
}

// GetZoo retrieve a specific zoo instance
func GetZoo(instance Instance) Zoo {
	if zoo, ok := zooInstances[instance]; ok {
		return zoo
	}
	panic(fmt.Sprintf("Unable to retrieve zoo %s.", instance))
}

// Instance Zoo instance types
type Instance string

const (
	// Artis amsterdam zoo instance identifier
	Artis Instance = "Artis"
)

// ListCrocodiles list all available crocodiles in the zoo
func (artis *ArtisZoo) ListCrocodiles() []*model.Crocodile {
	return artis.crocodileNest.listCrocs()
}

// AddCrocodile add a new crocodile to the zoo
func (artis *ArtisZoo) AddCrocodile(input *model.CrocodileInput) *model.Crocodile {
	panic("Not implemented.")
}

// GetCrocodile retrieve crocodile details from Artis
func (artis *ArtisZoo) GetCrocodile(id int) *model.Crocodile {
	return artis.crocodileNest.getCroc(id)
}

type crocodileNest interface {
	getCroc(id int) *model.Crocodile
	listCrocs() []*model.Crocodile
}

type artisCrocodileNest struct {
	base string
}

func (c *artisCrocodileNest) getBody(endpoint string) (*http.Response, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unable to retrieve crocs, service responded with %v", res.StatusCode)
	}
	return res, nil
}

func (c *artisCrocodileNest) getCroc(id int) *model.Crocodile {
	result, err := c.getBody(fmt.Sprintf(c.base+"/public/crocodiles/%+v/?format=json", id))
	if err != nil {
		panic(err)
	}
	if result == nil {
		return nil
	}
	defer result.Body.Close()
	var croc *model.Crocodile
	err = json.NewDecoder(result.Body).Decode(&croc)
	if err != nil {
		panic(err)
	}
	return croc
}

func (c *artisCrocodileNest) listCrocs() []*model.Crocodile {
	result, err := c.getBody(c.base + "/public/crocodiles/?format=json")
	if err != nil {
		panic(err)
	}
	if result == nil {
		panic("No crocs found.")
	}
	defer result.Body.Close()
	var crocs []*model.Crocodile
	err = json.NewDecoder(result.Body).Decode(&crocs)
	if err != nil {
		panic(err)
	}
	return crocs
}