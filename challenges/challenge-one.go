package challenges

type Developer struct {
  Name string
  Age int
}

func GetDeveloper(name interface{}, age interface{}) Developer {
  return Developer{Name: name.(string), Age: age.(int)}
}

