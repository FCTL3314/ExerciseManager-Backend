package errormapper

type Mapper interface {
	MapError(err error) (mappedErr error, ok bool)
}

type Chain interface {
	MapError(err error) error
	GetMappers() []Mapper
}

type MapperChain struct {
	mappers []Mapper
}

func NewChain() *MapperChain {
	return &MapperChain{}
}

func (mc *MapperChain) registerMapper(mapper Mapper) {
	mc.mappers = append(mc.mappers, mapper)
}

func (mc *MapperChain) MapError(err error) error {
	for _, mapper := range mc.mappers {
		if mappedErr, ok := mapper.MapError(err); ok {
			return mappedErr
		}
	}
	return err
}

func (mc *MapperChain) GetMappers() []Mapper {
	return mc.mappers
}

func BuildAllErrorsMapperChain() *MapperChain {
	mc := NewChain()
	GORMMapperChain := BuildGORMErrorsMapperChain()
	PostgresMapperChain := BuildPostgresErrorsMapperChain()

	allMapperChains := [2]Chain{
		GORMMapperChain,
		PostgresMapperChain,
	}

	for _, mapperChain := range allMapperChains {

		for _, mapper := range mapperChain.GetMappers() {
			mc.registerMapper(mapper)
		}
	}

	return mc
}
